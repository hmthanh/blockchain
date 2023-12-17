package main

import (
	bcn "github.com/hmthanh/blockchain/pkg/blockchain"
)

func main() {
	// Create a new blockchain
	blockchain := bcn.NewBlockChain()

	// Add transactions and blocks
	transaction1 := &bcn.Transaction{Data: []byte("Alex -> Bob, 100 BTC")}
	transaction2 := &bcn.Transaction{Data: []byte("Bob -> Alex, 50 BTC")}

	blockchain.AddBlock([]*bcn.Transaction{transaction1})
	blockchain.AddBlock([]*bcn.Transaction{transaction2})

	// Print the blockchain
	blockchain.PrintBlockchain()
}
