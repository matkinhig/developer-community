package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/database"
	"github.com/matkinhig/developer-community/handler"
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
		g.Writer.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.DONE(g, http.StatusCreated, rs)
	}(handler)
}
