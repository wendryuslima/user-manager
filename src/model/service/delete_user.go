package service

import (
	"github.com/wendryuslima/user-manager/src/configuration/logger"
	"github.com/wendryuslima/user-manager/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("Init delete user model", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		return err
	}

	return nil
}
