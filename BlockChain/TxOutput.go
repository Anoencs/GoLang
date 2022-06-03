package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

type TXOutputs struct {
	OutPuts []TXOutput
}

func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

func DeserializeOutput(data []byte) TXOutputs {
	var res TXOutputs
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&res)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func (outs TXOutputs) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(outs)
	if err != nil {
		log.Panic(err)
	}
	return res.Bytes()
}
