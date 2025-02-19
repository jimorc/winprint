//go:build windows

package wingdi

import (
	"syscall"
	"unsafe"
)

var (
	modgdi32 = syscall.NewLazyDLL("gdi32.dll")

	procEndDoc   = modgdi32.NewProc("EndDoc")
	procStartDoc = modgdi32.NewProc("StartDocW")
)

func EndDoc(handle uintptr) bool {
	r1, _, _ := procEndDoc.Call(handle)
	return r1 > 0
}

func StartDoc(handle uintptr, docInfo *DocInfo) uintptr {
	r1, _, _ := procStartDoc.Call(
		handle,
		uintptr(unsafe.Pointer(docInfo)))
	return r1
}
