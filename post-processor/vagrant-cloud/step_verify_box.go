package vagrantcloud

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type Box struct {
	Tag string `json:"tag"`
}

type stepVerifyBox struct {
}

func (s *stepVerifyBox) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*VagrantCloudClient)
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("config").(Config)

	ui.Say(fmt.Sprintf("Verifying box is accessible: %s", config.Tag))

	path := fmt.Sprintf("box/%s", config.Tag)
	resp, err := client.Get(path)

	if err != nil {
		state.Put("error", fmt.Errorf("Error retrieving box: %s", err))
		return multistep.ActionHalt
	}

	box := &Box{}

	if err = decodeBody(resp, box); err != nil {
		state.Put("error", fmt.Errorf("Error parsing box response: %s", err))
		return multistep.ActionHalt
	}

	if box.Tag != config.Tag {
		state.Put("error", fmt.Errorf("Could not verify box: %s", err))
		return multistep.ActionHalt
	}

	// Keep the box in state for later
	state.Put("box", box)

	// Box exists and is accessible
	return multistep.ActionContinue
}

func (s *stepVerifyBox) Cleanup(state multistep.StateBag) {
	// no cleanup needed
}
