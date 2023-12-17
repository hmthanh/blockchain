package main

import (
	"testing"

	"github.com/hmthanh/blockchain/pkg/blockchain"
)

func TestHashTrasaction(t *testing.T) {
	testBlock := &blockchain.Block{
		Timestamp:     123456789,
		Transactions:  []*blockchain.Transaction{{Data: []byte("Transaction1")}, {Data: []byte("Transaction2")}},
		PrevBlockHash: []byte("PreviousHash"),
		Hash:          []byte("CurrentHash")}

	result := blockchain.HashTransactions(testBlock.Transactions)

	expectedHash := "a748685c5eecc7b9a0e6c5e9d9c5ee179a7ecdd722ec97f8f3f48c5c4d0c1c8a"

	// Compare the result with the expected hash
	if result != expectedHash {
		t.Errorf("HashTransactions: Expected %s, got %s", expectedHash, result)
	}
}

func TestSetHash(t *testing.T) {
	// Create a block with transactions for testing
	testBlock := &blockchain.Block{
		Timestamp:     123456789,
		Transactions:  []*blockchain.Transaction{{Data: []byte("Transaction1")}, {Data: []byte("Transaction2")}},
		PrevBlockHash: []byte("PreviousHash"),
		Hash:          []byte("CurrentHash"),
	}

	// Call the SetHash function
	blockchain.SetHash(testBlock)

	// Provide the expected hash value based on your test data
	expectedHash := "a748685c5eecc7b9a0e6c5e9d9c5ee179a7ecdd722ec97f8f3f48c5c4d0c1c8a"

	// Compare the result with the expected hash
	if hex := string(testBlock.Hash); hex != expectedHash {
		t.Errorf("SetHash: Expected %s, got %s", expectedHash, hex)
	}
}
