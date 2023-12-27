package main

import (
	"fmt"

	btc "github.com/hmthanh/blockchain/pkg/blockchain"
	bcn "github.com/hmthanh/blockchain/pkg/simpleblockchain"
)

func TestSimpleBlockChain() {
	// Create a new blockchain
	blockchain := bcn.NewBlockChain()

	// Add transactions and blocks
	transaction1 := &bcn.Transaction{Data: []byte("Alex -> Bob, 100 BTC")}
	transaction2 := &bcn.Transaction{Data: []byte("Bob -> Alex, 50 BTC")}

	blockchain.AddBlock([]*bcn.Transaction{transaction1})
	blockchain.AddBlock([]*bcn.Transaction{transaction2})

	// Print the blockchain
	blockchain.PrintBlockchain()

	// Verify the blockchain
	isValid := blockchain.VerifyBlockchain()
	if isValid {
		println("Blockchain is valid")
	} else {
		println("Blockchain is invalid")
	}
}

func TestMerleTreeBlockChain() {
	coin := btc.NewBlockChain()

	// Add transactions and blocks
	transaction1 := &btc.Transaction{Data: []byte("Alex -> Bob, 100 BTC")}
	transaction2 := &btc.Transaction{Data: []byte("Bob -> Alex, 50 BTC")}
	trans1 := []*btc.Transaction{}
	trans1 = append(trans1, transaction1)
	trans1 = append(trans1, transaction2)

	coin.AddBlock(trans1)

	transaction3 := &btc.Transaction{Data: []byte("Mari -> Tom, 15 BTC")}
	trans2 := []*btc.Transaction{}
	trans2 = append(trans2, transaction3)
	coin.AddBlock(trans2)

	isValidMerleTree := coin.VerifyBlockchain()
	fmt.Println("isValidMerleTree : ", isValidMerleTree)
}

func main() {
	TestMerleTreeBlockChain()
}
