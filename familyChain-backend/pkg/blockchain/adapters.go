package blockchain

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DeployTaskRegistry 是一个占位符部署函数，实际应用中需要替换为真实的部署逻辑
func DeployTaskRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TaskRegistry, error) {
	// 这里应该是真实的部署逻辑
	return common.Address{}, nil, nil, fmt.Errorf("DeployTaskRegistry not implemented")
}

// DeployFamilyRegistry 是一个占位符部署函数，实际应用中需要替换为真实的部署逻辑
func DeployFamilyRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FamilyRegistry, error) {
	// 这里应该是真实的部署逻辑
	return common.Address{}, nil, nil, fmt.Errorf("DeployFamilyRegistry not implemented")
}

// DeployRewardToken 是一个占位符部署函数，实际应用中需要替换为真实的部署逻辑
func DeployRewardToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *RewardToken, error) {
	// 这里应该是真实的部署逻辑
	return common.Address{}, nil, nil, fmt.Errorf("DeployRewardToken not implemented")
}

// DeployRewardRegistry 是一个占位符部署函数，实际应用中需要替换为真实的部署逻辑
func DeployRewardRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address) (common.Address, *types.Transaction, *RewardRegistry, error) {
	// 这里应该是真实的部署逻辑
	return common.Address{}, nil, nil, fmt.Errorf("DeployRewardRegistry not implemented")
}

// 更新后的 GetTask 函数，适配新的合约调用方式
func (cm *ContractManager) GetTask(taskID uint64) (Task, error) {
	var task Task

	if cm.TaskRegistry == nil {
		return task, fmt.Errorf("task registry not initialized")
	}

	opts := &bind.CallOpts{Context: context.Background()}
	
	// 直接从 Tasks 映射获取任务
	taskData, err := cm.TaskRegistry.Tasks(opts, big.NewInt(int64(taskID)))
	if err != nil {
		return task, fmt.Errorf("failed to get task: %v", err)
	}

	// 将 taskregistry 包中的结构转换为我们自己的 Task 结构
	task = Task{
		ID:          taskData.Id.Uint64(),
		Creator:     taskData.Creator,
		AssignedTo:  taskData.AssignedTo,
		Title:       taskData.Title,
		Description: taskData.Description,
		Reward:      taskData.Reward,
		Completed:   taskData.Completed,
		Approved:    taskData.Approved,
	}

	return task, nil
}
