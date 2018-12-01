package application

import (
	"ShellToolKit/dump/infrastructure"
	"os"
)

// Config struct to hold settings
type Config struct {
	Address    bool
	NoBreak    bool
	Hex        bool
	Bin        bool
	Oct        bool
	Upper      bool
	InputFile  string
	OutputFile string
	Offset     int
	Seek       int64
	Len        int64
	Text       bool
}

func buildLineProcessor(cfg *Config) LineProcessor {
	var _chain []LineProcessor

	if cfg.Address {
		if cfg.Offset == 0 {
			_chain = append(_chain, WritePosition)
		} else {
			_chain = append(_chain, WritePositionWithOffset)
		}
	}

	if cfg.Text {
		_chain = append(_chain, WriteText)
	}

	if !cfg.NoBreak {
		_chain = append(_chain, WriteLn)
	}

	return func(cfg *Config, outFile *os.File, chunk []byte, offset int, pos int64) {
		for _, p := range _chain {
			p(cfg, outFile, chunk, offset, pos)
		}
	}
}

func buildProcessor(cfg *Config, processLine LineProcessor) infrastructure.BlockProcessor {
	return func(outFile *os.File, block []byte, count int, pos int64) int {
		m := 16
		for n := 0; n < count; n += 16 {
			if m > count {
				m = count
			}

			processLine(cfg, outFile, block[n:m], n, pos+int64(n))
			m += 16
		}
		return count
	}
}

// Process is the main entry point to dispatch processor functions
// Read file // Buffering // Transformation // Output
func Process(cfg *Config) {
	inFile := infrastructure.OpenInputFile(cfg.InputFile)
	outFile := infrastructure.OpenOutputFile(cfg.OutputFile)

	defer infrastructure.CloseFile(outFile)
	defer infrastructure.CloseFile(inFile)

	blockProcessor := buildProcessor(cfg, buildLineProcessor(cfg))

	infrastructure.ProcessAllFile(inFile, outFile, cfg.Seek, blockProcessor)

}
