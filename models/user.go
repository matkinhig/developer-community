package models

import "time"

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"ID"`
	UserLogin string `gorm:"size:60; not null; unique" json:"user_login"`
	UserPass string `gorm:"size:255; not null" json:"user_pass"`
	UserNicename string `gorm:"size:50; not null;" json:"user_nicename"`
	UserEmail string `gorm:"size:100; not null" json:"user_email"`
	UserURL string `gorm:"size:100" json:"user_url"`
	UserRegisterd time.Time `gorm:"default:current_timestamp" json:"user_registered"`

}