package util

import "testing"

func TestHitProbability(t *testing.T) {
	if HitProbability(128) ||
		HitProbability(0) ||
		!HitProbability(100) {
		t.Error("unexpected result")
	}
}
