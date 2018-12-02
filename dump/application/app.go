package application

import (
	"ShellToolKit/dump/infrastructure"
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

func buildLineProcessor(cfg *Config) LineProcessor {
	var _chain []LineProcessor

	if cfg.Address {
		if cfg.Offset == 0 {
			_chain = append(_chain, WritePosition)
		} else {
			_chain = append(_chain, WritePositionWithOffset)
		}
	}

	if cfg.Hex {
		if cfg.Upper {
			_chain = append(_chain, HexProcessorUpper)
		} else {
			_chain = append(_chain, HexProcessor)
		}
	}

	if cfg.Oct {
		_chain = append(_chain, OctProcessor)
	}

	if cfg.Bin {
		_chain = append(_chain, BinProcessor)
	}

	if cfg.Dec {
		_chain = append(_chain, DecProcessor)
	}

	if cfg.Text {
		_chain = append(_chain, WriteText)
	}

	if !cfg.NoBreak {
		_chain = append(_chain, WriteLn)
	}

	return func(cfg *Config, outFile *bufio.Writer, chunk []byte, offset int, pos int64) {
		for _, p := range _chain {
			p(cfg, outFile, chunk, offset, pos)
		}
	}
}

func buildProcessor(cfg *Config, chunkSize int, processLine LineProcessor) infrastructure.BlockProcessor {
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

// Process is the main entry point to dispatch processor functions
// Read file // Buffering // Transformation // Output
func Process(cfg *Config) {
	inFile := infrastructure.OpenInputFile(cfg.InputFile)
	outFile := infrastructure.OpenOutputFile(cfg.OutputFile)
	writer := bufio.NewWriter(outFile)

	defer infrastructure.CloseFile(inFile)
	defer infrastructure.CloseFile(outFile)
	defer writer.Flush()

	blockProcessor := buildProcessor(cfg, 8, buildLineProcessor(cfg))

	infrastructure.ProcessAllFile(inFile, writer, 1024, cfg.Seek, blockProcessor)

}
