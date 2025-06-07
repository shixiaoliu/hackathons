package services

import (
	"errors"

	"eth-for-babies-backend/internal/models"
	"eth-for-babies-backend/internal/repository"
	"eth-for-babies-backend/internal/utils"
)

type FamilyService struct {
	familyRepo *repository.FamilyRepository
	childRepo  *repository.ChildRepository
}

func NewFamilyService(familyRepo *repository.FamilyRepository, childRepo *repository.ChildRepository) *FamilyService {
	return &FamilyService{
		familyRepo: familyRepo,
		childRepo:  childRepo,
	}
}

// CreateFamily 创建家庭
func (s *FamilyService) CreateFamily(family *models.Family) error {
	// 验证家庭数据
	if family.Name == "" {
		return errors.New("family name is required")
	}

	// 验证家长地址格式
	if !utils.IsValidEthereumAddress(family.ParentAddress) {
		return errors.New("invalid parent address format")
	}

	// 检查家长是否已经有家庭
	existingFamily, err := s.familyRepo.GetByParentAddress(family.ParentAddress)
	if err == nil && existingFamily != nil {
		return errors.New("parent already has a family")
	}

	// 清理输入数据
	family.Name = utils.SanitizeString(family.Name)

	return s.familyRepo.Create(family)
}

// GetFamiliesByParent 获取家长的家庭
func (s *FamilyService) GetFamiliesByParent(parentAddress string) ([]*models.Family, error) {
	family, err := s.familyRepo.GetByParentAddress(parentAddress)
	if err != nil {
		return nil, err
	}
	if family == nil {
		return []*models.Family{}, nil
	}
	return []*models.Family{family}, nil
}

// GetFamilyByID 根据ID获取家庭
func (s *FamilyService) GetFamilyByID(id uint) (*models.Family, error) {
	return s.familyRepo.GetByID(id)
}

// UpdateFamily 更新家庭信息
func (s *FamilyService) UpdateFamily(id uint, updates map[string]interface{}) error {
	family, err := s.familyRepo.GetByID(id)
	if err != nil {
		return err
	}

	// 清理字符串输入
	if name, exists := updates["name"]; exists {
		if nameStr, ok := name.(string); ok {
			if nameStr == "" {
				return errors.New("family name cannot be empty")
			}
			updates["name"] = utils.SanitizeString(nameStr)
		}
	}

	// 不允许修改家长地址
	delete(updates, "parent_address")

	return s.familyRepo.Update(id, updates)
}

// DeleteFamily 删除家庭
func (s *FamilyService) DeleteFamily(id uint) error {
	// 检查是否有关联的孩子
	children, err := s.childRepo.GetByFamilyID(id)
	if err != nil {
		return err
	}
	if len(children) > 0 {
		return errors.New("cannot delete family with existing children")
	}

	return s.familyRepo.Delete(id)
}

// AddFamilyMember 添加家庭成员（孩子）
func (s *FamilyService) AddFamilyMember(familyID uint, child *models.Child) error {
	// 验证家庭是否存在
	family, err := s.familyRepo.GetByID(familyID)
	if err != nil {
		return err
	}

	// 设置孩子的家庭ID和家长地址
	child.FamilyID = familyID
	child.ParentAddress = family.ParentAddress

	// 验证孩子数据
	if child.Name == "" {
		return errors.New("child name is required")
	}
	if child.WalletAddress != "" && !utils.IsValidEthereumAddress(child.WalletAddress) {
		return errors.New("invalid child wallet address format")
	}
	if child.Age <= 0 || child.Age > 18 {
		return errors.New("child age must be between 1 and 18")
	}

	// 清理输入数据
	child.Name = utils.SanitizeString(child.Name)
	if child.Avatar != "" {
		child.Avatar = utils.SanitizeString(child.Avatar)
	}

	return s.childRepo.Create(child)
}

// GetFamilyMembers 获取家庭成员
func (s *FamilyService) GetFamilyMembers(familyID uint) ([]*models.Child, error) {
	// 验证家庭是否存在
	_, err := s.familyRepo.GetByID(familyID)
	if err != nil {
		return nil, err
	}

	return s.childRepo.GetByFamilyID(familyID)
}

// GetFamilyStatistics 获取家庭统计信息
func (s *FamilyService) GetFamilyStatistics(familyID uint) (map[string]interface{}, error) {
	// 验证家庭是否存在
	family, err := s.familyRepo.GetByID(familyID)
	if err != nil {
		return nil, err
	}

	// 获取家庭成员
	children, err := s.childRepo.GetByFamilyID(familyID)
	if err != nil {
		return nil, err
	}

	// 计算统计信息
	totalChildren := len(children)
	totalTasksCompleted := 0
	totalRewards := 0.0

	for _, child := range children {
		totalTasksCompleted += child.TasksCompleted
		totalRewards += child.TotalRewards
	}

	stats := map[string]interface{}{
		"family_name":           family.Name,
		"total_children":        totalChildren,
		"total_tasks_completed": totalTasksCompleted,
		"total_rewards":         totalRewards,
		"created_at":            family.CreatedAt,
	}

	return stats, nil
}

// ValidateFamilyAccess 验证用户是否有权限访问家庭
func (s *FamilyService) ValidateFamilyAccess(familyID uint, userAddress string, userRole string) error {
	family, err := s.familyRepo.GetByID(familyID)
	if err != nil {
		return err
	}

	// 家长可以访问自己的家庭
	if userRole == "parent" && family.ParentAddress == userAddress {
		return nil
	}

	// 孩子可以访问自己所属的家庭
	if userRole == "child" {
		children, err := s.childRepo.GetByFamilyID(familyID)
		if err != nil {
			return err
		}
		for _, child := range children {
			if child.WalletAddress == userAddress {
				return nil
			}
		}
	}

	return errors.New("access denied")
}