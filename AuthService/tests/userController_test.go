package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"AuthService/controllers"
	"alumniportal.com/shared/initializers"
	"alumniportal.com/shared/models"
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

func TestSignupPositive(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/signup", controllers.Signup)

	// Positive case: User signup
	body := map[string]interface{}{
		"name":     "Test",
		"lastName": "User",
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "{}")

	initializers.DB.Exec("DELETE FROM users WHERE email = 'testuser@example.com'")

}

func TestSignupNegative(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/signup", controllers.Signup)

	// Negative case: Missing email
	body := map[string]interface{}{
		"name":     "Test",
		"lastName": "User",
		"password": "testpassword",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Contains(t, resp.Body.String(), "Failed to read body")
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", controllers.Login)

	// Positive case: User login
	password, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), 10)
	user := models.User{Name: "Test", LastName: "User", Email: "testuser@example.com", Password: string(password)}
	initializers.DB.Create(&user)

	body := map[string]interface{}{
		"email":    "testuser@example.com",
		"password": "testpassword",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "{}")

	// Negative case: Invalid email
	body = map[string]interface{}{
		"email":    "invalid@example.com",
		"password": "testpassword",
	}
	jsonBody, _ = json.Marshal(body)
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid Email or Password")

	// Negative case: Invalid password
	body = map[string]interface{}{
		"email":    "testuser@example.com",
		"password": "wrongpassword",
	}
	jsonBody, _ = json.Marshal(body)
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "Invalid Email or Password")

	initializers.DB.Exec("DELETE FROM users WHERE id = $1", user.ID)
}

func TestGetInfo(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/info", controllers.GetInfo)

	// Positive case: User info
	user := models.User{Name: "Test", LastName: "User", Email: "testuser@example.com", Role: "admin", Specialization: "IT", PortfolioLink: "portfolio.com", SocialsLink: "socials.com", IsAlumni: false, IsAdmin: false}
	initializers.DB.Create(&user)
	
	req, _ := http.NewRequest(http.MethodGet, "/info", nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "valid_token"})

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test")
	assert.Contains(t, resp.Body.String(), "User")
	assert.Contains(t, resp.Body.String(), "testuser@example.com")

	// Negative case: User not found
	req, _ = http.NewRequest(http.MethodGet, "/info", nil)
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: "invalid_token"})

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	assert.Contains(t, resp.Body.String(), "User not found")

	initializers.DB.Exec("DELETE FROM users WHERE id = $1", user.ID)
}
