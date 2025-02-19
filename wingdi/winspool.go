//go:build windows

package wingdi

import (
	"syscall"
	"unsafe"
)

const (
	PRINTER_ENUM_DEFAULT     uint32 = 0x1
	PRINTER_ENUM_LOCAL       uint32 = 0x2
	PRINTER_ENUM_CONNECTIONS uint32 = 0x4
	PRINTER_ENUM_FAVORITE    uint32 = 0x4
	PRINTER_ENUM_NAME        uint32 = 0x8
	PRINTER_ENUM_REMOTE      uint32 = 0x10
	PRINTER_ENUM_SHARED      uint32 = 0x20
	PRINTER_ENUM_NETWORK     uint32 = 0x40
)

const (
	PRINTER_ACCESS_ADMINISTER uint32 = 0x4
	PRINTER_ACCESS_USE        uint32 = 0x8
)

var (
	modwinspool = syscall.NewLazyDLL("winspool.drv")

	procEnumPrinters      = modwinspool.NewProc("EnumPrintersW")
	procGetDefaultPrinter = modwinspool.NewProc("GetDefaultPrinterW")
	procOpenPrinter       = modwinspool.NewProc("OpenPrinterW")
)

func EnumPrinters(flags uint32,
	name string,
	level uint32,
	buf *byte,
	cbBuf uint32,
	needed *uint32,
	cReturned *uint32) (uintptr, error) {
	n, _ := syscall.UTF16FromString(name)
	r1, _, err := procEnumPrinters.Call(
		uintptr(flags),
		uintptr(unsafe.Pointer(&n[0])),
		uintptr(level),
		uintptr(unsafe.Pointer(buf)),
		uintptr(cbBuf),
		uintptr(unsafe.Pointer(needed)),
		uintptr(unsafe.Pointer(cReturned)))
	return r1, err
}

func OpenPrinter(prName string, defaults *PrinterDefaults) (uintptr, error) {
	var handle uint32 = 0
	name, _ := syscall.UTF16FromString(prName)
	r1, _, err := procOpenPrinter.Call(
		uintptr(unsafe.Pointer(&name[0])),
		uintptr(unsafe.Pointer(&handle)),
		uintptr(unsafe.Pointer(defaults)))
	if r1 == 0 {
		return 0, err
	}
	return uintptr(handle), err
}

func getDefaultPrinter(buf *uint16, bufN *uint32) error {
	_, _, err := procGetDefaultPrinter.Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(unsafe.Pointer(bufN)))
	return err
}
