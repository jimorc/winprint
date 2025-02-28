//go:build windows

package wingdi

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modgdiplus = syscall.NewLazyDLL("gdiplus.dll")

	procGdiplusStartup  = modgdiplus.NewProc("GdiplusStartup")
	procGdiplusShutdown = modgdiplus.NewProc("GdiplusShutdown")
)

func GdiplusStartup(token *uintptr, startupInput *gdiplusStartupInput, startupOuput *GdiplusStartupOutput) uint32 {
	r1, _, err := procGdiplusStartup.Call(uintptr(unsafe.Pointer(token)),
		uintptr(unsafe.Pointer(startupInput)),
		uintptr(unsafe.Pointer(startupOuput)))
	fmt.Printf("GdiplusStartup status: %d, err: %s\n", uint32(r1), err.Error())
	fmt.Printf("GdiplusStartup token: %d\n", *token)
	return uint32(r1)
}

func GdiplusShutdown(token uintptr) {
	fmt.Printf("GdiplusShutdown token: %d\n", token)
	r1, _, err := procGdiplusShutdown.Call(uintptr(unsafe.Pointer(token)))
	fmt.Printf("gidplusShutdown status: %d, err: %s\n", r1, err.Error())
}
