package firewall

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "firewall"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_filter", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_firewall_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FirewallRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["filter_id"] = config.Reference{
			TerraformName: "cloudflare_filter",
		}
	})

	p.AddResourceConfigurator("cloudflare_rate_limit", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RateLimit"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_rate_limiting_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RateLimitingRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_waf_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WAFRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_waf_override", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WAFOverride"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_waf_package", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WAFPackage"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_waf_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WAFGroup"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_ruleset", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Ruleset"
		// ruleset can be at zone or account level
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_managed_headers", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ManagedHeaders"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_url_normalization_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "URLNormalizationSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_user_agent_blocking_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UserAgentBlockingRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_page_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PageRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_bot_management", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BotManagement"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_leaked_credential_check", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LeakedCredentialCheck"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_leaked_credential_check_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LeakedCredentialCheckRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_cache_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CacheRule"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_cache_reserve", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CacheReserve"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})
}
