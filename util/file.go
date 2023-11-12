package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/duchoang206h/file-storage/config"
)

func RandomInt(length int) string {
	const digits = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}

func TimestampFileName() string {
	t := time.Now().Unix()
	return fmt.Sprintf("%d-%s", t, RandomInt(4))
}

func FormatFileUrl(fileName string) string {
	return fmt.Sprintf("%s/%s", config.Config("FILE_HOST"), fileName)
}
