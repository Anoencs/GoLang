package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Xlsx struct {
	filePath  string
	sheetName string
}

func (xlsx *Xlsx) read_xlsx() [][]string {
	f, err := excelize.OpenFile(xlsx.filePath)
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(xlsx.sheetName)
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }
	return rows
}
