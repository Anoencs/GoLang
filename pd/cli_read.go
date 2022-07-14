package main

func (cli *CommandLine) read(dbname, tbname string) {
	database := database{dbname, tbname}
	if tbname == "okr_org" {
		database.list_okr_org()
	} else if tbname == "okr_period" {
		database.list_okr_period()
	}
}
