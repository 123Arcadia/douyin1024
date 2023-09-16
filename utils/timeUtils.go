package utils

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/initConfig"
	"log"
	"strconv"
	"time"
)

func GetFormatTime(timeStamp string) string {
	if timeStamp != "" {
		timestampInt64, err := strconv.ParseInt(timeStamp, 10, 64)
		if err != nil {
			log.Fatalf("时间格式错误 timeStamp: %v\n", timeStamp)
			return time.Now().Format(initConfig.TimeFormat)
		}
		if timestampInt64 > 1000000000000 {
			timestampInt64 /= 1000
			return time.Unix(timestampInt64, 0).Format(initConfig.TimeFormat)
		}
	}
	return time.Now().Format(initConfig.TimeFormat)
}
