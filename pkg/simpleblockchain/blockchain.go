package simpleblockchain

import (
	"bytes"
	"fmt"
	"time"
)

type Blockchain struct {
	blocks []*Block
}

func NewBlockChain() *Blockchain {
	genesisBlock := CreateGenesisBlock()
	return &Blockchain{[]*Block{genesisBlock}}
}

func (bc *Blockchain) AddBlock(transaction []*Transaction) {
	// previous block
	prevBlock := bc.blocks[len(bc.blocks)-1]

	// new block
	newBlock := CreateBlock(transaction, prevBlock.Hash)
	newBlock.SetHash()

	// append new block to blockchain
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) VerifyBlockchain() bool {
	for i := 1; i < bc.GetBlockCount(); i++ {
		currBlock := bc.GetBlock(i)
		prevBlock := bc.GetBlock(i - 1)

		isEqual := bytes.Compare(currBlock.PrevBlockHash, prevBlock.Hash) == 0
		if isEqual == false {
			return false
		}
	}
	return true
}

// Prints the details of the blockchain.
func (bc *Blockchain) PrintBlockchain() {
	for _, block := range bc.blocks {
		fmt.Printf("Timestamp: %d (%s)\n", block.Timestamp, time.Unix(block.Timestamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println("Transactions:")
		for _, tran := range block.Transactions {
			fmt.Printf("  Data: %s\n", tran.Data)
		}
		fmt.Println("----------------------")
	}
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func (bc *Blockchain) GetBlocks() []*Block {
	return bc.blocks
}

func (bc *Blockchain) GetBlock(index int) *Block {
	return bc.blocks[index]
}

func (bc *Blockchain) GetBlockCount() int {
	return len(bc.blocks)
}

func (bc *Blockchain) GetLatestBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) GetGenesisBlock() *Block {
	return bc.blocks[0]
}

func (bc *Blockchain) GetBlockHash(index int) []byte {
	return bc.blocks[index].Hash
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
