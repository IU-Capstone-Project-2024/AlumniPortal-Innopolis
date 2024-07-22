package controllers_test

import (
	auth_user "AuthService/routes/user"
	"ProjectService/controllers"
	user "ProjectService/routes/user"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"alumniportal.com/shared/initializers"
	sharedModels "alumniportal.com/shared/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	testUserID = 3
	testUser   sharedModels.User
	authToken  string
)

func init() {
	// Establish database connection before tests
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	router := gin.Default()

	// Настройка маршрутов
	auth_user.SetupRouter(router)
	// Retrieve test data
	testUser = RetrieveUser(testUserID)
	body := struct {
		Name     string `json:"name"`
		LastName string `json:"lastName"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		IsAlumni bool   `json:"isAlumni"`
	}{
		Email:    "a.lopatov@innopolis.university",
		Password: "dep23m4e",
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

func RetrieveUser(id int) sharedModels.User {
	var user sharedModels.User
	initializers.DB.First(&user, id)

	return user
}

func TestCreateProject(t *testing.T) {

	// Настройка Gin роутера и маршрута
	router := gin.Default()
	user.SetupRouter(router)

	input := controllers.ProjectInput{
		Name:        "PPPPPP",
		Description: "DDDDDD",
	}
	jsonData, _ := json.Marshal(input)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/projects/create", bytes.NewBuffer(jsonData))
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBuffer(jsonData))
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user", testUser)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCurrentUserProjects(t *testing.T) {
	router := gin.Default()
	user.SetupRouter(router)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/projects/user", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user", testUser)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response []sharedModels.Project
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 4)
}

func TestGetProject(t *testing.T) {
	// Add this line to seed test data

	router := gin.Default()
	user.SetupRouter(router)

	w := httptest.NewRecorder()
	url := fmt.Sprintf("/projects/%d", 2)
	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user", testUser)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response sharedModels.Project
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 2, response.ID)
}
func TestUpdateProject(t *testing.T) {
	router := gin.Default()
	user.SetupRouter(router)
	input := controllers.ProjectInput{
		Name:        "New_Name",
		Description: "New_Descr",
	}
	jsonData, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/projects/%d/edit", 3)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user", testUser)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteProject(t *testing.T) {
	router := gin.Default()
	user.SetupRouter(router)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/projects/%d/delete", 6)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user", testUser)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
