package bitcoin

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
}

func CreateGenesisBlock() *Block {
	return &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  []*Transaction{},
		PrevBlockHash: []byte{},
		Hash:          []byte{},
	}
}

func CreateBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	block.SetMerkleRoot()
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

	// Transactions
	bytesData = append(bytesData, HashTransactions(b.Transactions)...)

	// Timestamp
	bytesData = append(bytesData, []byte(fmt.Sprint(b.Timestamp))...)

	// Calculate hash
	hashBytes := sha256.Sum256(bytesData)

	b.Hash = hashBytes[:]
}

func GetHashString(b *Block) string {
	return hex.EncodeToString(b.Hash)
}
