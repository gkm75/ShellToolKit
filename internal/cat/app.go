package cat

import (
	"ShellToolKit/pkg/infrastructure"
	"bufio"
	"io"
	"log"
)

// Config cat program config
type Config struct {
	OutputFile string
	InputFiles []string
	Line       struct {
		Start, Finish uint64
		Len           int64
	}
	TabSize int
	NumberNonBlank, Number,
	ShowEnds, ShowTabs, ShowNonPrinting,
	Squeeze bool
}

func processLine(cfg *Config, line []byte, isPrefix bool, writer *bufio.Writer) {

}

func processFile(cfg *Config, filepath string, writer *bufio.Writer) {
	inFile := infrastructure.OpenInputFile(filepath)
	defer infrastructure.CloseFile(inFile)
	defer writer.Flush()

	reader := bufio.NewReader(inFile)

	for {
		line, isPrefix, err := reader.ReadLine()
		if len(line) <= 0 && err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("file read failed: ", err)
			break
		}

		processLine(cfg, line, isPrefix, writer)
	}

}

// Process is the main entry point to dispatch processor functions
// Read file // Buffering // Transformation // Output
func (cfg *Config) Process() {
	outFile := infrastructure.OpenOutputFile(cfg.OutputFile)
	writer := bufio.NewWriter(outFile)
	defer infrastructure.CloseFile(outFile)

	if len(cfg.InputFiles) == 0 {
		processFile(cfg, "", writer)
	} else {
		for _, filepath := range cfg.InputFiles {
			processFile(cfg, filepath, writer)
		}
	}
}
