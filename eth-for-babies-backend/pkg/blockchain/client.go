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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient wraps ethclient.Client to provide additional functionality
type EthClient struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	address    common.Address
	chainID    *big.Int
}

// NewEthClient creates a new Ethereum client
func NewEthClient(rpcURL string, privateKeyHex string) (*EthClient, error) {
	// Connect to Ethereum node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %v", err)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	// Get public key and address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	return &EthClient{
		client:     client,
		privateKey: privateKey,
		address:    address,
		chainID:    chainID,
	}, nil
}

// GetClient returns the underlying ethclient.Client
func (ec *EthClient) GetClient() *ethclient.Client {
	return ec.client
}

// GetAddress returns the wallet address associated with the client
func (ec *EthClient) GetAddress() common.Address {
	return ec.address
}

// GetPrivateKey returns the private key associated with the client
func (ec *EthClient) GetPrivateKey() *ecdsa.PrivateKey {
	return ec.privateKey
}

// GetChainID returns the chain ID of the connected network
func (ec *EthClient) GetChainID() *big.Int {
	return ec.chainID
}

// GetAuth creates a new transactor for contract interactions
func (ec *EthClient) GetAuth() (*bind.TransactOpts, error) {
	nonce, err := ec.client.PendingNonceAt(context.Background(), ec.address)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ec.privateKey, ec.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create keyed transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

// CreateAuthWithPrivateKey creates a new transactor using the provided private key
func (ec *EthClient) CreateAuthWithPrivateKey(privateKeyHex string) (*bind.TransactOpts, error) {
	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	// Get public key and address from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get nonce for the address
	nonce, err := ec.client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// Create auth
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, ec.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create keyed transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

// SendTransaction sends a raw transaction to the blockchain
func (ec *EthClient) SendTransaction(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	nonce, err := ec.client.PendingNonceAt(context.Background(), ec.address)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	// Create transaction
	tx := types.NewTransaction(
		nonce,
		to,
		value,
		uint64(3000000), // Gas limit
		gasPrice,
		data,
	)

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(ec.chainID), ec.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send transaction
	err = ec.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx, nil
}

// WaitForTransaction waits for a transaction to be mined
func (ec *EthClient) WaitForTransaction(txHash common.Hash, timeout time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout waiting for transaction %s", txHash.Hex())
		case <-ticker.C:
			receipt, err := ec.client.TransactionReceipt(context.Background(), txHash)
			if err != nil {
				continue
			}
			log.Printf("Transaction %s mined in block %d", txHash.Hex(), receipt.BlockNumber)
			return receipt, nil
		}
	}
}

// GetBalance returns the balance of an address
func (ec *EthClient) GetBalance(address common.Address) (*big.Int, error) {
	return ec.client.BalanceAt(context.Background(), address, nil)
}

// Close closes the client connection
func (ec *EthClient) Close() {
	ec.client.Close()
}
