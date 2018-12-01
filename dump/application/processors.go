package application

import (
	"fmt"
	"os"
)

// LineProcessor function type
type LineProcessor func(*Config, *os.File, []byte, int, int64)

// WritePosition prints file position
func WritePosition(cfg *Config, outFile *os.File, block []byte, count int, pos int64) {
	fmt.Fprintf(outFile, "%08x", pos)
}

// WritePositionWithOffset WritePositionWithOffset
func WritePositionWithOffset(cfg *Config, outFile *os.File, block []byte, count int, pos int64) {
	fmt.Fprintf(outFile, "%08x", pos+int64(cfg.Offset))
}

// WriteText WriteText
func WriteText(cfg *Config, outFile *os.File, block []byte, count int, pos int64) {
	fmt.Fprint(outFile, " ")
	for _, b := range block {
		if b < 32 {
			fmt.Fprint(outFile, ".")
		} else {
			fmt.Fprintf(outFile, "%c", b)
		}
	}
}

// WriteLn WriteLn
func WriteLn(cfg *Config, outFile *os.File, block []byte, count int, pos int64) {
	fmt.Fprintln(outFile)
}

// HexProcessor HexProcessor
func HexProcessor(cfg *Config, outFile *os.File, block []byte, count int, pos int64) {
	for _, b := range block {
		fmt.Fprintf(outFile, " %02x", b)
	}
	for n := len(block); n < 16; n++ {
		fmt.Fprintf(outFile, "   ")
	}
}

// HexProcessorUpper HexProcessorUpper
func HexProcessorUpper(cfg *Config, outFile *os.File, block []byte, count int, pos int64) {
	for _, b := range block {
		fmt.Fprintf(outFile, " %02X", b)
	}
	for n := len(block); n < 16; n++ {
		fmt.Fprintf(outFile, "   ")
	}
}

// OctProcessor OctProcessor
func OctProcessor(cfg *Config, outFile *os.File, block []byte, count int, pos int64) {
	for _, b := range block {
		fmt.Fprintf(outFile, " %03o", b)
	}
	for n := len(block); n < 16; n++ {
		fmt.Fprintf(outFile, "    ")
	}
}
