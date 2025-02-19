//go:build windows

package wingdi

import (
	"syscall"
	"unsafe"
)

var (
	modcomdlg32 = syscall.NewLazyDLL("comdlg32.dll")

	procPrintDlg = modcomdlg32.NewProc("PrintDlgExW")
)

func PrintDlgEx(printDlg *PrintDlgExW) (pdResult, error) {
	r1, _, err := procPrintDlg.Call(uintptr(unsafe.Pointer(printDlg)))
	if err == syscall.Errno(0) {
		return pdResult(r1), nil
	}
	return PD_RESULT_CANCEL, err
}
