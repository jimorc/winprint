//go:build windows

package main

import (
	"fmt"

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
	/*	printJob := wingdi.StartDocPrinter(handle, docInfo)
		//	fmt.Printf("Print Job is %d\n", printJob)

		startPageOk := wingdi.StartPage(handle)
		fmt.Printf("StartPage OK: %t\n", startPageOk)

		printerDC := wingdi.CreateDC("WINSPOOL", pInfo2s[0].Name(), pInfo2s[0].DevMode)
		fmt.Printf("printerDC: %d\n", printerDC)

		//			elOk := wingdi.Ellipse(printerDC, 0, 0, 2000, 2500)
		// 			fmt.Printf("Ellipse OK: %t\n", elOk)

		text, _ := syscall.UTF16FromString("This is a test string")
		buf := uintptr(unsafe.Pointer(&text[0]))

		written, ok := wingdi.WritePrinter(handle, buf, 2*len(text))
		fmt.Printf("WritePrinter: chars written: %d, write OK: %t\n", written, ok)

		endPageOk := wingdi.EndPage(handle)
		fmt.Printf("EndPage OK: %t\n", endPageOk)

		endDocOk := wingdi.EndDoc(handle)
		fmt.Printf("EndDoc OK: %t\n", endDocOk)
	*/
	err := wingdi.ClosePrinter(handle)

	fmt.Printf("Result of ClosePrinter: %s\n", err.Error())
}
