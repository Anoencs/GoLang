package main

func (cli *CommandLine) import_xlsx(dbname, tbname string) {
	database := database{dbname, tbname}
	if tbname == "okr_org" {
		database.import_xlsx_okr_org()
	} else if tbname == "okr_period" {
		database.import_xlsx_okr_period()
	} else if tbname == "okr_obj" {
		database.import_xlsx_okr_obj()
	} else if tbname == "okr_user" {
		database.import_xlsx_okr_user()
	} else if tbname == "okr_kr" {
		database.import_xlsx_okr_kr()
	}
}
