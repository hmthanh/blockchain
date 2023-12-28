package blockchain

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

// func (bc *Blockchain) VerifyMerkleRoot(transactions []*Transaction, merkleRoot []byte) bool {

// 	// Calculate merkle root from transactions
// 	calculatedMerkleRoot := bc.CalculateMerkleRoot(transactions)

// 	// Compare calculated root to passed in root
// 	return bytes.Equal(calculatedMerkleRoot, merkleRoot)
// }

// func VerifyTransaction(transaction *Transaction, merkleRoot []byte) bool {

// 	var hashes [][]byte
// 	for _, t := range transactions {
// 		if bytes.Equal(t.Hash, transaction.Hash) {
// 			hashes = append(hashes, t.Hash)
// 			break
// 		}
// 	}

// 	// Sort hashes
// 	sort.Slice(hashes, func(i, j int) bool {
// 		return bytes.Compare(hashes[i], hashes[j]) < 0
// 	})

// 	// Calculate merkle root from transaction hash
// 	calculatedMerkleRoot := CalculateMerkleTree(hashes)

// 	return bytes.Equal(calculatedMerkleRoot, merkleRoot)
// }

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

// // verifyMerkleTree verifies if the given Merkle root is valid for the specified transaction.
// func VerifyMerkleTree(merkleRoot []byte, transactionHashes [][]byte, targetTransactionHash []byte, position int) bool {
// 	if position < 0 || position >= len(transactionHashes) {
// 		// Invalid position
// 		return false
// 	}

// 	// Check if the given Merkle root matches the calculated Merkle root for the specified transaction
// 	calculatedMerkleRoot := CalculateMerkleRoot(transactionHashes, position)
// 	return fmt.Sprintf("%x", merkleRoot) == fmt.Sprintf("%x", calculatedMerkleRoot)
// }

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
