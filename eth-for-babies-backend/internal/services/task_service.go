package services

import (
	"errors"
	"time"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/repository"
)

type TaskService struct {
	taskRepo   *repository.TaskRepository
	childRepo  *repository.ChildRepository
	familyRepo *repository.FamilyRepository
}

func NewTaskService(taskRepo *repository.TaskRepository, childRepo *repository.ChildRepository, familyRepo *repository.FamilyRepository) *TaskService {
	return &TaskService{
		taskRepo:   taskRepo,
		childRepo:  childRepo,
		familyRepo: familyRepo,
	}
}

// CreateTask 创建任务
func (s *TaskService) CreateTask(task *models.Task) error {
	// 验证任务数据
	if task.Title == "" {
		return errors.New("task title is required")
	}
	if task.RewardAmount <= 0 {
		return errors.New("reward amount must be positive")
	}

	// 如果指定了孩子，验证孩子是否存在且属于创建者
	if task.AssignedChildID != nil {
		child, err := s.childRepo.GetByID(*task.AssignedChildID)
		if err != nil {
			return errors.New("assigned child not found")
		}
		if child.ParentAddress != task.CreatedBy {
			return errors.New("cannot assign task to child not belonging to you")
		}
		task.Status = "in_progress"
	} else {
		task.Status = "pending"
	}

	return s.taskRepo.Create(task)
}

// GetTasksByParent 获取家长创建的任务
func (s *TaskService) GetTasksByParent(parentAddress string) ([]*models.Task, error) {
	return s.taskRepo.GetByCreator(parentAddress)
}

// GetTasksByChild 获取分配给孩子的任务
func (s *TaskService) GetTasksByChild(childID uint) ([]*models.Task, error) {
	return s.taskRepo.GetByAssignedChild(childID)
}

// GetTaskByID 根据ID获取任务
func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.taskRepo.GetByID(id)
}

// UpdateTask 更新任务
func (s *TaskService) UpdateTask(id uint, updates map[string]interface{}) error {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	// 如果任务已完成或已批准，不允许修改
	if task.Status == "completed" || task.Status == "approved" {
		return errors.New("cannot update completed or approved task")
	}

	// 如果更新了分配的孩子，需要验证
	if assignedChildID, exists := updates["assigned_child_id"]; exists {
		if assignedChildID != nil {
			childID := assignedChildID.(uint)
			child, err := s.childRepo.GetByID(childID)
			if err != nil {
				return errors.New("assigned child not found")
			}
			if child.ParentAddress != task.CreatedBy {
				return errors.New("cannot assign task to child not belonging to you")
			}
			updates["status"] = "in_progress"
		} else {
			updates["status"] = "pending"
		}
	}

	return s.taskRepo.Update(id, updates)
}

// CompleteTask 完成任务
func (s *TaskService) CompleteTask(id uint, proof string) error {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	if task.Status != "in_progress" {
		return errors.New("task is not in progress")
	}

	updates := map[string]interface{}{
		"status":          "completed",
		"completion_proof": proof,
		"completed_at":     time.Now(),
	}

	return s.taskRepo.Update(id, updates)
}

// ApproveTask 批准任务
func (s *TaskService) ApproveTask(id uint) error {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	if task.Status != "completed" {
		return errors.New("task is not completed")
	}

	// 开始事务
	err = s.taskRepo.WithTransaction(func(repo *repository.TaskRepository) error {
		// 更新任务状态
		updates := map[string]interface{}{
			"status":      "approved",
			"approved_at": time.Now(),
		}
		if err := repo.Update(id, updates); err != nil {
			return err
		}

		// 更新孩子的统计信息
		if task.AssignedChildID != nil {
			childUpdates := map[string]interface{}{
				"tasks_completed": "tasks_completed + 1",
				"total_rewards":   "total_rewards + ?",
			}
			return s.childRepo.UpdateWithRaw(*task.AssignedChildID, childUpdates, task.RewardAmount)
		}

		return nil
	})

	return err
}

// RejectTask 拒绝任务
func (s *TaskService) RejectTask(id uint, reason string) error {
	task, err := s.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	if task.Status != "completed" {
		return errors.New("task is not completed")
	}

	updates := map[string]interface{}{
		"status":          "rejected",
		"completion_proof": reason, // 使用completion_proof字段存储拒绝原因
		"rejected_at":      time.Now(),
	}

	return s.taskRepo.Update(id, updates)
}

// GetTaskStatistics 获取任务统计信息
func (s *TaskService) GetTaskStatistics(parentAddress string) (map[string]interface{}, error) {
	tasks, err := s.taskRepo.GetByCreator(parentAddress)
	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"total":     len(tasks),
		"pending":   0,
		"in_progress": 0,
		"completed": 0,
		"approved":  0,
		"rejected":  0,
	}

	for _, task := range tasks {
		if count, exists := stats[task.Status]; exists {
			stats[task.Status] = count.(int) + 1
		}
	}

	return stats, nil
}