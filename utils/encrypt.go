package utils

import (
	"crypto/md5"
	"fmt"
)

// Md5Encode MD5加密
func Md5Encode(str string) string {
	sum := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}
