package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wendryuslima/user-manager/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	
	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindByUserEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
