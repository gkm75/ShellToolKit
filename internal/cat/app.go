package cat

import "ShellToolKit/internal/common"

// Process is the main entry point to dispatch processor functions
// Read file // Buffering // Transformation // Output
func Process(cfg *common.Config) {
	println("Processing %s", cfg.InputFile)
}
