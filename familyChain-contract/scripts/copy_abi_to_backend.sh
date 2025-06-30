#!/bin/bash

# 创建目标目录
mkdir -p ../familyChain-backend/contracts/abi

# 复制ABI文件
cp ./abi/FamilyRegistry.json ../familyChain-backend/contracts/abi/
cp ./abi/RewardToken.json ../familyChain-backend/contracts/abi/
cp ./abi/TaskRegistry.json ../familyChain-backend/contracts/abi/

echo "已成功将ABI文件复制到后端项目" 