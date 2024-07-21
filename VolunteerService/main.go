package main

import (
	"VolunteerService/routes/admin"
	"VolunteerService/routes/alumni"
	"VolunteerService/routes/student"
	"net/http"

	"alumniportal.com/shared/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://alumni-inno.netlify.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "access-control-allow-origin", "access-control-allow-headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
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

	httpServer := &http.Server{
		Addr:    ":8086",
		Handler: r,
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatal("Failed to start HTTP Volunteer Service:", err)
		return
	}
	logrus.Info("HTTP Volunteer Service started")
}
