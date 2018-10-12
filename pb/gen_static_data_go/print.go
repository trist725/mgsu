package main

import "fmt"

func panicf(f string, args ...interface{}) {
	panic(fmt.Sprintf(f, args...))
}

func errorf(f string, args ...interface{}) {
	fmt.Printf("[error] "+f+"\n", args...)
}

func infof(f string, args ...interface{}) {
	fmt.Printf("[info] "+f+"\n", args...)
}

func debugf(f string, args ...interface{}) {
	if gCfg.Debug {
		fmt.Printf("[debug] "+f+"\n", args...)
	}
}
