package models

import (
	"crud_app/xlsx"
	"errors"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Okr_obj struct {
	Id               uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Org_id           uuid.UUID `gorm:"type:uuid; not null"`
	User_id          uuid.UUID `gorm:"type:uuid; not null"`
	Period_id        uuid.UUID `gorm:"type:uuid;<-"`
	Name             string    `gorm:"type:varchar(500)"`
	Status           uint64    `gorm:"<-"`
	Review_date      time.Time `gorm:"type:date"`
	Create_date      time.Time `gorm:"type:date;default:null"`
	Create_by        uuid.UUID `gorm:"<-"`
	Last_modified    time.Time `gorm:"type:date;default:null"`
	Last_modified_by uuid.UUID `gorm:"<-;default:null"`
	Okr_kr           []Okr_kr  `gorm:"foreignkey:Obj_id;references:Id"`
}

type Okr_kr struct {
	Id               uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Obj_id           uuid.UUID `gorm:"type:uuid; not null"`
	User_id          uuid.UUID `gorm:"type:uuid; not null"`
	Name             string    `gorm:"<-"`
	Itype            string    `gorm:"<-"`
	Criterias        uint64
	Start            float64
	Target           string
	Self_grade       float64
	Grade            float64
	Duedate          time.Time `gorm:"type:date"`
	Create_date      time.Time `gorm:"type:date;default:null"`
	Create_by        uuid.UUID `gorm:"type:uuid"`
	Last_modified    time.Time `gorm:"type:date;default:null"`
	Last_modified_by uuid.UUID `gorm:"type:uuid;default:null"`
}

// Model okr_user
type Okr_user struct {
	User_id             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique"`
	Manager_id          uuid.UUID `gorm:"type:uuid;default:null;<-"`
	Org_id              uuid.UUID `gorm:"type:uuid;<-"`
	Email               string    `gorm:"type:varchar(500);<-"`
	Manager_email       string    `gorm:"type:varchar(500);<-"`
	Name                string    `gorm:"type:varchar(100);<-"`
	Role                string    `gorm:"type:varchar(50);<-"`
	Department          string    `gorm:"type:varchar(50);<-"`
	Manager             *Okr_user `gorm:"foreignkey:Manager_id;references:User_id"`
	Okr_kr              Okr_kr    `gorm:"foreignkey:User_id;references:User_id"`
	Okr_kr_create_by    Okr_kr    `gorm:"foreignkey:Create_by;references:User_id"`
	Okr_kr_modified_by  Okr_kr    `gorm:"foreignkey:Last_modified_by;references:User_id"`
	Okr_obj_create_by   Okr_obj   `gorm:"foreignkey:Create_by;references:User_id"`
	Okr_obj_modified_by Okr_obj   `gorm:"foreignkey:Last_modified_by;references:User_id"`
	Okr_obj_user        Okr_obj   `gorm:"foreignkey:User_id;references:User_id"`
}

type Okr_org struct {
	Id       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string
	Okr_obj  Okr_obj  `gorm:"foreignkey:Org_id;references:Id"`
	Okr_user Okr_user `gorm:"foreignkey:Org_id;references:Id"`
}

type Okr_period struct {
	Id      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Month   uint64
	Year    uint64
	Quarter uint64
	Name    string
	Okr_obj Okr_obj `gorm:"foreignkey:Period_id;references:Id"`
}

func (org *Okr_org) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		return errors.New("NOT CHANGE")
	}
	return nil
}

func (org *Okr_period) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		return errors.New("NOT CHANGE")
	}
	return nil
}
func (org *Okr_obj) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		return errors.New("NOT CHANGE")
	}
	return nil
}
func (org *Okr_kr) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		return errors.New("NOT CHANGE")
	}
	return nil
}
func (org *Okr_user) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed() {
		return errors.New("NOT CHANGE")
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////// Read Excel ////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Read_Excel interface {
	Read(excel xlsx.Xlsx) (interface{}, interface{})
}

func (org *Okr_org) Read(excel xlsx.Xlsx) {
	org.Id = uuid.New()
	org.Name = "ITD Corporation"
}

func (period *Okr_period) Read(excel xlsx.Xlsx) {
	period.Id = uuid.New()
	cell_reader := excel.Read_cell_xlsx()
	myq, err := cell_reader.GetCellValue(excel.SheetName, "H3")
	if err != nil {
		log.Panic(err)
	}
	myq_split := strings.Split(myq, "/")
	if len(myq_split) >= 2 {
		period.Month, _ = strconv.ParseUint(myq_split[1], 10, 64)
		period.Quarter = uint64(math.Ceil(float64(period.Month / 3)))
		period.Year, _ = strconv.ParseUint(myq_split[2], 10, 64)
	} else {
		period.Month = 0
		period.Quarter = 0
		period.Year = 0
	}
	name, err := cell_reader.GetCellValue(excel.SheetName, "F3")
	if err != nil {
		log.Panic(err)
	}
	period.Name = name
}

func (user *Okr_user) Read(excel xlsx.Xlsx) {
	user.User_id = uuid.New()
	cell_reader := excel.Read_cell_xlsx()
	name, err := cell_reader.GetCellValue(excel.SheetName, "C5")
	if err != nil {
		log.Panic(err)
	}
	role, err := cell_reader.GetCellValue(excel.SheetName, "C6")
	if err != nil {
		log.Panic(err)
	}
	deparment, err := cell_reader.GetCellValue(excel.SheetName, "C7")
	if err != nil {
		log.Panic(err)
	}
	user.Name = name
	user.Role = role
	user.Department = deparment
	////////////////////////manager/////////////////////////////////////////////////////
	manager := Okr_user{}
	manager.User_id = uuid.New()
	mngName, err := cell_reader.GetCellValue(excel.SheetName, "E5")
	if err != nil {
		log.Panic(err)
	}
	mngRole, err := cell_reader.GetCellValue(excel.SheetName, "E6")
	if err != nil {
		log.Panic(err)
	}
	mngDeparment, err := cell_reader.GetCellValue(excel.SheetName, "E7")
	if err != nil {
		log.Panic(err)
	}
	manager.Name = mngName
	manager.Role = mngRole
	manager.Department = mngDeparment
	user.Manager = &manager
	user.Manager_id = manager.User_id
}

func (obj *Okr_obj) Read(excel xlsx.Xlsx) ([]Okr_obj, []Okr_kr) {
	okr_obj := []Okr_obj{}
	okr_kr := []Okr_kr{}
	row_reader := excel.Read_row_xlsx()
	cell_reader := excel.Read_cell_xlsx()
	var err error
	for i := 11; i < len(row_reader); i++ {
		if len(row_reader[i]) <= 2 {
			break
		}
		for {
			if len(row_reader[i]) >= 8 {
				break
			}
			row_reader[i] = append(row_reader[i], "")
		}

		if row_reader[i][1] == "" {
			kr_temp := Okr_kr{}
			kr_temp.Obj_id = okr_kr[len(okr_kr)-1].Obj_id
			kr_temp.Name = row_reader[i][2]
			kr_temp.Itype = row_reader[i][3]
			kr_temp.Start, _ = strconv.ParseFloat(row_reader[i][4], 64)
			kr_temp.Target = row_reader[i][5]
			kr_temp.Self_grade, err = strconv.ParseFloat(row_reader[i][6], 64)
			if err != nil {
				kr_temp.Grade = 0
			}
			kr_temp.Grade, err = strconv.ParseFloat(row_reader[i][7], 64)
			if err != nil {
				kr_temp.Grade = 0
			}
			rv_date, _ := cell_reader.GetCellValue(excel.SheetName, "H10")
			kr_temp.Duedate, _ = time.Parse("02/01/2006", rv_date)
			//////////////Append/////////////////////////////////
			okr_kr = append(okr_kr, kr_temp)
			continue
		}
		//////////////////OBJ FORMATE/////////////////////////
		obj_temp := Okr_obj{}
		obj_temp.Id = uuid.New()
		obj_temp.Name = row_reader[i][1]
		rv_date, _ := cell_reader.GetCellValue(excel.SheetName, "H10")
		rv_date_time, _ := time.Parse("02/01/2006", rv_date)
		obj_temp.Review_date = rv_date_time
		/////////////////KR FORMATE////////////////////////////
		kr_temp := Okr_kr{}
		kr_temp.Obj_id = obj_temp.Id
		kr_temp.Name = row_reader[i][2]
		kr_temp.Itype = row_reader[i][3]
		kr_temp.Start, _ = strconv.ParseFloat(row_reader[i][4], 64)
		kr_temp.Target = row_reader[i][5]
		kr_temp.Self_grade, err = strconv.ParseFloat(row_reader[i][6], 64)
		if err != nil {
			kr_temp.Grade = 0
		}
		kr_temp.Grade, err = strconv.ParseFloat(row_reader[i][7], 64)
		if err != nil {
			kr_temp.Grade = 0
		}
		kr_temp.Duedate = rv_date_time
		//////////////Append/////////////////////////////////
		okr_obj = append(okr_obj, obj_temp)
		okr_kr = append(okr_kr, kr_temp)
	}
	return okr_obj, okr_kr
}
