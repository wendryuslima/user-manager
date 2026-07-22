package service

import (
	"github.com/wendryuslima/user-manager/src/configuration/logger"
	"github.com/wendryuslima/user-manager/src/configuration/rest_err"
	"github.com/wendryuslima/user-manager/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	logger.Info(
		"User data",
		zap.String("email", userDomain.GetEmail()),
		zap.String("name", userDomain.GetName()),
		zap.Int8("age", userDomain.GetAge()),
	)

	return nil
}
