package loadbalancer

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "loadbalancer"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_load_balancer", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LoadBalancer"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["default_pool_ids"] = config.Reference{
			TerraformName: "cloudflare_load_balancer_pool",
		}
		r.References["fallback_pool_id"] = config.Reference{
			TerraformName: "cloudflare_load_balancer_pool",
		}
	})

	p.AddResourceConfigurator("cloudflare_load_balancer_pool", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LoadBalancerPool"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
		r.References["monitor"] = config.Reference{
			TerraformName: "cloudflare_load_balancer_monitor",
		}
	})

	p.AddResourceConfigurator("cloudflare_load_balancer_monitor", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "LoadBalancerMonitor"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
