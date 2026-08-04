package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/wendryuslima/user-manager/src/configuration/rest_err"
	"github.com/wendryuslima/user-manager/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		erroMessage := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(int(erroMessage.Code), erroMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		c.JSON(int(err.Code), err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))

}

func (uc *userControllerInterface) FindByUserEmail(c *gin.Context) {

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid user email")
		c.JSON(int(errorMessage.Code), errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		c.JSON(int(err.Code), err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))

}
