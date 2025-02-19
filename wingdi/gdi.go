//go:build windows

package wingdi

import (
	"syscall"
	"unsafe"
)

var (
	modgdi32 = syscall.NewLazyDLL("gdi32.dll")

	procStartDoc = modgdi32.NewProc(("StartDocW"))
)

func StartDoc(handle uintptr, docInfo *DocInfo) uintptr {
	r1, _, _ := procStartDoc.Call(
		handle,
		uintptr(unsafe.Pointer(docInfo)))
	return r1
}
