package bitcoin

import (
	"bytes"
	"crypto/sha256"
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

func (b *Block) SetMerkleRoot() {
	transactionCount := len(b.Transactions)
	var merkleTree [][]byte

	// Create leaves of the Merkle Tree
	for _, tran := range b.Transactions {
		merkleTree = append(merkleTree, HashTransaction(tran))
	}

	// Build the Merkle Tree
	for transactionCount > 1 {
		var level [][]byte
		for i := 0; i < transactionCount; i += 2 {
			// Concatenate and hash pairs of nodes
			hash := sha256.Sum256(append(merkleTree[i], merkleTree[i+1]...))
			level = append(level, hash[:])
		}
		// If the number of nodes is odd, duplicate the last one
		if transactionCount%2 == 1 {
			level = append(level, merkleTree[transactionCount-1])
		}
		merkleTree = level
		transactionCount = (transactionCount + 1) / 2
	}

	b.MerkleRoot = merkleTree[0]
}

func HashTransaction(tran *Transaction) []byte {
	bytesData := sha256.Sum256(tran.Data[:])
	return bytesData[:]
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
