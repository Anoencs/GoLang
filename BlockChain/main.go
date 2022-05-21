package main

import(
	"fmt"
)

func main(){
	bc := NewBlockchain()
	bc.AddBlock("An sent Tien 10$")
	bc.AddBlock("Tien sent An 100$")
	for _,block :=range(bc.blocks){
		fmt.Printf("Prev Hash: %x\n",block.prevHash)
		fmt.Printf("Data: %s\n",block.data)
		fmt.Printf("Hash: %x\n",block.hash)
		
	}
}