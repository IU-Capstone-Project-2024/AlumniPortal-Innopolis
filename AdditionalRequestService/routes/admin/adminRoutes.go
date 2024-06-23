package admin

import (
	"AdditionalRequestService/controllers"
	"github.com/gin-gonic/gin"
	"shared/middleware"
)

func SetupRouter(route *gin.Engine) {
	protected := route.Group("/pass_requests")

	protected.Use(middleware.RequireAdminRights)
	{
		protected.GET("/unverified", controllers.GetUnverifiedRequests)
		protected.DELETE("/:id", controllers.DeletePassRequest)
		protected.POST("/:id/approve", controllers.ApprovePassRequest)
		protected.POST("/:id/decline", controllers.DeclinePassRequest)
	}
}
