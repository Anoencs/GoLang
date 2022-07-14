package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

type database struct {
	dbName    string
	tableName string
}

func (database *database) db_connect() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://postgres:1@localhost:5439/%s?sslmode=disable", database.dbName))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (database *database) import_xlsx_okr_period() {
	db := database.db_connect()
	xlsx := Xlsx{"data.xlsx", database.tableName}
	xlsx_reader := xlsx.read_xlsx()
	_, err := db.Query(fmt.Sprintf("DELETE FROM %s;", database.tableName))
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range xlsx_reader {
		_, err = db.Query(fmt.Sprintf("INSERT INTO %s (id,month,year,quarter,name) VALUES ('%s','%s','%s','%s','%s');", database.tableName,
			row[0], row[1], row[2], row[3], row[4]))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (database *database) import_xlsx_okr_org() {
	db := database.db_connect()
	xlsx := Xlsx{"data.xlsx", database.tableName}
	xlsx_reader := xlsx.read_xlsx()
	_, err := db.Query(fmt.Sprintf("DELETE FROM %s;", database.tableName))
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range xlsx_reader {
		_, err = db.Query(fmt.Sprintf("INSERT INTO %s (id,name) VALUES ('%s','%s');", database.tableName, row[0], row[1]))
		if err != nil {
			log.Fatal(err)
		}
	}

}
func (database *database) import_xlsx_okr_obj() {
	db := database.db_connect()
	xlsx := Xlsx{"data.xlsx", database.tableName}
	xlsx_reader := xlsx.read_xlsx()
	_, err := db.Query(fmt.Sprintf("DELETE FROM %s;", database.tableName))
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range xlsx_reader {
		review_date, _ := time.Parse("00:00:00", row[7])
		last_modified, _ := time.Parse("00:00:00", row[9])
		_, err = db.Query(fmt.Sprintf("INSERT INTO %s (id,org_id,user_id,period_id,name,status,review_date,create_date,create_by,last_modified,last_modified_by) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');", database.tableName,
			row[0], row[1], row[2], row[3], row[4], row[5], row[6], review_date, row[8], last_modified, row[10]))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (database *database) import_xlsx_okr_user() {
	db := database.db_connect()
	xlsx := Xlsx{"data.xlsx", database.tableName}
	xlsx_reader := xlsx.read_xlsx()
	_, err := db.Query(fmt.Sprintf("DELETE FROM %s;", database.tableName))
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range xlsx_reader {
		_, err = db.Query(fmt.Sprintf("INSERT INTO %s (user_id,manager_id,org_id,email,manager_email,name,role,department) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s');", database.tableName,
			row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7]))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (database *database) import_xlsx_okr_kr() {
	db := database.db_connect()
	xlsx := Xlsx{"data.xlsx", database.tableName}
	xlsx_reader := xlsx.read_xlsx()
	_, err := db.Query(fmt.Sprintf("DELETE FROM %s;", database.tableName))
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range xlsx_reader {
		_, err = db.Query(fmt.Sprintf("INSERT INTO %s (id,obj_id,user_id,name,itype,criterias,start,target,target_date,self_grade,grade,duedate,create,create_by,last_modified,last_modified_by) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');", database.tableName,
			row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14], row[15]))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (database *database) delete_id(id string) {
	db := database.db_connect()
	_, err := db.Query(fmt.Sprintf("DELETE FROM %s WHERE id= '%s';", database.tableName, id))
	if err != nil {
		log.Fatal(err)
	}
}

func (database *database) delete_userid(id string) {
	db := database.db_connect()
	_, err := db.Query(fmt.Sprintf("DELETE FROM %s WHERE user_id= '%s';", database.tableName, id))
	if err != nil {
		log.Fatal(err)
	}
}
func (database *database) list_okr_org() {
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

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user_id.String(), name)
	}
}
func (database *database) list_okr_period() {
	db := database.db_connect()
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s;", database.tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id      uuid.UUID
			month   uint
			year    uint
			quarter uint
			name    string
		)

		err = rows.Scan(&id, &month, &year, &quarter, &name)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id.String(), month, year, quarter, name)
	}
}

func (database *database) update_okr_org(id, name string) {
	db := database.db_connect()
	_, err := db.Query(fmt.Sprintf("UPDATE %s SET name = '%s' WHERE id = '%s'", database.tableName, name, id))
	if err != nil {
		log.Fatal(err)
	}
}

func (database *database) update_okr_period(id, month, year, quarter, name string) {
	db := database.db_connect()
	_, err := db.Query(fmt.Sprintf("UPDATE %s SET month = %s,year = %s, quarter = %s,name = '%s' WHERE id = '%s';", database.tableName, month, year, quarter, name, id))
	if err != nil {
		log.Fatal(err)
	}
}

func (database *database) update_okr_obj(id, name, org_id, user_id, period_id, status, review_date, create_date, create_by, last_modified, last_modified_by string) {
	db := database.db_connect()
	_, err := db.Query(fmt.Sprintf("UPDATE %s SET org_id = '%s',user_id ='%s',period_id ='%s',name ='%s',status ='%s',review_date ='%s',create_date ='%s',create_by ='%s',last_modified ='%s',last_modified_by ='%s' WHERE id = '%s'", database.tableName,
		org_id, user_id, period_id, name, status, review_date, create_date, create_by, last_modified, last_modified_by, id))
	if err != nil {
		log.Fatal(err)
	}
}
