package database

import (
	"fmt"
	"log"
	"path/filepath"

	"crud_app/models"
	"crud_app/util"
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
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dbURL := config.DBSource
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
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	dbURL := config.DBSource
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
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
func (database *Database) Import2db(excel_fp xlsx.Xlsx) {
	db := database.Connect()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	okr_user := models.Okr_user{}
	okr_period := models.Okr_period{}
	obj := models.Okr_obj{}
	//read
	okr_obj, okr_kr := obj.Read(excel_fp)
	okr_user.Read(excel_fp)
	okr_period.Read(excel_fp)

	//Create
	var exists bool
	// create org
	org := models.Okr_org{}
	db.First(&org, "Name = ?", filepath.Base(filepath.Dir(excel_fp.FilePath)))

	okr_user.Org_id = org.Id
	okr_user.Manager.Org_id = org.Id
	db.Create(&okr_period) // create period
	//create user
	//manager exists ???????
	_ = db.Model(okr_user.Manager).
		Select("count(*) > 0").
		Where("Name = ? AND Role = ?", okr_user.Manager.Name, okr_user.Manager.Role).
		Find(&exists).Error
	if exists {
		res := models.Okr_user{}
		db.First(&res, "Name = ?", okr_user.Manager.Name)
		okr_user.Manager_id = res.User_id
	}
	//exits ork_users ?????
	_ = db.Model(okr_user).
		Select("count(*) > 0").
		Where("Name = ? AND Role = ?", okr_user.Name, okr_user.Role).
		Find(&exists).Error
	if !exists { //non exist
		if okr_user.Name != "" && okr_user.Manager.Name != "" {
			db.Create(&okr_user)
		} else {
			log.Printf("Error username xlsxFile: %s, sheet: %s\n", excel_fp.FilePath, excel_fp.SheetName)
			return
		}

	} else { //exist
		res := models.Okr_user{}
		db.First(&res, "Name = ? AND Role = ?", okr_user.Name, okr_user.Role)
		if res.Manager_id == uuid.Nil {
			if okr_user.Manager_id != uuid.Nil {
				res.Manager_id = okr_user.Manager_id
				db.Create(&okr_user.Manager)
				db.Save(&res)
			}
		}
		okr_user.User_id = res.User_id
	}

	// update attribute for objtive and create obj
	for i := 0; i < len(okr_obj); i++ {
		okr_obj[i].Org_id = org.Id
		okr_obj[i].Period_id = okr_period.Id
		okr_obj[i].User_id = okr_user.User_id
		okr_obj[i].Create_by = okr_user.User_id
		db.Create(&okr_obj[i])
	}
	// update attribute for key result and create kr
	for i := 0; i < len(okr_kr)-1; i++ {
		okr_kr[i].Create_by = okr_user.User_id
		okr_kr[i].User_id = okr_user.User_id
		db.Create(&okr_kr[i])
	}
}
func (database *Database) Query_list_user() {
	db := database.Connect()
	var users []models.Okr_user
	db.Raw("SELECT DISTINCT  User_id,Name FROM okr_users ").Scan(&users)
	fmt.Println(len(users))
	for _, user := range users {
		fmt.Printf("Id: %s,Name: %s \n", user.User_id, user.Name)
	}
}

func (database *Database) Query_list_numsobj_user() {
	var users []models.Okr_user
	db := database.Connect()
	db.Table("okr_users").
		Select("okr_users.user_id,okr_users.name, count(okr_objs.id) as numObjs").
		Joins("inner join okr_objs on okr_users.user_id = okr_objs.user_id").
		Group("okr_users.user_id").Scan(&users)
	println(len(users))
	for _, user := range users {
		fmt.Printf("Id: %s Name : %s , Objs: %d\n", user.User_id, user.Name, user.Numobjs)
	}

}
