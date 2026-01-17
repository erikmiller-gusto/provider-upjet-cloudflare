package argo

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "argo"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_argo", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Argo"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_argo_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ArgoTunnel"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_tiered_cache", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TieredCache"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_regional_tiered_cache", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RegionalTieredCache"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_static_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StaticRoute"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_gre_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "GRETunnel"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_ipsec_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "IPSecTunnel"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Magic WAN
	p.AddResourceConfigurator("cloudflare_magic_firewall_ruleset", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicFirewallRuleset"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_wan_gre_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicWANGRETunnel"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_wan_ipsec_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicWANIPSecTunnel"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_magic_wan_static_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "MagicWANStaticRoute"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
