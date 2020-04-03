package handler

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/matkinhig/developer-community/channels"
	"github.com/matkinhig/developer-community/helper"
	"github.com/matkinhig/developer-community/models"
)

type handlerUser struct {
	db *gorm.DB
}

func NewHandlerUser(db *gorm.DB) *handlerUser {
	return &handlerUser{db}
}

func (h *handlerUser) Save(u models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err := h.db.Model(&models.User{}).Table("wp_users").Create(&u).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return u, nil
	}
	return u, err
}

func (h *handlerUser) FindAll(p helper.Pagination) ([]models.User, error) {
	offset := p.GetOffset()
	limit := p.GetLimit()
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = h.db.Model(&models.User{}).Table("wp_users").Limit(limit).Offset(offset).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

func (h *handlerUser) FindByUserLogin(str string) (models.User, error) {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = h.db.Model(&models.User{}).Table("wp_users").Where("user_login = ? ", str).Find(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("user not found")
	}
	return models.User{}, errors.New("user not found")
}

func (h *handlerUser) UpdateByUserLogin(str string, user models.User) error {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = h.db.Model(&models.User{}).Where("user_login = ?", str).Take(&models.User{}).UpdateColumn(
			map[string]interface{}{
				"user_login": user.UserLogin,
				"user_email": user.UserEmail,
				"user_pass":  user.UserPass,
				"status":     user.UserStatus,
			},
		)
		ch <- true
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return errors.New("Cant Update Record")
		}
		return nil
	}
	return rs.Error
}
