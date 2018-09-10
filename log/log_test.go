package log

import (
	"testing"
)

func Test(t *testing.T) {
	l := New(".", "test")
	defer l.Close()

	l.SetLevel(LevelInformational)

	l.Debug("hello debug")
	l.Info("hello info")
	l.Warn("hello warn")
	l.Error("hello error")
}

func TestConsoleLogger(t *testing.T) {
	l := NewConsoleLogger()
	defer l.Close()

	l.SetLevel(LevelInformational)

	l.Debug("hello debug")
	l.Info("hello info")
	l.Warn("hello warn")
	l.Error("hello error")
}
