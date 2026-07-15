package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wendryuslima/user-manager/src/configuration/validation"
	"github.com/wendryuslima/user-manager/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := validation.ValidateUserError(err)
	c.JSON(int(restError.Code), restError)
	return
	}

	fmt.Println(userRequest)
}