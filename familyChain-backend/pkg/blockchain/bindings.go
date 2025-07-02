package blockchain

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"eth-for-babies-backend/pkg/blockchain/contracts/taskregistry"
	"eth-for-babies-backend/pkg/blockchain/contracts/familyregistry"
	"eth-for-babies-backend/pkg/blockchain/contracts/rewardtoken"
	"eth-for-babies-backend/pkg/blockchain/contracts/rewardregistry"
)

// 类型别名，使其与原有代码兼容
type TaskRegistry = taskregistry.TaskRegistry
type FamilyRegistry = familyregistry.FamilyRegistry
type RewardToken = rewardtoken.RewardToken
type RewardRegistry = rewardregistry.RewardRegistry

// 合约实例化函数
func NewTaskRegistry(address common.Address, backend bind.ContractBackend) (*TaskRegistry, error) {
	return taskregistry.NewTaskRegistry(address, backend)
}

func NewFamilyRegistry(address common.Address, backend bind.ContractBackend) (*FamilyRegistry, error) {
	return familyregistry.NewFamilyRegistry(address, backend)
}

func NewRewardToken(address common.Address, backend bind.ContractBackend) (*RewardToken, error) {
	return rewardtoken.NewRewardToken(address, backend)
}

func NewRewardRegistry(address common.Address, backend bind.ContractBackend) (*RewardRegistry, error) {
	return rewardregistry.NewRewardRegistry(address, backend)
}

// 注意：由于我们没有生成 Deploy 函数，所以我们这里使用占位符函数
// 在实际使用时，你可能需要手动编写部署逻辑或重新生成带有 bin 参数的绑定
