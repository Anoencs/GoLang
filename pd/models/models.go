package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Okr_obj struct {
	gorm.Model

	Id               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
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
}

type Okr_kr struct {
	gorm.Model

	Id               uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Obj_id           uuid.UUID `gorm:"type:uuid; not null"`
	User_id          uuid.UUID `gorm:"type:uuid; not null"`
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

	User_id       uuid.UUID `gorm:"primary key; type:uuid;not null; default:uuid_generate_v4()"`
	Manager_id    uuid.UUID `gorm:"type:uuid;<-"`
	Org_id        uuid.UUID `gorm:"type:uuid;<-"`
	Email         string    `gorm:"type:varchar(500);<-"`
	Manager_email string    `gorm:"type:varchar(500);<-"`
	Name          string    `gorm:"type:varchar(100);<-"`
	Role          string    `gorm:"type:varchar(50);<-"`
	Department    string    `gorm:"type:varchar(50);<-"`
}

type Okr_org struct {
	gorm.Model

	Id   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name string
}

type Okr_period struct {
	gorm.Model

	Id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Month   uint64
	Year    uint64
	Quarter uint64
	Name    string
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
