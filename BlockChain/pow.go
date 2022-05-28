package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 12

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-targetBits)
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the new block \n")
	for nonce < math.MaxInt64 {
		data := bytes.Join([][]byte{
			pow.block.PrevHash,
			pow.block.HashTransaction(),
			IntToHex(pow.block.TimeStamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		}, []byte{})

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}
