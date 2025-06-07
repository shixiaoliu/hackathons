package repository

import (
	"eth-for-babies-backend/internal/models"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// Create 创建任务
func (r *TaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

// GetByID 根据ID获取任务
func (r *TaskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// GetByCreator 根据创建者获取任务列表
func (r *TaskRepository) GetByCreator(creatorAddress string) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").Where("created_by = ?", creatorAddress).Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// GetByAssignedChild 根据分配的孩子获取任务列表
func (r *TaskRepository) GetByAssignedChild(childID uint) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").Where("assigned_child_id = ?", childID).Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// GetByStatus 根据状态获取任务列表
func (r *TaskRepository) GetByStatus(status string) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").Where("status = ?", status).Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// GetByCreatorAndStatus 根据创建者和状态获取任务列表
func (r *TaskRepository) GetByCreatorAndStatus(creatorAddress, status string) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").Where("created_by = ? AND status = ?", creatorAddress, status).Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// GetByChildAndStatus 根据孩子和状态获取任务列表
func (r *TaskRepository) GetByChildAndStatus(childID uint, status string) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").Where("assigned_child_id = ? AND status = ?", childID, status).Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// Update 更新任务
func (r *TaskRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.Task{}).Where("id = ?", id).Updates(updates).Error
}

// Delete 删除任务
func (r *TaskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}

// List 获取任务列表
func (r *TaskRepository) List(limit, offset int) ([]*models.Task, error) {
	var tasks []*models.Task
	query := r.db.Preload("Creator").Preload("AssignedChild").Order("created_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	err := query.Find(&tasks).Error
	return tasks, err
}

// GetPendingTasks 获取待分配的任务
func (r *TaskRepository) GetPendingTasks(creatorAddress string) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Where("created_by = ? AND status = ? AND assigned_child_id IS NULL", creatorAddress, "pending").Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// GetActiveTasks 获取进行中的任务
func (r *TaskRepository) GetActiveTasks(creatorAddress string) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").Where("created_by = ? AND status = ?", creatorAddress, "in_progress").Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// GetCompletedTasks 获取已完成待审核的任务
func (r *TaskRepository) GetCompletedTasks(creatorAddress string) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").Where("created_by = ? AND status = ?", creatorAddress, "completed").Order("completed_at DESC").Find(&tasks).Error
	return tasks, err
}

// GetTasksByDifficulty 根据难度获取任务列表
func (r *TaskRepository) GetTasksByDifficulty(difficulty string) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").Where("difficulty = ?", difficulty).Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// Count 获取任务总数
func (r *TaskRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&models.Task{}).Count(&count).Error
	return count, err
}

// CountByCreator 根据创建者获取任务数量
func (r *TaskRepository) CountByCreator(creatorAddress string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Task{}).Where("created_by = ?", creatorAddress).Count(&count).Error
	return count, err
}

// CountByStatus 根据状态获取任务数量
func (r *TaskRepository) CountByStatus(status string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Task{}).Where("status = ?", status).Count(&count).Error
	return count, err
}

// CountByChild 根据孩子获取任务数量
func (r *TaskRepository) CountByChild(childID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Task{}).Where("assigned_child_id = ?", childID).Count(&count).Error
	return count, err
}

// GetTaskStatistics 获取任务统计信息
func (r *TaskRepository) GetTaskStatistics(creatorAddress string) (map[string]interface{}, error) {
	// 统计各状态任务数量
	var statusStats []struct {
		Status string
		Count  int64
	}
	r.db.Model(&models.Task{}).Select("status, COUNT(*) as count").Where("created_by = ?", creatorAddress).Group("status").Scan(&statusStats)

	// 转换为map
	statusMap := make(map[string]int64)
	for _, stat := range statusStats {
		statusMap[stat.Status] = stat.Count
	}

	// 统计难度分布
	var difficultyStats []struct {
		Difficulty string
		Count      int64
	}
	r.db.Model(&models.Task{}).Select("difficulty, COUNT(*) as count").Where("created_by = ?", creatorAddress).Group("difficulty").Scan(&difficultyStats)

	// 转换为map
	difficultyMap := make(map[string]int64)
	for _, stat := range difficultyStats {
		difficultyMap[stat.Difficulty] = stat.Count
	}

	// 计算总奖励
	var totalRewards float64
	r.db.Model(&models.Task{}).Where("created_by = ? AND status = ?", creatorAddress, "approved").Select("COALESCE(SUM(reward_amount), 0)").Scan(&totalRewards)

	stats := map[string]interface{}{
		"creator_address":   creatorAddress,
		"status_breakdown":  statusMap,
		"difficulty_breakdown": difficultyMap,
		"total_rewards":     totalRewards,
	}

	return stats, nil
}

// GetOverdueTasks 获取过期任务
func (r *TaskRepository) GetOverdueTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Preload("Creator").Preload("AssignedChild").Where("due_date < NOW() AND status IN ?", []string{"pending", "in_progress"}).Find(&tasks).Error
	return tasks, err
}

// WithTransaction 在事务中执行操作
func (r *TaskRepository) WithTransaction(fn func(*TaskRepository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		txRepo := &TaskRepository{db: tx}
		return fn(txRepo)
	})
}