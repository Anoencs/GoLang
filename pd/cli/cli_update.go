package cli

import (
	"crud_app/database"
	"crud_app/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (cli *CommandLine) Update_okr_period(dbname, tbname, id, name string, month, year, quarter uint64) *gorm.DB {
	database := database.Database{DbName: dbname, TableName: tbname}
	db := database.Connect()
	uuid, _ := uuid.Parse(id)
	temp := models.Okr_period{Id: uuid, Month: month, Year: year, Quarter: quarter, Name: name}
	return db.Model(&temp).Updates(&temp)

}

func (cli *CommandLine) Update_okr_org(dbname, tbname, id, name string) *gorm.DB {
	database := database.Database{DbName: dbname, TableName: tbname}
	db := database.Connect()
	uuid, _ := uuid.Parse(id)
	temp := models.Okr_org{Id: uuid, Name: name}
	return db.Model(&temp).Updates(&temp)
}

func (cli *CommandLine) Update_okr_obj(dbname, tbname, id, name, org_id, user_id, period_id, review_date, create_date, create_by, last_modified, last_modified_by string, status uint) *gorm.DB {
	database := database.Database{DbName: dbname, TableName: tbname}
	db := database.Connect()
	id_uuid, _ := uuid.Parse(id)
	org_id_uuid, _ := uuid.Parse(org_id)
	user_id_uuid, _ := uuid.Parse(user_id)
	period_id_uuid, _ := uuid.Parse(period_id)
	create_by_uuid, _ := uuid.Parse(create_by)
	last_modified_by_uuid, _ := uuid.Parse(last_modified_by)
	review_date_time, _ := time.Parse("00:00:00", review_date)
	create_date_time, _ := time.Parse("00:00:00", create_date)
	last_modified_time, _ := time.Parse("00:00:00", last_modified)
	temp := models.Okr_obj{Id: id_uuid, Org_id: org_id_uuid, User_id: user_id_uuid, Period_id: period_id_uuid, Name: name, Status: status, Review_date: review_date_time,
		Create_date: create_date_time, Create_by: create_by_uuid, Last_modified: last_modified_time, Last_modified_by: last_modified_by_uuid}
	return db.Model(&temp).Updates(&temp)
}

func (cli *CommandLine) Update_okr_user(dbname, tbname, org_id, name, user_id, manager_id, email, manager_email, role, department string) *gorm.DB {
	database := database.Database{DbName: dbname, TableName: tbname}
	db := database.Connect()
	org_id_uuid, _ := uuid.Parse(org_id)
	manager_id_uuid, _ := uuid.Parse(manager_id)
	user_id_uuid, _ := uuid.Parse(user_id)
	temp := models.Okr_user{User_id: user_id_uuid, Manager_id: manager_id_uuid, Org_id: org_id_uuid, Email: email, Manager_email: manager_email, Name: name, Role: role, Department: department}
	return db.Model(&temp).Updates(&temp)
}

func (cli *CommandLine) Update_okr_kr(dbname, tbname, create, last_modified, duedate, obj_id, name, id, user_id, last_modified_by, target_date, create_by, itype, target string, criterias uint64, start, self_grade, grade float64) *gorm.DB {
	database := database.Database{DbName: dbname, TableName: tbname}
	db := database.Connect()
	id_uuid, _ := uuid.Parse(id)
	obj_id_uuid, _ := uuid.Parse(obj_id)
	user_id_uuid, _ := uuid.Parse(user_id)
	create_by_uuid, _ := uuid.Parse(create_by)
	last_modified_by_uuid, _ := uuid.Parse(last_modified_by)
	duedate_time, _ := time.Parse("00:00:00", duedate)
	create_time, _ := time.Parse("00:00:00", create)
	last_modified_time, _ := time.Parse("00:00:00", last_modified)
	temp := models.Okr_kr{Id: id_uuid, Obj_id: obj_id_uuid, User_id: user_id_uuid, Name: name, Itype: itype, Criterias: criterias, Start: start, Target: target,
		Self_grade: self_grade, Grade: grade, Duedate: duedate_time, Create_date: create_time, Create_by: create_by_uuid, Last_modified: last_modified_time, Last_modified_by: last_modified_by_uuid}

	return db.Model(&temp).Updates(&temp)
}
