package admin

import (
	"ProjectService/controllers"
	"alumniportal.com/shared/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(route *gin.Engine) {
	admin := route.Group("/projects")
	admin.Use(middleware.RequireAdminRights)
	{
		admin.GET("/admin/:id", controllers.GetProjectAdmin)
		admin.GET("/unverified", controllers.GetUnverifiedProjects)
		admin.PATCH("/edit/:id", controllers.UpdateProjectAdmin)
		admin.DELETE("/delete/:id", controllers.DeleteProject)
		admin.POST("/:id/approve", controllers.ApproveProject)
		admin.POST("/:id/decline", controllers.DeclineProject)
	}
}
