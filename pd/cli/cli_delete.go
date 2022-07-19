package cli

import (
	"crud_app/database"
)

func (cli *CommandLine) delete(dbname, tbname, id string) {
	database := database.Database{DbName: dbname, TableName: tbname}
	database.Delete_by_id(id)
}
