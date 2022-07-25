package database

import (
	"fmt"
	"log"

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
	okr_org := models.Okr_org{}
	obj := models.Okr_obj{}
	//read
	okr_obj, okr_kr := obj.Read(excel_fp)
	okr_user.Read(excel_fp)
	okr_period.Read(excel_fp)
	okr_org.Read(excel_fp)

	//Create
	var exists bool
	okr_org.Name = okr_user.Department
	//exist org of okr_org ??
	_ = db.Model(okr_org.Org).
		Select("count(*) > 0").
		Where("Name = ?", okr_org.Org.Name).
		Find(&exists).Error
	if exists {
		res := models.Okr_org{}
		db.First(&res, "Name = ?", okr_org.Org.Name)
		okr_org.Org_id = res.Id
		okr_org.Org.Id = res.Id
	}
	//exists okr_org ?????

	_ = db.Model(okr_org).
		Select("count(*) > 0").
		Where("Name = ?", okr_org.Name). //Where("Name = ?", okr_org.Name).
		Find(&exists).Error
	if !exists { // non exists
		db.Create(&okr_org)
	} else { // exists
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
		okr_user.Manager_id = res.User_id
	}
	////////// exits ork_users ????? ///////////////////////
	_ = db.Model(okr_user).
		Select("count(*) > 0").
		Where("Name = ? AND Role = ?", okr_user.Name, okr_user.Role).
		Find(&exists).Error
	if !exists { //non exist

		// if okr_user.Name == "" {
		// 	okr_user.Name = "UNKNOWN"
		// }
		// if okr_user.Manager.Name == "" {
		// 	okr_user.Manager.Name = "UNKNOWN"
		// }
		// db.Create(&okr_user)
		if okr_user.Name != "" && okr_user.Manager.Name != "" {
			db.Create(&okr_user)
		} else {
			log.Printf("Error xlsxFile: %s, sheet: %s\n", excel_fp.FilePath, excel_fp.SheetName)
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

	// update attribute for objtive
	for i := 0; i < len(okr_obj); i++ {
		okr_obj[i].Org_id = okr_org.Id
		okr_obj[i].Period_id = okr_period.Id
		okr_obj[i].User_id = okr_user.User_id
		okr_obj[i].Create_by = okr_user.User_id
		db.Create(&okr_obj[i])
	}
	// update attribute for key result
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
