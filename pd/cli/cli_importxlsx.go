package cli

import (
	"crud_app/database"
	"crud_app/xlsx"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (cli *CommandLine) import_xlsx(path, sheet string) {
	db := database.Database{DbName: "okr"}
	xlsx := xlsx.Xlsx{FilePath: fmt.Sprintf("%s.xlsx", path), SheetName: sheet}
	cell_reader := xlsx.Read_cell_xlsx()
	check, _ := cell_reader.GetCellValue(sheet, "A2")
	if check == "BẢN BÀN GIAO VÀ ĐÁNH GIÁ TRÁCH NHIỆM CÁ NHÂN (OKR)" {
		db.Import2db(xlsx)
	}

}
func (cli *CommandLine) import_all_xlsx(path string) {
	db := database.Database{DbName: "okr"}
	xlsx := xlsx.Xlsx{FilePath: fmt.Sprintf("%s.xlsx", path)}
	listSheet := xlsx.GetListSheet()
	for _, sheet_name := range listSheet {
		xlsx.SheetName = sheet_name
		cell_reader := xlsx.Read_cell_xlsx()
		check, _ := cell_reader.GetCellValue(sheet_name, "A2")
		if check == "BẢN BÀN GIAO VÀ ĐÁNH GIÁ TRÁCH NHIỆM CÁ NHÂN (OKR)" {
			db.Import2db(xlsx)
		}
	}
}

func (cli *CommandLine) import_all_xlsx_folder(path string) {
	db := database.Database{DbName: "okr"}
	xlsx := xlsx.Xlsx{}
	//////////////// read dir ///////////////
	list_file_path, _ := cli.FilePathWalkDir(path)
	for _, file_path := range list_file_path {
		xlsx.FilePath = file_path
		listSheet := xlsx.GetListSheet()
		for _, sheet_name := range listSheet {
			xlsx.SheetName = sheet_name
			cell_reader := xlsx.Read_cell_xlsx()
			check, _ := cell_reader.GetCellValue(sheet_name, "A2")
			if check == "BẢN BÀN GIAO VÀ ĐÁNH GIÁ TRÁCH NHIỆM CÁ NHÂN (OKR)" {
				db.Import2db(xlsx)
			}

		}
	}
	////////////////////////////////////////
}

func (cli *CommandLine) FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		if err != nil {
			log.Panic(err)
		}
		return nil
	})
	return files, err
}
