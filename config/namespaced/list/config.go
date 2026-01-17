package list

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "list"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_list", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "List"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_list_item", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ListItem"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["list_id"] = config.Reference{
			TerraformName: "cloudflare_list",
		}
	})
}
