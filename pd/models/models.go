package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Okr_obj struct {
	gorm.Model

	Id               uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Org_id           uuid.UUID `gorm:"type:uuid; not null"`
	User_id          uuid.UUID `gorm:"type:uuid; not null"`
	Period_id        uuid.UUID `gorm:"<-"`
	Name             string    `gorm:"type:varchar(500)"`
	Status           uint64    `gorm:"<-"`
	Review_date      time.Time `gorm:"type:date"`
	Create_date      time.Time `gorm:"type:date"`
	Create_by        uuid.UUID `gorm:"<-"`
	Last_modified    time.Time `gorm:"type:date"`
	Last_modified_by uuid.UUID `gorm:"<-"`
	Okr_kr           Okr_kr    `gorm:"foreignkey:Obj_id;references:Id"`
}

type Okr_kr struct {
	gorm.Model

	Id               uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Obj_id           uuid.UUID `gorm:"type:uuid; not null;"`
	User_id          uuid.UUID `gorm:"type:uuid; not null;"`
	Name             string    `gorm:"<-"`
	Itype            uint64    `gorm:"<-"`
	Criterias        uint64
	Start            float64
	Target           float64
	Target_date      time.Time `gorm:"type:date"`
	Self_grade       float64
	Grade            float64
	Duedate          time.Time `gorm:"type:date"`
	Create           time.Time `gorm:"type:date"`
	Create_by        uuid.UUID
	Last_modified    time.Time `gorm:"type:date"`
	Last_modified_by uuid.UUID
}

// Model okr_user
type Okr_user struct {
	gorm.Model

	User_id       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();unique"`
	Manager_id    uuid.UUID `gorm:"type:uuid;<-"`
	Org_id        uuid.UUID `gorm:"type:uuid;<-"`
	Email         string    `gorm:"type:varchar(500);<-"`
	Manager_email string    `gorm:"type:varchar(500);<-"`
	Name          string    `gorm:"type:varchar(100);<-"`
	Role          string    `gorm:"type:varchar(50);<-"`
	Department    string    `gorm:"type:varchar(50);<-"`
	Manager       *Okr_user `gorm:"foreignkey:Manager_id;references:User_id"`
	Okr_kr        Okr_kr    `gorm:"foreignkey:User_id;references:User_id"`
}

type Okr_org struct {
	gorm.Model

	Id       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string
	Okr_obj  Okr_obj  `gorm:"foreignkey:Org_id;references:Id"`
	Okr_user Okr_user `gorm:"foreignkey:Org_id;references:Id"`
}

type Okr_period struct {
	gorm.Model

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
