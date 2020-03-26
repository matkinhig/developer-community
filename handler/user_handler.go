package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/matkinhig/developer-community/channels"
	"github.com/matkinhig/developer-community/models"
)

type handlerUser struct {
	db *gorm.DB
}

func NewHandlerUser(db *gorm.DB) *handlerUser {
	return &handerUser{db}
}

func (h *handlerUser) Save(u models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err := h.db.Model(&models.User{}).Create(&u).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return u, nil
	}
	return nil, err
}
