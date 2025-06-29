package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/services"
)

// Mock repositories
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetByID(id int64) (*models.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAll() ([]*models.Task, error) {
	args := m.Called()
	return args.Get(0).([]*models.Task), args.Error(1)
}

func (m *MockTaskRepository) GetByFamilyID(familyID int64) ([]*models.Task, error) {
	args := m.Called(familyID)
	return args.Get(0).([]*models.Task), args.Error(1)
}

func (m *MockTaskRepository) GetByChildID(childID int64) ([]*models.Task, error) {
	args := m.Called(childID)
	return args.Get(0).([]*models.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockFamilyRepository struct {
	mock.Mock
}

func (m *MockFamilyRepository) Create(family *models.Family) error {
	args := m.Called(family)
	return args.Error(0)
}

func (m *MockFamilyRepository) GetByID(id int64) (*models.Family, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Family), args.Error(1)
}

func (m *MockFamilyRepository) GetByParentID(parentID int64) ([]*models.Family, error) {
	args := m.Called(parentID)
	return args.Get(0).([]*models.Family), args.Error(1)
}

func (m *MockFamilyRepository) Update(family *models.Family) error {
	args := m.Called(family)
	return args.Error(0)
}

func (m *MockFamilyRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockChildRepository struct {
	mock.Mock
}

func (m *MockChildRepository) Create(child *models.Child) error {
	args := m.Called(child)
	return args.Error(0)
}

func (m *MockChildRepository) GetByID(id int64) (*models.Child, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Child), args.Error(1)
}

func (m *MockChildRepository) GetByFamilyID(familyID int64) ([]*models.Child, error) {
	args := m.Called(familyID)
	return args.Get(0).([]*models.Child), args.Error(1)
}

func (m *MockChildRepository) GetByWalletAddress(walletAddress string) (*models.Child, error) {
	args := m.Called(walletAddress)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Child), args.Error(1)
}

func (m *MockChildRepository) Update(child *models.Child) error {
	args := m.Called(child)
	return args.Error(0)
}

func (m *MockChildRepository) Delete(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

// Tests for TaskService
func TestTaskService_CreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := services.NewTaskService(mockRepo)

	task := &models.Task{
		Title:       "Test Task",
		Description: "Test Description",
		Reward:      "0.01",
		Status:      "available",
		CreatorID:   1,
		FamilyID:    1,
	}

	// Set up expectations
	mockRepo.On("Create", task).Return(nil)

	// Call the service
	err := service.CreateTask(task)

	// Assert expectations
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTaskService_GetTaskByID(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := services.NewTaskService(mockRepo)

	task := &models.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Reward:      "0.01",
		Status:      "available",
		CreatorID:   1,
		FamilyID:    1,
	}

	// Set up expectations
	mockRepo.On("GetByID", int64(1)).Return(task, nil)

	// Call the service
	result, err := service.GetTaskByID(1)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskService_AssignTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	service := services.NewTaskService(mockRepo)

	task := &models.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Reward:      "0.01",
		Status:      "available",
		CreatorID:   1,
		FamilyID:    1,
	}

	updatedTask := &models.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
		Reward:      "0.01",
		Status:      "in-progress",
		CreatorID:   1,
		FamilyID:    1,
		AssignedTo:  2,
	}

	// Set up expectations
	mockRepo.On("GetByID", int64(1)).Return(task, nil)
	mockRepo.On("Update", updatedTask).Return(nil)

	// Call the service
	err := service.AssignTask(1, 2)

	// Assert expectations
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// Tests for FamilyService
func TestFamilyService_CreateFamily(t *testing.T) {
	mockRepo := new(MockFamilyRepository)
	service := services.NewFamilyService(mockRepo)

	family := &models.Family{
		Name:     "Test Family",
		ParentID: 1,
	}

	// Set up expectations
	mockRepo.On("Create", family).Return(nil)

	// Call the service
	err := service.CreateFamily(family)

	// Assert expectations
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFamilyService_GetFamilyByID(t *testing.T) {
	mockRepo := new(MockFamilyRepository)
	service := services.NewFamilyService(mockRepo)

	family := &models.Family{
		ID:       1,
		Name:     "Test Family",
		ParentID: 1,
	}

	// Set up expectations
	mockRepo.On("GetByID", int64(1)).Return(family, nil)

	// Call the service
	result, err := service.GetFamilyByID(1)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, family, result)
	mockRepo.AssertExpectations(t)
}

// Tests for ChildService
func TestChildService_CreateChild(t *testing.T) {
	mockRepo := new(MockChildRepository)
	service := services.NewChildService(mockRepo)

	child := &models.Child{
		Name:          "Test Child",
		Age:           10,
		FamilyID:      1,
		WalletAddress: "0x1234567890",
	}

	// Set up expectations
	mockRepo.On("Create", child).Return(nil)

	// Call the service
	err := service.CreateChild(child)

	// Assert expectations
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestChildService_GetChildByID(t *testing.T) {
	mockRepo := new(MockChildRepository)
	service := services.NewChildService(mockRepo)

	child := &models.Child{
		ID:            1,
		Name:          "Test Child",
		Age:           10,
		FamilyID:      1,
		WalletAddress: "0x1234567890",
	}

	// Set up expectations
	mockRepo.On("GetByID", int64(1)).Return(child, nil)

	// Call the service
	result, err := service.GetChildByID(1)

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, child, result)
	mockRepo.AssertExpectations(t)
}

func TestChildService_GetChildByWalletAddress(t *testing.T) {
	mockRepo := new(MockChildRepository)
	service := services.NewChildService(mockRepo)

	child := &models.Child{
		ID:            1,
		Name:          "Test Child",
		Age:           10,
		FamilyID:      1,
		WalletAddress: "0x1234567890",
	}

	// Set up expectations
	mockRepo.On("GetByWalletAddress", "0x1234567890").Return(child, nil)

	// Call the service
	result, err := service.GetChildByWalletAddress("0x1234567890")

	// Assert expectations
	assert.NoError(t, err)
	assert.Equal(t, child, result)
	mockRepo.AssertExpectations(t)
}
