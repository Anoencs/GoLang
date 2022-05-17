package main

import(
	"fmt"
)

func main(){
	bc := NewBlockChain()
	bc.AddBlock("An sent Tien 10$")
	bc.AddBlock("Tien sent An 100$")
	for _,block :=range(bc.blocks){
		fmt.Printf("Prev Hash: %x\n",block.PrevHash)
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Println()
	}
}