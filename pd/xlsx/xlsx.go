package xlsx

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

type Xlsx struct {
	FilePath  string
	SheetName string
}

func (xlsx *Xlsx) Read_row_xlsx() [][]string {
	f, err := excelize.OpenFile(xlsx.FilePath)
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := f.GetRows(xlsx.SheetName)
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}
	return rows
}

func (xlsx *Xlsx) Read_cell_xlsx() excelize.File {
	f, err := excelize.OpenFile(xlsx.FilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return *f
}
func (xlsx *Xlsx) Close_file() {
	f, _ := excelize.OpenFile(xlsx.FilePath)
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}
func (xlsx *Xlsx) GetListSheet() []string {
	f, err := excelize.OpenFile(xlsx.FilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheet_list := f.GetSheetList()
	var res []string
	for _, sheet := range sheet_list {
		if f.GetSheetVisible(sheet) {
			res = append(res, sheet)
		}
	}
	return res
}
