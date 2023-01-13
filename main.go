package main

import (
	"fmt"
	"github.com/Loner1024/go-blockchain/chain"
)

func main() {
	blockChain := chain.NewBlockChain()
	blockChain.AddBlock("Send 1 Btc to Ivan")
	blockChain.AddBlock("Send 2 Btc to Ivan")
	
	for _, block := range blockChain.Blocks {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
