package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ContractManager handles interactions with smart contracts
type ContractManager struct {
	client           *EthClient
	privateKey       *ecdsa.PrivateKey
	chainID          *big.Int
	TaskRegistry     *TaskRegistry
	FamilyRegistry   *FamilyRegistry
	RewardToken      *RewardToken
	RewardRegistry   *RewardRegistry
	taskAddress      common.Address
	familyAddress    common.Address
	tokenAddress     common.Address
	rewardRegAddress common.Address
}

// NewContractManager creates a new contract manager instance
func NewContractManager(client *EthClient, contractAddresses map[string]string) (*ContractManager, error) {
	cm := &ContractManager{
		client:  client,
		chainID: client.GetChainID(),
	}

	// 如果提供了合约地址，就使用它们
	if len(contractAddresses) > 0 {
		// 设置任务合约地址
		if addr, ok := contractAddresses["task"]; ok {
			cm.taskAddress = common.HexToAddress(addr)
		}
		// 设置家庭合约地址
		if addr, ok := contractAddresses["family"]; ok {
			cm.familyAddress = common.HexToAddress(addr)
		}
		// 设置代币合约地址
		if addr, ok := contractAddresses["token"]; ok {
			cm.tokenAddress = common.HexToAddress(addr)
		}
		// 设置奖励合约地址
		if addr, ok := contractAddresses["reward"]; ok {
			cm.rewardRegAddress = common.HexToAddress(addr)
		}
	}

	// 初始化合约实例
	err := cm.initializeContracts()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize contracts: %v", err)
	}

	return cm, nil
}

// initializeContracts initializes contract instances based on addresses
func (cm *ContractManager) initializeContracts() error {
	// 如果有任务合约地址，则初始化任务合约
	if cm.taskAddress != (common.Address{}) {
		taskRegistry, err := NewTaskRegistry(cm.taskAddress, cm.client.GetClient())
		if err != nil {
			return fmt.Errorf("failed to instantiate task registry contract: %v", err)
		}
		cm.TaskRegistry = taskRegistry
		log.Printf("TaskRegistry initialized at address: %s", cm.taskAddress.Hex())
	}

	// 如果有家庭合约地址，则初始化家庭合约
	if cm.familyAddress != (common.Address{}) {
		familyRegistry, err := NewFamilyRegistry(cm.familyAddress, cm.client.GetClient())
		if err != nil {
			return fmt.Errorf("failed to instantiate family registry contract: %v", err)
		}
		cm.FamilyRegistry = familyRegistry
		log.Printf("FamilyRegistry initialized at address: %s", cm.familyAddress.Hex())
	}

	// 如果有代币合约地址，则初始化代币合约
	if cm.tokenAddress != (common.Address{}) {
		rewardToken, err := NewRewardToken(cm.tokenAddress, cm.client.GetClient())
		if err != nil {
			return fmt.Errorf("failed to instantiate reward token contract: %v", err)
		}
		cm.RewardToken = rewardToken
		log.Printf("RewardToken initialized at address: %s", cm.tokenAddress.Hex())
	}

	// 如果有奖励合约地址，则初始化奖励合约
	if cm.rewardRegAddress != (common.Address{}) {
		rewardRegistry, err := NewRewardRegistry(cm.rewardRegAddress, cm.client.GetClient())
		if err != nil {
			return fmt.Errorf("failed to instantiate reward registry contract: %v", err)
		}
		cm.RewardRegistry = rewardRegistry
		log.Printf("RewardRegistry initialized at address: %s", cm.rewardRegAddress.Hex())
	}

	return nil
}

// DeployContracts deploys all contracts
func (cm *ContractManager) DeployContracts() error {
	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	// 部署任务合约
	taskAddress, tx, taskInstance, err := DeployTaskRegistry(auth, cm.client.GetClient())
	if err != nil {
		return fmt.Errorf("failed to deploy task registry: %v", err)
	}
	cm.taskAddress = taskAddress
	cm.TaskRegistry = taskInstance

	log.Printf("TaskRegistry deployed with transaction: %s", tx.Hash().Hex())
	log.Printf("TaskRegistry address: %s", taskAddress.Hex())

	receipt, err := cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for TaskRegistry deploy: %v", err)
	}
	log.Printf("TaskRegistry deployed at block: %d", receipt.BlockNumber.Uint64())

	// 部署家庭合约
	familyAddress, tx, familyInstance, err := DeployFamilyRegistry(auth, cm.client.GetClient())
	if err != nil {
		return fmt.Errorf("failed to deploy family registry: %v", err)
	}
	cm.familyAddress = familyAddress
	cm.FamilyRegistry = familyInstance

	log.Printf("FamilyRegistry deployed with transaction: %s", tx.Hash().Hex())
	log.Printf("FamilyRegistry address: %s", familyAddress.Hex())

	receipt, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for FamilyRegistry deploy: %v", err)
	}
	log.Printf("FamilyRegistry deployed at block: %d", receipt.BlockNumber.Uint64())

	// 部署代币合约
	tokenAddress, tx, tokenInstance, err := DeployRewardToken(auth, cm.client.GetClient(), "TaskReward", "TRW")
	if err != nil {
		return fmt.Errorf("failed to deploy reward token: %v", err)
	}
	cm.tokenAddress = tokenAddress
	cm.RewardToken = tokenInstance

	log.Printf("RewardToken deployed with transaction: %s", tx.Hash().Hex())
	log.Printf("RewardToken address: %s", tokenAddress.Hex())

	receipt, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for RewardToken deploy: %v", err)
	}
	log.Printf("RewardToken deployed at block: %d", receipt.BlockNumber.Uint64())

	return nil
}

// GetContractAddresses returns a map of contract addresses
func (cm *ContractManager) GetContractAddresses() map[string]string {
	return map[string]string{
		"task":   cm.taskAddress.Hex(),
		"family": cm.familyAddress.Hex(),
		"token":  cm.tokenAddress.Hex(),
		"reward": cm.rewardRegAddress.Hex(),
	}
}

// CreateTask creates a new task
func (cm *ContractManager) CreateTask(title, description string, reward *big.Int) (uint64, error) {
	if cm.TaskRegistry == nil {
		return 0, fmt.Errorf("task registry not initialized")
	}

	// 获取交易选项
	auth, err := cm.client.GetAuth()
	if err != nil {
		return 0, fmt.Errorf("failed to create auth: %v", err)
	}

	// 调用合约创建任务
	tx, err := cm.TaskRegistry.CreateTask(auth, title, description, reward)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %v", err)
	}

	log.Printf("CreateTask transaction submitted: %s", tx.Hash().Hex())

	// 等待交易确认
	receipt, err := cm.client.WaitForTransaction(tx.Hash(), 2*time.Minute)
	if err != nil {
		return 0, fmt.Errorf("error waiting for transaction: %v", err)
	}

	// 解析事件获取任务ID
	for _, log := range receipt.Logs {
		event, err := cm.TaskRegistry.ParseTaskCreated(*log)
		if err == nil && event != nil {
			return event.TaskId.Uint64(), nil
		}
	}

	return 0, fmt.Errorf("could not find task ID in transaction logs")
}

// AssignTask assigns a task to a child
func (cm *ContractManager) AssignTask(taskID uint64, childAddress common.Address) error {
	if cm.TaskRegistry == nil {
		return fmt.Errorf("task registry not initialized")
	}

	// 获取交易选项
	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	// 调用合约分配任务
	tx, err := cm.TaskRegistry.AssignTask(auth, big.NewInt(int64(taskID)), childAddress)
	if err != nil {
		return fmt.Errorf("failed to assign task: %v", err)
	}

	// 等待交易确认
	_, err = cm.client.WaitForTransaction(tx.Hash(), 2*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}

	return nil
}

// CompleteTask marks a task as completed by a child
func (cm *ContractManager) CompleteTask(taskID uint64) error {
	if cm.TaskRegistry == nil {
		return fmt.Errorf("task registry not initialized")
	}

	// 获取交易选项
	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	// 调用合约完成任务
	tx, err := cm.TaskRegistry.CompleteTask(auth, big.NewInt(int64(taskID)))
	if err != nil {
		return fmt.Errorf("failed to complete task: %v", err)
	}

	// 等待交易确认
	_, err = cm.client.WaitForTransaction(tx.Hash(), 2*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}

	return nil
}

// ApproveTask approves a completed task and transfers reward
func (cm *ContractManager) ApproveTask(taskID uint64, reward *big.Int) error {
	if cm.TaskRegistry == nil {
		return fmt.Errorf("task registry not initialized")
	}

	// 获取交易选项
	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	// 设置交易值为奖励金额
	auth.Value = reward

	// 增加gas限制，确保交易有足够的gas执行
	auth.GasLimit = 500000

	// 增加gas价格，确保交易能够被优先处理
	if auth.GasPrice != nil {
		increasedGasPrice := new(big.Int).Mul(auth.GasPrice, big.NewInt(200))
		increasedGasPrice = increasedGasPrice.Div(increasedGasPrice, big.NewInt(100)) // 增加100%
		auth.GasPrice = increasedGasPrice
		log.Printf("[DEBUG] ApproveTask: 增加gas价格至原价的200%%: %s", auth.GasPrice.String())
	}

	// 调用合约批准任务
	tx, err := cm.TaskRegistry.ApproveTask(auth, big.NewInt(int64(taskID)))
	if err != nil {
		return fmt.Errorf("failed to approve task: %v", err)
	}

	log.Printf("[DEBUG] ApproveTask: 交易已提交，哈希: %s", tx.Hash().Hex())

	// 等待交易确认
	receipt, err := cm.client.WaitForTransaction(tx.Hash(), 2*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}

	log.Printf("[DEBUG] ApproveTask: 交易已确认，区块号: %d，状态: %d", receipt.BlockNumber.Uint64(), receipt.Status)

	return nil
}

// RejectTask rejects a completed task
func (cm *ContractManager) RejectTask(taskID uint64) error {
	if cm.TaskRegistry == nil {
		return fmt.Errorf("task registry not initialized")
	}

	// 获取交易选项
	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	// 调用合约拒绝任务
	tx, err := cm.TaskRegistry.RejectTask(auth, big.NewInt(int64(taskID)))
	if err != nil {
		return fmt.Errorf("failed to reject task: %v", err)
	}

	// 等待交易确认
	_, err = cm.client.WaitForTransaction(tx.Hash(), 2*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}

	return nil
}

// TransferETH transfers ETH to the given address
func (cm *ContractManager) TransferETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	// 获取交易选项
	auth, err := cm.client.GetAuth()
	if err != nil {
		return nil, fmt.Errorf("failed to create auth: %v", err)
	}

	// 设置交易值
	auth.Value = amount

	// 创建交易对象
	tx := types.NewTransaction(
		auth.Nonce.Uint64(),
		to,
		auth.Value,
		21000, // gas limit for standard ETH transfer
		auth.GasPrice,
		[]byte{}, // 没有额外数据
	)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(cm.chainID), cm.client.GetPrivateKey())
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	// 发送交易
	err = cm.client.GetClient().SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx, nil
}

// CreateFamily creates a new family
func (cm *ContractManager) CreateFamily(name string) (uint64, error) {
	if cm.FamilyRegistry == nil {
		return 0, fmt.Errorf("family registry not initialized")
	}

	// 获取交易选项
	auth, err := cm.client.GetAuth()
	if err != nil {
		return 0, fmt.Errorf("failed to create auth: %v", err)
	}

	// 调用合约创建家庭
	tx, err := cm.FamilyRegistry.CreateFamily(auth, name)
	if err != nil {
		return 0, fmt.Errorf("failed to create family: %v", err)
	}

	log.Printf("CreateFamily transaction submitted: %s", tx.Hash().Hex())

	// 等待交易确认
	receipt, err := cm.client.WaitForTransaction(tx.Hash(), 2*time.Minute)
	if err != nil {
		return 0, fmt.Errorf("error waiting for transaction: %v", err)
	}

	// 解析事件获取家庭ID
	for _, log := range receipt.Logs {
		event, err := cm.FamilyRegistry.ParseFamilyCreated(*log)
		if err == nil && event != nil {
			return event.FamilyId.Uint64(), nil
		}
	}

	return 0, fmt.Errorf("could not find family ID in transaction logs")
}

// AddChild adds a child to a family
func (cm *ContractManager) AddChild(familyID uint64, childAddress common.Address, name string, age uint8) error {
	if cm.FamilyRegistry == nil {
		return fmt.Errorf("family registry not initialized")
	}

	// 获取交易选项
	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	// 调用合约添加孩子
	tx, err := cm.FamilyRegistry.AddChild(auth, big.NewInt(int64(familyID)), childAddress, name, age)
	if err != nil {
		return fmt.Errorf("failed to add child: %v", err)
	}

	// 等待交易确认
	_, err = cm.client.WaitForTransaction(tx.Hash(), 2*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}

	return nil
}

// Task represents a task from the TaskRegistry contract
type Task struct {
	ID          uint64
	Creator     common.Address
	AssignedTo  common.Address
	Title       string
	Description string
	Reward      *big.Int
	Completed   bool
	Approved    bool
}

// GetChildTransactOpts 获取孩子账户的交易选项
func (m *ContractManager) GetChildTransactOpts(ctx context.Context, childAddress common.Address) (*bind.TransactOpts, error) {
	// 这个方法应该通过孩子的地址获取他们的私钥或使用特定的方法来签署交易
	// 在此实现中，我们使用管理员账户模拟孩子账户的交易

	// 详细记录当前请求信息
	log.Printf("[nonce管理] 开始为地址获取交易选项: %s", childAddress.Hex())

	client := m.client.GetClient()

	// 1. 首先获取账户的已确认nonce (已经被打包进区块的交易数量)
	stateNonce, err := client.NonceAt(ctx, childAddress, nil)
	if err != nil {
		log.Printf("[nonce管理] 获取已确认nonce失败: %v", err)
		return nil, fmt.Errorf("failed to get confirmed nonce: %w", err)
	}
	log.Printf("[nonce管理] 已确认nonce: %d", stateNonce)

	// 2. 获取待处理nonce (包括内存池中的交易)
	pendingNonce, err := client.PendingNonceAt(ctx, childAddress)
	if err != nil {
		log.Printf("[nonce管理] 获取待处理nonce失败: %v", err)
		return nil, fmt.Errorf("failed to get pending nonce: %w", err)
	}
	log.Printf("[nonce管理] 待处理nonce: %d", pendingNonce)

	// 3. 选择正确的nonce值 (应该使用最大的那个)
	var finalNonce uint64
	if pendingNonce > stateNonce {
		finalNonce = pendingNonce
		log.Printf("[nonce管理] 使用待处理nonce: %d", finalNonce)
	} else {
		finalNonce = stateNonce
		log.Printf("[nonce管理] 使用已确认nonce: %d", finalNonce)
	}

	// 为解决持续的nonce错误，添加一个随机增量
	// 这可以避免多个交易使用相同的nonce
	randomIncrement := uint64(time.Now().UnixNano() % 1000)
	finalNonce = finalNonce + randomIncrement
	log.Printf("[nonce管理] 添加随机增量 %d，最终nonce: %d", randomIncrement, finalNonce)

	// 4. 获取gas价格，添加重试机制
	var gasPrice *big.Int
	for retries := 0; retries < 3; retries++ {
		gasPrice, err = client.SuggestGasPrice(ctx)
		if err == nil {
			break
		}
		log.Printf("[nonce管理] 获取gas价格失败(尝试 %d/3): %v", retries+1, err)
		time.Sleep(time.Second) // 等待1秒后重试
	}
	if err != nil {
		log.Printf("[nonce管理] 获取gas价格失败，所有尝试均失败: %v", err)
		return nil, fmt.Errorf("failed to get gas price: %w", err)
	}

	// 增加gas价格，确保交易能够被优先处理
	increasedGasPrice := new(big.Int).Mul(gasPrice, big.NewInt(150))
	increasedGasPrice = increasedGasPrice.Div(increasedGasPrice, big.NewInt(100)) // 增加50%
	log.Printf("[nonce管理] 原始gas价格: %s, 增加后: %s", gasPrice.String(), increasedGasPrice.String())

	// 5. 创建交易签名器
	auth, err := bind.NewKeyedTransactorWithChainID(m.client.GetPrivateKey(), m.client.GetChainID())
	if err != nil {
		log.Printf("[nonce管理] 创建交易签名器失败: %v", err)
		return nil, fmt.Errorf("failed to create transaction signer: %w", err)
	}

	// 6. 设置交易参数
	auth.Nonce = big.NewInt(int64(finalNonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(800000)    // 进一步增加gas限制
	auth.GasPrice = increasedGasPrice // 使用增加后的gas价格

	// 7. 记录最终交易参数
	log.Printf("[nonce管理] 交易参数设置完成 - Nonce: %d, GasPrice: %s, GasLimit: %d, Chain ID: %s",
		auth.Nonce.Uint64(), auth.GasPrice.String(), auth.GasLimit, m.client.GetChainID().String())

	return auth, nil
}

// syncNonce 尝试从区块链获取更准确的nonce值
// 这是一个临时解决方案，用于处理nonce同步问题
func (m *ContractManager) syncNonce(ctx context.Context, address common.Address, currentNonce *uint64) error {
	// 如果有错误，我们只记录但不中断流程
	client := m.client.GetClient()

	// 尝试获取地址的交易数量作为nonce参考
	count, err := client.PendingTransactionCount(ctx)
	if err != nil {
		return fmt.Errorf("获取待处理交易数量失败: %v", err)
	}
	log.Printf("当前待处理交易总数: %d", count)

	// 获取账户状态nonce (这是账户已确认的交易数量)
	stateNonce, err := client.NonceAt(ctx, address, nil)
	if err != nil {
		log.Printf("获取账户状态nonce失败: %v", err)
	} else {
		log.Printf("账户状态nonce: %d", stateNonce)
		if stateNonce > *currentNonce {
			*currentNonce = stateNonce
			log.Printf("使用账户状态nonce: %d", *currentNonce)
			return nil
		}
	}

	// 获取待处理nonce (包括内存池中的交易)
	pendingNonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		log.Printf("获取待处理nonce失败: %v", err)
	} else {
		log.Printf("待处理nonce: %d", pendingNonce)
		if pendingNonce > *currentNonce {
			*currentNonce = pendingNonce
			log.Printf("使用待处理nonce: %d", *currentNonce)
			return nil
		}
	}

	// 额外安全措施: 如果当前nonce看起来太低，增加一个合理的值
	// 这可以避免反复使用太低的nonce值
	if *currentNonce < 30 && count > 100 {
		// 如果账户交易总数很高但nonce很低，可能是严重不同步
		// 增加nonce到一个较高的值
		newNonce := *currentNonce + 150
		log.Printf("检测到可能的严重nonce不同步，尝试增加nonce: %d -> %d", *currentNonce, newNonce)
		*currentNonce = newNonce
	}

	return nil
}

// GetTransactOpts 获取交易选项
func (cm *ContractManager) GetTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := cm.client.GetAuth()
	if err != nil {
		return nil, fmt.Errorf("failed to create auth: %v", err)
	}

	// 增加gas价格，确保交易能够被优先处理
	if auth.GasPrice != nil {
		increasedGasPrice := new(big.Int).Mul(auth.GasPrice, big.NewInt(150))
		increasedGasPrice = increasedGasPrice.Div(increasedGasPrice, big.NewInt(100)) // 增加50%
		auth.GasPrice = increasedGasPrice
		log.Printf("[DEBUG] 增加gas价格至原价的150%%: %s", auth.GasPrice.String())
	}

	return auth, nil
}

// WaitForTxReceipt 等待交易确认并返回收据
func (cm *ContractManager) WaitForTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	log.Printf("[交易确认] 开始等待交易确认: %s", txHash.Hex())

	// 创建一个带有取消的上下文
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 60*time.Second) // 增加超时时间到60秒
	defer cancel()

	// 创建一个ticker，每秒检查一次
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 记录开始等待时间
	startTime := time.Now()

	// 设置最大重试次数
	maxRetries := 60
	retryCount := 0

	for {
		select {
		case <-ctxWithTimeout.Done():
			// 上下文已取消或超时
			if ctx.Err() == context.DeadlineExceeded {
				log.Printf("[交易确认] 等待交易确认超时: %s, 已等待 %v 秒",
					txHash.Hex(), time.Since(startTime).Seconds())
				return nil, fmt.Errorf("waiting for transaction receipt timed out: %w", ctx.Err())
			}
			log.Printf("[交易确认] 上下文已取消: %s", txHash.Hex())
			return nil, ctx.Err()
		case <-ticker.C:
			retryCount++
			if retryCount > maxRetries {
				log.Printf("[交易确认] 达到最大重试次数 %d，放弃等待: %s", maxRetries, txHash.Hex())
				return nil, fmt.Errorf("reached maximum retry count waiting for transaction receipt")
			}

			// 每10次尝试记录一次日志
			if retryCount%10 == 0 {
				log.Printf("[交易确认] 尝试 %d/%d, 已等待 %.1f 秒: %s",
					retryCount, maxRetries, time.Since(startTime).Seconds(), txHash.Hex())
			}

			// 尝试获取交易收据
			receipt, err := cm.client.GetClient().TransactionReceipt(ctx, txHash)
			if err != nil {
				if err == ethereum.NotFound {
					// 交易尚未被打包，继续等待
					continue
				}
				// 其他错误
				log.Printf("[交易确认] 获取交易收据失败: %v", err)
				return nil, fmt.Errorf("failed to get transaction receipt: %w", err)
			}

			// 获取到收据，检查状态
			if receipt.Status == types.ReceiptStatusFailed {
				log.Printf("[交易确认] 交易失败: %s, 区块号: %d", txHash.Hex(), receipt.BlockNumber.Uint64())
				return receipt, fmt.Errorf("transaction failed with status: %d", receipt.Status)
			}

			log.Printf("[交易确认] 交易成功确认: %s, 区块号: %d, Gas使用: %d, 确认用时: %.1f 秒",
				txHash.Hex(), receipt.BlockNumber.Uint64(), receipt.GasUsed, time.Since(startTime).Seconds())
			return receipt, nil
		}
	}
}

// InitRewardToken initializes or reinitializes the reward token contract
func (cm *ContractManager) InitRewardToken(tokenAddress string) error {
	// 使用地址初始化代币合约
	if tokenAddress == "" {
		return fmt.Errorf("token address is empty")
	}

	address := common.HexToAddress(tokenAddress)
	if address == (common.Address{}) {
		return fmt.Errorf("invalid token address format")
	}

	cm.tokenAddress = address

	// 初始化代币合约实例
	rewardToken, err := NewRewardToken(address, cm.client.GetClient())
	if err != nil {
		return fmt.Errorf("failed to instantiate reward token contract: %v", err)
	}

	// 更新合约实例
	cm.RewardToken = rewardToken
	log.Printf("RewardToken re-initialized at address: %s", address.Hex())

	return nil
}
