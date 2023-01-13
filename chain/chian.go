// package chain is about block chain.
package chain

import "github.com/Loner1024/go-blockchain/block"

// BlockChain is a BlockChain.
type BlockChain struct {
	Blocks []*block.Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*block.Block{block.NewGenesisBlock()}}
}

// AddBlock add block to chain.
func (bc *BlockChain) AddBlock(data string) {
	prevBlockHash := bc.Blocks[len(bc.Blocks)-1].PrevBlockHash
	bc.Blocks = append(bc.Blocks, block.NewBlock(data, prevBlockHash))
}
