package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"crud_app/models"
	"crud_app/xlsx"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DbName    string
	TableName string
}

func (database *Database) Init() {
	dbURL := "postgres://postgres:1@localhost:5439/okr?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.Migrator().CreateTable(&models.Okr_org{})
	db.Migrator().CreateTable(&models.Okr_period{})
	db.Migrator().CreateTable(&models.Okr_user{})
	db.Migrator().CreateTable(&models.Okr_obj{})
	db.Migrator().CreateTable(&models.Okr_kr{})
}
func (database *Database) Connect() *gorm.DB {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		"localhost", "5439", "postgres", database.DbName, "1",
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (database *Database) Import_xlsx_okr_period() {
	db := database.Connect()
	if db.Migrator().HasTable(&models.Okr_period{}) {
		db.Migrator().DropTable(&models.Okr_period{})
	}
	db.Migrator().CreateTable(&models.Okr_period{})
	xlsx := xlsx.Xlsx{FilePath: "data.xlsx", SheetName: database.TableName}
	xlsx_reader := xlsx.Read_row_xlsx()
	for _, row := range xlsx_reader {
		id, _ := uuid.Parse(row[0])
		month, _ := strconv.ParseUint(row[1], 10, 4)
		year, _ := strconv.ParseUint(row[2], 10, 4)
		quarter, _ := strconv.ParseUint(row[3], 10, 4)
		var okr_period = models.Okr_period{Id: id, Month: month, Year: year, Quarter: quarter, Name: row[4]}
		db.Create(&okr_period)
	}
}

func (database *Database) Import_xlsx_okr_org() {
	db := database.Connect()
	if db.Migrator().HasTable(&models.Okr_org{}) {
		db.Migrator().DropTable(&models.Okr_org{})
	}
	//db.Migrator().CreateTable(&models.Okr_org{})
	xlsx := xlsx.Xlsx{FilePath: "data.xlsx", SheetName: database.TableName}
	xlsx_reader := xlsx.Read_row_xlsx()
	for _, row := range xlsx_reader {
		uuid, _ := uuid.Parse(row[0])
		name := row[1]
		var okr_org = models.Okr_org{Id: uuid, Name: name}
		db.Create(&okr_org)
	}
}

func (database *Database) Import_xlsx_okr_obj() {
	db := database.Connect()
	if db.Migrator().HasTable(&models.Okr_obj{}) {
		db.Migrator().DropTable(&models.Okr_obj{})
	}
	db.Migrator().CreateTable(&models.Okr_obj{})
	xlsx := xlsx.Xlsx{FilePath: "data.xlsx", SheetName: database.TableName}
	xlsx_reader := xlsx.Read_row_xlsx()
	for _, row := range xlsx_reader {
		id, _ := uuid.Parse(row[0])
		org_id, _ := uuid.Parse(row[1])
		user_id, _ := uuid.Parse(row[2])
		period_id, _ := uuid.Parse(row[3])
		name := row[4]
		status, _ := strconv.ParseUint(row[5], 10, 64)
		review_date, _ := time.Parse("00:00:00", row[6])
		create_date, _ := time.Parse("00:00:00", row[7])
		create_by, _ := uuid.Parse(row[8])
		last_modified, _ := time.Parse("00:00:00", row[9])
		last_modified_by, _ := uuid.Parse(row[10])
		var okr_obj = models.Okr_obj{Id: id, Org_id: org_id, User_id: user_id, Period_id: period_id, Name: name, Status: status, Review_date: review_date,
			Create_date: create_date, Create_by: create_by, Last_modified: last_modified, Last_modified_by: last_modified_by}
		db.Create(&okr_obj)
	}
}

func (database *Database) Import_xlsx_okr_user() {
	db := database.Connect()
	if db.Migrator().HasTable(&models.Okr_user{}) {
		db.Migrator().DropTable(&models.Okr_user{})
	}
	db.Migrator().CreateTable(&models.Okr_user{})
	xlsx := xlsx.Xlsx{FilePath: "data.xlsx", SheetName: database.TableName}
	xlsx_reader := xlsx.Read_row_xlsx()

	for _, row := range xlsx_reader {
		user_id, _ := uuid.Parse(row[0])
		manager_id, _ := uuid.Parse(row[1])
		org_id, _ := uuid.Parse(row[2])
		email := row[3]
		manager_email := row[4]
		name := row[5]
		role := row[6]
		department := row[7]
		var okr_user = models.Okr_user{User_id: user_id, Manager_id: manager_id, Org_id: org_id, Email: email, Manager_email: manager_email, Name: name, Role: role, Department: department}
		db.Create(&okr_user)
	}
}

func (database *Database) Import_xlsx_okr_kr() {
	db := database.Connect()
	if db.Migrator().HasTable(&models.Okr_kr{}) {
		db.Migrator().DropTable(&models.Okr_kr{})
	}
	db.Migrator().CreateTable(&models.Okr_kr{})
	xlsx := xlsx.Xlsx{FilePath: "data.xlsx", SheetName: database.TableName}
	xlsx_reader := xlsx.Read_row_xlsx()
	for _, row := range xlsx_reader {
		id, _ := uuid.Parse(row[0])
		obj_id, _ := uuid.Parse(row[1])
		user_id, _ := uuid.Parse(row[2])
		name := row[3]
		itype := row[4]
		criterias, _ := strconv.ParseUint(row[5], 10, 64)
		start, _ := strconv.ParseFloat(row[6], 64)
		target := row[7]
		self_grade, _ := strconv.ParseFloat(row[9], 64)
		grade, _ := strconv.ParseFloat(row[10], 64)
		duedate, _ := time.Parse("00:00:00", row[11])
		create, _ := time.Parse("00:00:00", row[12])
		create_by, _ := uuid.Parse(row[13])
		last_modified, _ := time.Parse("00:00:00", row[14])
		last_modified_by, _ := uuid.Parse(row[15])
		var okr_kr = models.Okr_kr{Id: id, Obj_id: obj_id, User_id: user_id, Name: name, Itype: itype, Criterias: criterias, Start: start, Target: target,
			Self_grade: self_grade, Grade: grade, Duedate: duedate, Create_date: create, Create_by: create_by, Last_modified: last_modified, Last_modified_by: last_modified_by}
		db.Create(&okr_kr)
	}
}

func (database *Database) Delete_by_id(id string) {
	db := database.Connect()
	uuid, _ := uuid.Parse(id)
	if database.TableName == "okr_org" {
		db.Unscoped().Delete(&models.Okr_org{}, uuid)
	} else if database.TableName == "okr_period" {
		db.Unscoped().Delete(&models.Okr_period{}, uuid)
	} else if database.TableName == "okr_obj" {
		db.Unscoped().Delete(&models.Okr_obj{}, uuid)
	} else if database.TableName == "okr_kr" {
		db.Unscoped().Delete(&models.Okr_kr{}, uuid)
	} else {
		db.Unscoped().Delete(&models.Okr_user{}, uuid)
	}

}

func (database *Database) Import2db(xlsx xlsx.Xlsx) {
	db := database.Connect()
	okr_user := models.Okr_user{}
	okr_period := models.Okr_period{}
	okr_org := models.Okr_org{}
	obj := models.Okr_obj{}
	///////////// read ///////////////////////////////////////
	okr_obj, okr_kr := obj.Read(xlsx)
	okr_user.Read(xlsx)
	okr_period.Read(xlsx)
	okr_org.Read(xlsx)

	////////////////////// Create ////////////////////////////

	//////////////exists okr_org ????? ///////////////////////
	var exists bool
	_ = db.Model(okr_org).
		Select("count(*) > 0").
		Where("Name = ?", okr_org.Name).
		Find(&exists).Error
	if !exists {
		db.Create(&okr_org)
	} else {
		res := models.Okr_org{}
		db.First(&res, "Name = ?", okr_org.Name)
		okr_org.Id = res.Id
	}
	///////////////////////////////////////////////////////////

	okr_user.Org_id = okr_org.Id
	okr_user.Manager.Org_id = okr_org.Id
	db.Create(&okr_period)
	//////////////////////// manager exists ???????//////////////////////////////////////
	_ = db.Model(okr_user.Manager).
		Select("count(*) > 0").
		Where("Name = ? AND Role = ?", okr_user.Manager.Name, okr_user.Manager.Role).
		Find(&exists).Error
	if exists {
		res := models.Okr_user{}
		db.First(&res, "Name = ?", okr_user.Manager.Name)
		okr_user.Manager.User_id = res.User_id
	}
	////////// exits ork_users ????? ///////////////////////
	_ = db.Model(okr_user).
		Select("count(*) > 0").
		Where("Name = ? AND Role = ?", okr_user.Name, okr_user.Role).
		Find(&exists).Error
	if !exists {
		db.Create(&okr_user)
	} else {
		res := models.Okr_user{}
		db.First(&res, "Name = ?", okr_user.Name)
		okr_user.User_id = res.User_id
		if res.Manager_id == uuid.Nil {
			if okr_user.Manager_id != uuid.Nil {
				res.Manager_id = okr_user.Manager_id
				db.Create(&okr_user.Manager)
				db.Save(&res)
			}
		}
	}

	////////////////////////////////////////////////////////
	for i := 0; i < len(okr_obj); i++ {
		okr_obj[i].Org_id = okr_org.Id
		okr_obj[i].Period_id = okr_period.Id
		okr_obj[i].User_id = okr_user.User_id
		okr_obj[i].Create_by = okr_user.User_id
		db.Create(&okr_obj[i])
	}
	for i := 0; i < len(okr_kr)-1; i++ {
		okr_kr[i].Create_by = okr_user.User_id
		okr_kr[i].User_id = okr_user.User_id
		db.Create(&okr_kr[i])
	}

}
