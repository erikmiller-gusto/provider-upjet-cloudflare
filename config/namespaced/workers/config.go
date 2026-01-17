package workers

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "workers"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_workers_script", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersScript"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_worker_route", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkerRoute"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["script_name"] = config.Reference{
			TerraformName: "cloudflare_workers_script",
		}
	})

	p.AddResourceConfigurator("cloudflare_worker_cron_trigger", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkerCronTrigger"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["script_name"] = config.Reference{
			TerraformName: "cloudflare_workers_script",
		}
	})

	p.AddResourceConfigurator("cloudflare_worker_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkerDomain"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_worker_secret", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkerSecret"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["script_name"] = config.Reference{
			TerraformName: "cloudflare_workers_script",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_for_platforms_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersForPlatformsNamespace"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_kv", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersKV"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["namespace_id"] = config.Reference{
			TerraformName: "cloudflare_workers_kv_namespace",
		}
	})

	p.AddResourceConfigurator("cloudflare_workers_kv_namespace", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WorkersKVNamespace"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	// Pages
	p.AddResourceConfigurator("cloudflare_pages_project", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PagesProject"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_pages_domain", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PagesDomain"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["project_name"] = config.Reference{
			TerraformName: "cloudflare_pages_project",
		}
	})
}
