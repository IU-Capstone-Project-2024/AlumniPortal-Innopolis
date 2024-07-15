package user

import (
	"ProjectService/controllers"
	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	student := route.Group("/projects")
	student.Use(middleware.RequireStudent)
	{
		student.GET("/user", controllers.GetCurrentUserProjects)
		student.GET("/:id", controllers.GetProject)
		student.PATCH("/:id/edit", controllers.UpdateProject)
		student.DELETE("/:id/delete", controllers.DeleteProject)
		student.POST("/create", controllers.CreateProject)
	}
	student.Use(middleware.RequireVerify)
	{
		student.GET("", controllers.GetProjects)
	}
}
