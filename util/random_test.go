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
