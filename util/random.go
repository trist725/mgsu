package util

import (
	"math/rand"
	"sync"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// GenRandomByteArray 生成一定长度的随机字节数组
func GenRandomByteArray(size int) []byte {
	diceLock.Lock()
	defer diceLock.Unlock()

	b := make([]byte, size)
	for i, cache, remain := size-1, dice.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = dice.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

// GenRandomString 生成一定长度的随机字符串
func GenRandomString(size int) string {
	return string(GenRandomByteArray(size))
}

// rand不是线程/协程安全，必须加锁
var dice = rand.New(rand.NewSource(time.Now().UnixNano()))
var diceLock = &sync.Mutex{}

// RandomTimeDuration 注意一下RandomXX函数族得到的是 [min, max)
func RandomTimeDuration(min time.Duration, max time.Duration) time.Duration {
	if min == max {
		return min
	}

	diceLock.Lock()
	defer diceLock.Unlock()

	if min < max {
		return min + time.Duration(dice.Int63n(int64(max-min)))
	}
	return max + time.Duration(dice.Int63n(int64(min-max)))
}

func RandomInt32(min int32, max int32) int32 {
	if min == max {
		return min
	}

	diceLock.Lock()
	defer diceLock.Unlock()

	if min < max {
		return min + dice.Int31n(max-min)
	}
	return max + dice.Int31n(min-max)
}

func RandomInt(min int, max int) int {
	if min == max {
		return min
	}

	diceLock.Lock()
	defer diceLock.Unlock()

	if min < max {
		return min + dice.Intn(max-min)
	}
	return max + dice.Intn(min-max)
}

func RandomInt64(min int64, max int64) int64 {
	if min == max {
		return min
	}

	diceLock.Lock()
	defer diceLock.Unlock()

	if min < max {
		return min + dice.Int63n(max-min)
	}
	return max + dice.Int63n(min-max)
}

// HitProbability 百分比概率
func HitProbability(prob uint8) bool {
	if prob > 100 || prob == 0 {
		return false
	}

	if prob >= uint8(RandomInt(1, int(100+1))) {
		return true
	}

	return false
}

// HitProbabilityThousands 万分比概率
func HitProbabilityThousands(prob uint16) bool {
	if prob > 10000 || prob == 0 {
		return false
	}

	if prob >= uint16(RandomInt(1, 10000+1)) {
		return true
	}

	return false
}

// RandArrElemNoRepeat 随机从arr选count个位置不重复的元素
// 空间复杂度O(1) 会改变arr元素顺序
func RandArrElemNoRepeat[T any](arr []T, count int) (chosen []T) {
	if count > len(arr) {
		count = len(arr)
	}
	for i := 0; i < count; i++ {
		rd := RandomInt(0, len(arr)-i)
		chosen = append(chosen, arr[rd])
		arr[rd], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[rd]
	}
	return chosen
}
