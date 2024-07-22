package controllers_test

// import (
// 	auth_user "AuthService/routes/user"
// 	"ProjectService/routes/admin"
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"alumniportal.com/shared/initializers"
// 	sharedModels "alumniportal.com/shared/models"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	testAdminID = 4
// 	testAdmin   sharedModels.User
// 	adminTocken string
// )

// func init() {
// 	// Establish database connection before tests
// 	initializers.LoadEnvVariables()
// 	initializers.ConnectToDb()
// 	router := gin.Default()

// 	// Настройка маршрутов
// 	auth_user.SetupRouter(router)
// 	// Retrieve test data
// 	testAdmin = RetrieveAdmin(testAdminID)
// 	body := struct {
// 		Name     string `json:"name"`
// 		LastName string `json:"lastName"`
// 		Email    string `json:"email" binding:"required"`
// 		Password string `json:"password" binding:"required"`
// 		IsAlumni bool   `json:"isAlumni"`
// 	}{
// 		Email:    "g.grigoryev@innopolis.university",
// 		Password: "3kmd2k2l",
// 	}
// 	jsonData, _ := json.Marshal(body)
// 	// req, _ := http.NewRequest("POST", "158.160.145.1/auth/login", bytes.NewBuffer(jsonData))
// 	// req.Header.Set("Content-Type", "application/json")

// 	// // Send the request
// 	// client := &http.Client{}
// 	// resp, err := client.Do(req)
// 	// if err != nil {
// 	// 	fmt.Println("Error sending request:", err)
// 	// 	return
// 	// }
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
// 	req.Header.Set("Content-Type", "application/json")

// 	router.ServeHTTP(w, req)

// 	for _, cookie := range w.Result().Cookies() {
// 		if cookie.Name == "Authorization" {
// 			adminTocken = cookie.Value
// 			fmt.Println(adminTocken)
// 		}
// 	}
// 	//defer resp.Body.Close()

// 	// Check the response status code
// 	statusCode := w.Code
// 	fmt.Println("Status Code:", statusCode)
// }
// func RetrieveAdmin(id int) sharedModels.User {
// 	var user sharedModels.User
// 	initializers.DB.First(&user, id)

// 	return user
// }

// //	func TestApproveProject(t *testing.T) {
// //		router := gin.Default()
// //		admin.SetupRouter(router)
// //		w := httptest.NewRecorder()
// //		url := fmt.Sprintf("/projects/%d/approve", 12)
// //		req, _ := http.NewRequest("PATCH", url, nil)
// //		req.AddCookie(&http.Cookie{Name: "Authorization", Value: adminTocken})
// //		router.ServeHTTP(w, req)
// //		assert.Equal(t, http.StatusOK, w.Code)
// //	}
// // func TestDeclineProject(t *testing.T) {
// // 	router := gin.Default()
// // 	admin.SetupRouter(router)
// // 	w := httptest.NewRecorder()
// // 	url := fmt.Sprintf("/projects/%d/decline", 11)
// // 	req, _ := http.NewRequest("PATCH", url, nil)
// // 	req.AddCookie(&http.Cookie{Name: "Authorization", Value: adminTocken})
// // 	router.ServeHTTP(w, req)
// // 	assert.Equal(t, http.StatusOK, w.Code)
// // }

// func TestGetAdminProject(t *testing.T) {
// 	router := gin.Default()
// 	admin.SetupRouter(router)
// 	w := httptest.NewRecorder()
// 	url := fmt.Sprintf("/projects/admin/%d", 10)
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.AddCookie(&http.Cookie{Name: "Authorization", Value: adminTocken})
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)

// }
// func TestGetUnverifiedProjects(t *testing.T) {
// 	router := gin.Default()
// 	admin.SetupRouter(router)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/projects/unverified", nil)
// 	req.AddCookie(&http.Cookie{Name: "Authorization", Value: adminTocken})

// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response []sharedModels.Project
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)

// }
// func TestAdminDeleteProject(t *testing.T) {
// 	router := gin.Default()
// 	admin.SetupRouter(router)
// 	w := httptest.NewRecorder()
// 	url := fmt.Sprintf("/projects/delete/%d", 10)
// 	req, _ := http.NewRequest("DELETE", url, nil)
// 	req.AddCookie(&http.Cookie{Name: "Authorization", Value: adminTocken})
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = req
// 	c.Set("user", testAdmin)
// 	router.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }
