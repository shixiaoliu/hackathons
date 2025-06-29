package utils

import (
	"regexp"
	"strings"
)

// IsValidEthereumAddress 验证以太坊地址格式
func IsValidEthereumAddress(address string) bool {
	if len(address) != 42 {
		return false
	}
	if !strings.HasPrefix(address, "0x") {
		return false
	}
	matched, _ := regexp.MatchString("^0x[a-fA-F0-9]{40}$", address)
	return matched
}

// IsValidRole 验证用户角色
func IsValidRole(role string) bool {
	return role == "parent" || role == "child"
}

// IsValidDifficulty 验证任务难度
func IsValidDifficulty(difficulty string) bool {
	validDifficulties := []string{"easy", "medium", "hard"}
	for _, d := range validDifficulties {
		if d == difficulty {
			return true
		}
	}
	return false
}

// IsValidTaskStatus 验证任务状态
func IsValidTaskStatus(status string) bool {
	validStatuses := []string{"pending", "in_progress", "completed", "approved", "rejected"}
	for _, s := range validStatuses {
		if s == status {
			return true
		}
	}
	return false
}

// SanitizeString 清理字符串输入
func SanitizeString(input string) string {
	return strings.TrimSpace(input)
}