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
OUTPUT_DIR="${BACKEND_DIR}/pkg/blockchain"

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
    local abi_file="${ABI_DIR}/${contract_name}.json"
    local output_file="${OUTPUT_DIR}/${package_name}_generated.go"
    
    if [ -f "${abi_file}" ]; then
        echo -e "${YELLOW}正在生成 ${contract_name} 的 Go 绑定...${NC}"
        abigen --abi=${abi_file} --pkg=blockchain --out=${output_file}
        
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

# 生成所有合约的绑定
generate_binding "TaskRegistry" "task_registry" || echo -e "${RED}TaskRegistry 绑定生成失败${NC}"
generate_binding "FamilyRegistry" "family_registry" || echo -e "${RED}FamilyRegistry 绑定生成失败${NC}"
generate_binding "RewardToken" "reward_token" || echo -e "${RED}RewardToken 绑定生成失败${NC}"
generate_binding "RewardRegistry" "reward_registry" || echo -e "${RED}RewardRegistry 绑定生成失败${NC}"

echo -e "${GREEN}Go 语言绑定文件生成完成!${NC}"
echo "文件保存在: ${OUTPUT_DIR}" 