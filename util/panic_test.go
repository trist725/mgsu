package util

import (
	"testing"
)

func TestRewriteStderrFile(t *testing.T) {
	RewriteStderrFile("./panic_test.log")
	panic("xxx")
}
