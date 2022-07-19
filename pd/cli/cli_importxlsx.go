package cli

import (
	"crud_app/database"
	"crud_app/xlsx"
	"fmt"
)

func (cli *CommandLine) import_xlsx(path, sheet string) {
	db := database.Database{DbName: "okr"}
	xlsx := xlsx.Xlsx{FilePath: fmt.Sprintf("%s.xlsx", path), SheetName: sheet}
	db.Import2db(xlsx)

}
func (cli *CommandLine) import_all_xlsx(path string) {
	db := database.Database{DbName: "okr"}
	xlsx := xlsx.Xlsx{FilePath: fmt.Sprintf("%s.xlsx", path)}
	listSheet := xlsx.GetListSheet()
	for _, sheet_name := range listSheet {
		xlsx.SheetName = sheet_name
		db.Import2db(xlsx)

	}
}
