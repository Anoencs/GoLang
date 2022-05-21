package main 

import(
	"os"
	"fmt"
	"strconv"
)

func main(){
	if len(os.Args) <= 1{
		return
	}
	numRow,_ := strconv.Atoi(os.Args[1])
	db := make(map[string]interface{})
	for i:= 1; i < len(os.Args[2:]); i = i + 2{
		switch os.Args[i+1]{
			case "int":
				db[os.Args[i]] = make([]int,numRow)
			case "string":
				db[os.Args[i]] = make([]string,numRow)
			case "bool":
				db[os.Args[i]] = make([]bool,numRow)
			case "uint":
				db[os.Args[i]] = make([]uint,numRow)
			case "float64":
				db[os.Args[i]] = make([]float64,numRow)
			case "rune":
				db[os.Args[i]] = make([]rune,numRow)
			case "byte":
				db[os.Args[i]] = make([]byte,numRow)
			case "complex64":
				db[os.Args[i]] = make([]complex64,numRow)
		}
		
	}
}