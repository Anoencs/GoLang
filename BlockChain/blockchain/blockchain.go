package main


type blockchain struct{
	blocks []*block
}

func (bc *blockchain) AddBlock(data string){
	prevhash := bc.blocks[len(bc.blocks)-1].Hash
	new_block := NewBlock(data,prevhash)
	bc.blocks = append(bc.blocks,new_block)
}

func NewBlockChain() *blockchain{
	return &blockchain{[]*block{NewGenesisBlock()}}
}

