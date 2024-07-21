package main

import (
	"ProjectService/routes/admin"
	"ProjectService/routes/user"
	"alumniportal.com/shared/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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

	r.ForwardedByClientIP = true
	if r.SetTrustedProxies([]string{"127.0.0.1"}) != nil {
		panic("SetTrustedProxies failed")
	}

	user.SetupRouter(r)
	admin.SetupRouter(r)

	logrus.Info("Starting Project Service")

	httpServer := &http.Server{
		Addr:    ":8085",
		Handler: r,
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatal("Failed to start HTTP Project Service:", err)
		return
	}
	logrus.Info("HTTP Project Service started")
}
