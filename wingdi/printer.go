//go:build windows

package wingdi

type Printer struct {
	handle          uintptr
	printerDefaults *PrinterDefaults
}

func NewPrinter(prName string) *Printer {
	//	_, _ = openPrinter(prName)
	return nil
}
