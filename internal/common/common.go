package common

import (
	"ShellToolKit/pkg/infrastructure"
	"bufio"
)

// Config struct to hold settings
type Config struct {
	Address    bool
	NoBreak    bool
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
func BuildProcessor(cfg *Config, chunkSize int, processLine LineProcessor) infrastructure.BlockProcessor {
	return func(outFile *bufio.Writer, block []byte, count int, pos int64) int {
		m := chunkSize
		for n := 0; n < count; n += chunkSize {
			if m > count {
				m = count
			}

			processLine(cfg, outFile, block[n:m], n, pos+int64(n))
			m += chunkSize
		}
		return count
	}
}
