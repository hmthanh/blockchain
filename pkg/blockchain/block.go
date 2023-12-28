package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	MerkleRoot    []byte
	Hash          []byte
}

type Transaction struct {
	Data []byte
	Hash []byte
}

func CreateGenesisBlock() *Block {
	return &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  []*Transaction{},
		PrevBlockHash: []byte{},
		MerkleRoot:    []byte{},
		Hash:          []byte{},
	}
}

func CreateBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		MerkleRoot:    []byte{},
		Hash:          []byte{},
	}

	block.SetHash()
	return block
}

func HashTransactions(trans []*Transaction) []byte {
	var transactionData []byte
	for _, tran := range trans {
		transactionData = append(transactionData, tran.Data...)
	}

	hash := sha256.Sum256(transactionData)
	return hash[:]
}

func (b *Block) SetHash() {
	var bytesData []byte

	// PrevBlockHash
	bytesData = append(bytesData, b.PrevBlockHash...)

	// MerkleRoot
	bytesData = append(bytesData, b.MerkleRoot...)

	// Transactions
	// bytesData = append(bytesData, HashTransactions(b.Transactions)...)

	// Timestamp
	bytesData = append(bytesData, []byte(fmt.Sprint(b.Timestamp))...)

	// Calculate hash
	hashBytes := sha256.Sum256(bytesData)

	// block.SetMerkleRoot()
	b.Hash = hashBytes[:]
}

func CalculateMerkleRoot(hashes [][]byte) []byte {
	if len(hashes) == 0 {
		return nil
	}

	if len(hashes) == 1 {
		return hashes[0]
	}

	var newHashes [][]byte

	// Combine pairs of hashes and hash them together
	for i := 0; i < len(hashes)-1; i += 2 {
		combined := append(hashes[i], hashes[i+1]...)
		hash := sha256.Sum256(combined)
		newHashes = append(newHashes, hash[:])
	}

	// If there's an odd number of hashes, duplicate the last one
	if len(hashes)%2 == 1 {
		newHashes = append(newHashes, hashes[len(hashes)-1])
	}

	return CalculateMerkleRoot(newHashes)
}

func GetHashString(b *Block) string {
	return hex.EncodeToString(b.Hash)
}
