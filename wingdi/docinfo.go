//go:build windows

package wingdi

import (
	"syscall"
	"unsafe"
)

type DocInfo struct {
	size       int
	docName    uintptr
	outputFile uintptr
	dataType   uintptr
	fwType     uint32
}

func NewDocInfo(dName string) *DocInfo {
	n, _ := syscall.UTF16FromString(dName)
	return &DocInfo{docName: uintptr(unsafe.Pointer(&n[0]))}
}
