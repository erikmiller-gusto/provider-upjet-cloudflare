package misc

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "misc"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	// Web3 Hostname
	p.AddResourceConfigurator("cloudflare_web3_hostname", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Web3Hostname"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	// Address Map & BYO IP
	// Note: r.Kind is set to avoid generating a directory called "map" which is a Go reserved keyword
	p.AddResourceConfigurator("cloudflare_address_map", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AddressMap"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_byo_ip_prefix", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "BYOIPPrefix"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Regional Hostname
	p.AddResourceConfigurator("cloudflare_regional_hostname", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RegionalHostname"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	// Snippet
	p.AddResourceConfigurator("cloudflare_snippet", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Snippet"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_snippet_rules", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SnippetRules"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	// Turnstile Widget
	p.AddResourceConfigurator("cloudflare_turnstile_widget", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "TurnstileWidget"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Web Analytics
	p.AddResourceConfigurator("cloudflare_web_analytics_site", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WebAnalyticsSite"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_web_analytics_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WebAnalyticsRule"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Queue
	p.AddResourceConfigurator("cloudflare_queue", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Queue"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_queue_consumer", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "QueueConsumer"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["queue_id"] = config.Reference{
			TerraformName: "cloudflare_queue",
		}
	})

	// Hyperdrive
	p.AddResourceConfigurator("cloudflare_hyperdrive_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "HyperdriveConfig"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// AI Gateway
	p.AddResourceConfigurator("cloudflare_ai_gateway", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AIGateway"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// AI Search
	p.AddResourceConfigurator("cloudflare_ai_search", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AISearch"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// D1 Database
	p.AddResourceConfigurator("cloudflare_d1_database", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "D1Database"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Infrastructure Access Target
	p.AddResourceConfigurator("cloudflare_infrastructure_access_target", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "InfrastructureAccessTarget"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Observatory
	p.AddResourceConfigurator("cloudflare_observatory_scheduled_test", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ObservatoryScheduledTest"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	// Cloudforce One
	p.AddResourceConfigurator("cloudflare_cloudforce_one_request", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudforceOneRequest"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_cloudforce_one_request_asset", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudforceOneRequestAsset"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_cloudforce_one_request_message", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudforceOneRequestMessage"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_cloudforce_one_request_priority", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudforceOneRequestPriority"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Images
	p.AddResourceConfigurator("cloudflare_images_variant", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ImagesVariant"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Calls
	p.AddResourceConfigurator("cloudflare_calls_sfu_app", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CallsSFUApp"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_calls_turn_app", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CallsTURNApp"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Vectorize
	p.AddResourceConfigurator("cloudflare_vectorize_index", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "VectorizeIndex"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Content Scanning
	p.AddResourceConfigurator("cloudflare_content_scanning", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ContentScanning"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_content_scanning_expression", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ContentScanningExpression"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	// Cloud Connector
	p.AddResourceConfigurator("cloudflare_cloud_connector_rules", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CloudConnectorRules"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	// Zaraz
	p.AddResourceConfigurator("cloudflare_zaraz_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZarazConfig"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zaraz_default_config", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZarazDefaultConfig"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zaraz_history_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZarazHistorySettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_zaraz_workflow", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ZarazWorkflow"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})
}
