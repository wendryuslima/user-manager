package repository

import (
	"context"
	"os"

	"github.com/wendryuslima/user-manager/src/configuration/logger"
	"github.com/wendryuslima/user-manager/src/configuration/rest_err"
	"github.com/wendryuslima/user-manager/src/model"
	"github.com/wendryuslima/user-manager/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser repository")
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConverteDomainToEntity(userDomain)
	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return rest_err.NewBadRequestError("invalid user ID")
	}
	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
