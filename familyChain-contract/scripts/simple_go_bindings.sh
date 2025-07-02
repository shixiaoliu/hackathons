#!/bin/bash

# 设置目录路径
CONTRACT_DIR="$(pwd)"
BACKEND_DIR="../familyChain-backend"
ABI_DIR="${CONTRACT_DIR}/abi"
OUTPUT_DIR="${BACKEND_DIR}/pkg/blockchain"

# 确保输出目录存在
mkdir -p ${OUTPUT_DIR}

# 检查是否安装了abigen
if ! command -v abigen &> /dev/null; then
    echo "错误: abigen 工具未安装"
    echo "请先安装: go install github.com/ethereum/go-ethereum/cmd/abigen@latest"
    exit 1
fi

# 为每个合约生成绑定
echo "生成 TaskRegistry 绑定..."
abigen --abi=${ABI_DIR}/TaskRegistry.json --pkg=blockchain --out=${OUTPUT_DIR}/task_registry.go --type=TaskRegistry

echo "生成 FamilyRegistry 绑定..."
abigen --abi=${ABI_DIR}/FamilyRegistry.json --pkg=blockchain --out=${OUTPUT_DIR}/family_registry.go --type=FamilyRegistry

echo "生成 RewardToken 绑定..."
abigen --abi=${ABI_DIR}/RewardToken.json --pkg=blockchain --out=${OUTPUT_DIR}/reward_token.go --type=RewardToken

echo "生成 RewardRegistry 绑定..."
abigen --abi=${ABI_DIR}/RewardRegistry.json --pkg=blockchain --out=${OUTPUT_DIR}/reward_registry.go --type=RewardRegistry

echo "ABI 绑定生成完成！"
