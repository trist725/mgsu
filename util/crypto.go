package util

import (
	"crypto/md5"
	"fmt"
	"sync"
)

var (
	md5Hasher     = md5.New()
	md5HasherLock = &sync.Mutex{}
)

// 生成md串
func MD5Sum(input []byte) (output []byte) {
	md5HasherLock.Lock()
	defer md5HasherLock.Unlock()

	output = md5Hasher.Sum(input)

	return
}

// 生成md串
func MD5Sumf(format string, args ...interface{}) (output []byte) {
	md5HasherLock.Lock()
	defer md5HasherLock.Unlock()

	input := []byte(fmt.Sprintf(format, args...))

	output = md5Hasher.Sum(input)

	return
}
