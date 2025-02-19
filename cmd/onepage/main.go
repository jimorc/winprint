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

	printerDefs := wingdi.NewPrinterDefaults(pInfo2s[0].DataType(), pInfo2s[0].DevMode, wingdi.PRINTER_ACCESS_USE)
	handle, _ := wingdi.OpenPrinter(defPrinter, printerDefs)
	fmt.Printf("Default printer handle: %x\n", handle)

	docInfo := wingdi.NewDocInfo("Test Print Doc")
	printJob := wingdi.StartDoc(handle, docInfo)
	fmt.Printf("Print Job is %d\n", printJob)

	err := wingdi.ClosePrinter(handle)

	fmt.Printf("Result of ClosePrinter: %s\n", err.Error())
}
