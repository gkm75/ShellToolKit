package main

import (
	"ShellToolKit/internal/cat"
	"ShellToolKit/internal/common"
	"math"

	"github.com/spf13/pflag"
)

func parseConfigArgumentsSetup(cfg *cat.Config) {
	pflag.BoolVarP(&cfg.NumberNonBlank, "number-nonblank", "b", false, "number nonempty output lines, overrides -n")
	pflag.BoolVarP(&cfg.ShowEnds, "show-ends", "E", false, "display $ at end of each line")
	pflag.BoolVarP(&cfg.Number, "number", "n", false, "number all output lines")
	pflag.BoolVarP(&cfg.Squeeze, "squeeze", "s", false, "suppress empty output lines")
	pflag.BoolVarP(&cfg.ShowTabs, "show-tabs", "T", false, "display TAB characters as \t")
	pflag.BoolVarP(&cfg.ShowNonPrinting, "show-nonprinting", "V", false, "use \\x notation, except for LFD and TAB")
	pflag.Uint64VarP(&cfg.Line.Start, "start-at", "S", 0, "start at line ")
	pflag.Int64VarP(&cfg.Line.Len, "lines", "L", math.MaxInt64, "process l lines")
	pflag.Uint64VarP(&cfg.Line.Finish, "finish-at", "F", math.MaxUint64, "finish at line")
}

func main() {
	var cfg cat.Config
	version := pflag.BoolP("version", "v", false, "prints the version number")
	help := pflag.BoolP("help", "h", false, "shows the usage help")

	parseConfigArgumentsSetup(&cfg)

	pflag.Parse()

	if *help {
		println("cat [options] [files]")
		pflag.Usage()
	} else if *version {
		println("cat version ", common.Version)
	} else {
		cfg.InputFiles = pflag.Args()
		cfg.Process()
	}
}
