package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	humanize "github.com/dustin/go-humanize"
	pieceio "github.com/filecoin-project/go-fil-markets/pieceio"
	sectorbuilder "github.com/filecoin-project/go-sectorbuilder"
)

func generateCommPDirect(path string) ([]byte, uint64, uint64, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, err
	}

	fstat, err := fi.Stat()
	if err != nil {
		return nil, 0, 0, err
	}

	size := uint64(fstat.Size())

	reader := bufio.NewReader(fi)

	pr, psize := NewPadreader(reader, size)

	commp, err := sectorbuilder.GeneratePieceCommitment(pr, psize)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("Error generating CommP: %w", err)
	}

	return commp[:], size, psize, nil
}

func generateCommP(path string) ([]byte, uint64, uint64, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, err
	}

	fstat, err := fi.Stat()
	if err != nil {
		return nil, 0, 0, err
	}

	size := uint64(fstat.Size())

	reader := bufio.NewReader(fi)

	commp, psize, err := pieceio.GeneratePieceCommitment(reader, size)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("Error generating CommP: %w", err)
	}

	return commp[:], size, psize, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: commp <file>")
		os.Exit(1)
	}

	commp, size, psize, err := generateCommP(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("(markets) File '%v'\n\tSize: %v\n\tPadded: %v\n\tcommP: %v\n",
		os.Args[1],
		humanize.BigIBytes(big.NewInt(int64(size))),
		humanize.BigIBytes(big.NewInt(int64(psize))),
		hex.EncodeToString(commp))

	commp, size, psize, err = generateCommPDirect(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("(direct) File '%v'\n\tSize: %v\n\tPadded: %v\n\tcommP: %v\n",
		os.Args[1],
		humanize.BigIBytes(big.NewInt(int64(size))),
		humanize.BigIBytes(big.NewInt(int64(psize))),
		hex.EncodeToString(commp))
}
