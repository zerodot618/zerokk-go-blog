package utils

import "testing"

func TestMd5Crypt(t *testing.T) {

	t.Log(Md5Crypt("123456", "zerokk"))
}
