package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wendryuslima/user-manager/src/configuration/logger"
	"github.com/wendryuslima/user-manager/src/configuration/validation"
	"github.com/wendryuslima/user-manager/src/controller/model/request"
	"github.com/wendryuslima/user-manager/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init Create controller", 
	zap.String("journey", "createUser"),
)
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"), )
		restError := validation.ValidateUserError(err)
	c.JSON(int(restError.Code), restError)
	return
	}

	

	response := response.UserResponse{
		
		ID: "test",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,

	}

	logger.Info("User create succesfully", 
	zap.String("journey", "createUser"))
	c.JSON(http.StatusOK, response)
}