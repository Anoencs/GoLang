package database

import (
	"fmt"
	"log"

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
func (database *Database) Delete_by_id(id string) {
	db := database.Connect()
	uuid, _ := uuid.Parse(id)
	if database.TableName == "okr_orgs" {
		db.Unscoped().Delete(&models.Okr_org{}, uuid)
	} else if database.TableName == "okr_periods" {
		db.Unscoped().Delete(&models.Okr_period{}, uuid)
	} else if database.TableName == "okr_objs" {
		db.Unscoped().Delete(&models.Okr_obj{}, uuid)
	} else if database.TableName == "okr_krs" {
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
