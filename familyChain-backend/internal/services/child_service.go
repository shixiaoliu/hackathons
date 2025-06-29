package services

import (
	"errors"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/repository"
	"eth-for-babies-backend/internal/utils"
)

type ChildService struct {
	childRepo  *repository.ChildRepository
	familyRepo *repository.FamilyRepository
	taskRepo   *repository.TaskRepository
}

func NewChildService(childRepo *repository.ChildRepository, familyRepo *repository.FamilyRepository, taskRepo *repository.TaskRepository) *ChildService {
	return &ChildService{
		childRepo:  childRepo,
		familyRepo: familyRepo,
		taskRepo:   taskRepo,
	}
}

// CreateChild 创建孩子
func (s *ChildService) CreateChild(child *models.Child) error {
	// 验证孩子数据
	if child.Name == "" {
		return errors.New("child name is required")
	}
	if child.Age <= 0 || child.Age > 18 {
		return errors.New("child age must be between 1 and 18")
	}
	if child.ParentAddress == "" {
		return errors.New("parent address is required")
	}

	// 验证地址格式
	if !utils.IsValidEthereumAddress(child.ParentAddress) {
		return errors.New("invalid parent address format")
	}
	if child.WalletAddress != "" && !utils.IsValidEthereumAddress(child.WalletAddress) {
		return errors.New("invalid child wallet address format")
	}

	// 验证家庭是否存在
	if child.FamilyID != 0 {
		family, err := s.familyRepo.GetByID(child.FamilyID)
		if err != nil {
			return errors.New("family not found")
		}
		if family.ParentAddress != child.ParentAddress {
			return errors.New("child parent address does not match family parent")
		}
	}

	// 清理输入数据
	child.Name = utils.SanitizeString(child.Name)
	if child.Avatar != "" {
		child.Avatar = utils.SanitizeString(child.Avatar)
	}

	// 初始化统计数据
	child.TasksCompleted = 0
	child.TotalRewards = 0.0

	return s.childRepo.Create(child)
}

// GetChildrenByParent 获取家长的所有孩子
func (s *ChildService) GetChildrenByParent(parentAddress string) ([]*models.Child, error) {
	return s.childRepo.GetByParentAddress(parentAddress)
}

// GetChildByID 根据ID获取孩子
func (s *ChildService) GetChildByID(id uint) (*models.Child, error) {
	return s.childRepo.GetByID(id)
}

// GetChildByWalletAddress 根据钱包地址获取孩子
func (s *ChildService) GetChildByWalletAddress(walletAddress string) (*models.Child, error) {
	if !utils.IsValidEthereumAddress(walletAddress) {
		return nil, errors.New("invalid wallet address format")
	}
	return s.childRepo.GetByWalletAddress(walletAddress)
}

// UpdateChild 更新孩子信息
func (s *ChildService) UpdateChild(id uint, updates map[string]interface{}) error {
	child, err := s.childRepo.GetByID(id)
	if err != nil {
		return err
	}

	// 验证更新数据
	if name, exists := updates["name"]; exists {
		if nameStr, ok := name.(string); ok {
			if nameStr == "" {
				return errors.New("child name cannot be empty")
			}
			updates["name"] = utils.SanitizeString(nameStr)
		}
	}

	if age, exists := updates["age"]; exists {
		if ageInt, ok := age.(int); ok {
			if ageInt <= 0 || ageInt > 18 {
				return errors.New("child age must be between 1 and 18")
			}
		}
	}

	if walletAddress, exists := updates["wallet_address"]; exists {
		if addrStr, ok := walletAddress.(string); ok && addrStr != "" {
			if !utils.IsValidEthereumAddress(addrStr) {
				return errors.New("invalid wallet address format")
			}
		}
	}

	if avatar, exists := updates["avatar"]; exists {
		if avatarStr, ok := avatar.(string); ok {
			updates["avatar"] = utils.SanitizeString(avatarStr)
		}
	}

	// 不允许修改某些字段
	delete(updates, "parent_address")
	delete(updates, "family_id")
	delete(updates, "tasks_completed")
	delete(updates, "total_rewards")

	return s.childRepo.Update(id, updates)
}

// DeleteChild 删除孩子
func (s *ChildService) DeleteChild(id uint) error {
	// 检查是否有关联的任务
	tasks, err := s.taskRepo.GetByAssignedChild(id)
	if err != nil {
		return err
	}

	// 检查是否有进行中或已完成但未处理的任务
	for _, task := range tasks {
		if task.Status == "in_progress" || task.Status == "completed" {
			return errors.New("cannot delete child with active or completed tasks")
		}
	}

	return s.childRepo.Delete(id)
}

// GetChildProgress 获取孩子的进度信息
func (s *ChildService) GetChildProgress(id uint) (map[string]interface{}, error) {
	child, err := s.childRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// 获取孩子的任务
	tasks, err := s.taskRepo.GetByAssignedChild(id)
	if err != nil {
		return nil, err
	}

	// 统计任务状态
	taskStats := map[string]int{
		"total":      len(tasks),
		"pending":    0,
		"in_progress": 0,
		"completed":  0,
		"approved":   0,
		"rejected":   0,
	}

	for _, task := range tasks {
		if count, exists := taskStats[task.Status]; exists {
			taskStats[task.Status] = count + 1
		}
	}

	// 计算完成率
	completionRate := 0.0
	if taskStats["total"] > 0 {
		completionRate = float64(taskStats["approved"]) / float64(taskStats["total"]) * 100
	}

	progress := map[string]interface{}{
		"child_info": map[string]interface{}{
			"id":             child.ID,
			"name":           child.Name,
			"age":            child.Age,
			"avatar":         child.Avatar,
			"wallet_address": child.WalletAddress,
		},
		"statistics": map[string]interface{}{
			"tasks_completed": child.TasksCompleted,
			"total_rewards":   child.TotalRewards,
			"completion_rate": completionRate,
		},
		"task_breakdown": taskStats,
		"recent_tasks":   s.getRecentTasks(tasks, 5),
	}

	return progress, nil
}

// getRecentTasks 获取最近的任务
func (s *ChildService) getRecentTasks(tasks []*models.Task, limit int) []map[string]interface{} {
	// 按创建时间排序（最新的在前）
	recentTasks := make([]map[string]interface{}, 0)
	count := 0

	for i := len(tasks) - 1; i >= 0 && count < limit; i-- {
		task := tasks[i]
		recentTasks = append(recentTasks, map[string]interface{}{
			"id":          task.ID,
			"title":       task.Title,
			"status":      task.Status,
			"reward":      task.RewardAmount,
			"difficulty":  task.Difficulty,
			"created_at":  task.CreatedAt,
			"due_date":    task.DueDate,
		})
		count++
	}

	return recentTasks
}

// ValidateChildAccess 验证用户是否有权限访问孩子信息
func (s *ChildService) ValidateChildAccess(childID uint, userAddress string, userRole string) error {
	child, err := s.childRepo.GetByID(childID)
	if err != nil {
		return err
	}

	// 家长可以访问自己的孩子
	if userRole == "parent" && child.ParentAddress == userAddress {
		return nil
	}

	// 孩子可以访问自己的信息
	if userRole == "child" && child.WalletAddress == userAddress {
		return nil
	}

	return errors.New("access denied")
}

// UpdateChildStatistics 更新孩子的统计信息
func (s *ChildService) UpdateChildStatistics(childID uint, tasksCompleted int, rewardAmount float64) error {
	updates := map[string]interface{}{
		"tasks_completed": "tasks_completed + ?",
		"total_rewards":   "total_rewards + ?",
	}
	return s.childRepo.UpdateWithRaw(childID, updates, tasksCompleted, rewardAmount)
}

// GetChildrenByFamily 获取家庭的所有孩子
func (s *ChildService) GetChildrenByFamily(familyID uint) ([]*models.Child, error) {
	return s.childRepo.GetByFamilyID(familyID)
}