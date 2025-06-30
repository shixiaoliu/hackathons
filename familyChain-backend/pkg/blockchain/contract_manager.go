package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ContractManager manages interactions with smart contracts
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

// NewContractManager creates a new contract manager
func NewContractManager(client *EthClient, contractAddresses map[string]string) (*ContractManager, error) {
	manager := &ContractManager{
		client: client,
	}

	// Set contract addresses if provided
	if taskAddr, ok := contractAddresses["TaskRegistry"]; ok && taskAddr != "" {
		manager.taskAddress = common.HexToAddress(taskAddr)
	}

	if familyAddr, ok := contractAddresses["FamilyRegistry"]; ok && familyAddr != "" {
		manager.familyAddress = common.HexToAddress(familyAddr)
	}

	if tokenAddr, ok := contractAddresses["RewardToken"]; ok && tokenAddr != "" {
		manager.tokenAddress = common.HexToAddress(tokenAddr)
	}

	if rewardRegAddr, ok := contractAddresses["RewardRegistry"]; ok && rewardRegAddr != "" {
		rewardRegAddress := common.HexToAddress(rewardRegAddr)
		rewardRegistry, err := NewRewardRegistry(rewardRegAddress, client.GetClient())
		if err != nil {
			return nil, fmt.Errorf("failed to initialize reward registry contract: %v", err)
		}
		manager.RewardRegistry = rewardRegistry
		manager.rewardRegAddress = rewardRegAddress
	}

	// Initialize contract instances
	err := manager.initializeContracts()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize contracts: %v", err)
	}

	return manager, nil
}

// initializeContracts initializes contract instances
func (cm *ContractManager) initializeContracts() error {
	ethClient := cm.client.GetClient()

	// Initialize TaskRegistry if address is set
	if cm.taskAddress != (common.Address{}) {
		taskRegistry, err := NewTaskRegistry(cm.taskAddress, ethClient)
		if err != nil {
			return fmt.Errorf("failed to initialize TaskRegistry contract: %v", err)
		}
		cm.TaskRegistry = taskRegistry
	}

	// Initialize FamilyRegistry if address is set
	if cm.familyAddress != (common.Address{}) {
		familyRegistry, err := NewFamilyRegistry(cm.familyAddress, ethClient)
		if err != nil {
			return fmt.Errorf("failed to initialize FamilyRegistry contract: %v", err)
		}
		cm.FamilyRegistry = familyRegistry
	}

	// Initialize RewardToken if address is set
	if cm.tokenAddress != (common.Address{}) {
		rewardToken, err := NewRewardToken(cm.tokenAddress, ethClient)
		if err != nil {
			return fmt.Errorf("failed to initialize RewardToken contract: %v", err)
		}
		cm.RewardToken = rewardToken
	}

	return nil
}

// DeployContracts deploys all contracts and initializes them
func (cm *ContractManager) DeployContracts() error {
	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	ethClient := cm.client.GetClient()

	// Deploy TaskRegistry
	log.Println("Deploying TaskRegistry contract...")
	taskRegistryAddress, tx, taskRegistry, err := DeployTaskRegistry(auth, ethClient)
	if err != nil {
		return fmt.Errorf("failed to deploy TaskRegistry: %v", err)
	}
	cm.taskAddress = taskRegistryAddress
	cm.TaskRegistry = taskRegistry
	log.Printf("TaskRegistry deployed at: %s", taskRegistryAddress.Hex())

	// Wait for transaction to be mined
	_, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		log.Printf("Error waiting for TaskRegistry deployment: %v", err)
	}

	// Deploy FamilyRegistry
	log.Println("Deploying FamilyRegistry contract...")
	familyRegistryAddress, tx, familyRegistry, err := DeployFamilyRegistry(auth, ethClient)
	if err != nil {
		return fmt.Errorf("failed to deploy FamilyRegistry: %v", err)
	}
	cm.familyAddress = familyRegistryAddress
	cm.FamilyRegistry = familyRegistry
	log.Printf("FamilyRegistry deployed at: %s", familyRegistryAddress.Hex())

	// Wait for transaction to be mined
	_, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		log.Printf("Error waiting for FamilyRegistry deployment: %v", err)
	}

	// Deploy RewardToken
	log.Println("Deploying RewardToken contract...")
	rewardTokenAddress, tx, rewardToken, err := DeployRewardToken(auth, ethClient, "EthForBabiesToken", "EFBT")
	if err != nil {
		return fmt.Errorf("failed to deploy RewardToken: %v", err)
	}
	cm.tokenAddress = rewardTokenAddress
	cm.RewardToken = rewardToken
	log.Printf("RewardToken deployed at: %s", rewardTokenAddress.Hex())

	// Wait for transaction to be mined
	_, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		log.Printf("Error waiting for RewardToken deployment: %v", err)
	}

	return nil
}

// GetContractAddresses returns the addresses of all contracts
func (cm *ContractManager) GetContractAddresses() map[string]string {
	return map[string]string{
		"TaskRegistry":   cm.taskAddress.Hex(),
		"FamilyRegistry": cm.familyAddress.Hex(),
		"RewardToken":    cm.tokenAddress.Hex(),
	}
}

// Task Management Functions

// CreateTask creates a new task in the TaskRegistry contract
func (cm *ContractManager) CreateTask(title, description string, reward *big.Int) (uint64, error) {
	if cm.TaskRegistry == nil {
		return 0, fmt.Errorf("TaskRegistry contract not initialized")
	}

	auth, err := cm.client.GetAuth()
	if err != nil {
		return 0, fmt.Errorf("failed to create auth: %v", err)
	}

	// Set the value to be sent with the transaction (must equal reward)
	auth.Value = reward

	// Add log for task creation
	log.Printf("Creating task with title: %s, description: %s, reward: %s",
		title, description, reward.String())

	tx, err := cm.TaskRegistry.CreateTask(auth, title, description, reward)
	log.Printf("create task	tx: %v", tx)
	if err != nil {
		return 0, fmt.Errorf("failed to create task: %v", err)
	}

	// Wait for transaction to be mined
	log.Printf("Waiting for transaction %s to be mined", tx.Hash().Hex())
	receipt, err := cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return 0, fmt.Errorf("error waiting for transaction: %v", err)
	}

	// Parse task ID from logs
	for _, log := range receipt.Logs {
		event, err := cm.TaskRegistry.ParseTaskCreated(*log)
		if err == nil {
			return event.TaskId.Uint64(), nil
		}
	}

	return 0, fmt.Errorf("failed to parse task ID from transaction logs")
}

// AssignTask assigns a task to a child
func (cm *ContractManager) AssignTask(taskID uint64, childAddress common.Address) error {
	if cm.TaskRegistry == nil {
		return fmt.Errorf("TaskRegistry contract not initialized")
	}

	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := cm.TaskRegistry.AssignTask(auth, big.NewInt(int64(taskID)), childAddress)
	if err != nil {
		return fmt.Errorf("failed to assign task: %v", err)
	}

	// Wait for transaction to be mined
	_, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}

	return nil
}

// CompleteTask marks a task as completed by the assigned child
func (cm *ContractManager) CompleteTask(taskID uint64) error {
	if cm.TaskRegistry == nil {
		return fmt.Errorf("TaskRegistry contract not initialized")
	}

	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := cm.TaskRegistry.CompleteTask(auth, big.NewInt(int64(taskID)))
	if err != nil {
		return fmt.Errorf("failed to complete task: %v", err)
	}

	// Wait for transaction to be mined
	_, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}

	return nil
}

// ApproveTask approves a completed task and transfers the reward
func (cm *ContractManager) ApproveTask(taskID uint64, reward *big.Int) error {
	if cm.TaskRegistry == nil {
		return fmt.Errorf("TaskRegistry contract not initialized")
	}

	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	// Do not set auth.Value as the reward is already locked in the contract
	// The smart contract will transfer the reward from its own balance to the child
	auth.Value = big.NewInt(0)

	tx, err := cm.TaskRegistry.ApproveTask(auth, big.NewInt(int64(taskID)))
	if err != nil {
		return fmt.Errorf("failed to approve task: %v", err)
	}

	// Wait for transaction to be mined
	_, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}

	return nil
}

// RejectTask rejects a completed task and refunds the reward to the parent
func (cm *ContractManager) RejectTask(taskID uint64) error {
	if cm.TaskRegistry == nil {
		return fmt.Errorf("TaskRegistry contract not initialized")
	}

	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	// Do not set auth.Value as the reward is already locked in the contract
	// The smart contract will refund the reward from its balance to the parent
	auth.Value = big.NewInt(0)

	// 直接使用生成的RejectTask方法
	tx, err := cm.TaskRegistry.RejectTask(auth, big.NewInt(int64(taskID)))
	if err != nil {
		return fmt.Errorf("failed to reject task: %v", err)
	}

	// Wait for transaction to be mined
	_, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return fmt.Errorf("error waiting for transaction: %v", err)
	}

	return nil
}

// TransferETH transfers ETH from the contract manager's account to a specified address
func (cm *ContractManager) TransferETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	if cm.client == nil {
		return nil, fmt.Errorf("eth client not initialized")
	}

	// Send ETH transaction
	tx, err := cm.client.SendTransaction(to, amount, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send ETH transaction: %v", err)
	}

	// Wait for transaction to be mined
	_, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("error waiting for ETH transfer transaction: %v", err)
	}

	return tx, nil
}

// GetTask retrieves a task by ID
func (cm *ContractManager) GetTask(taskID uint64) (Task, error) {
	if cm.TaskRegistry == nil {
		return Task{}, fmt.Errorf("TaskRegistry contract not initialized")
	}

	id, creator, assignedTo, title, description, reward, completed, approved, err := cm.TaskRegistry.GetTask(&bind.CallOpts{}, big.NewInt(int64(taskID)))
	if err != nil {
		return Task{}, fmt.Errorf("failed to get task: %v", err)
	}

	return Task{
		ID:          id.Uint64(),
		Creator:     creator,
		AssignedTo:  assignedTo,
		Title:       title,
		Description: description,
		Reward:      reward,
		Completed:   completed,
		Approved:    approved,
	}, nil
}

// Family Management Functions

// CreateFamily creates a new family in the FamilyRegistry contract
func (cm *ContractManager) CreateFamily(name string) (uint64, error) {
	if cm.FamilyRegistry == nil {
		return 0, fmt.Errorf("FamilyRegistry contract not initialized")
	}

	auth, err := cm.client.GetAuth()
	if err != nil {
		return 0, fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := cm.FamilyRegistry.CreateFamily(auth, name)
	if err != nil {
		return 0, fmt.Errorf("failed to create family: %v", err)
	}

	// Wait for transaction to be mined
	receipt, err := cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
	if err != nil {
		return 0, fmt.Errorf("error waiting for transaction: %v", err)
	}

	// Parse family ID from logs
	for _, log := range receipt.Logs {
		event, err := cm.FamilyRegistry.ParseFamilyCreated(*log)
		if err == nil {
			return event.FamilyId.Uint64(), nil
		}
	}

	return 0, fmt.Errorf("failed to parse family ID from transaction logs")
}

// AddChild adds a child to a family
func (cm *ContractManager) AddChild(familyID uint64, childAddress common.Address, name string, age uint8) error {
	if cm.FamilyRegistry == nil {
		return fmt.Errorf("FamilyRegistry contract not initialized")
	}

	auth, err := cm.client.GetAuth()
	if err != nil {
		return fmt.Errorf("failed to create auth: %v", err)
	}

	tx, err := cm.FamilyRegistry.AddChild(auth, big.NewInt(int64(familyID)), childAddress, name, age)
	if err != nil {
		return fmt.Errorf("failed to add child: %v", err)
	}

	// Wait for transaction to be mined
	_, err = cm.client.WaitForTransaction(tx.Hash(), 5*time.Minute)
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

// Note: TaskRegistry type is generated by abigen
// Placeholder types for FamilyRegistry and RewardToken (to be generated by abigen)
type FamilyRegistry struct{}
type RewardToken struct{}

// Note: NewTaskRegistry function is generated by abigen
// Placeholder functions for FamilyRegistry and RewardToken
func NewFamilyRegistry(address common.Address, client bind.ContractBackend) (*FamilyRegistry, error) {
	return &FamilyRegistry{}, nil
}

func NewRewardToken(address common.Address, client bind.ContractBackend) (*RewardToken, error) {
	return &RewardToken{}, nil
}

// Placeholder deploy functions
func DeployTaskRegistry(auth *bind.TransactOpts, client bind.ContractBackend) (common.Address, *types.Transaction, *TaskRegistry, error) {
	return common.Address{}, &types.Transaction{}, &TaskRegistry{}, fmt.Errorf("TaskRegistry deployment not implemented")
}

func DeployFamilyRegistry(auth *bind.TransactOpts, client bind.ContractBackend) (common.Address, *types.Transaction, *FamilyRegistry, error) {
	return common.Address{}, &types.Transaction{}, &FamilyRegistry{}, fmt.Errorf("FamilyRegistry deployment not implemented")
}

func DeployRewardToken(auth *bind.TransactOpts, client bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *RewardToken, error) {
	return common.Address{}, &types.Transaction{}, &RewardToken{}, fmt.Errorf("RewardToken deployment not implemented")
}

// Note: TaskRegistry methods are now generated in task_registry_generated.go

// Placeholder methods for FamilyRegistry
func (fr *FamilyRegistry) CreateFamily(auth *bind.TransactOpts, name string) (*types.Transaction, error) {
	return &types.Transaction{}, nil
}

func (fr *FamilyRegistry) AddChild(auth *bind.TransactOpts, familyID *big.Int, childAddress common.Address, name string, age uint8) (*types.Transaction, error) {
	return &types.Transaction{}, nil
}

// Placeholder method for parsing events
func (fr *FamilyRegistry) ParseFamilyCreated(log types.Log) (*struct{ FamilyId *big.Int }, error) {
	return &struct{ FamilyId *big.Int }{FamilyId: big.NewInt(0)}, nil
}

// GetChildTransactOpts 获取孩子账户的交易选项
func (m *ContractManager) GetChildTransactOpts(ctx context.Context, childAddress common.Address) (*bind.TransactOpts, error) {
	// 这个方法应该通过孩子的地址获取他们的私钥或使用特定的方法来签署交易
	// 这里是一个简化的实现，实际使用中需要根据应用架构调整
	// 在此简化实现中，我们使用管理员账户模拟孩子账户的交易

	nonce, err := m.client.GetClient().PendingNonceAt(ctx, childAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := m.client.GetClient().SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(m.client.GetPrivateKey(), m.client.GetChainID())
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return auth, nil
}

// GetTransactOpts 获取交易选项
func (cm *ContractManager) GetTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := cm.client.GetAuth()
	if err != nil {
		return nil, fmt.Errorf("failed to create auth: %v", err)
	}

	return auth, nil
}

// WaitForTxReceipt 等待交易确认并返回收据
func (cm *ContractManager) WaitForTxReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	receipt, err := cm.client.WaitForTransaction(txHash, 5*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("error waiting for transaction: %v", err)
	}

	return receipt, nil
}
