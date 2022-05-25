package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	timeStamp int64
	data      []byte
	prevHash  []byte
	hash      []byte
	nonce     int
}

/////Serialize struct
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}
func Deserialize(d []byte) *Block {
	var block Block
	var decoder = gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func NewBlock(data string, prevHash []byte) *Block {
	block := Block{time.Now().Unix(), []byte(data), prevHash, []byte{}, 0}
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.hash = hash[:]
	block.nonce = nonce
	return &block
}

func NewGenesisBlock() *Block {
	return NewBlock("GenesisBlock", []byte{})
}
