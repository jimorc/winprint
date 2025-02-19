//go:build windows

package wingdi

import (
	"syscall"
	"unsafe"
)

var (
	modgdi32 = syscall.NewLazyDLL("gdi32.dll")

	procCreateDC  = modgdi32.NewProc("CreateDCW")
	procEndDoc    = modgdi32.NewProc("EndDoc")
	procEndPage   = modgdi32.NewProc("EndPage")
	procStartDoc  = modgdi32.NewProc("StartDocW")
	procStartPage = modgdi32.NewProc("StartPage")
)

func CreateDC(driver string, printerName string, devMode *PrinterDevMode) uintptr {
	dr, _ := syscall.UTF16FromString(driver)
	pName, _ := syscall.UTF16FromString(printerName)
	r1, _, _ := procCreateDC.Call(uintptr(unsafe.Pointer(&dr[0])),
		uintptr(unsafe.Pointer(&pName[0])),
		0,
		uintptr(unsafe.Pointer(devMode)))
	return r1
}

func EndDoc(handle uintptr) bool {
	r1, _, _ := procEndDoc.Call(handle)
	return r1 > 0
}

func EndPage(handle uintptr) bool {
	r1, _, _ := procEndPage.Call(handle)
	return r1 > 0
}

func StartDoc(handle uintptr, docInfo *DocInfo) uintptr {
	r1, _, _ := procStartDoc.Call(
		handle,
		uintptr(unsafe.Pointer(docInfo)))
	return r1
}

func StartPage(handle uintptr) bool {
	r1, _, _ := procStartPage.Call(handle)
	return r1 > 0
}
