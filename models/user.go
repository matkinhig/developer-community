package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/matkinhig/developer-community/security"
)

type User struct {
	ID                uint32    `gorm:"primary_key;auto_increment" json:"ID"`
	UserLogin         string    `gorm:"size:60; not null; unique" json:"user_login"`
	UserPass          string    `gorm:"size:255; not null" json:"user_pass"`
	UserNicename      string    `gorm:"size:50; not null;" json:"user_nicename"`
	UserEmail         string    `gorm:"size:100; not null" json:"user_email"`
	UserURL           string    `gorm:"size:100" json:"user_url"`
	UserRegisterd     time.Time `gorm:"default:current_timestamp" json:"user_registered"`
	UserActivationKey string    `gorm:"size:255" json:"user_activation_key"`
	UserStatus        uint32    `json:"user_status"`
	DisplayName       string    `gorm:"size:250" json:"display_name"`
}

func (u *User) BeforeSave() error {
	hp, err := security.Hash(u.UserPass)
	if err != nil {
		return err
	}
	u.UserPass = string(hp)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.UserLogin = html.EscapeString(strings.TrimSpace(u.UserLogin))
	u.UserEmail = html.EscapeString(strings.TrimSpace(u.UserEmail))
	u.UserRegisterd = time.Now()
}

func (u *User) Validate(str string) error {
	switch strings.ToLower(str) {
	case "update":
		if u.UserLogin == "" {
			return errors.New("Required Nickname")
		}
		if u.UserEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.UserEmail); err != nil {
			return errors.New("Invalid email")
		}
		return nil
		break
	case "login":
		if u.UserLogin == "" {
			return errors.New("Required user login")
		}
		if err := checkmail.ValidateFormat(u.UserEmail); err != nil {
			return errors.New("Invalid email")
		}
		if u.UserPass == "" {
			return errors.New("Required Password")
		}
		return nil
		break
	default:
		if u.UserLogin == "" {
			return errors.New("Required Nickname")
		}
		if u.UserPass == "" {
			return errors.New("Required Password")
		}
		if u.UserEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.UserEmail); err != nil {
			return errors.New("Invalid email")
		}
		return nil
		break
	}
	return nil
}
