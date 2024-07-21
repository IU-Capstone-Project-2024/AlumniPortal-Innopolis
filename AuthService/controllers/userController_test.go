package controllers_test

import (
	"AuthService/routes/user"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"alumniportal.com/shared/helpers"
	"alumniportal.com/shared/initializers"
	sharedModels "alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	testUser  sharedModels.User
	authToken string
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	// body := struct {
	// 	Name     string `json:"name"`
	// 	LastName string `json:"lastName"`
	// 	Email    string `json:"email" binding:"required"`
	// 	Password string `json:"password" binding:"required"`
	// 	IsAlumni bool   `json:"isAlumni"`
	// }{
	// 	Email:    testUser.Email,
	// 	Password: "apek2mesa",
	// }
	// jsonData, _ := json.Marshal(body)
	// req, _ := http.NewRequest("POST", "158.160.145.1/auth/login", bytes.NewBuffer(jsonData))
	// req.Header.Set("Content-Type", "application/json")

	// // Send the request
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error sending request:", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// // Check the response status code
	// statusCode := resp.StatusCode
	// fmt.Println("Status Code:", statusCode)
}
func RetrieveUser(id int) sharedModels.User {
	var user sharedModels.User
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to retrieve user: %v", result.Error))
	}
	return user
}

func TestSignUp(t *testing.T) {

	router := gin.Default()
	user.SetupRouter(router)

	// Определение тестов
	body := struct {
		Name     string           `json:"name"`
		LastName string           `json:"lastName"`
		Email    string           `json:"email" binding:"required"`
		Password string           `json:"password" binding:"required"`
		Role     helpers.UserRole `json:"role"`
	}{
		Name:     "Alex",
		LastName: "Kul",
		Email:    "e.kul@innopolis.university",
		Password: "fkek3kf4k",
		Role:     helpers.Student,
	}
	jsonData, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestSuccessfulLogin(t *testing.T) {
	// Создание нового роута Gin
	router := gin.Default()

	// Настройка маршрутов
	user.SetupRouter(router)

	// Подготовка данных для запроса
	body := struct {
		Name     string           `json:"name"`
		LastName string           `json:"lastName"`
		Email    string           `json:"email" binding:"required"`
		Password string           `json:"password" binding:"required"`
		Role     helpers.UserRole `json:"role"`
	}{
		Email:    "m.lan@innopolis.university",
		Password: "djwk2md2",
	}

	// Преобразование данных в JSON
	jsonData, _ := json.Marshal(body)

	// Создание нового запроса
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Создание записи ответа
	w := httptest.NewRecorder()

	// Вызов обработчика маршрута
	router.ServeHTTP(w, req)

	// Проверка ответа
	assert.Equal(t, http.StatusOK, w.Code)
	for _, cookie := range w.Result().Cookies() {
		if cookie.Name == "Authorization" {
			authToken = cookie.Value
			fmt.Println(authToken)
		}
	}

}

func TestErrorLogin(t *testing.T) {
	router := gin.Default()

	// Настройка маршрутов
	user.SetupRouter(router)

	// Подготовка данных для запроса
	body := struct {
		Name     string           `json:"name"`
		LastName string           `json:"lastName"`
		Email    string           `json:"email" binding:"required"`
		Password string           `json:"password" binding:"required"`
		Role     helpers.UserRole `json:"role"`
	}{
		Email:    "a.alexandrov@innopolis.university",
		Password: "adwa",
	}

	// Преобразование данных в JSON
	jsonData, _ := json.Marshal(body)

	// Создание нового запроса
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Создание записи ответа
	w := httptest.NewRecorder()

	// Вызов обработчика маршрута
	router.ServeHTTP(w, req)

	// Проверка ответа
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	fmt.Println(w.Code)

}

func TestValidate(t *testing.T) {
	router := gin.Default()

	user.SetupRouter(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/auth/validate", nil)

	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})

	c, _ := gin.CreateTestContext(w)

	c.Request = req
	c.Set("user", testUser)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetInfo(t *testing.T) {
	router := gin.Default()

	user.SetupRouter(router)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/auth/user", nil)

	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})

	c, _ := gin.CreateTestContext(w)

	c.Request = req

	c.Set("user", testUser)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
