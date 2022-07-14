package main

func (cli *CommandLine) update_okr_org(dbname, tbname, id string, name string) {
	database := database{dbname, tbname}
	database.update_okr_org(id, name)

}
func (cli *CommandLine) update_okr_period(dbname, tbname, id, month, year, quarter, name string) {
	database := database{dbname, tbname}
	database.update_okr_period(id, month, year, quarter, name)

}

func (cli *CommandLine) update_okr_obj(dbname, tbname, id, name, org_id, user_id, period_id, status, review_date, create_date, create_by, last_modified, last_modified_by string) {
	database := database{dbname, tbname}
	database.update_okr_obj(id, name, org_id, user_id, period_id, status, review_date, create_date, create_by, last_modified, last_modified_by)
}
