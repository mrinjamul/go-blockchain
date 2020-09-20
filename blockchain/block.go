package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block is a block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock will create new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis is the init
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize : encode blocks to bytes
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()

}

// Deserialize : decodes bytes to block
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

// Handle : handles err
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
