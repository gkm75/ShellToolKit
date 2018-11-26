package main

import (
	app "ShellToolKit/dump/application"

	"github.com/spf13/pflag"
)

func main() {
	var cfg app.Config
	println("Dump")
	flgNoOffset := pflag.BoolP("no-offset", "n", false, "'true' to hide file offsets")
	pflag.BoolVarP(&cfg.Hexy, "hex", "h", true, "prints hex values")
	pflag.StringVarP(&cfg.InputFile, "inputfile", "i", "", "input file name")

	pflag.Parse()

	if !(*flgNoOffset) {
		println("offset on")
	} else {
		println("offset off")
	}

	if cfg.Hexy {
		println("Hex mode")
	}

	println("Using ", cfg.InputFile)
}
