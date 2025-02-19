//go:build windows

package wingdi

import (
	"syscall"
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGlobalFree = modkernel32.NewProc("GlobalFree")
	procGlobalLock = modkernel32.NewProc("GlobalLock")
)

func GlobalLock(memLocation uintptr) (uintptr, error) {
	r1, _, err := procGlobalLock.Call(memLocation)
	if err == syscall.Errno(0) {
		return r1, nil
	}
	return uintptr(0), err
}

func GlobalFree(memLocation uintptr) (uintptr, error) {
	r1, _, err := procGlobalFree.Call(memLocation)
	if err == syscall.Errno(0) {
		return uintptr(0), nil
	}
	return r1, err
}
