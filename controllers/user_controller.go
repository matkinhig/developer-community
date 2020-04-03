package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/database"
	"github.com/matkinhig/developer-community/handler"
	"github.com/matkinhig/developer-community/helper"
	"github.com/matkinhig/developer-community/models"
	"github.com/matkinhig/developer-community/repository"
	"github.com/matkinhig/developer-community/responses"
)

func CreateUser(g *gin.Context) {
	uIn := models.User{}
	if err := g.ShouldBind(&uIn); err != nil {
		responses.FAIL(g, http.StatusPartialContent, err)
		return
	}

	uIn.Prepare()
	if err := uIn.Validate(""); err != nil {
		responses.FAIL(g, http.StatusConflict, err)
		return
	}

	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		responses.FAIL(g, http.StatusInternalServerError, err)
		return
	}
	handler := handler.NewHandlerUser(db)
	func(ur repository.UserRepository) {
		rs, err := ur.Save(uIn)
		if err != nil {
			responses.FAIL(g, http.StatusUnprocessableEntity, err)
			return
		}
		g.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", g.Request.Host, g.Request.URL, rs.ID))
		responses.DONE(g, http.StatusCreated, rs)
	}(handler)
}

//code here tonight

func GetUsers(g *gin.Context) {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		responses.FAIL(g, http.StatusInternalServerError, err)
		return
	}
	handler := handler.NewHandlerUser(db)
	var p helper.Pagination
	func(u repository.UserRepository) {
		us, err := u.FindAll(p)
		if err != nil {
			responses.FAIL(g, http.StatusUnprocessableEntity, err)
			return
		}
		responses.DONE(g, http.StatusOK, us)
	}(handler)
}

func GetUser(g *gin.Context) {
	str := g.Param("userLogin")
	if str == "" {
		responses.FAIL(g, http.StatusUnprocessableEntity, errors.New("Required User Login"))
		return
	}
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		responses.FAIL(g, http.StatusInternalServerError, err)
		return
	}
	handler := handler.NewHandlerUser(db)
	func(repo repository.UserRepository) {
		us, err := repo.FindByUserLogin(str)
		if err != nil {
			responses.FAIL(g, http.StatusUnprocessableEntity, err)
			return
		}
		responses.DONE(g, http.StatusOK, us)
	}(handler)
}

func UpdateUser(g *gin.Context) {
	str := g.Param("userLogin")
	if str == "" {
		responses.FAIL(g, http.StatusUnprocessableEntity, errors.New("Required User Login"))
		return
	}
	body, err := ioutil.ReadAll(g.Request.Body)
	if err != nil {
		responses.FAIL(g, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	if err = json.Unmarshal(body, &user); err != nil {
		responses.FAIL(g, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	if err = user.Validate("update"); err != nil {
		responses.FAIL(g, http.StatusConflict, err)
		return
	}

	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		responses.FAIL(g, http.StatusInternalServerError, err)
		return
	}

	handler := handler.NewHandlerUser(db)
	func(repo repository.UserRepository) {
		err := repo.UpdateByUserLogin(str, user)
		if err != nil {
			responses.FAIL(g, http.StatusConflict, err)
			return
		}
		responses.FAIL(g, http.StatusOK, errors.New("Update Successful"))
	}(handler)
}
