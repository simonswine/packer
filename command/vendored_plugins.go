package command

import (
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"

	// Previously core-bundled components, split into their own plugins but
	// still vendored with Packer for now. Importing as library instead of
	// forcing use of packer init, until packer v1.8.0
	exoscaleimportpostprocessor "github.com/exoscale/packer-plugin-exoscale/post-processor/exoscale-import"
)

// VendoredBuilders are builder components that were once bundle with Packer core, but are now being shim with there multi-component counterparts.
var VendoredBuilders = map[string]packersdk.Builder{}

// VendoredProvisioners are components that were once bundle with Packer core, but are now being shim with there multi-component counterparts.
var VendoredProvisioners = map[string]packersdk.Provisioner{}

// VendoredPostProcessors are components that were once bundle with Packer core, but are now being shim with there multi-component counterparts.
var VendoredPostProcessors = map[string]packersdk.PostProcessor{
	"exoscale-import": new(exoscaleimportpostprocessor.PostProcessor),
}

// Upon init lets load up any plugins that were vendored manually into the default
// set of plugins.
func init() {
	for k, v := range VendoredBuilders {
		if _, ok := Builders[k]; ok {
			continue
		}
		Builders[k] = v
	}

	for k, v := range VendoredProvisioners {
		if _, ok := Provisioners[k]; ok {
			continue
		}
		Provisioners[k] = v
	}

	for k, v := range VendoredPostProcessors {
		if _, ok := PostProcessors[k]; ok {
			continue
		}
		PostProcessors[k] = v
	}
}
