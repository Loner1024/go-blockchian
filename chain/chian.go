// package chain is about block chain.
package chain

import "github.com/Loner1024/go-blockchain/block"

// BlockChain is a BlockChain.
type BlockChain struct {
	blocks []*block.Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*block.Block{block.NewGenesisBlock()}}
}

// AddBlock add block to chain.
func (bc *BlockChain) AddBlock(data string) {
	prevBlockHash := bc.blocks[len(bc.blocks)-1].PrevBlockHash
	bc.blocks = append(bc.blocks, block.NewBlock(data, prevBlockHash))
}
