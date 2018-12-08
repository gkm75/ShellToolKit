package hex

import (
	"ShellToolKit/internal/common"
	"ShellToolKit/pkg/infrastructure"
	"bufio"
	"fmt"
)

func octProcessor(cfg *common.Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	for _, b := range block {
		fmt.Fprintf(outFile, "%03o", b)
	}
}

func buildLineProcessor(cfg *common.Config) common.LineProcessor {
	var _chain []common.LineProcessor

	_chain = append(_chain, octProcessor)

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

	blockProcessor := common.BuildProcessor(cfg, buildLineProcessor(cfg))

	infrastructure.ProcessFileBlocks(inFile, outFile, 1024, cfg.Seek, cfg.Len, blockProcessor)
}
