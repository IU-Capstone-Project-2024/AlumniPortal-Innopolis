package admin

import (
	"AuthService/controllers"
	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	admin := route.Group("/auth")
	admin.Use(middleware.RequireAdminRights)
	{
		admin.PATCH("verify", controllers.Verify)
		admin.DELETE("delete_user", controllers.DeleteUser)
	}
}
