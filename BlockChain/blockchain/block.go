package main

import(
	"crypto/sha256"
	"time"
	"strconv"
	"bytes"
)

type block struct {
	TimeStamp int64
	Data []byte
	PrevHash []byte
	Hash []byte
}

func (b *block) setHash(){
	time := []byte(strconv.FormatInt(b.TimeStamp, 10))
	headers := bytes.Join([][]byte{time,b.Data,b.PrevHash},[]byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *block {
	block := block{time.Now().Unix(),[]byte(data),prevHash,[]byte{}}
	block.setHash()
	return &block
}

func NewGenesisBlock() *block{
	return NewBlock("GenesisBlock",[]byte{})
}