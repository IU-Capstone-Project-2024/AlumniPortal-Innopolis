package controllers_test

import (
	"AuthService/routes/admin"
	"AuthService/routes/user"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"alumniportal.com/shared/initializers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var authTocken string

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	router := gin.Default()
	user.SetupRouter(router)
	body := struct {
		Name     string `json:"name"`
		LastName string `json:"lastName"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		IsAlumni bool   `json:"isAlumni"`
	}{
		Email:    "a.kobalt@innopolis.university",
		Password: "adel23l4",
	}
	jsonData, _ := json.Marshal(body)
	// req, _ := http.NewRequest("POST", "158.160.145.1/auth/login", bytes.NewBuffer(jsonData))
	// req.Header.Set("Content-Type", "application/json")

	// // Send the request
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error sending request:", err)
	// 	return
	// }
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	for _, cookie := range w.Result().Cookies() {
		if cookie.Name == "Authorization" {
			authToken = cookie.Value
			fmt.Println(authToken)
		}
	}
	//defer resp.Body.Close()

	// Check the response status code
	statusCode := w.Code
	fmt.Println("Status Code:", statusCode)
}

func TestVerify(t *testing.T) {
	userId := struct {
		UserID uint `json:"user_id" binding:"required"`
	}{
		UserID: 13,
	}
	router := gin.Default()
	admin.SetupRouter(router)
	jsonData, _ := json.Marshal(userId)
	req, _ := http.NewRequest("PATCH", "/auth/verify", bytes.NewBuffer(jsonData))
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	req.Header.Set("Content-Type", "application/json")

	// Создание записи ответа
	w := httptest.NewRecorder()

	// Вызов обработчика маршрута
	router.ServeHTTP(w, req)

	// Проверка ответа
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Code)
}
func TestDelete(t *testing.T) {
	userId := struct {
		UserID uint `json:"user_id" binding:"required"`
	}{
		UserID: 13,
	}
	router := gin.Default()
	admin.SetupRouter(router)
	jsonData, _ := json.Marshal(userId)
	req, _ := http.NewRequest("DELETE", "/auth/delete_user", bytes.NewBuffer(jsonData))
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	req.Header.Set("Content-Type", "application/json")

	// Создание записи ответа
	w := httptest.NewRecorder()

	// Вызов обработчика маршрута
	router.ServeHTTP(w, req)

	// Проверка ответа
	assert.Equal(t, http.StatusOK, w.Code)
	fmt.Println(w.Code)
}
