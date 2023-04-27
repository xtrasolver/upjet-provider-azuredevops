/*
Copyright 2021 Upbound Inc.
*/

package environment

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("azuredevops_environment", func(r *ujconfig.Resource) {
		r.Kind = "Environment"
		// And other overrides.
	})
}
