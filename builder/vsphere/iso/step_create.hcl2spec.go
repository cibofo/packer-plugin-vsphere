// Code generated by "mapstructure-to-hcl2 -type NIC,CreateConfig"; DO NOT EDIT.

package iso

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-vsphere/builder/vsphere/common"
	"github.com/zclconf/go-cty/cty"
)

// FlatCreateConfig is an auto-generated flat version of CreateConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatCreateConfig struct {
	Version            *uint                   `mapstructure:"vm_version" cty:"vm_version" hcl:"vm_version"`
	GuestOSType        *string                 `mapstructure:"guest_os_type" cty:"guest_os_type" hcl:"guest_os_type"`
	DiskControllerType []string                `mapstructure:"disk_controller_type" cty:"disk_controller_type" hcl:"disk_controller_type"`
	Storage            []common.FlatDiskConfig `mapstructure:"storage" cty:"storage" hcl:"storage"`
	NICs               []FlatNIC               `mapstructure:"network_adapters" cty:"network_adapters" hcl:"network_adapters"`
	USBController      []string                `mapstructure:"usb_controller" cty:"usb_controller" hcl:"usb_controller"`
	Notes              *string                 `mapstructure:"notes" cty:"notes" hcl:"notes"`
}

// FlatMapstructure returns a new FlatCreateConfig.
// FlatCreateConfig is an auto-generated flat version of CreateConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*CreateConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatCreateConfig)
}

// HCL2Spec returns the hcl spec of a CreateConfig.
// This spec is used by HCL to read the fields of CreateConfig.
// The decoded values from this spec will then be applied to a FlatCreateConfig.
func (*FlatCreateConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"vm_version":           &hcldec.AttrSpec{Name: "vm_version", Type: cty.Number, Required: false},
		"guest_os_type":        &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
		"disk_controller_type": &hcldec.AttrSpec{Name: "disk_controller_type", Type: cty.List(cty.String), Required: false},
		"storage":              &hcldec.BlockListSpec{TypeName: "storage", Nested: hcldec.ObjectSpec((*common.FlatDiskConfig)(nil).HCL2Spec())},
		"network_adapters":     &hcldec.BlockListSpec{TypeName: "network_adapters", Nested: hcldec.ObjectSpec((*FlatNIC)(nil).HCL2Spec())},
		"usb_controller":       &hcldec.AttrSpec{Name: "usb_controller", Type: cty.List(cty.String), Required: false},
		"notes":                &hcldec.AttrSpec{Name: "notes", Type: cty.String, Required: false},
	}
	return s
}

// FlatNIC is an auto-generated flat version of NIC.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatNIC struct {
	Network     *string `mapstructure:"network" cty:"network" hcl:"network"`
	NetworkCard *string `mapstructure:"network_card" required:"true" cty:"network_card" hcl:"network_card"`
	MacAddress  *string `mapstructure:"mac_address" cty:"mac_address" hcl:"mac_address"`
	Passthrough *bool   `mapstructure:"passthrough" cty:"passthrough" hcl:"passthrough"`
}

// FlatMapstructure returns a new FlatNIC.
// FlatNIC is an auto-generated flat version of NIC.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*NIC) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatNIC)
}

// HCL2Spec returns the hcl spec of a NIC.
// This spec is used by HCL to read the fields of NIC.
// The decoded values from this spec will then be applied to a FlatNIC.
func (*FlatNIC) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"network":      &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"network_card": &hcldec.AttrSpec{Name: "network_card", Type: cty.String, Required: false},
		"mac_address":  &hcldec.AttrSpec{Name: "mac_address", Type: cty.String, Required: false},
		"passthrough":  &hcldec.AttrSpec{Name: "passthrough", Type: cty.Bool, Required: false},
	}
	return s
}
