package main

import (
	app "ShellToolKit/internal/cat"
	"ShellToolKit/internal/common"
	"math"

	"github.com/spf13/pflag"
)

func parseConfigArgumentsSetup(cfg *common.Config) {
	pflag.BoolVarP(&cfg.NoBreak, "no-break", "n", false, "prints the dump in one line")
	pflag.IntVarP(&cfg.Cols, "cols", "c", 8, "number of bytes per line")

	pflag.StringVarP(&cfg.InputFile, "inputfile", "i", "", "input file name (default stdin)")
	pflag.StringVarP(&cfg.OutputFile, "outputfile", "o", "", "input file name (default stdout)")

	pflag.IntVarP(&cfg.Offset, "offset", "O", 0, "adds offset to displayed position")
	pflag.Int64VarP(&cfg.Seek, "seek", "s", 0, "starts from the offset")
	pflag.Int64VarP(&cfg.Len, "len", "l", math.MaxInt64, "reads up to len bytes")
}

func main() {
	var cfg common.Config
	version := pflag.BoolP("version", "v", false, "prints the version number")
	help := pflag.BoolP("help", "h", false, "shows the usage help")

	parseConfigArgumentsSetup(&cfg)

	pflag.Parse()

	if *help {
		println("cat [options] [files]")
		pflag.Usage()
	} else if *version {
		println("cat version %s", common.Version)
	} else {
		if pflag.NArg() == 0 {
			app.Process(&cfg)
		} else {
			for _, filepath := range pflag.Args() {
				cfg.InputFile = filepath
				app.Process(&cfg)
			}
		}
	}
}