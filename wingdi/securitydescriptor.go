//go:build windows

package wingdi

type SecurityDescriptor struct {
	revision byte
	sbz1     byte
	control  uint32
	owner    uintptr
	group    uintptr
	sacl     uintptr
	dacl     uintptr
}
