package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Md5Crypt 给字符串生成 md5
// @param str 需要加密的字符串
// @param salt 附加的盐
// @return string
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
