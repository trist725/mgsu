package util

// KnuthDurstenfeldShuffle 泛型洗牌
// 时间复杂度为O(n)
// 空间复杂度为O(1)
// 必须知道数组长度
func KnuthDurstenfeldShuffle[T any](arr []T) []T {
	for i := len(arr) - 1; i >= 0; i-- {
		p := RandomInt64(0, int64(i))
		a := arr[i]
		arr[i] = arr[p]
		arr[p] = a
	}
	return arr
}
