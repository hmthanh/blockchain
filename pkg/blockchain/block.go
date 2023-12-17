package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

type Transaction struct {
	Data []byte
}

func CreateGenesisBlock() *Block {
	currentTimestamp := time.Now().Unix()

	return &Block{
		Timestamp:     currentTimestamp,
		Transactions:  []*Transaction{},
		PrevBlockHash: []byte{},
		Hash:          []byte{},
	}
}

func CreateNewBlock(prev []byte, trans []*Transaction) *Block {
	currentTimestamp := time.Now().Unix()

	return &Block{
		Timestamp:     currentTimestamp,
		Transactions:  trans,
		PrevBlockHash: prev,
		Hash:          []byte{},
	}
}

func SetHash(b *Block) {
	var bytesData []byte
	// fmt.Printf("\nkkkkkkkkkkkkkknkkkkkkkkkkkkkk: %x\n\n", b.PrevBlockHash)

	// PrevBlockHash
	bytesData = append(bytesData, b.PrevBlockHash...)

	// Transactions
	for _, tran := range b.Transactions {
		bytesData = append(bytesData, tran.Data...)
	}

	// Timestamp
	bytesData = append(bytesData, []byte(strconv.FormatInt(b.Timestamp, 10))...)

	// Calculate hash
	hashBytes := sha256.Sum256(bytesData)

	b.Hash = hashBytes[:]
}

func GetHashString(b *Block) string {
	return hex.EncodeToString(b.Hash)
}
