package dump

import (
	"ShellToolKit/internal/common"
	"ShellToolKit/pkg/infrastructure"
	"bufio"
)

func buildLineProcessor(cfg *common.Config) common.LineProcessor {
	var _chain []common.LineProcessor

	if cfg.Address {
		if cfg.Offset == 0 {
			_chain = append(_chain, common.WritePosition)
		} else {
			_chain = append(_chain, common.WritePositionWithOffset)
		}
	}

	if cfg.Hex {
		if cfg.Upper {
			_chain = append(_chain, common.HexProcessorUpper)
		} else {
			_chain = append(_chain, common.HexProcessor)
		}
	}

	if cfg.Oct {
		_chain = append(_chain, common.OctProcessor)
	}

	if cfg.Bin {
		_chain = append(_chain, common.BinProcessor)
	}

	if cfg.Dec {
		_chain = append(_chain, common.DecProcessor)
	}

	if cfg.Text {
		_chain = append(_chain, common.WriteText)
	}

	if !cfg.NoBreak {
		_chain = append(_chain, common.WriteLn)
	}

	return func(cfg *common.Config, outFile *bufio.Writer, chunk []byte, offset int, pos int64) {
		for _, p := range _chain {
			p(cfg, outFile, chunk, offset, pos)
		}
	}
}

// Process is the main entry point to dispatch processor functions
// Read file // Buffering // Transformation // Output
func Process(cfg *common.Config) {
	inFile := infrastructure.OpenInputFile(cfg.InputFile)
	outFile := infrastructure.OpenOutputFile(cfg.OutputFile)

	defer infrastructure.CloseFile(inFile)
	defer infrastructure.CloseFile(outFile)

	blockProcessor := common.BuildProcessor(cfg, 8, buildLineProcessor(cfg))

	infrastructure.ProcessFileBlocks(inFile, outFile, 1024, cfg.Seek, cfg.Len, blockProcessor)

}
