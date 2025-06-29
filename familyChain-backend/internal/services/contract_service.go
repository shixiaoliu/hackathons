package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"strings"

	"eth-for-babies-backend/internal/config"
	"eth-for-babies-backend/internal/utils"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractService struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	config     *config.BlockchainConfig
}

func NewContractService(cfg *config.BlockchainConfig) (*ContractService, error) {
	// 连接到以太坊节点
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, errors.New("failed to connect to Ethereum client: " + err.Error())
	}

	// 解析私钥
	var privateKey *ecdsa.PrivateKey
	if cfg.PrivateKey != "" {
		privateKey, err = crypto.HexToECDSA(strings.TrimPrefix(cfg.PrivateKey, "0x"))
		if err != nil {
			return nil, errors.New("failed to parse private key: " + err.Error())
		}
	}

	return &ContractService{
		client:     client,
		privateKey: privateKey,
		config:     cfg,
	}, nil
}

// GetBalance 获取地址的ETH余额
func (s *ContractService) GetBalance(address string) (*big.Int, error) {
	if !utils.IsValidEthereumAddress(address) {
		return nil, errors.New("invalid Ethereum address")
	}

	account := common.HexToAddress(address)
	balance, err := s.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, errors.New("failed to get balance: " + err.Error())
	}

	return balance, nil
}

// GetTokenBalance 获取代币余额（需要实现具体的代币合约调用）
func (s *ContractService) GetTokenBalance(tokenAddress, walletAddress string) (*big.Int, error) {
	if !utils.IsValidEthereumAddress(tokenAddress) || !utils.IsValidEthereumAddress(walletAddress) {
		return nil, errors.New("invalid address format")
	}

	// TODO: 实现ERC20代币余额查询
	// 这里需要根据具体的代币合约ABI来实现
	// 暂时返回模拟数据
	return big.NewInt(1000), nil
}

// TransferETH 转移ETH
func (s *ContractService) TransferETH(to string, amount *big.Int) (*types.Transaction, error) {
	if s.privateKey == nil {
		return nil, errors.New("private key not configured")
	}

	if !utils.IsValidEthereumAddress(to) {
		return nil, errors.New("invalid recipient address")
	}

	// 获取发送者地址
	fromAddress := crypto.PubkeyToAddress(s.privateKey.PublicKey)

	// 获取nonce
	nonce, err := s.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, errors.New("failed to get nonce: " + err.Error())
	}

	// 获取gas价格
	gasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.New("failed to get gas price: " + err.Error())
	}

	// 创建交易
	toAddress := common.HexToAddress(to)
	tx := types.NewTransaction(nonce, toAddress, amount, 21000, gasPrice, nil)

	// 签名交易
	chainID := big.NewInt(s.config.ChainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), s.privateKey)
	if err != nil {
		return nil, errors.New("failed to sign transaction: " + err.Error())
	}

	// 发送交易
	err = s.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, errors.New("failed to send transaction: " + err.Error())
	}

	return signedTx, nil
}

// TransferToken 转移代币
func (s *ContractService) TransferToken(tokenAddress, to string, amount *big.Int) (*types.Transaction, error) {
	if s.privateKey == nil {
		return nil, errors.New("private key not configured")
	}

	if !utils.IsValidEthereumAddress(tokenAddress) || !utils.IsValidEthereumAddress(to) {
		return nil, errors.New("invalid address format")
	}

	// TODO: 实现ERC20代币转账
	// 这里需要根据具体的代币合约ABI来实现
	// 暂时返回模拟交易
	return &types.Transaction{}, nil
}

// GetTransactionReceipt 获取交易收据
func (s *ContractService) GetTransactionReceipt(txHash string) (*types.Receipt, error) {
	if len(txHash) != 66 || !strings.HasPrefix(txHash, "0x") {
		return nil, errors.New("invalid transaction hash format")
	}

	hash := common.HexToHash(txHash)
	receipt, err := s.client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		return nil, errors.New("failed to get transaction receipt: " + err.Error())
	}

	return receipt, nil
}

// GetTransactionByHash 根据哈希获取交易信息
func (s *ContractService) GetTransactionByHash(txHash string) (*types.Transaction, bool, error) {
	if len(txHash) != 66 || !strings.HasPrefix(txHash, "0x") {
		return nil, false, errors.New("invalid transaction hash format")
	}

	hash := common.HexToHash(txHash)
	tx, isPending, err := s.client.TransactionByHash(context.Background(), hash)
	if err != nil {
		return nil, false, errors.New("failed to get transaction: " + err.Error())
	}

	return tx, isPending, nil
}

// EstimateGas 估算gas费用
func (s *ContractService) EstimateGas(to string, data []byte) (uint64, error) {
	if !utils.IsValidEthereumAddress(to) {
		return 0, errors.New("invalid recipient address")
	}

	fromAddress := crypto.PubkeyToAddress(s.privateKey.PublicKey)
	toAddress := common.HexToAddress(to)

	msg := ethereum.CallMsg{
		From: fromAddress,
		To:   &toAddress,
		Data: data,
	}

	gasLimit, err := s.client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0, errors.New("failed to estimate gas: " + err.Error())
	}

	return gasLimit, nil
}

// GetChainID 获取链ID
func (s *ContractService) GetChainID() (*big.Int, error) {
	chainID, err := s.client.ChainID(context.Background())
	if err != nil {
		return nil, errors.New("failed to get chain ID: " + err.Error())
	}
	return chainID, nil
}

// GetBlockNumber 获取最新区块号
func (s *ContractService) GetBlockNumber() (uint64, error) {
	blockNumber, err := s.client.BlockNumber(context.Background())
	if err != nil {
		return 0, errors.New("failed to get block number: " + err.Error())
	}
	return blockNumber, nil
}

// CreateTransactor 创建交易器
func (s *ContractService) CreateTransactor() (*bind.TransactOpts, error) {
	if s.privateKey == nil {
		return nil, errors.New("private key not configured")
	}

	chainID := big.NewInt(s.config.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, chainID)
	if err != nil {
		return nil, errors.New("failed to create transactor: " + err.Error())
	}

	return auth, nil
}

// Close 关闭客户端连接
func (s *ContractService) Close() {
	if s.client != nil {
		s.client.Close()
	}
}