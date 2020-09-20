package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Block is a block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// BlockChain : chaining all the blocks
type BlockChain struct {
	blocks []*Block
}

// DeriveHash will create hash
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock will create new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock will add new block to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Genesis is the init
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain will init the blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	for _, Block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", Block.PrevHash)
		fmt.Printf("Data in Block: %s\n", Block.Data)
		fmt.Printf("Hash: %x\n", Block.Hash)
	}
}
