package dns

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "dns"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_dns_record", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSRecord"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_record", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Record"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_dns_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_dns_firewall", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "DNSFirewall"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_account_dns_firewall_cluster", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AccountDNSFirewallCluster"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_secondary_dns_primary_zone", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_secondary_dns_secondary_zone", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_secondary_dns_tsig", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
