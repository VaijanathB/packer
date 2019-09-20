package dtl

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/devtestlabs/mgmt/2018-09-15/dtl"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"github.com/hashicorp/packer/builder/azure/common/template"
)

type templateFactoryFuncDtl func(*Config) (*dtl.LabVirtualMachineCreationParameter, error)

func GetKeyVaultDeployment(config *Config) (*resources.Deployment, error) {
	params := &template.TemplateParameters{
		KeyVaultName:        &template.TemplateParameter{Value: config.tmpKeyVaultName},
		KeyVaultSecretValue: &template.TemplateParameter{Value: config.winrmCertificate},
		ObjectId:            &template.TemplateParameter{Value: config.ObjectID},
		TenantId:            &template.TemplateParameter{Value: config.TenantID},
	}

	builder, _ := template.NewTemplateBuilder(template.KeyVault)
	builder.SetTags(&config.AzureTags)

	doc, _ := builder.ToJSON()
	return createDeploymentParameters(*doc, params)
}

func newBool(val bool) *bool {
	b := true
	if val == b {
		return &b
	} else {
		b = false
		return &b
	}
}

func getCustomImageId (config *Config) *string {
	if config.CustomManagedImageName != "" && config.CustomManagedImageResourceGroupName != "" {
		customManagedImageID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/images/%s",
			config.SubscriptionID,
			config.CustomManagedImageResourceGroupName,
			config.CustomManagedImageName)
		return &customManagedImageID
	}
	return nil
}

func GetVirtualMachineDeployment(config *Config) (*dtl.LabVirtualMachineCreationParameter, error) {
	
	galleryImageRef := dtl.GalleryImageReference{
		Offer:     &config.ImageOffer,
		Publisher: &config.ImagePublisher,
		Sku:       &config.ImageSku,
		OsType:    &config.OSType,
		Version:   &config.ImageVersion,
	}

	labVirtualNetworkID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/virtualnetworks/%s",
		config.SubscriptionID,
		config.tmpResourceGroupName,
		config.LabName,
		config.LabVirtualNetworkName)

	dtlArtifacts := []dtl.ArtifactInstallProperties{}

	if config.DtlArtifacts != nil {
		for i := range config.DtlArtifacts {
			if config.DtlArtifacts[i].RepositoryName == "" {
				config.DtlArtifacts[i].RepositoryName = "public repo"
			}
			config.DtlArtifacts[i].ArtifactId = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/artifactSources/%s/artifacts/%s",
				config.SubscriptionID,
				config.tmpResourceGroupName,
				config.LabName,
				config.DtlArtifacts[i].RepositoryName,
				config.DtlArtifacts[i].ArtifactName)

			dparams := []dtl.ArtifactParameterProperties{}
			for j := range config.DtlArtifacts[i].Parameters {

				dp := &dtl.ArtifactParameterProperties{}
				dp.Name = &config.DtlArtifacts[i].Parameters[j].Name
				dp.Value = &config.DtlArtifacts[i].Parameters[j].Value

				dparams = append(dparams, *dp)
			}
			dtlArtifact := &dtl.ArtifactInstallProperties{
				ArtifactTitle: &config.DtlArtifacts[i].ArtifactName,
				ArtifactID:    &config.DtlArtifacts[i].ArtifactId,
				Parameters:    &dparams,
			}
			dtlArtifacts = append(dtlArtifacts, *dtlArtifact)
		}
	}

	if strings.ToLower(config.OSType) == "windows" {
		// Add mandatory Artifact
		var winrma = "windows-winrm"
		var artifactid = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/artifactSources/public repo/artifacts/%s",
			config.SubscriptionID,
			config.tmpResourceGroupName,
			config.LabName,
			winrma)

		var hostname = "hostName"
		var hostNameValue = fmt.Sprintf("%s.%s.cloudapp.azure.com", config.VMName, config.Location)
		dparams := []dtl.ArtifactParameterProperties{}
		dp := &dtl.ArtifactParameterProperties{}
		dp.Name = &hostname
		dp.Value = &hostNameValue
		dparams = append(dparams, *dp)

		winrmArtifact := &dtl.ArtifactInstallProperties{
			ArtifactTitle: &winrma,
			ArtifactID:    &artifactid,
			Parameters:    &dparams,
		}
		dtlArtifacts = append(dtlArtifacts, *winrmArtifact)
	}

	labMachineProps := &dtl.LabVirtualMachineCreationParameterProperties{
		CreatedByUserID:            &config.ClientConfig.ClientID,
		OwnerObjectID:              &config.ClientConfig.ObjectID,
		OsType:                     &config.OSType,
		Size:                       &config.VMSize,
		UserName:                   &config.UserName,
		Password:                   &config.Password,
		SSHKey:                     &config.sshAuthorizedKey,
		IsAuthenticationWithSSHKey: newBool(true),
		LabSubnetName:              &config.LabSubnetName,
		LabVirtualNetworkID:        &labVirtualNetworkID,
		DisallowPublicIPAddress:    newBool(false),
		GalleryImageReference:      &galleryImageRef,
		CustomImageID:              getCustomImageId(config),

		AllowClaim:                   newBool(false),
		StorageType:                  &config.StorageType,
		VirtualMachineCreationSource: dtl.FromGalleryImage,
		Artifacts:                    &dtlArtifacts,
	}

	labMachine := &dtl.LabVirtualMachineCreationParameter{
		Name:     &config.tmpComputeName,
		Location: &config.Location,
		Tags:     config.AzureTags,
		LabVirtualMachineCreationParameterProperties: labMachineProps,
	}

	return labMachine, nil
}

func createDeploymentParameters(doc string, parameters *template.TemplateParameters) (*resources.Deployment, error) {
	var template map[string]interface{}
	err := json.Unmarshal(([]byte)(doc), &template)
	if err != nil {
		return nil, err
	}

	bs, err := json.Marshal(*parameters)
	if err != nil {
		return nil, err
	}

	var templateParameters map[string]interface{}
	err = json.Unmarshal(bs, &templateParameters)
	if err != nil {
		return nil, err
	}

	return &resources.Deployment{
		Properties: &resources.DeploymentProperties{
			Mode:       resources.Incremental,
			Template:   &template,
			Parameters: &templateParameters,
		},
	}, nil
}
