package cli

import (
	"crud_app/database"
	"crud_app/models"
	"crud_app/util"
	"crud_app/xlsx"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (cli *CommandLine) import_xlsx(path, sheet string) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db := database.Database{DbName: config.DBNAME}
	xlsx := xlsx.Xlsx{FilePath: path, SheetName: sheet}
	cell_reader := xlsx.Read_cell_xlsx()
	check, _ := cell_reader.GetCellValue(sheet, "A2")
	if check == config.OKR_CHECK {
		if sheet != config.SHEET_SKIP {
			db.Import2db(xlsx)
		}
	}

}
func (cli *CommandLine) import_all_xlsx(path string) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db := database.Database{DbName: config.DBNAME}
	xlsx := xlsx.Xlsx{FilePath: path}
	listSheet := xlsx.GetListSheet()
	for _, sheet_name := range listSheet {
		xlsx.SheetName = sheet_name
		fmt.Printf("Reading file %s, sheet name: %s\n", path, sheet_name)
		cell_reader := xlsx.Read_cell_xlsx()
		check, _ := cell_reader.GetCellValue(sheet_name, "A2")
		if check == config.OKR_CHECK {
			if sheet_name != config.SHEET_SKIP {
				db.Import2db(xlsx)
			}
		}
	}
}

func (cli *CommandLine) import_all_xlsx_folder(path string) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	db := database.Database{DbName: config.DBNAME}
	database := db.Connect()
	okr_org := models.Okr_org{}
	okr_org_list := okr_org.Read(path)
	for _, org := range okr_org_list {
		database.Create(&org)
	}
	xlsx := xlsx.Xlsx{}
	// read dir cursively
	list_file_path, _ := cli.FilePathWalkDir(path)
	for _, file_path := range list_file_path {
		xlsx.FilePath = file_path
		listSheet := xlsx.GetListSheet()
		for _, sheet_name := range listSheet {
			//fmt.Printf("Reading file %s, sheet name: %s\n", file_path, sheet_name)
			xlsx.SheetName = sheet_name
			cell_reader := xlsx.Read_cell_xlsx()
			check, _ := cell_reader.GetCellValue(sheet_name, "A2")
			if check == config.OKR_CHECK {
				if sheet_name != config.SHEET_SKIP {
					db.Import2db(xlsx)
				}
			}

		}
	}
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
