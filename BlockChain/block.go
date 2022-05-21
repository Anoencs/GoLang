package main

import(
	"time"
)

type Block struct {
	timeStamp int64
	data []byte
	prevHash []byte
	hash []byte
	nonce int
}


func NewBlock(data string, prevHash []byte) *Block {
	block := Block{time.Now().Unix(),[]byte(data),prevHash,[]byte{},0}
	pow := NewProofOfWork(&block)
	nonce,hash := pow.Run()
	block.hash = hash[:]
	block.nonce = nonce
	return &block
}

func NewGenesisBlock() *Block{
	return NewBlock("GenesisBlock",[]byte{})
}