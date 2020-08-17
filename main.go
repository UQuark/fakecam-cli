package main

import (
	"os"
	"strconv"

	"github.com/UQuark0/fcbe"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	v4l2, err := fcbe.IsV4L2Loaded()
	check(err)
	if !v4l2 {
		panic("'v4l2loopback' kernel module is not loaded")
	}

	var filename string = ""
	var loop int = 1
	var device int = -1

	for i := 1; i < len(os.Args); i++ {
		if i == len(os.Args)-1 {
			panic("Expected value after '" + os.Args[i] + "'")
		}

		switch os.Args[i] {
		case "-l", "--loop":
			i++
			loop, err = strconv.Atoi(os.Args[i])
			check(err)
		case "-d", "--device":
			i++
			device, err = strconv.Atoi(os.Args[i])
			check(err)
		case "-f", "--file":
			i++
			filename = os.Args[i]
		default:
			panic("Unknown argument: '" + os.Args[i] + "'")
		}
	}

	if filename == "" {
		panic("No file specified")
	}

	if device == -1 {
		panic("No device specified")
	}

	proc, err := fcbe.PlayFile(filename, loop, device)
	check(err)
	proc.Wait()
}
