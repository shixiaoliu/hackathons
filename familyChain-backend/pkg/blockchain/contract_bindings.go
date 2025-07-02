// Package blockchain provides contract bindings for blockchain interactions
package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"eth-for-babies-backend/pkg/blockchain/contracts/taskregistry"
	"eth-for-babies-backend/pkg/blockchain/contracts/familyregistry"
	"eth-for-babies-backend/pkg/blockchain/contracts/rewardtoken"
	"eth-for-babies-backend/pkg/blockchain/contracts/rewardregistry"
)

// Contract bindings
type (
	TaskRegistry   = taskregistry.TaskRegistry
	FamilyRegistry = familyregistry.FamilyRegistry
	RewardToken    = rewardtoken.RewardToken
	RewardRegistry = rewardregistry.RewardRegistry
)

// NewTaskRegistry creates a new task registry contract instance
func NewTaskRegistry(address common.Address, backend bind.ContractBackend) (*TaskRegistry, error) {
	return taskregistry.NewTaskRegistry(address, backend)
}

// NewFamilyRegistry creates a new family registry contract instance
func NewFamilyRegistry(address common.Address, backend bind.ContractBackend) (*FamilyRegistry, error) {
	return familyregistry.NewFamilyRegistry(address, backend)
}

// NewRewardToken creates a new reward token contract instance
func NewRewardToken(address common.Address, backend bind.ContractBackend) (*RewardToken, error) {
	return rewardtoken.NewRewardToken(address, backend)
}

// NewRewardRegistry creates a new reward registry contract instance
func NewRewardRegistry(address common.Address, backend bind.ContractBackend) (*RewardRegistry, error) {
	return rewardregistry.NewRewardRegistry(address, backend)
}

// DeployTaskRegistry deploys a new task registry contract
func DeployTaskRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *taskregistry.TaskRegistry, error) {
	return taskregistry.DeployTaskRegistry(auth, backend)
}

// DeployFamilyRegistry deploys a new family registry contract
func DeployFamilyRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *familyregistry.FamilyRegistry, error) {
	return familyregistry.DeployFamilyRegistry(auth, backend)
}

// DeployRewardToken deploys a new reward token contract
func DeployRewardToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *rewardtoken.RewardToken, error) {
	return rewardtoken.DeployRewardToken(auth, backend, name, symbol)
}

// DeployRewardRegistry deploys a new reward registry contract
func DeployRewardRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address) (common.Address, *types.Transaction, *rewardregistry.RewardRegistry, error) {
	return rewardregistry.DeployRewardRegistry(auth, backend, tokenAddress)
}
