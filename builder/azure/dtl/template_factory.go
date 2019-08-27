package dtl

import (
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/devtestlabs/mgmt/2018-09-15/dtl"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"github.com/hashicorp/packer/builder/azure/common/template"
)

type templateFactoryFuncDtl func(*Config) (*dtl.LabVirtualMachine, error)
type templateFactoryFunc func(*Config) (*resources.Deployment, error)

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
func GetVirtualMachineDeployment(config *Config) (*dtl.LabVirtualMachine, error) {
	galleryImageRef := &dtl.GalleryImageReference{
		Offer:     &config.ImageOffer,
		Publisher: &config.ImagePublisher,
		Sku:       &config.ImageSku,
		OsType:    &config.OSType,
		Version:   &config.ImageVersion,
	}

	// /subscriptions/cba4e087-aceb-44f0-970e-65e96eff4081/resourcegroups/packerrg/providers/microsoft.devtestlab/labs/packerlab/virtualnetworks/dtlpackerlab
	labVirtualNetworkID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/virtualnetworks/%s",
		config.SubscriptionID,
		config.tmpResourceGroupName,
		config.LabName,
		config.LabVirtualNetworkName)

	dtlArtifacts := []dtl.ArtifactInstallProperties{}

	//config.DtlArtifacts[0].Parameters[1].Value = config.winrmCertificate
	if config.DtlArtifacts != nil {
		for i := range config.DtlArtifacts {
			///subscriptions/cba4e087-aceb-44f0-970e-65e96eff4081/resourceGroups/packerrg/providers/Microsoft.DevTestLab/labs/packerlab/artifactSources/public repo/artifacts/linux-apt-package"
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

	labMachineProps := &dtl.LabVirtualMachineProperties{
		CreatedByUserID:              &config.ClientConfig.ClientID,
		OwnerObjectID:                &config.ClientConfig.ObjectID,
		OsType:                       &config.OSType,
		Size:                         &config.VMSize,
		UserName:                     &config.UserName,
		Password:                     &config.Password,
		SSHKey:                       &config.sshAuthorizedKey,
		IsAuthenticationWithSSHKey:   newBool(true),
		LabSubnetName:                &config.LabSubnetName,
		LabVirtualNetworkID:          &labVirtualNetworkID,
		DisallowPublicIPAddress:      newBool(false),
		GalleryImageReference:        galleryImageRef,
		AllowClaim:                   newBool(false),
		StorageType:                  &config.StorageType,
		VirtualMachineCreationSource: dtl.FromGalleryImage,
		Artifacts:                    &dtlArtifacts,
	}

	labMachine := &dtl.LabVirtualMachine{
		Location:                    &config.Location,
		Tags:                        config.AzureTags,
		LabVirtualMachineProperties: labMachineProps,
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
