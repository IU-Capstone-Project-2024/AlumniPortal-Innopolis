package main

import (
	"VolunteerService/routes/admin"
	"VolunteerService/routes/alumni"
	"VolunteerService/routes/student"

	"alumniportal.com/shared/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	// r.Use(middleware.AuthenticateToken())

	r.ForwardedByClientIP = true
	if r.SetTrustedProxies([]string{"127.0.0.1"}) != nil {
		panic("SetTrustedProxies failed")
	}

	alumni.SetupRouter(r)
	admin.SetupRouter(r)
	student.SetupRouter(r)

	logrus.Info("Starting Volunteer Service")

	if err := r.Run(":8086"); err != nil {
		logrus.Fatal("Error starting Volunteer Service")
		panic(err)
	}
}
