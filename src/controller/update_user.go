package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wendryuslima/user-manager/src/configuration/logger"
	"github.com/wendryuslima/user-manager/src/configuration/rest_err"
	"github.com/wendryuslima/user-manager/src/configuration/validation"
	"github.com/wendryuslima/user-manager/src/controller/model/request"
	"github.com/wendryuslima/user-manager/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info(
		"Init UpdateUser controller",
		zap.String("journey", "updateUser"),
	)

	userID := strings.TrimSpace(c.Param("userId"))
	if userID == "" {
		restError := rest_err.NewBadRequestError("User ID is required")
		c.JSON(int(restError.Code), restError)
		return
	}

	var userRequest request.UpdateUserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validate user info",
			err,
			zap.String("journey", "updateUser"),
		)

		restError := validation.ValidateUserError(err)
		c.JSON(int(restError.Code), restError)
		return
	}

	domain := model.NewUserUpdateDomain(
		"",
		userID,
		"",
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.UpdateUser(userID, domain); err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info(
		"User updated successfully",
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)
}
