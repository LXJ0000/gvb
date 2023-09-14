package utils

import (
	"fmt"
	"time"
)

// InList 判断key是否存在于list列表中
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

func GetUniqueFileName(imageName string) string {
	now := time.Now().Unix()
	return fmt.Sprintf("%d__%s", now, imageName)
}
