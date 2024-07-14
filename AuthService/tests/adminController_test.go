package tests

import (
	"AuthService/controllers"
	"alumniportal.com/shared/initializers"
	"alumniportal.com/shared/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize database connection
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	initializers.DB = db

	// Perform any other setup needed
}

func TestVerify(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/verify", controllers.Verify)

	// Positive case: User exists
	user := models.User{Name: "Test", Email: "test@example.com", IsAdmin: true}
	initializers.DB.Create(&user)

	body := map[string]interface{}{
		"user_id": user.ID,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/verify", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "User verified successfully")

	initializers.DB.Exec("DELETE FROM users WHERE users.id = $1", user.ID)

	// Negative case: User does not exist
	body = map[string]interface{}{
		"user_id": 9999,
	}
	jsonBody, _ = json.Marshal(body)
	req, _ = http.NewRequest(http.MethodPost, "/verify", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Contains(t, resp.Body.String(), "User not found")

}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/delete", controllers.DeleteUser)

	// Positive case: User exists
	user := models.User{Name: "Test", Email: "test@example.com", IsAdmin: true}
	initializers.DB.Create(&user)
	body := map[string]interface{}{
		"user_id": user.ID,
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/delete", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "User deleted successfully")

	initializers.DB.Exec("DELETE FROM users WHERE users.id = $1", user.ID)

	// Negative case: User does not exist
	body = map[string]interface{}{
		"user_id": 9999,
	}
	jsonBody, _ = json.Marshal(body)
	req, _ = http.NewRequest(http.MethodPost, "/delete", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
	assert.Contains(t, resp.Body.String(), "User not found")
}
