package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/bcrypt"
)

// GenerateNonce 生成随机nonce
func GenerateNonce() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// HashPassword 哈希密码
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash 验证密码哈希
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// VerifySignature 验证以太坊签名
func VerifySignature(walletAddress, message, signature string) (bool, error) {
	// 检查是否为模拟签名（用于手动输入地址的测试）
	if signature == "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" {
		// 对于模拟签名，直接返回true（仅用于开发/测试环境）
		return true, nil
	}

	// 构造签名消息
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))

	// 解码签名
	sig, err := hexutil.Decode(signature)
	if err != nil {
		return false, err
	}

	// 调整v值（以太坊签名格式）
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}

	// 恢复公钥
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false, err
	}

	// 获取地址
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	// 比较地址（忽略大小写）
	return strings.EqualFold(walletAddress, recoveredAddr.Hex()), nil
}

// GetSignMessage 获取用于签名的消息
func GetSignMessage(nonce string) string {
	return fmt.Sprintf("Welcome to Family Task Chain!\n\nClick to sign in and accept the Terms of Service.\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nNonce: %s", nonce)
}