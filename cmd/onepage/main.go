//go:build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/jimorc/winprint/wingdi"
)

func main() {
	pInfo2s, _ := wingdi.GetPrintersInfo()
	for _, pi2 := range pInfo2s {
		pi2.Print()
		fmt.Println()
	}

	defPrinter, _ := wingdi.GetDefaultPrinter()
	fmt.Printf("Default printer: %s\n", defPrinter)

	printerDefs := wingdi.NewPrinterDefaults("RAW", pInfo2s[0].DevMode, wingdi.PRINTER_ACCESS_USE)
	handle, _ := wingdi.OpenPrinter(defPrinter, printerDefs)
	fmt.Printf("Default printer handle: %x\n", handle)

	docInfo := wingdi.NewDocInfo1("Test Print Doc")

	printJob := wingdi.StartDocPrinter(handle, docInfo)
	fmt.Printf("Print Job is %d\n", printJob)

	startPagePrinterOk := wingdi.StartPagePrinter(handle)
	fmt.Printf("StartPagePrinter status: %t\n", startPagePrinterOk)

	text, _ := syscall.UTF16FromString("This is a test string")
	buf := uintptr(unsafe.Pointer(&text[0]))

	written, ok := wingdi.WritePrinter(handle, buf, 2*len(text))
	fmt.Printf("WritePrinter: chars written: %d, write OK: %t\n", written, ok)

	endPagePrinterOk := wingdi.EndPagePrinter(handle)
	fmt.Printf("EndPagePrinter status: %t\n", endPagePrinterOk)

	endDocPrinterOk := wingdi.EndDocPrinter(handle)
	fmt.Printf("EndDocPrinter status: %t\n", endDocPrinterOk)

	err := wingdi.ClosePrinter(handle)

	fmt.Printf("Result of ClosePrinter: %s\n", err.Error())
}
