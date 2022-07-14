package main

import (
	_ "github.com/lib/pq"
)

func main() {

	// xlsx := Xlsx{"data.xlsx", "okr_org"}
	// rows := xlsx.read_xlsx()
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }
	cli := CommandLine{}
	cli.Run()
}
