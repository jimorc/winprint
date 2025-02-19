//go:build windows

package wingdi

import (
	"syscall"
	"unsafe"
)

type Printers struct {
	printers []Printer
}

func GetDefaultPrinter() (string, error) {
	buf := make([]uint16, 3)
	bufN := uint32(len(buf))
	err := getDefaultPrinter(&buf[0], &bufN)
	if err != syscall.ERROR_INSUFFICIENT_BUFFER {
		return "", err
	}
	buf = make([]uint16, bufN)
	err = getDefaultPrinter(&buf[0], &bufN)
	if err != syscall.Errno(0) {
		return "", err
	}
	return syscall.UTF16ToString(buf), nil
}

func GetPrintersInfo() ([]PrinterInfo2, error) {
	pInfo2, _, _ := getAllPrinter2Info()
	return pInfo2, nil
}

const maxUTF16Len = 1024

// StringFromUTF16 converts a uint16 pointer to a string.
// The maximum length of the UTF16 slice is maxUTF16Len as that should be more than enough for any
// UTF16 string returned from win32 printer functions.
func StringFromUTF16(utf16 *uint16) string {
	p := unsafe.Pointer(utf16)
	pSlice := (*[1 << 30]uint16)(p)[0:maxUTF16Len]

	return syscall.UTF16ToString(pSlice)

}

func getAllPrinter2Info() ([]PrinterInfo2, uint32, error) {
	flags := PRINTER_ENUM_LOCAL |
		PRINTER_ENUM_CONNECTIONS
		// buffer is a slice of bytes that will contain an array of PrinterInfo1 structs
	var buffer = make([]byte, 1)
	// size of *info1 in bytes
	var info2Size uint32 = 0
	// number of bytes needed in *info1
	var info2Needed uint32 = 0
	// number of info1 structs returned.
	var info2Count uint32 = 0

	_, err := EnumPrinters(flags,
		"",
		2,
		&buffer[0],
		info2Size,
		&info2Needed,
		&info2Count)

	if err != syscall.ERROR_INSUFFICIENT_BUFFER {
		return nil, 0, err
	}
	info2Size = info2Needed
	buffer = make([]byte, info2Size)
	_, err = EnumPrinters(flags,
		"",
		2,
		&buffer[0],
		info2Size,
		&info2Needed,
		&info2Count)
	if err != syscall.Errno(0) {
		return nil, 0, err
	}
	pInfo2 := (*[1024]PrinterInfo2)(unsafe.Pointer(&buffer[0]))[:info2Count:info2Count]
	return pInfo2, info2Count, nil
}
