package image

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	_ "golang.org/x/image/webp"

	"github.com/HugoSmits86/nativewebp"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ImageConverter struct {
	ctx        context.Context
	Path       string
	Data       bytes.Buffer
	TargetType string
}

func NewImageConverter() *ImageConverter {
	return &ImageConverter{}
}

func (c *ImageConverter) ConvertImage(target string) error {
	c.Data.Reset()
	c.TargetType = target
	file, err := os.Open(c.Path)
	if err != nil {
		return fmt.Errorf("Error : %v", err.Error())
	}

	image, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("Error : %v", err.Error())
	}

	switch target {
	case "jpeg":
		err = jpeg.Encode(&c.Data, image, nil)
		if err != nil {
			return fmt.Errorf("Error : %v", err.Error())
		}
	case "png":
		err = png.Encode(&c.Data, image)
		if err != nil {
			return fmt.Errorf("Error : %v", err.Error())
		}
	case "gif":
		err = gif.Encode(&c.Data, image, nil)
		if err != nil {
			return fmt.Errorf("Error : %v", err.Error())
		}
	case "webp":
		err = nativewebp.Encode(&c.Data, image, nil)
		if err != nil {
			return fmt.Errorf("Error : %v", err.Error())
		}
	}

	return nil
}

func (c *ImageConverter) OpenFile() (string, error) {
	file, err := runtime.OpenFileDialog(c.ctx, runtime.OpenDialogOptions{
		Title: "Select a File to be Converted",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Image Files (JPG, PNG, GIF, WEBP)",
				Pattern:     "*.jpg;*.jpeg;*.png;*.gif;*.webp",
			},
		},
	})

	if err != nil {
		return "", fmt.Errorf("Error : %v", err.Error())
	}
	c.Path = file
	return file, nil
}

func (c *ImageConverter) SaveFile() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error : %v", err.Error())
	}
	saveDir := filepath.Join(homeDir, "Downloads")

	savePath, err := runtime.SaveFileDialog(c.ctx, runtime.SaveDialogOptions{
		Title:            "Save Your Converted Image",
		DefaultDirectory: saveDir,
		DefaultFilename:  fmt.Sprintf("converted.%v", c.TargetType),
		Filters: []runtime.FileFilter{
			{
				DisplayName: fmt.Sprintf("%v Files (*.%v)", strings.ToUpper(c.TargetType), c.TargetType),
				Pattern:     fmt.Sprintf("*.%v", c.TargetType),
			},
		},
	})
	if savePath == "" {
		return fmt.Errorf("Save Canceled by User")
	}
	if err != nil {
		return fmt.Errorf("Error : %v", err.Error())
	}

	err = os.WriteFile(savePath, c.Data.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("Error : %v", err.Error())
	}
	return nil
}

func (c *ImageConverter) SetContext(ctx context.Context) {
	c.ctx = ctx
}
