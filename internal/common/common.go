package common

import (
	"ShellToolKit/pkg/infrastructure"
	"bufio"
)

// Config struct to hold settings
type Config struct {
	Address    bool
	NoBreak    bool
	Cols       int
	Hex        bool
	Bin        bool
	Oct        bool
	Dec        bool
	Upper      bool
	InputFile  string
	OutputFile string
	Offset     int
	Seek       int64
	Len        int64
	Text       bool
}

// LineProcessor function type
type LineProcessor func(*Config, *bufio.Writer, []byte, int, int64)

// BuildProcessor builds a processor which calls the LineProcessor passed as param
func BuildProcessor(cfg *Config, processLine LineProcessor) infrastructure.BlockProcessor {
	if cfg.Cols < 0 {
		cfg.Cols = 8
	}

	return func(outFile *bufio.Writer, block []byte, count int, pos int64) int {
		m := cfg.Cols
		for n := 0; n < count; n += cfg.Cols {
			if m > count {
				m = count
			}

			processLine(cfg, outFile, block[n:m], n, pos+int64(n))
			m += cfg.Cols
		}
		return count
	}
}
