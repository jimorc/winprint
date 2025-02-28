//go:build windows

package wingdi

type win32Bool uint32

const (
	False win32Bool = 0
	True  win32Bool = 1
)

type gdiplusStartupInput struct {
	gdiplusVersion           uint32
	debugEvtCB               debugEventProc
	suppressBackgroundThread win32Bool
	suppressExternalCodecs   win32Bool
}

func NewGdiplusStartupInput(debugEventCallback debugEventProc, suppressBGThread win32Bool, suppressCodecs win32Bool) *gdiplusStartupInput {
	return &gdiplusStartupInput{gdiplusVersion: 1,
		debugEvtCB:               debugEventCallback,
		suppressBackgroundThread: suppressBGThread,
		suppressExternalCodecs:   suppressCodecs,
	}
}

func NewArgumentFreeGdiplusStartupInput() *gdiplusStartupInput {
	//	noMsg, _ := syscall.UTF16FromString("")
	return &gdiplusStartupInput{gdiplusVersion: 1, debugEvtCB: debugEventProc{level: 0,
		//		msg: uintptr(unsafe.Pointer(&noMsg[0]))},
		msg: uintptr(0)},
		suppressBackgroundThread: False,
		suppressExternalCodecs:   False}
}

type debugEventProc struct {
	level uint32
	msg   uintptr
}

type GdiplusStartupOutput struct {
	notificationHook   uintptr
	notificationUnhook uintptr
}
