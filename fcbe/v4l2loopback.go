package fcbe

import (
	"bytes"
	"os/exec"
)

// IsV4L2Loaded checks if v4l2loopback kernel module is loaded
func IsV4L2Loaded() (bool, error) {
	out, err := exec.Command("lsmod").Output()
	if err != nil {
		return false, err
	}

	return bytes.Contains(out, []byte("v4l2loopback")), nil
}
