package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: commp-cid <file>")
		os.Exit(1)
	}

	commp, err := hex.DecodeString(os.Args[1])
	if err != nil {
		panic(err)
	}

	mh, err := multihash.Encode(commp, multihash.SHA2_256_TRUNC254_PADDED)
	if err != nil {
		panic(err)
	}

	ccid := cid.NewCidV1(cid.FilCommitmentUnsealed, mh)

	println(ccid.String())
}
