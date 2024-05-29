# 定义变量
APP_NAME := hebbian-net
CMD_DIR := cmd
MAIN_FILE := $(CMD_DIR)/main.go
BUILD_DIR := bin
BUILD_FILE := $(BUILD_DIR)/$(APP_NAME)

# 默认目标
.PHONY: all
all: build

# 编译目标
.PHONY: build
build:
	@echo "==> Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_FILE) $(MAIN_FILE)
	@echo "==> Build complete: $(BUILD_FILE)"

# 运行目标
.PHONY: run
run: build
	@echo "==> Running $(APP_NAME)..."
	@$(BUILD_FILE)

# 测试目标
.PHONY: test
test:
	@echo "==> Running tests..."
	@go test ./...

# 清理目标
.PHONY: clean
clean:
	@echo "==> Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "==> Clean complete."

# 帮助信息
.PHONY: help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo ""
	@echo "Usage:"
	@echo "  make build    - Build the application"
	@echo "  make run      - Build and run the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean the build directory"
	@echo "  make help     - Display this help message"
