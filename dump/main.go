package main

import (
	app "ShellToolKit/dump/application"
	"math"

	"github.com/spf13/pflag"
)

func parseConfigArgumentsSetup(cfg *app.Config) {
	pflag.BoolVarP(&cfg.Address, "show-position", "p", true, "'false' to hide file offsets")
	pflag.BoolVarP(&cfg.NoBreak, "no-break", "n", false, "prints the dump in one line")

	pflag.BoolVarP(&cfg.Hex, "hex", "x", true, "prints hex values")
	pflag.BoolVarP(&cfg.Oct, "oct", "8", false, "prints oct values")
	pflag.BoolVarP(&cfg.Bin, "bin", "b", false, "prints bin values")
	pflag.BoolVarP(&cfg.Upper, "upper", "u", false, "prints uppercase hex values")

	pflag.StringVarP(&cfg.InputFile, "inputfile", "i", "", "input file name (default stdin)")
	pflag.StringVarP(&cfg.OutputFile, "outputfile", "o", "", "input file name (default stdout)")

	pflag.IntVarP(&cfg.Offset, "offset", "O", 0, "adds offset to displayed position")
	pflag.IntVarP(&cfg.Seek, "seek", "s", 0, "starts from the offset")
	pflag.IntVarP(&cfg.Len, "len", "l", math.MaxInt32, "reads up to len bytes")

	pflag.BoolVarP(&cfg.Text, "text", "t", false, "prints text dump")
}

func main() {
	var cfg app.Config
	version := pflag.BoolP("version", "v", false, "prints the version number")
	help := pflag.BoolP("help", "h", false, "shows the usage help")

	parseConfigArgumentsSetup(&cfg)

	pflag.Parse()

	if *help {
		println("Dump [options]")
		pflag.Usage()
	} else if *version {
		println("Dump version 1.0.0")
	} else {
		app.Process(&cfg)
	}
}
