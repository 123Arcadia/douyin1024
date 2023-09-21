package utils

import (
	"fmt"
)

// GetVideoNewName 获取存储的文件名称userId + i(该用户第几个视频) + fileName
func GetVideoNewName(userId uint, num int64, filename string, title string) string {
	return fmt.Sprintf("%d_%d_%s_%s", userId, num, filename, title)
}
