#!/bin/bash

# 设置颜色
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[0;33m'
NC='\033[0m' # 无颜色

# 检查 Go 是否已安装
if ! command -v go &> /dev/null; then
    echo -e "${RED}错误: 未找到 Go 语言环境!${NC}"
    echo "请先安装 Go 语言环境，可以从 https://golang.org/dl/ 下载"
    exit 1
fi

echo -e "${YELLOW}检查 Go 版本...${NC}"
go version

echo -e "${YELLOW}安装 go-ethereum...${NC}"
go get -u github.com/ethereum/go-ethereum

echo -e "${YELLOW}安装 abigen 工具...${NC}"
go install github.com/ethereum/go-ethereum/cmd/abigen@latest

# 检查是否安装成功
if command -v abigen &> /dev/null; then
    echo -e "${GREEN}✅ abigen 安装成功!${NC}"
    echo -e "${YELLOW}abigen 版本信息:${NC}"
    abigen --version
else
    echo -e "${RED}❌ abigen 安装失败!${NC}"
    echo -e "${YELLOW}可能需要添加 GOPATH 到环境变量 PATH 中:${NC}"
    echo -e "export PATH=\$PATH:\$GOPATH/bin"
    
    # 尝试自动添加到环境变量并执行
    GOPATH=$(go env GOPATH)
    echo -e "${YELLOW}自动添加 GOPATH 到当前会话的 PATH...${NC}"
    export PATH=$PATH:$GOPATH/bin
    
    if command -v abigen &> /dev/null; then
        echo -e "${GREEN}✅ 现在 abigen 可用了!${NC}"
        echo -e "${YELLOW}abigen 版本信息:${NC}"
        abigen --version
    else
        echo -e "${RED}❌ 仍然找不到 abigen，请手动将 GOPATH 添加到 PATH 环境变量中${NC}"
        exit 1
    fi
fi 