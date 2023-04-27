/*
Copyright 2021 Upbound Inc.
*/

package null

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("azuredevops_variable_group", func(r *ujconfig.Resource) {
		r.Kind = "VariableGroup"
		// And other overrides.
	})
}
