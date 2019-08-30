package dtl

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer/builder/azure/common/constants"
	"github.com/hashicorp/packer/helper/multistep"
	"github.com/hashicorp/packer/packer"
)

type StepDeleteVirtualMachine struct {
	client *AzureClient
	config *Config
	delete func(ctx context.Context, resourceGroupName string, computeName string, state multistep.StateBag) error
	say    func(message string)
	error  func(e error)
}

func NewStepDeleteVirtualMachine(client *AzureClient, ui packer.Ui, config *Config) *StepDeleteVirtualMachine {
	var step = &StepDeleteVirtualMachine{
		client: client,
		config: config,
		say:    func(message string) { ui.Say(message) },
		error:  func(e error) { ui.Error(e.Error()) },
	}

	step.delete = step.deleteVirtualMachine
	return step
}

func (s *StepDeleteVirtualMachine) deleteVirtualMachine(ctx context.Context, resourceGroupName string, vmName string, state multistep.StateBag) error {
	if state.Get(constants.ArmIsExistingResourceGroup).(bool) {
		f, err := s.client.DtlVirtualMachineClient.Delete(ctx, resourceGroupName, s.config.LabName, vmName)
		if err == nil {
			err = f.WaitForCompletionRef(ctx, s.client.VirtualMachinesClient.Client)
		}
		if err != nil {
			s.say(s.client.LastError.Error())
		}

		return err
	} else {
		s.say("\nThe resource group was created by Packer, deleting ...")
		f, err := s.client.GroupsClient.Delete(ctx, resourceGroupName)
		if err == nil {
			if state.Get(constants.ArmAsyncResourceGroupDelete).(bool) {
				// No need to wait for the completion for delete if request is Accepted
				s.say(fmt.Sprintf("\nResource Group is being deleted, not waiting for deletion due to config. Resource Group Name '%s'", resourceGroupName))
			} else {
				f.WaitForCompletionRef(ctx, s.client.GroupsClient.Client)
			}

		}

		if err != nil {
			s.say(s.client.LastError.Error())
		}
		return err
	}
}

func (s *StepDeleteVirtualMachine) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	s.say("Deleting the virtual machine ...")

	var resourceGroupName = state.Get(constants.ArmResourceGroupName).(string)
	var computeName = state.Get(constants.ArmComputeName).(string)

	s.say(fmt.Sprintf(" -> ResourceGroupName : '%s'", resourceGroupName))
	s.say(fmt.Sprintf(" -> ComputeName       : '%s'", computeName))

	err := s.deleteVirtualMachine(ctx, resourceGroupName, computeName, state)

	s.say("Deleting virtual machine ...Complete")
	return processStepResult(err, s.error, state)
}

func (*StepDeleteVirtualMachine) Cleanup(multistep.StateBag) {
}
