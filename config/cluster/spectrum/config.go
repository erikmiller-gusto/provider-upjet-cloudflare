package spectrum

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "spectrum"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_spectrum_application", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SpectrumApplication"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})
}
