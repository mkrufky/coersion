package main

import (
	"fmt"
	"os/exec"
	"time"
)

// ScaleImageTask uses ffmpeg to perform the scaling of a still image from a video source
type ScaleImageTask struct {
	source string
	cmd    *exec.Cmd
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// NewScaleImageTask creates a new image Scaleion task
func NewScaleImageTask(source string, w, h uint, o time.Duration) *ScaleImageTask {

	offset := fmtDuration(o)

	return &ScaleImageTask{
		source: source,
		cmd: exec.Command(ffmpegCmd, []string{
			"-ss", offset,
			"-i", source,
			"-vframes", "1",
			"-vf", fmt.Sprintf("scale=%d:%d", w, h),
			"-f", "image2pipe",
			"-vcodec", "png",
			"-q:v", "2",
			"-",
		}...),
	}
}

// Start begins the task
func (j *ScaleImageTask) getCmd() *exec.Cmd {
	return j.cmd
}
