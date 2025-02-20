//go:build windows

package wingdi

import (
	"syscall"
	"unsafe"
)

type DocInfo1 struct {
	docName    uintptr
	outputFile uintptr
	dataType   uintptr
}

func NewDocInfo1(dName string) *DocInfo1 {
	n, _ := syscall.UTF16FromString(dName)
	raw, _ := syscall.UTF16FromString("RAW")
	return &DocInfo1{docName: uintptr(unsafe.Pointer(&n[0])),
		dataType: uintptr(unsafe.Pointer(&raw[0]))}
}
