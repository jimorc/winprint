//go:build windows

package wingdi

import (
	"unsafe"
)

type PrintPageRange struct {
	nFromPage uint32
	mToPage   uint32
}

type PrintDlgExW struct {
	lStructSize         uintptr
	hwndOwner           uintptr
	hDevMode            *PrinterDevMode
	hDevNames           uintptr
	hDC                 uintptr
	Flags               uint32
	Flags2              uint32
	ExclusionFlags      uint32
	nPageRanges         uint32
	nMaxPageRanges      uint32
	lpPageRanges        uintptr
	nFromPage           uint32
	nMinPage            uint32
	nMaxPage            uint32
	nCopies             uint32
	hInstance           uintptr
	lpPrintTemplateName uintptr
	lpCallback          uintptr
	lphPropertyPages    uintptr
	nStartPage          uint32
	dwResultAction      uint32
}

func (pd *PrintDlgExW) DevMode() *PrinterDevMode {
	p := uintptr(unsafe.Pointer(pd.hDevMode))
	GlobalLock(p)
	dm := CopyDM((*PrinterDevMode)(pd.hDevMode))
	GlobalFree(p)

	return dm
}

func NewPrinteDlgExW(flags uint32) *PrintDlgExW {
	pd := &PrintDlgExW{Flags: flags}
	pd.lStructSize = unsafe.Sizeof(*pd)
	return pd
}
