package view

import (
	"github.com/wendryuslima/user-manager/src/controller/model/response"
	"github.com/wendryuslima/user-manager/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID: "",
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
		
	}


}