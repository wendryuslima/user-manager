package main

import (
	"github.com/wendryuslima/user-manager/src/controller"
	"github.com/wendryuslima/user-manager/src/model/repository"
	"github.com/wendryuslima/user-manager/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
