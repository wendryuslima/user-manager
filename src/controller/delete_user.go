package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wendryuslima/user-manager/src/configuration/logger"
	"github.com/wendryuslima/user-manager/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser Controller")

	userId := strings.TrimSpace(c.Param("userId"))
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		restError := rest_err.NewBadRequestError("invalid user ID")
		c.JSON(int(restError.Code), restError)
		return
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		c.JSON(int(err.Code), err)
		return
	}

	logger.Info(
		"User deleted successfully",
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusOK)

}
