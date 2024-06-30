package admin

import (
	"AdditionalRequestService/controllers"
	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	protected := route.Group("/requests")

	protected.Use(middleware.RequireAdminRights)
	{
		protected.GET("/unverified", controllers.GetUnverifiedRequests)
		protected.GET("/unverified/:id", controllers.GetAdminPassRequest)
		protected.DELETE("/delete/:id", controllers.DeletePassRequest)
		protected.POST("/:id/approve", controllers.ApprovePassRequest)
		protected.POST("/:id/decline", controllers.DeclinePassRequest)
	}
}
