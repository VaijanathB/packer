// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package saltmasterless

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	SkipBootstrap       *bool             `mapstructure:"skip_bootstrap" cty:"skip_bootstrap"`
	BootstrapArgs       *string           `mapstructure:"bootstrap_args" cty:"bootstrap_args"`
	DisableSudo         *bool             `mapstructure:"disable_sudo" cty:"disable_sudo"`
	CustomState         *string           `mapstructure:"custom_state" cty:"custom_state"`
	MinionConfig        *string           `mapstructure:"minion_config" cty:"minion_config"`
	GrainsFile          *string           `mapstructure:"grains_file" cty:"grains_file"`
	LocalStateTree      *string           `mapstructure:"local_state_tree" cty:"local_state_tree"`
	LocalPillarRoots    *string           `mapstructure:"local_pillar_roots" cty:"local_pillar_roots"`
	RemoteStateTree     *string           `mapstructure:"remote_state_tree" cty:"remote_state_tree"`
	RemotePillarRoots   *string           `mapstructure:"remote_pillar_roots" cty:"remote_pillar_roots"`
	TempConfigDir       *string           `mapstructure:"temp_config_dir" cty:"temp_config_dir"`
	NoExitOnFailure     *bool             `mapstructure:"no_exit_on_failure" cty:"no_exit_on_failure"`
	LogLevel            *string           `mapstructure:"log_level" cty:"log_level"`
	SaltCallArgs        *string           `mapstructure:"salt_call_args" cty:"salt_call_args"`
	SaltBinDir          *string           `mapstructure:"salt_bin_dir" cty:"salt_bin_dir"`
	CmdArgs             *string           `cty:"cmd_args"`
	GuestOSType         *string           `mapstructure:"guest_os_type" cty:"guest_os_type"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a FlatConfig.
// This spec is used by HCL to read the fields of FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"skip_bootstrap":             &hcldec.AttrSpec{Name: "skip_bootstrap", Type: cty.Bool, Required: false},
		"bootstrap_args":             &hcldec.AttrSpec{Name: "bootstrap_args", Type: cty.String, Required: false},
		"disable_sudo":               &hcldec.AttrSpec{Name: "disable_sudo", Type: cty.Bool, Required: false},
		"custom_state":               &hcldec.AttrSpec{Name: "custom_state", Type: cty.String, Required: false},
		"minion_config":              &hcldec.AttrSpec{Name: "minion_config", Type: cty.String, Required: false},
		"grains_file":                &hcldec.AttrSpec{Name: "grains_file", Type: cty.String, Required: false},
		"local_state_tree":           &hcldec.AttrSpec{Name: "local_state_tree", Type: cty.String, Required: false},
		"local_pillar_roots":         &hcldec.AttrSpec{Name: "local_pillar_roots", Type: cty.String, Required: false},
		"remote_state_tree":          &hcldec.AttrSpec{Name: "remote_state_tree", Type: cty.String, Required: false},
		"remote_pillar_roots":        &hcldec.AttrSpec{Name: "remote_pillar_roots", Type: cty.String, Required: false},
		"temp_config_dir":            &hcldec.AttrSpec{Name: "temp_config_dir", Type: cty.String, Required: false},
		"no_exit_on_failure":         &hcldec.AttrSpec{Name: "no_exit_on_failure", Type: cty.Bool, Required: false},
		"log_level":                  &hcldec.AttrSpec{Name: "log_level", Type: cty.String, Required: false},
		"salt_call_args":             &hcldec.AttrSpec{Name: "salt_call_args", Type: cty.String, Required: false},
		"salt_bin_dir":               &hcldec.AttrSpec{Name: "salt_bin_dir", Type: cty.String, Required: false},
		"cmd_args":                   &hcldec.AttrSpec{Name: "cmd_args", Type: cty.String, Required: false},
		"guest_os_type":              &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
	}
	return s
}
