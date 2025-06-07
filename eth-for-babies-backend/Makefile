# Family Task Chain Backend Makefile

# 变量定义
APP_NAME=family-task-chain-backend
MAIN_PATH=cmd/server/main.go
BUILD_DIR=build
DATA_DIR=data

# Go相关变量
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# 默认目标
.PHONY: all
all: clean deps build

# 安装依赖
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# 构建应用
.PHONY: build
build:
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

# 运行应用
.PHONY: run
run:
	@echo "Running application..."
	@mkdir -p $(DATA_DIR)
	$(GOCMD) run $(MAIN_PATH)

# 开发模式运行（带热重载）
.PHONY: dev
dev:
	@echo "Running in development mode..."
	@mkdir -p $(DATA_DIR)
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "Air not found. Installing..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

# 运行测试
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# 运行测试并生成覆盖率报告
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# 代码格式化
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

# 代码检查
.PHONY: vet
vet:
	@echo "Running go vet..."
	$(GOCMD) vet ./...

# 代码质量检查（需要安装golangci-lint）
.PHONY: lint
lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Please install it first."; \
		echo "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2"; \
	fi

# 清理构建文件
.PHONY: clean
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

# 创建必要的目录
.PHONY: setup
setup:
	@echo "Setting up project..."
	@mkdir -p $(DATA_DIR)
	@mkdir -p $(BUILD_DIR)
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo "Created .env file from .env.example"; \
		echo "Please edit .env file with your configuration"; \
	fi

# 数据库相关
.PHONY: db-reset
db-reset:
	@echo "Resetting database..."
	@rm -f $(DATA_DIR)/*.db
	@echo "Database reset complete"

# 生产构建
.PHONY: build-prod
build-prod:
	@echo "Building for production..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)

# Docker相关
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME) .

.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env $(APP_NAME)

# 健康检查
.PHONY: health
health:
	@echo "Checking application health..."
	@curl -f http://localhost:8080/health || echo "Application is not running or unhealthy"

# 安装开发工具
.PHONY: install-tools
install-tools:
	@echo "Installing development tools..."
	$(GOCMD) install github.com/cosmtrek/air@latest
	@echo "Installing golangci-lint..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.54.2

# 显示帮助信息
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make deps          - Install dependencies"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make dev           - Run in development mode with hot reload"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make fmt           - Format code"
	@echo "  make vet           - Run go vet"
	@echo "  make lint          - Run linter"
	@echo "  make clean         - Clean build files"
	@echo "  make setup         - Setup project (create directories and .env)"
	@echo "  make db-reset      - Reset database"
	@echo "  make build-prod    - Build for production"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make health        - Check application health"
	@echo "  make install-tools - Install development tools"
	@echo "  make help          - Show this help message"