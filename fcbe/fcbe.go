package fcbe

import (
	"os"
	"os/exec"
	"strconv"
)

// PlayFile - play file to fake webcam
func PlayFile(filename string, loop int, device int) (*os.Process, error) {
	var cmd *exec.Cmd
	var err error

	deviceName := "/dev/video" + strconv.Itoa(device)

	const (
		cmdffmpeg     = "ffmpeg"
		argre         = "-re"
		argi          = "-i"
		argmap        = "-map"
		argzcv        = "0:v"
		argf          = "-f"
		argv4l2       = "v4l2"
		argstreamloop = "-stream_loop"
		argmo         = "-1"
	)

	cmd = exec.Command(cmdffmpeg, argstreamloop, strconv.Itoa(loop), argre, argi, filename, argmap, argzcv, argf, argv4l2, deviceName)
	err = cmd.Start()

	if err != nil {
		return nil, err
	}
	return cmd.Process, nil
}
