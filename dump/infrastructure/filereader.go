package infrastructure

// BlockProcessor is a function type that processes the block part of the file passed
type BlockProcessor func([]byte) int

// ProcessFile opens the passed filename, reads chunks and calles the processor function passed as param
func ProcessFile(filename string, processor BlockProcessor) {
	// open file name
	// loop
	// need some form of bufferring
}
