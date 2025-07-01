# 修复 RewardRegistry ABI 问题

## 问题说明

我们发现合约兑换功能出现 "failed to exchange reward on chain: method 'exchangeReward' not found" 错误。这是因为后端的 `reward_registry_generated.go` 文件中 `RewardRegistryABI` 常量包含的 ABI 字符串被截断了，导致后端无法识别 `exchangeReward` 方法。

## 解决方案

我已创建了一个修复后的 Go 文件 `RewardRegistryFixed.go`，包含了完整的 ABI 定义和必要的函数实现。

## 修复步骤

1. 打开终端，进入后端项目目录：
   ```bash
   cd ../familyChain-backend
   ```

2. 备份原始文件：
   ```bash
   cp pkg/blockchain/reward_registry_generated.go pkg/blockchain/reward_registry_generated.go.bak
   ```

3. 复制修复后的文件：
   ```bash
   cp ../familyChain-contract/RewardRegistryFixed.go pkg/blockchain/reward_registry_generated.go
   ```

4. 重新启动后端服务：
   ```bash
   go run cmd/server/main.go
   ```

## 验证

完成上述步骤后，再次尝试兑换操作，应该不会再出现 "method 'exchangeReward' not found" 错误。

## 技术细节

1. 问题原因：后端使用的 ABI 定义不完整，被截断了
2. 修复方法：用合约编译后生成的完整 ABI 替换原始的截断 ABI
3. 关键改动：
   - 完整的 RewardRegistryABI 常量定义
   - 正确实现 ExchangeReward 方法
   - 修复其他相关方法

## 注意事项

此修复仅解决 ABI 不完整的问题。如果后端还有其他配置问题（如合约地址不正确），可能需要额外调整。 