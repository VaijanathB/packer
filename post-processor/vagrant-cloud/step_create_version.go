package vagrantcloud

import (
	"fmt"
    "github.com/mitchellh/multistep"
    "github.com/mitchellh/packer/packer"
)

type Version struct {
	Version string `json:"version"`
	Number  uint `json:"number"`
}

type NewVersion struct {
    Version  string `json:"version"`
}

type stepCreateVersion struct {
	number uint // number of the version, if needed in cleanup
}

func (s *stepCreateVersion) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*VagrantCloudClient)
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("config").(Config)
	box := state.Get("box").(*Box)

    path := fmt.Sprintf("box/%s/versions", box.Tag)

	// Wrap the version in a version object for the API
	wrapper := make(map[string]interface{})
	wrapper["version"] = NewVersion{Version: config.Version}

    ui.Say(fmt.Sprintf("Creating version: %s", config.Version))

	resp, err := client.Post(path, wrapper)
    version := &Version{}

	if err != nil || (resp.StatusCode != 200) {
        cloudErrors := &VagrantCloudErrors{};
        err = decodeBody(resp, cloudErrors);
        state.Put("error", fmt.Errorf("Error creating version: %s", cloudErrors.FormatErrors()))
		return multistep.ActionHalt
	}

	if err = decodeBody(resp, version); err != nil {
		state.Put("error", fmt.Errorf("Error parsing version response: %s", err))
        return multistep.ActionHalt
	}

    // Save the number for cleanup
    s.number = version.Number

    state.Put("version", version)

	return multistep.ActionContinue
}

func (s *stepCreateVersion) Cleanup(state multistep.StateBag) {
    // If we didn't save the version number, it likely doesn't exist
    if (s.number == 0) {
        return
    }

    client := state.Get("client").(*VagrantCloudClient)
    ui := state.Get("ui").(packer.Ui)
    box := state.Get("box").(*Box)

    path := fmt.Sprintf("box/%s/version/%s", box.Tag, s.number)

    // No need for resp from the cleanup DELETE
    _, err := client.Delete(path)

    if err != nil {
        ui.Error(fmt.Sprintf("Error destroying version: %s", err))
    }
}
