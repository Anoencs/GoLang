package cli

import (
	"crud_app/database"
)

func (cli *CommandLine) import_xlsx(dbname, tbname string) {
	database := database.Database{DbName: dbname, TableName: tbname}
	if tbname == "okr_org" {
		database.Import_xlsx_okr_org()
	} else if tbname == "okr_period" {
		database.Import_xlsx_okr_period()
	} else if tbname == "okr_obj" {
		database.Import_xlsx_okr_obj()
	} else if tbname == "okr_user" {
		database.Import_xlsx_okr_user()
	} else if tbname == "okr_kr" {
		database.Import_xlsx_okr_kr()
	}

}
