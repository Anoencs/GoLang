package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	TimeStamp    int64
	Transactions []*Transaction
	PrevHash     []byte
	Hash         []byte
	Nonce        int
	Height       int
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
func DeserializeBlock(d []byte) *Block {
	var block Block
	var decoder = gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
func (b *Block) HashTransaction() []byte {

	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}

	mTree := NewMerkleTree(transactions)

	return mTree.Root.Data
}

func NewBlock(tx []*Transaction, prevHash []byte, height int) *Block {
	block := Block{time.Now().Unix(), tx, prevHash, []byte{}, 0, height}
	pow := NewProofOfWork(&block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return &block
}

func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{}, 0)
}
