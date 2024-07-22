package controllers_test

import (
	"AuthService/routes/user"
	"VolunteerService/routes/alumni"
	"bytes"
	"encoding/json"
	"fmt"
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
	testUserID    = 2
	testProjectID = 13
	testUser      sharedModels.User
	testProject   sharedModels.Project
	authToken     string
)

func init() {
	// Establish database connection before tests
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

	// Retrieve test data
	testUser = RetrieveUser(testUserID)
	testProject = RetrieveProject(testProjectID)
	fmt.Println("test_Project:", testProject.Name)

	fmt.Println(testUser.Email)
	router := gin.Default()
	user.SetupRouter(router)
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
	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to retrieve user: %v", result.Error))
	}
	return user
}

func RetrieveProject(id int) sharedModels.Project {
	var project sharedModels.Project
	result := initializers.DB.First(&project, id)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to retrieve project: %v", result.Error))
	}
	return project
}

func TestCreateVolunteerRequest(t *testing.T) {
	router := gin.Default()
	//router.Use(middleware.AuthenticateToken())
	alumni.SetupRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/volunteer/create", nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user", testUser)
	c.Set("project", testProject)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCurrentUserVolunteerRequests(t *testing.T) {
	router := gin.Default()
	alumni.SetupRouter(router)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/volunteer")
	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user", testUser)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetVolunteerRequest(t *testing.T) {
	router := gin.Default()
	alumni.SetupRouter(router)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/volunteer/%d", 3)
	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user", testUser)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
func TestDeleteVolunteerRequest(t *testing.T) {
	router := gin.Default()
	alumni.SetupRouter(router)
	w := httptest.NewRecorder()
	url := fmt.Sprintf("/volunteer/%d/delete", 5)
	req, _ := http.NewRequest("DELETE", url, nil)
	req.AddCookie(&http.Cookie{Name: "Authorization", Value: authToken})
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("user", testUser)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}
