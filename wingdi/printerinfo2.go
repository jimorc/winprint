//go:build windows

package wingdi

import (
	"fmt"
)

type PrinterInfo2 struct {
	serverName      *uint16
	printerName     *uint16
	shareName       *uint16
	portName        *uint16
	driverName      *uint16
	comment         *uint16
	location        *uint16
	DevMode         *PrinterDevMode
	sepFile         *uint16
	printProcessor  *uint16
	dataType        *uint16
	parameters      *uint16
	secDescriptor   *SecurityDescriptor
	attributes      uint32
	priority        uint32
	defaultPriority uint32
	startTime       uint32
	untilTime       uint32
	status          uint32
	cJobs           uint32
	averagePPMs     uint32
}

func (pi2 *PrinterInfo2) DataType() string {
	return StringFromUTF16(pi2.dataType)
}

func (pi2 *PrinterInfo2) Name() string {
	return StringFromUTF16(pi2.printerName)
}

func (pi2 *PrinterInfo2) ServerName() string {
	if pi2.serverName == nil {
		return "nil"
	} else {
		return StringFromUTF16(pi2.serverName)
	}
}

func (pi2 *PrinterInfo2) Print() {
	fmt.Println("PrinterInfo2")
	fmt.Printf("    Server Name: %s\n", pi2.ServerName())
	fmt.Printf("    Printer Name: %s\n", pi2.Name())
	fmt.Printf("    Data Type: %s\n", pi2.DataType())
	pi2.DevMode.Print()
}
