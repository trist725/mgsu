package util

import (
	"bufio"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"strings"
	"syscall"
)

func MakeSignalChannel(sig ...os.Signal) chan os.Signal {
	ch := make(chan os.Signal)
	signal.Notify(ch, sig...)
	return ch
}

func WaitExitSignal() os.Signal {
	ch := MakeSignalChannel(syscall.SIGINT, syscall.SIGTERM)
	return <-ch
}

func ReadLineFromConsole() (string, error) {
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimRight(text, "\n\r")
	return text, nil
}

// GetTypeName 获取接口中存放的实例的类型名
func GetTypeName(i interface{}) string {
	if i == nil {
		return ""
	}
	rt := reflect.TypeOf(i)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	return rt.Name()
}

// GetFuncName 获取当前函数名
// layer: 函数调用层数
func GetFuncName(layer int) string {
	pc, _, _, _ := runtime.Caller(layer)
	name := runtime.FuncForPC(pc).Name()
	split := strings.Split(name, ".")
	if layer == 0 {
		return name
	} else {
		return split[len(split)-1]
	}
}
