package stream

import "github.com/crossplane/upjet/v2/pkg/config"

const shortGroup = "stream"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_stream", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Stream"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_key", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamKey"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_live_input", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamLiveInput"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_webhook", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamWebhook"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_audio_track", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamAudioTrack"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_caption", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamCaption"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_download", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamDownload"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})

	p.AddResourceConfigurator("cloudflare_stream_watermark", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StreamWatermark"
		r.References["account_id"] = config.Reference{
			TerraformName: "cloudflare_account",
		}
	})
}
