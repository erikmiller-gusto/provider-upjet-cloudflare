package logpush

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "logpush"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_logpush_job", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogpushJob"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_logpull_retention", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LogpullRetention"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})
}
