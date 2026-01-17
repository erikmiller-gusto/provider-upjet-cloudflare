package access

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "access"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_access_application", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessApplication"
		// access_application can be at zone or account level
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessPolicy"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["application_id"] = config.Reference{
			TerraformName: "cloudflare_access_application",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessGroup"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_identity_provider", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessIdentityProvider"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_mutual_tls_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessMutualTLSCertificate"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_service_token", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessServiceToken"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_ca_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessCACertificate"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_keys_configuration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessKeysConfiguration"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_organization", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessOrganization"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_tag", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessTag"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_custom_page", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessCustomPage"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_access_bookmark", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccessBookmark"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
