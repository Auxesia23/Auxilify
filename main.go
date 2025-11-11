package main

import (
	"Auxilify/pkg/image"
	"Auxilify/pkg/video"
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	image := image.NewImageConverter()
	video := video.NewVideoConverter()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Auxilify",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: func(ctx context.Context) {
			image.SetContext(ctx)
			video.SetContext(ctx)
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			image,
			video,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
