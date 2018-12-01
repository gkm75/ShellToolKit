package infrastructure

import (
	"io"
	"log"
	"os"
)

// BlockProcessor is a function type that processes the block part of the file passed
type BlockProcessor func([]byte, int) int

// ProcessFile opens the passed filename, reads chunks and calles the processor function passed as param
func ProcessFile(filepath string, processor BlockProcessor) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("file open failed: ", err)
	}

	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if n <= 0 && err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("file read failed: ", err)
			break
		}
		processor(buffer, n)
	}
}
