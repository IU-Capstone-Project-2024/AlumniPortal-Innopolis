package user

import (
	"EventService/controllers"
	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	student := route.Group("/events")
	student.Use(middleware.RequireStudent)
	{
		student.GET("", controllers.GetCurrentUserEvents)
		student.GET("/:id", controllers.GetEvent)
		student.PATCH("/:id/edit", controllers.UpdateEvent)
		student.DELETE("/:id/delete", controllers.DeleteEvent)
		student.POST("/create", controllers.CreateEvent)
	}
}
