//go:build windows

package wingdi

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modgdi32 = syscall.NewLazyDLL("gdi32.dll")

	procBitblt                 = modgdi32.NewProc("BitBlt")
	procCreateBitmap           = modgdi32.NewProc("CreateBitmap")
	procCreateBitmapIndirect   = modgdi32.NewProc(("CreateBitmapIndirect"))
	procCreateCompatibleBitmap = modgdi32.NewProc("CreateCompatibleBitmap")
	procCreateCompatibleDC     = modgdi32.NewProc("CreateCompatibleDC")
	procCreateDC               = modgdi32.NewProc("CreateDCW")
	procEllipse                = modgdi32.NewProc("Ellipse")
	procEndDoc                 = modgdi32.NewProc("EndDoc")
	procEndPage                = modgdi32.NewProc("EndPage")
	procGetObject              = modgdi32.NewProc("GetObjectW")
	procSelectObject           = modgdi32.NewProc("SelectObject")
	procStartDoc               = modgdi32.NewProc("StartDocW")
	procStartPage              = modgdi32.NewProc("StartPage")
)

func Bitblt(destDC uintptr, x, y, cx, cy int32, srcDC uintptr, x1, y1 int32, rop uint32) bool {
	r1, _, err := procBitblt.Call(destDC,
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		srcDC,
		uintptr(x1),
		uintptr(y1),
		uintptr(rop))
	fmt.Printf("Bitblt err: %s\n", err.Error())
	return r1 != 0
}

func CreateBitmap(w, h int, planes, bitcount uint32, b *byte) uintptr {
	r1, _, err := procCreateBitmap.Call(uintptr(w),
		uintptr(h),
		uintptr(planes),
		uintptr(bitcount),
		uintptr(unsafe.Pointer(b)))
	fmt.Printf("CreateBitmap: err: %s\n", err.Error())
	return r1
}

func CreateBitmapIndirect(bmp *[]byte) uintptr {
	b := uintptr(unsafe.Pointer(bmp))
	r1, _, err := procCreateBitmapIndirect.Call(b)
	fmt.Println(err.Error())
	return r1
}

func CreateCompatibleBitmap(hdc uintptr, cx, cy int) uintptr {
	r1, _, err := procCreateCompatibleBitmap.Call(hdc,
		uintptr(unsafe.Pointer(&cx)),
		uintptr(unsafe.Pointer(&cy)))
	fmt.Printf("CreateCompatibleBitmap err: %s\n", err.Error())
	return r1
}

func CreateCompatibleDC(dc uintptr) uintptr {
	r1, _, err := procCreateCompatibleDC.Call(dc)
	fmt.Printf("CreateCompatibleDC err: %s\n", err.Error())
	return r1
}

func CreateDC(driver string, printerName string, devMode *PrinterDevMode) uintptr {
	dr, _ := syscall.UTF16FromString(driver)
	pName, _ := syscall.UTF16FromString(printerName)
	r1, _, _ := procCreateDC.Call(uintptr(unsafe.Pointer(&dr[0])),
		uintptr(unsafe.Pointer(&pName[0])),
		0,
		uintptr(unsafe.Pointer(devMode)))
	return r1
}

func Ellipse(dc uintptr, left, top, right, bottom uint32) bool {
	r1, _, _ := procEllipse.Call(dc,
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))
	return r1 > 0
}

func EndDoc(dc uintptr) bool {
	r1, _, _ := procEndDoc.Call(dc)
	return r1 > 0
}

func EndPage(dc uintptr) bool {
	r1, _, _ := procEndPage.Call(dc)
	return r1 > 0
}

func GetObject(handle uintptr, objSize uint32, obj uintptr) int32 {
	r1, _, err := procGetObject.Call(handle,
		uintptr(objSize),
		obj)
	fmt.Printf("GetObject err: %s\n", err.Error())
	return int32(r1)
}

func SelectObject(dc uintptr, hObject uintptr) uintptr {
	r1, _, err := procSelectObject.Call(dc, hObject)
	fmt.Printf("SelectObject: err: %s\n", err.Error())
	return r1
}

func StartDoc(dc uintptr, docInfo *DocInfo) uintptr {
	r1, _, _ := procStartDoc.Call(
		dc,
		uintptr(unsafe.Pointer(docInfo)))
	return r1
}

func StartPage(dc uintptr) bool {
	r1, _, _ := procStartPage.Call(dc)
	return r1 > 0
}
