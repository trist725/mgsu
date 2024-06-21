package util

import (
	"fmt"
	"testing"
)

func TestHitProbability(t *testing.T) {
	if HitProbability(128) ||
		HitProbability(0) ||
		!HitProbability(100) {
		t.Error("unexpected result")
	}
	count := 0
	for i := 0; i < 10000; i++ {
		if HitProbabilityThousands(uint16(16)) {
			count++
		}
	}
	t.Log(count)
}

func BenchmarkHitProbability(b *testing.B) {
	var s, f uint64
	for i := 0; i < b.N; i++ {
		if HitProbability(50) {
			s++
		} else {
			f++
		}

	}
	fmt.Println("s:......", s)
	fmt.Println("f:......", f)
}

func TestRandNoRepeat(t *testing.T) {
	t.Log(RandArrElemNoRepeat([]int{0, 1, 2, 3, 4, 5, 6}, -100))
	t.Log(RandArrElemNoRepeat([]int{0, 1, 2, 3, 4, 5, 6}, -1))
	t.Log(RandArrElemNoRepeat([]int{0, 1, 2, 3, 4, 5, 6}, 0))
	t.Log(RandArrElemNoRepeat([]int{0, 1, 2, 3, 4, 5, 6}, 1))
	t.Log(RandArrElemNoRepeat([]int{0, 1, 2, 3, 4, 5, 6}, 2))
	t.Log(RandArrElemNoRepeat([]int{0, 1, 2, 3, 4, 5, 6}, 3))
	t.Log(RandArrElemNoRepeat([]int{0, 1, 2, 3, 4, 5, 6}, 7))
	t.Log(RandArrElemNoRepeat([]int{0, 1, 2, 3, 4, 5, 6}, 99))
}
