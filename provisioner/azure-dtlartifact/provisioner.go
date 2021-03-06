// This package implements a provisioner for Packer that uses
package devtestlabsartifacts

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/devtestlabs/mgmt/2018-09-15/dtl"
	dtlBuilder "github.com/hashicorp/packer/builder/azure/dtl"
	"github.com/hashicorp/packer/packer"

	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/provisioner"
	"github.com/hashicorp/packer/template/interpolate"
)

type DtlArtifact struct {
	ArtifactName string              `mapstructure:"artifact_name"`
	ArtifactId   string              `mapstructure:"artifact_id"`
	Parameters   []ArtifactParameter `mapstructure:"parameters"`
}

type ArtifactParameter struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
	Type  string `mapstructure:"type"`
}

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	// Authentication via OAUTH
	ClientConfig dtlBuilder.ClientConfig `mapstructure:",squash"`

	DtlArtifacts []DtlArtifact `mapstructure:"dtl_artifacts"`
	LabName      string        `mapstructure:"lab_name"`

	ResourceGroupName string `mapstructure:"resource_group_name"`

	VMName string `mapstructure:"vm_name"`

	AzureTags map[string]*string `mapstructure:"azure_tags"`

	Json map[string]interface{}

	ctx interpolate.Context
}

type Provisioner struct {
	config        Config
	communicator  packer.Communicator
	guestCommands *provisioner.GuestCommands
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	// // Create passthrough for winrm password so we can fill it in once we know
	// // it
	// p.config.ctx.Data = &EnvVarsTemplate{
	// 	WinRMPassword: `{{.WinRMPassword}}`,
	// }
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{
				"execute_command",
			},
		},
	}, raws...)
	if err != nil {
		return err
	}

	p.config.ClientConfig.CloudEnvironmentName = "Public"

	return nil
}

func (p *Provisioner) Communicator() packer.Communicator {
	return p.communicator
}

func (p *Provisioner) Provision(ctx context.Context, ui packer.Ui, comm packer.Communicator) error {

	p.communicator = comm

	err := p.config.ClientConfig.SetCloudEnvironment()
	if err != nil {
		ui.Say(fmt.Sprintf("Error saving debug key: %s", err))
		return nil
	}

	// User's intent to use MSI is indicated with empty subscription id, tenant, client id, client cert, client secret and jwt.
	// FillParameters function will set subscription and tenant id here. Therefore getServicePrincipalTokens won't select right auth type.
	// If we run this after getServicePrincipalTokens call then getServicePrincipalTokens won't have tenant id.
	if !p.config.ClientConfig.UseMSI() {
		if err := newConfigRetriever().FillParameters(&p.config); err != nil {
			return err
		}
	}

	spnCloud, err := p.config.ClientConfig.GetServicePrincipalTokens(ui.Say)

	ui.Message("Creating Azure Resource Manager (ARM) client ...")
	azureClient, err := dtlBuilder.NewAzureClient(
		p.config.ClientConfig.SubscriptionID,
		"",
		p.config.ClientConfig.CloudEnvironment,
		0,
		spnCloud)

	if err != nil {
		ui.Say(fmt.Sprintf("Error saving debug key: %s", err))
		return err
	}

	ui.Say("Installing Artifact DTL")
	dtlArtifacts := []dtl.ArtifactInstallProperties{}

	if p.config.DtlArtifacts != nil {
		for i := range p.config.DtlArtifacts {
			p.config.DtlArtifacts[i].ArtifactId = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/artifactSources/public repo/artifacts/%s",
				p.config.ClientConfig.SubscriptionID,
				p.config.ResourceGroupName,
				p.config.LabName,
				p.config.DtlArtifacts[i].ArtifactName)

			dparams := []dtl.ArtifactParameterProperties{}
			for j := range p.config.DtlArtifacts[i].Parameters {
				dp := &dtl.ArtifactParameterProperties{}
				dp.Name = &p.config.DtlArtifacts[i].Parameters[j].Name
				dp.Value = &p.config.DtlArtifacts[i].Parameters[j].Value

				dparams = append(dparams, *dp)
			}
			Aip := dtl.ArtifactInstallProperties{
				ArtifactID:    &p.config.DtlArtifacts[i].ArtifactId,
				Parameters:    &dparams,
				ArtifactTitle: &p.config.DtlArtifacts[i].ArtifactName,
			}
			dtlArtifacts = append(dtlArtifacts, Aip)
		}
	}

	dtlApplyArifactRequest := dtl.ApplyArtifactsRequest{
		Artifacts: &dtlArtifacts,
	}

	ui.Say("Applying artifact ")
	f, err := azureClient.DtlVirtualMachineClient.ApplyArtifacts(ctx, p.config.ResourceGroupName, p.config.LabName, p.config.VMName, dtlApplyArifactRequest)

	if err == nil {
		err = f.WaitForCompletionRef(ctx, azureClient.DtlVirtualMachineClient.Client)
	}
	if err != nil {
		ui.Say(fmt.Sprintf("Error Applying artifact: %s", err))
	}
	ui.Say("Aftifact installed")
	return err
}
