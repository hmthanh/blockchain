package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"

	bcn "github.com/hmthanh/blockchain/pkg/blockchain"
)

func TestSetHash(t *testing.T) {
	// Create a sample block
	block := &bcn.Block{
		Timestamp:     1702810000,
		Transactions:  []*bcn.Transaction{{Data: []byte("test data")}},
		PrevBlockHash: []byte("previous_hash"),
	}

	// Call SetHash to set the Hash field
	bcn.SetHash(block)

	// Calculate the expected hash manually
	var bytesData []byte
	bytesData = append(bytesData, block.PrevBlockHash...)

	for _, tran := range block.Transactions {
		bytesData = append(bytesData, tran.Data...)
	}
	bytesData = append(bytesData, []byte(strconv.FormatInt(block.Timestamp, 10))...)
	expectedHashBytes := sha256.Sum256(bytesData)
	expectedHash := hex.EncodeToString(expectedHashBytes[:])

	// Check if the calculated hash matches the expected hash
	if block.Hash == nil || hex.EncodeToString(block.Hash) != expectedHash {
		t.Errorf("SetHash failed. Expected hash: %s, actual hash: %s", expectedHash, hex.EncodeToString(block.Hash))
	}
}
