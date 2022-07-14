package main

func (cli *CommandLine) delete(dbname, tbname, id string) {
	database := database{dbname, tbname}
	if tbname == "okr_org" || tbname == "okr_obj" || tbname == "okr_kr" || tbname == "okr_period" {
		database.delete_id(id)
	} else {
		database.delete_userid(id)
	}

}
