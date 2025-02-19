//go:build windows

package wingdi

import (
	"syscall"
)

type PrinterDefaults struct {
	dataType      *uint16
	devMode       *PrinterDevMode
	desiredAccess uint32
}

func NewPrinterDefaults(dataType string, devMode *PrinterDevMode, accessMask uint32) *PrinterDefaults {
	t, _ := syscall.UTF16FromString(dataType)
	return &PrinterDefaults{
		dataType:      &t[0],
		devMode:       devMode,
		desiredAccess: accessMask,
	}
}
