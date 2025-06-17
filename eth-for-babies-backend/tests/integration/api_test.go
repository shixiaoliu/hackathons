package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"eth-for-babies-backend/internal/api/routes"
)

// Setup test environment
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Initialize routes (simplified version for testing)
	routes.SetupRoutes(r)

	return r
}

// Test user registration
func TestUserRegistration(t *testing.T) {
	router := setupTestRouter()

	// Create test user data
	userData := map[string]interface{}{
		"username":      "testuser",
		"password":      "password123",
		"email":         "test@example.com",
		"walletAddress": "0x1234567890abcdef",
	}

	jsonData, _ := json.Marshal(userData)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response
	assert.Nil(t, err)
	assert.Equal(t, true, response["success"])
	assert.NotNil(t, response["user"])
}

// Test user login
func TestUserLogin(t *testing.T) {
	router := setupTestRouter()

	// Create login data
	loginData := map[string]interface{}{
		"username": "testuser",
		"password": "password123",
	}

	jsonData, _ := json.Marshal(loginData)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response
	assert.Nil(t, err)
	assert.Equal(t, true, response["success"])
	assert.NotNil(t, response["token"])
}

// Test create family
func TestCreateFamily(t *testing.T) {
	router := setupTestRouter()

	// Create test token (would normally come from login)
	token := "test-token"

	// Create family data
	familyData := map[string]interface{}{
		"name": "Test Family",
	}

	jsonData, _ := json.Marshal(familyData)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/families", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response
	assert.Nil(t, err)
	assert.Equal(t, true, response["success"])
	assert.NotNil(t, response["family"])
}

// Test create child
func TestCreateChild(t *testing.T) {
	router := setupTestRouter()

	// Create test token (would normally come from login)
	token := "test-token"

	// Create child data
	childData := map[string]interface{}{
		"name":          "Test Child",
		"age":           10,
		"familyId":      1,
		"walletAddress": "0x0987654321abcdef",
	}

	jsonData, _ := json.Marshal(childData)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/children", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response
	assert.Nil(t, err)
	assert.Equal(t, true, response["success"])
	assert.NotNil(t, response["child"])
}

// Test create task
func TestCreateTask(t *testing.T) {
	router := setupTestRouter()

	// Create test token (would normally come from login)
	token := "test-token"

	// Create task data
	taskData := map[string]interface{}{
		"title":       "Test Task",
		"description": "This is a test task",
		"reward":      "0.01",
		"familyId":    1,
	}

	jsonData, _ := json.Marshal(taskData)

	// Create request
	req, _ := http.NewRequest("POST", "/api/v1/tasks", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response
	assert.Nil(t, err)
	assert.Equal(t, true, response["success"])
	assert.NotNil(t, response["task"])
}

// Test get tasks by family
func TestGetTasksByFamily(t *testing.T) {
	router := setupTestRouter()

	// Create test token (would normally come from login)
	token := "test-token"

	// Create request
	req, _ := http.NewRequest("GET", "/api/v1/families/1/tasks", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response
	assert.Nil(t, err)
	assert.Equal(t, true, response["success"])
	assert.NotNil(t, response["tasks"])
}

// Test assign task to child
func TestAssignTask(t *testing.T) {
	router := setupTestRouter()

	// Create test token (would normally come from login)
	token := "test-token"

	// Create assignment data
	assignData := map[string]interface{}{
		"childId": 1,
	}

	jsonData, _ := json.Marshal(assignData)

	// Create request
	req, _ := http.NewRequest("PUT", "/api/v1/tasks/1/assign", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response
	assert.Nil(t, err)
	assert.Equal(t, true, response["success"])
}

// Test complete task
func TestCompleteTask(t *testing.T) {
	router := setupTestRouter()

	// Create test token (would normally come from login)
	token := "test-token"

	// Create completion data
	completeData := map[string]interface{}{
		"proof": "Task completed proof",
	}

	jsonData, _ := json.Marshal(completeData)

	// Create request
	req, _ := http.NewRequest("PUT", "/api/v1/tasks/1/complete", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response
	assert.Nil(t, err)
	assert.Equal(t, true, response["success"])
}

// Test approve task
func TestApproveTask(t *testing.T) {
	router := setupTestRouter()

	// Create test token (would normally come from login)
	token := "test-token"

	// Create request
	req, _ := http.NewRequest("PUT", "/api/v1/tasks/1/approve", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// Perform request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response
	assert.Nil(t, err)
	assert.Equal(t, true, response["success"])
}
