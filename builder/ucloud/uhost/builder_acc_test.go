package uhost

import (
	"fmt"
	"github.com/hashicorp/packer/packer"
	"os"
	"testing"

	builderT "github.com/hashicorp/packer/helper/builder/testing"
)

func TestBuilderAcc_validateRegion(t *testing.T) {
	t.Parallel()

	if os.Getenv(builderT.TestEnvVar) == "" {
		t.Skip(fmt.Sprintf("Acceptance tests skipped unless env '%s' set", builderT.TestEnvVar))
		return
	}

	testAccPreCheck(t)

	access := &AccessConfig{Region: "cn-bj2"}
	err := access.Config()
	if err != nil {
		t.Fatalf("Error on initing UCloud AccessConfig, %s", err)
	}

	err = access.ValidateRegion("cn-sh2")
	if err != nil {
		t.Fatalf("Expected pass with valid region but failed: %s", err)
	}

	err = access.ValidateRegion("invalidRegion")
	if err == nil {
		t.Fatal("Expected failure due to invalid region but passed")
	}
}

func TestBuilderAcc_basic(t *testing.T) {
	t.Parallel()
	builderT.Test(t, builderT.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Builder:  &Builder{},
		Template: testBuilderAccBasic,
	})
}

const testBuilderAccBasic = `
{	"builders": [{
		"type": "test",
		"region": "cn-bj2",
		"availability_zone": "cn-bj2-02",
		"instance_type": "n-basic-2",
		"source_image_id":"uimage-f1chxn",
		"ssh_username":"root",
		"image_name": "packer-test-basic_{{timestamp}}"
	}]
}`

func TestBuilderAcc_ubuntu(t *testing.T) {
	t.Parallel()
	builderT.Test(t, builderT.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Builder:  &Builder{},
		Template: testBuilderAccUbuntu,
	})
}

const testBuilderAccUbuntu = `
{	"builders": [{
		"type": "test",
		"region": "cn-bj2",
		"availability_zone": "cn-bj2-02",
		"instance_type": "n-basic-2",
		"source_image_id":"uimage-irofn4",
		"ssh_username":"ubuntu",
		"image_name": "packer-test-ubuntu_{{timestamp}}"
	}]
}`

func TestBuilderAcc_regionCopy(t *testing.T) {
	t.Parallel()
	projectId := os.Getenv("UCLOUD_PROJECT_ID")
	builderT.Test(t, builderT.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Builder:  &Builder{},
		Template: testBuilderAccRegionCopy(projectId),
		Check: checkRegionCopy(
			projectId,
			[]ImageDestination{
				{projectId, "cn-sh2", "packer-test-regionCopy-sh", "test"},
			}),
	})
}

func testBuilderAccRegionCopy(projectId string) string {
	return fmt.Sprintf(`
{
	"builders": [{
		"type": "test",
		"region": "cn-bj2",
		"availability_zone": "cn-bj2-02",
		"instance_type": "n-basic-2",
		"source_image_id":"uimage-f1chxn",
		"ssh_username":"root",
		"image_name": "packer-test-regionCopy-bj",
		"image_copy_to_mappings": [{
			"project_id":  	%q,
			"region":		"cn-sh2",
			"name":			"packer-test-regionCopy-sh",
			"description": 	"test"
		}]
	}]
}`, projectId)
}

func checkRegionCopy(projectId string, imageDst []ImageDestination) builderT.TestCheckFunc {
	return func(artifacts []packer.Artifact) error {
		if len(artifacts) > 1 {
			return fmt.Errorf("more than 1 artifact")
		}

		artifactSet := artifacts[0]
		artifact, ok := artifactSet.(*Artifact)
		if !ok {
			return fmt.Errorf("unknown artifact: %#v", artifactSet)
		}

		destSet := newImageInfoSet(nil)
		for _, dest := range imageDst {
			destSet.Set(imageInfo{
				Region:    dest.Region,
				ProjectId: dest.ProjectId,
			})
		}

		for _, r := range artifact.UCloudImages.GetAll() {
			if r.ProjectId == projectId && r.Region == "cn-bj2" {
				destSet.Remove(r.Id())
				continue
			}

			if destSet.Get(r.ProjectId, r.Region) == nil {
				return fmt.Errorf("project%s : region%s is not the target but found in artifacts", r.ProjectId, r.Region)
			}

			destSet.Remove(r.Id())
		}

		if len(destSet.GetAll()) > 0 {
			return fmt.Errorf("the following copying targets not found in corresponding artifacts : %#v", destSet.GetAll())
		}

		client, _ := testUCloudClient()
		for _, r := range artifact.UCloudImages.GetAll() {
			if r.ProjectId == projectId && r.Region == "cn-bj2" {
				continue
			}
			imageSet, err := client.describeImageByInfo(r.ProjectId, r.Region, r.ImageId)
			if err != nil {
				if isNotFoundError(err) {
					return fmt.Errorf("image %s in artifacts can not be found", r.ImageId)
				}
				return err
			}

			if r.Region == "cn-sh2" && imageSet.ImageName != "packer-test-regionCopy-sh" {
				return fmt.Errorf("the name of image %q in artifacts should be %s, got %s", r.ImageId, "packer-test-regionCopy-sh", imageSet.ImageName)
			}
		}

		return nil
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("UCLOUD_PUBLIC_KEY"); v == "" {
		t.Fatal("UCLOUD_PUBLIC_KEY must be set for acceptance tests")
	}

	if v := os.Getenv("UCLOUD_PRIVATE_KEY"); v == "" {
		t.Fatal("UCLOUD_PRIVATE_KEY must be set for acceptance tests")
	}

	if v := os.Getenv("UCLOUD_PROJECT_ID"); v == "" {
		t.Fatal("UCLOUD_PROJECT_ID must be set for acceptance tests")
	}
}

func testUCloudClient() (*UCloudClient, error) {
	access := &AccessConfig{Region: "cn-bj2"}
	err := access.Config()
	if err != nil {
		return nil, err
	}
	client, err := access.Client()
	if err != nil {
		return nil, err
	}

	return client, nil
}
