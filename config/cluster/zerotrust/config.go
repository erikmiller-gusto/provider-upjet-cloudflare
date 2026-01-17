package zerotrust

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "zerotrust"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_zero_trust_list", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustList"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_gateway_policy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustGatewayPolicy"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_gateway_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustGatewaySettings"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_tunnel_cloudflared", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustTunnelCloudflared"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_tunnel_cloudflared_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustTunnelCloudflaredConfig"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["tunnel_id"] = config.Reference{
			TerraformName: "cloudflare_zero_trust_tunnel_cloudflared",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_tunnel_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustTunnelRoute"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["tunnel_id"] = config.Reference{
			TerraformName: "cloudflare_zero_trust_tunnel_cloudflared",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_posture_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDevicePostureRule"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_posture_integration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDevicePostureIntegration"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_profiles", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceProfiles"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_managed_networks", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceManagedNetworks"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_certificates", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceCertificates"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dex_test", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDEXTest"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_device_dex_test", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDeviceDEXTest"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPProfile"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_custom_profile", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPCustomProfile"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_entry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPEntry"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_dlp_custom_entry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustDLPCustomEntry"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_risk_behavior", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustRiskBehavior"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_risk_scoring_integration", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustRiskScoringIntegration"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_access_short_lived_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustAccessShortLivedCertificate"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_gateway_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustGatewayCertificate"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_local_domain_fallback", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustLocalDomainFallback"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_zero_trust_split_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZeroTrustSplitTunnel"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Legacy Teams resources (being replaced by Zero Trust)
	p.AddResourceConfigurator("cloudflare_teams_account", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TeamsAccount"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_teams_list", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TeamsList"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_teams_location", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TeamsLocation"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_teams_proxy_endpoint", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TeamsProxyEndpoint"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_teams_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TeamsRule"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_fallback_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FallbackDomain"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_split_tunnel", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SplitTunnel"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Browser Isolation
	p.AddResourceConfigurator("cloudflare_browser_isolation_banner", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BrowserIsolationBanner"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_browser_isolation_certificate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BrowserIsolationCertificate"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_browser_isolation_permissions", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BrowserIsolationPermissions"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
