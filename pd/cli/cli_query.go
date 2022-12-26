package cli

import (
	"crud_app/database"
)

func (cli *CommandLine) list_user_okr() {
	db := database.Database{DbName: "okr"}
	db.Query_list_user()
}
