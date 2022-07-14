package main

import (
	"fmt"
	"log"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestIMPORT_okr_org(t *testing.T) {
	database := database{"okr", "okr_org"}
	database.import_xlsx_okr_org()

	db := database.db_connect()
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s;", database.tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			user_id uuid.UUID
			name    string
		)

		err = rows.Scan(&user_id, &name)
		assert.NoError(t, err)
	}
}
func TestIMPORT_okr_period(t *testing.T) {
	database := database{"okr", "okr_period"}
	database.import_xlsx_okr_period()

	db := database.db_connect()
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s;", database.tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			user_id uuid.UUID
			month   uint
			year    uint
			quarter uint
			name    string
		)

		err = rows.Scan(&user_id, &month, &year, &quarter, &name)
		assert.NoError(t, err)
	}
}
