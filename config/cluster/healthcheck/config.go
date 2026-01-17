package healthcheck

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "healthcheck"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_healthcheck", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Healthcheck"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})
}
