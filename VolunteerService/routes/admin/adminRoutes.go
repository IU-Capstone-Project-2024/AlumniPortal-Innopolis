package admin

import (
	"VolunteerService/controllers"

	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	protected := route.Group("/volunteer")

	protected.Use(middleware.RequireAdminRights)
	{
		protected.GET("/:id", controllers.GetVolunteerRequest)
		protected.DELETE("/:id/delete", controllers.DeleteVolunteerRequest)
	}
}
