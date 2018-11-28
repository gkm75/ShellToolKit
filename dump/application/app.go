package application

import (
	"ShellToolKit/dump/infrastructure"

	"fmt"
)

// Config struct to hold settings
type Config struct {
	Offset     bool
	Stream     bool
	Hex        bool
	Bin        bool
	Oct        bool
	InputFile  string
	OutputFile string
}

func hexProcessor(block []byte, count int) int {
	fmt.Println("block.", count)
	return 0
}

// Process is the main entry point to dispatch processor functions
func Process(cfg *Config) {
	// Read file
	// Buffering
	// Transformation
	// Output

	infrastructure.ProcessFile(cfg.InputFile, hexProcessor)
}
