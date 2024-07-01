package admin

import (
	"EventService/controllers"
	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	admin := route.Group("/event")
	admin.Use(middleware.RequireAdminRights)
	{
		admin.GET("/:id", controllers.GetEventAdmin)
		admin.GET("/unverified", controllers.GetUnverifiedEvents)
		admin.PATCH("/edit/:id", controllers.UpdateEventAdmin)
		admin.DELETE("/delete/:id", controllers.DeleteEvent)
		admin.POST("/:id/approve", controllers.ApproveEvent)
		admin.POST("/:id/decline", controllers.DeclineEvent)
	}
}
