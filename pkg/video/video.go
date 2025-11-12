package video

import (
	"context"
	"fmt"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type VideoConverter struct {
	ctx        context.Context
	Path       string
	TargetType string
	DestDir    string
}

func NewVideoConverter() *VideoConverter {
	return &VideoConverter{}
}

func (c *VideoConverter) OpenFile() (string, error) {
	file, err := runtime.OpenFileDialog(c.ctx, runtime.OpenDialogOptions{
		Title: "Select a Video File to be Converted",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Video Files (MP4, MOV, MKV, AVI, WMV, WebM)",
				Pattern:     "*.mp4;*.mov;*.mkv;*.avi;*.wmv;*.webm;*.flv;*.3gp",
			},
		},
	})

	if err != nil {
		return "", fmt.Errorf("Error : %v", err.Error())
	}
	c.Path = file
	return file, nil
}

func (c *VideoConverter) ConvertVideo(target string) error {
	c.TargetType = target

	args := ffmpeg_go.KwArgs{}

	switch target {
	case "mp3":
		args = ffmpeg_go.KwArgs{
			"b:a": "128k",
		}
	case "opus":
		args = ffmpeg_go.KwArgs{
			"b:a": "128k",
		}
	}

	err := ffmpeg_go.Input(c.Path).
		Output(fmt.Sprintf("%v/converted.%v", c.DestDir, c.TargetType), args).
		OverWriteOutput().
		Run()
	if err != nil {
		return fmt.Errorf("Error : %v", err.Error())
	}

	return nil
}

func (c *VideoConverter) SelectDestDir() (string, error) {
	path, err := runtime.OpenDirectoryDialog(c.ctx, runtime.OpenDialogOptions{
		Title: "Select a Destination Folder",
	})
	if err != nil {
		return "", fmt.Errorf("Error : %v", err.Error())
	}
	c.DestDir = path

	return path, nil
}

func (c *VideoConverter) SetContext(ctx context.Context) {
	c.ctx = ctx
}
