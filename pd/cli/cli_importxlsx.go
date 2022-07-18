package cli

import (
	"crud_app/database"
	"crud_app/xlsx"
)

func (cli *CommandLine) import_xlsx(path, sheet string) {
	db := database.Database{DbName: "okr"}
	xlsx := xlsx.Xlsx{FilePath: path, SheetName: sheet}
	db.Import2db(xlsx)

}
