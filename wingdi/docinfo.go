//go:build windows

package wingdi

import (
	"syscall"
	"unsafe"
)

type DocInfo struct {
	size       uint32
	docName    uintptr
	outputFile uintptr
	dataType   uintptr
	fwType     uint32
}

func NewDocInfo(dName string) *DocInfo {
	n, _ := syscall.UTF16FromString(dName)
	raw, _ := syscall.UTF16FromString("RAW")
	return &DocInfo{
		size:     uint32(unsafe.Sizeof(DocInfo{})),
		docName:  uintptr(unsafe.Pointer(&n[0])),
		dataType: uintptr(unsafe.Pointer(&raw[0])),
	}
}

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
