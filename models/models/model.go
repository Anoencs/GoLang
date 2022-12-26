package models

import (
	"github.com/google/uuid"
)

type StaffEntity struct {
	staff_id             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	user_name            string    `gorm:"type:varchar(512)"`
	password             string    `gorm:"type:varchar(64)"`
	staff_name           string    `gorm:"type:varchar(64)"`
	address              string    `gorm:"typevarchar(128)"`
	date_of_birth        string    `gorm:"type:varchar(32)"`
	sex                  string    `gorm:"type:varchar(8)"`
	home_town            string    `gorm:"type:varchar(512)"`
	phone_number         string    `gorm:"type:varchar(512)"`
	create_date          string    `gorm:"type:varchar(512)"`
	state                string    `gorm:"type:varchar(512)"`
	identity             string    `gorm:"type:varchar(512)"`
	identify_issue_place string    `gorm:"type:varchar(512)"`
	email                string    `gorm:"type:varchar(512)"`
	department_id        uuid.UUID `gorm:"type:uuid"`
}

type StaffRoute struct {
	id                uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	staff_accept_name string    `gorm:"type:varchar(512)"`
	start_date        string    `gorm:"type:varchar(32)"`
	end_day           string    `gorm:"type:varchar(32)"`
	staff_assign_code string    `gorm:"type:varchar(512)"`
	assign_date       string    `gorm:"type:varchar(32)"`
}

type Token struct {
	id           uuid.UUID `gorm:type:uuid;primaryKey;default:uuid_generate_v4()`
	access_token string    `gorm:"type:varchar(512)"`
	device_token string    `gorm:"type:varchar(512)"`
	staff_id     uuid.UUID `gorm:type:uuid`
}
