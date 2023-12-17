package main

import (
	"testing"

	btn "github.com/hmthanh/blockchain/pkg/blockchain"
)

func TestNewBlockchain(t *testing.T) {
	blockchain := btn.NewBlockChain()

	// Check if the blockchain is not nil
	if blockchain == nil {
		t.Error("Expected a valid Blockchain instance, got nil")
	}

	// Check if the blockchain has one block (genesis block)
	if len(blockchain.GetBlocks()) != 1 {
		t.Errorf("Expected 1 block in the blockchain, got %d", len(blockchain.GetBlocks()))
	}

	// Check if the first block is the genesis block
	genesisBlock := blockchain.GetGenesisBlock()
	if genesisBlock == nil {
		t.Error("Expected a valid Genesis Block, got nil")
	} else {
		if len(genesisBlock.PrevBlockHash) != 0 {
			t.Error("Expected the Genesis Block to have no previous block hash")
		}
	}
}

func TestAddBlock(t *testing.T) {
	blockchain := btn.NewBlockChain()

	// Create a transaction
	transaction := &btn.Transaction{
		Data: []byte("Test Transaction"),
	}

	// Add a new block with the transaction
	blockchain.AddBlock([]*btn.Transaction{transaction})

	// Check if the blockchain now has two blocks
	if len(blockchain.GetBlocks()) != 2 {
		t.Errorf("Expected 2 blocks in the blockchain, got %d", len(blockchain.GetBlocks()))
	}

	// Check if the new block has the correct previous block hash
	newBlock := blockchain.GetBlock(1)
	if string(newBlock.PrevBlockHash) != string(blockchain.GetBlockHash(0)) {
		t.Error("Previous block hash mismatch in the new block")
	}

	// Check if the new block has a valid hash
	if len(newBlock.Hash) == 0 {
		t.Error("Expected a valid hash for the new block")
	}
}
