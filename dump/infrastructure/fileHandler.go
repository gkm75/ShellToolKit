package infrastructure

import (
	"io"
	"log"
	"os"
)

// BlockProcessor is a function type that processes the block part of the file passed
type BlockProcessor func(*os.File, []byte, int, int64) int

// OpenInputFile opens a file or stdin if "" is passed
func OpenInputFile(path string) *os.File {
	if path == "" {
		return os.Stdin
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("file open failed: ", err)
	}

	return file
}

// OpenOutputFile creates a file for writing or uses stdout if "" is passed
func OpenOutputFile(path string) *os.File {
	if path == "" {
		return os.Stdout
	}

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

// CloseFile closes the passed file
func CloseFile(file *os.File) {
	if file == nil {
		log.Fatal("file for closing is nil")
	}

	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// ProcessAllFile opens the passed filename, reads chunks and calls the processor function passed as param
func ProcessAllFile(inFile, outFile *os.File, startAt int64, processor BlockProcessor) {
	var pos int64
	var err error
	if startAt > 0 {
		pos, err = inFile.Seek(startAt, 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	buffer := make([]byte, 1024)

	for {
		n, err := inFile.Read(buffer)
		if n <= 0 && err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("file read failed: ", err)
			break
		}
		processor(outFile, buffer, n, pos)
		pos += int64(n)
	}
}
