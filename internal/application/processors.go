package application

import (
	"bufio"
	"fmt"
)

// LineProcessor function type
type LineProcessor func(*Config, *bufio.Writer, []byte, int, int64)

// WritePosition prints file position
func WritePosition(cfg *Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	fmt.Fprintf(outFile, "%08x", pos)
}

// WritePositionWithOffset WritePositionWithOffset
func WritePositionWithOffset(cfg *Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	fmt.Fprintf(outFile, "%08x", pos+int64(cfg.Offset))
}

// WriteText WriteText
func WriteText(cfg *Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	fmt.Fprint(outFile, " ")
	for _, b := range block {
		if b < 32 {
			fmt.Fprint(outFile, ".")
		} else if b > 126 {
			fmt.Fprint(outFile, "#")
		} else {
			fmt.Fprintf(outFile, "%c", b)
		}
	}
}

// WriteLn WriteLn
func WriteLn(cfg *Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	fmt.Fprintln(outFile)
}

// HexProcessor HexProcessor
func HexProcessor(cfg *Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	for _, b := range block {
		fmt.Fprintf(outFile, " %02x", b)
	}
	for n := len(block); n < 8; n++ {
		fmt.Fprintf(outFile, "   ")
	}
}

// HexProcessorUpper HexProcessorUpper
func HexProcessorUpper(cfg *Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	for _, b := range block {
		fmt.Fprintf(outFile, " %02X", b)
	}
	for n := len(block); n < 8; n++ {
		fmt.Fprintf(outFile, "   ")
	}
}

// OctProcessor OctProcessor
func OctProcessor(cfg *Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	for _, b := range block {
		fmt.Fprintf(outFile, " %03o", b)
	}
	for n := len(block); n < 8; n++ {
		fmt.Fprintf(outFile, "    ")
	}
}

// BinProcessor BinProcessor
func BinProcessor(cfg *Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	for _, b := range block {
		fmt.Fprintf(outFile, " %08b", b)
	}
	for n := len(block); n < 8; n++ {
		fmt.Fprintf(outFile, "         ")
	}
}

// DecProcessor DecProcessor
func DecProcessor(cfg *Config, outFile *bufio.Writer, block []byte, count int, pos int64) {
	for _, b := range block {
		fmt.Fprintf(outFile, " %3v", b)
	}
	for n := len(block); n < 8; n++ {
		fmt.Fprintf(outFile, "    ")
	}
}
