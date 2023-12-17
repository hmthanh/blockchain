package main

import (
	"fmt"

	"github.com/hmthanh/blockchain/pkg/blockchain"
)

func main() {
	data := "Hello, SH33A"

	// hash := sha256.New()
	// hash.Write([]byte(data))

	// hashValue := hash.Sum(nil)

	// hashString := hex.EncodeToString(hashValue)

	dataBytes := []byte(data)
	trans := []*blockchain.Transaction{
		&blockchain.Transaction{Data: dataBytes}, &blockchain.Transaction{Data: dataBytes},
	}
	res := blockchain.HashTransactions(trans)
	fmt.Print(res)

}
