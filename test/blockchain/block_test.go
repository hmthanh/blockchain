package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"testing"

	bcn "github.com/hmthanh/blockchain/pkg/simpleblockchain"
)

func TestBlockchain(t *testing.T) {
	// Create Genesis Block
	genesisBlock := bcn.CreateGenesisBlock()

	// Ensure the Genesis Block has the expected values
	if genesisBlock.Timestamp <= 0 {
		t.Error("Genesis Block timestamp should be greater than 0")
	}

	if len(genesisBlock.Transactions) != 0 {
		t.Error("Genesis Block should have no transactions")
	}

	if len(genesisBlock.PrevBlockHash) != 0 {
		t.Error("Genesis Block's Previous Block Hash should be empty")
	}

	if len(genesisBlock.Hash) != 0 {
		t.Error("Genesis Block's Hash should be empty")
	}

	// Create Transactions
	transactions := []*bcn.Transaction{
		{Data: []byte("Transaction 1")},
		{Data: []byte("Transaction 2")},
	}

	// Create a new Block
	prevBlockHash := genesisBlock.Hash
	newBlock := bcn.CreateBlock(transactions, prevBlockHash)

	// Ensure the new Block has the expected values
	if newBlock.Timestamp <= 0 {
		t.Error("New Block timestamp should be greater than 0")
	}

	if !reflect.DeepEqual(newBlock.Transactions, transactions) {
		t.Error("New Block should have the same transactions as provided")
	}

	if !reflect.DeepEqual(newBlock.PrevBlockHash, prevBlockHash) {
		t.Error("New Block's Previous Block Hash should match the provided value")
	}

	if len(newBlock.Hash) == 0 {
		t.Error("New Block's Hash should not be empty")
	}

	// Ensure the Hash of the Block is correctly set
	hashTrans := bcn.HashTransactions(transactions)
	var bytesData []byte
	bytesData = append(bytesData, prevBlockHash...)
	bytesData = append(bytesData, hashTrans...)
	bytesData = append(bytesData, []byte(fmt.Sprint(newBlock.Timestamp))...)

	// append(prevBlockHash, append(expectedHash, []byte(fmt.Sprint(newBlock.Timestamp)))...)

	expectedHash := sha256.Sum256(bytesData)
	if !reflect.DeepEqual(newBlock.Hash, expectedHash[:]) {
		t.Error("New Block's Hash is not set correctly")
	}
}
