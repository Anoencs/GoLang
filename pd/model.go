package main

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type okr_obj struct {
	gorm.Model

	id               uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	org_id           uuid.UUID `gorm:"type:uuid; not null"`
	user_id          uuid.UUID `gorm:"type:uuid; not null"`
	period_id        uuid.UUID `gorm:"<-"`
	name             string    `gorm:"type:varchar(500)"`
	status           uint      `gorm:"<-"`
	review_date      time.Time `gorm:"type:date"`
	create_date      time.Time `gorm:"type:date"`
	create_by        uuid.UUID `gorm:"<-"`
	last_modified    time.Time `gorm:"type:date"`
	last_modified_by uuid.UUID `gorm:"<-"`
}

type okr_kr struct {
	gorm.Model

	id               uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	obj_id           uuid.UUID `gorm:"type:uuid; not null"`
	user_id          uuid.UUID `gorm:"type:uuid; not null"`
	name             string    `gorm:"<-"`
	itype            uint      `gorm:"<-"`
	criterias        uint
	start            float32
	target           float32
	target_date      time.Time `gorm:"type:date"`
	self_grade       float32
	grade            float32
	duedate          time.Time `gorm:"type:date"`
	create           time.Time `gorm:"type:date"`
	create_by        uuid.UUID
	last_modified    time.Time `gorm:"type:date"`
	last_modified_by uuid.UUID
}

// Model okr_user
type okr_user struct {
	gorm.Model

	user_id       uuid.UUID `gorm:"primary key; type:uuid;not null; default:uuid_generate_v4()"`
	manager_id    uuid.UUID `gorm:"type:uuid;<-"`
	org_id        uuid.UUID `gorm:"type:uuid;<-"`
	email         string    `gorm:"type:varchar(500);<-"`
	manager_email string    `gorm:"type:varchar(500);<-"`
	name          string    `gorm:"type:varchar(100);<-"`
	role          string    `gorm:"type:varchar(50);<-"`
	department    string    `gorm:"type:varchar(50);<-"`
}

type okr_org struct {
	gorm.Model

	id   uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	name string
}

type okr_period struct {
	gorm.Model

	id      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	month   uint
	year    uint
	quarter uint
	name    string
}
