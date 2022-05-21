package main 

import(


)

type Blockchain struct {
	blocks []*Block	
}

func (bc *Blockchain) AddBlock(data string){
	preHash := bc.blocks[len(bc.blocks)-1].hash
	new_block := NewBlock(data,preHash)
	bc.blocks = append(bc.blocks,new_block)
}

func NewBlockchain() *Blockchain{
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}