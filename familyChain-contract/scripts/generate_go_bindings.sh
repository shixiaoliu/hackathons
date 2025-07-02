#!/bin/bash

# 设置颜色
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
NC='\033[0m' # 无颜色

# 设置目录路径
CONTRACT_DIR="$(pwd)"
BACKEND_DIR="../familyChain-backend"
ABI_DIR="${CONTRACT_DIR}/abi"
OUTPUT_DIR="${BACKEND_DIR}/pkg/blockchain/contracts"

# 确保输出目录存在
mkdir -p ${OUTPUT_DIR}

echo -e "${GREEN}开始生成 Go 语言绑定文件...${NC}"

# 检查 abigen 是否已安装
if ! command -v abigen &> /dev/null; then
    echo -e "${RED}错误: abigen 工具未安装${NC}"
    echo -e "请先安装 abigen 工具:"
    echo -e "${YELLOW}go install github.com/ethereum/go-ethereum/cmd/abigen@latest${NC}"
    exit 1
fi

# 从 ABI 生成 Go 绑定函数
generate_binding() {
    local contract_name=$1
    local package_name=${2:-$(echo $1 | tr '[:upper:]' '[:lower:]')}
    local type_name=$3  # 合约的类型前缀
    local abi_file="${ABI_DIR}/${contract_name}.json"
    local output_dir="${OUTPUT_DIR}/${package_name}"
    local output_file="${output_dir}/${package_name}.go"
    
    # 确保输出目录存在
    mkdir -p ${output_dir}
    
    if [ -f "${abi_file}" ]; then
        echo -e "${YELLOW}正在生成 ${contract_name} 的 Go 绑定...${NC}"
        
        # 使用包名作为包，合约名称作为类型前缀
        abigen --abi=${abi_file} --pkg=${package_name} --type=${type_name} --out=${output_file}
        
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}✅ 成功生成 ${contract_name} 的 Go 绑定: ${output_file}${NC}"
        else
            echo -e "${RED}❌ 生成 ${contract_name} 的 Go 绑定失败${NC}"
            return 1
        fi
    else
        echo -e "${RED}❌ ABI 文件不存在: ${abi_file}${NC}"
        return 1
    fi
}

# 生成所有合约的绑定，为每个合约指定唯一的包名
generate_binding "TaskRegistry" "taskregistry" "TaskRegistry" || echo -e "${RED}TaskRegistry 绑定生成失败${NC}"
generate_binding "FamilyRegistry" "familyregistry" "FamilyRegistry" || echo -e "${RED}FamilyRegistry 绑定生成失败${NC}"
generate_binding "RewardToken" "rewardtoken" "RewardToken" || echo -e "${RED}RewardToken 绑定生成失败${NC}"
generate_binding "RewardRegistry" "rewardregistry" "RewardRegistry" || echo -e "${RED}RewardRegistry 绑定生成失败${NC}"

echo -e "${GREEN}Go 语言绑定文件生成完成!${NC}"
echo "文件保存在: ${OUTPUT_DIR}"

# 生成包装器文件
echo -e "${YELLOW}生成包装器文件...${NC}"
cat > ${BACKEND_DIR}/pkg/blockchain/contract_bindings.go << 'EOF2'
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
EOF2

echo -e "${GREEN}✅ 成功生成包装器文件${NC}"
