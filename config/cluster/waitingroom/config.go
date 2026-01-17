package waitingroom

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "waitingroom"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_waiting_room", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WaitingRoom"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})

	p.AddResourceConfigurator("cloudflare_waiting_room_event", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WaitingRoomEvent"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["waiting_room_id"] = config.Reference{
			TerraformName: "cloudflare_waiting_room",
		}
	})

	p.AddResourceConfigurator("cloudflare_waiting_room_rules", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WaitingRoomRules"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
		r.References["waiting_room_id"] = config.Reference{
			TerraformName: "cloudflare_waiting_room",
		}
	})

	p.AddResourceConfigurator("cloudflare_waiting_room_settings", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "WaitingRoomSettings"
		r.References["zone_id"] = config.Reference{
			TerraformName: "cloudflare_zone",
		}
	})
}
