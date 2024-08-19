//go:build windows
// +build windows

package util

import (
	"os"
	"runtime"
	"syscall"
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
	f                *os.File
)

func setStdHandle(stdhandle int32, handle syscall.Handle) error {
	r0, _, e1 := syscall.SyscallN(procSetStdHandle.Addr(), uintptr(stdhandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}

func RewriteStderrFile(path string) (err error) {
	f, err = os.Create(path)
	if err != nil {
		return err
	}
	err = setStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(f.Fd()))
	if err != nil {
		return err
	}
	runtime.SetFinalizer(f, func(fd *os.File) {
		fd.Close()
	})
	return nil
}
