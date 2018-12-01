package application

import (
	"fmt"
	"os"
)

// LineProcessor function type
type LineProcessor func(*os.File, []byte, int, int64)

func hexProcessor(outFile *os.File, block []byte, count int) int {
	fmt.Println("block.", count)
	return 0
}
