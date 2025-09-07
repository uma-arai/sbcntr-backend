package utils

import (
	"log"
)

// LogError はエラーメッセージをログに記録します
func LogError(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

// LogInfo は情報メッセージをログに記録します
func LogInfo(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}
