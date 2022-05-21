package main

import(
	"math"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"encoding/binary"
	"log"
)

const targetBits = 24 

type ProofOfWork struct{
	block *Block
	target *big.Int
}

func IntToHex(num int64)[]byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func NewProofOfWork(b *Block) *ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target,256-24)
	return &ProofOfWork{b,target}
}

func (pow *ProofOfWork) Run() (int,[]byte){
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.data)
	for nonce < math.MaxInt64 {
		data := bytes.Join([][]byte{
			pow.block.prevHash,
			pow.block.data,
			IntToHex(pow.block.timeStamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},[]byte{})

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		}else{
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}



