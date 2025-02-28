//go:build windows

package wingdi

const (
	CCHDEVICENAME = 32
	CCHFORMNAME   = 32
)

// The following constants specify the PrinterDevMode fields in use. Some of the
// specified values are only valid for a display device, but are included here
// for completeness.
const (
	DM_ORIENTATION        = 0x1
	DM_PAPERSIZE          = 0x2
	DM_PAPERLENGTH        = 0x4
	DM_PAPERWIDTH         = 0x8
	DM_SCALE              = 0x10
	DM_POSITION           = 0x20
	DM_NUP                = 0x40
	DM_DISPLAYORIENTATION = 0x80
	DM_COPIES             = 0x100
	DM_DEFAULTSOURCE      = 0x200
	DM_PRINTQUALITY       = 0x400
	DM_COLOR              = 0x800
	DM_DUPLEX             = 0x1000
	DM_YRESOLUTION        = 0x2000
	DM_TTOPTION           = 0x400
	DM_COLLATE            = 0x8000
	DM_FORMNAME           = 0x10000
	DM_LOGPIXELS          = 0x20000
	DM_BITSPERPEL         = 0x40000
	DM_PELSWIDTH          = 0x80000
	DM_PELSHEIGHT         = 0x100000
	DM_DISPLAYFLAGS       = 0x200000
	DM_DISPLAYFREQUENCY   = 0x400000
	DM_ICMMETHOD          = 0x800000
	DM_ICMINTENT          = 0x1000000
	DM_MEDIATYPE          = 0x2000000
	DM_DITHERTYPE         = 0x4000000
	DM_PANNINGWIDTH       = 0x8000000
	DM_PANNINGHEIGHT      = 0x10000000
	DM_DISPLAYFIXEDOUTPUT = 0x20000000
)

// Flags for PrintDlg
const (
	PD_ALLPAGES                   uint32 = 0x0
	PD_SELECTION                  uint32 = 0x1
	PD_PAGENUMS                   uint32 = 0x2
	PD_NOSELECTION                uint32 = 0x4
	PD_NOPAGENUMS                 uint32 = 0x8
	PD_COLLATE                    uint32 = 0x10
	PD_PRINTTOFILE                uint32 = 0x20
	PD_PRINTSETUP                 uint32 = 0x40
	PD_NOWARNING                  uint32 = 0x80
	PD_RETURNDC                   uint32 = 0x100
	PD_RETURNIC                   uint32 = 0x200
	PD_RETURNDEFAULT              uint32 = 0x400
	PD_SHOWHELP                   uint32 = 0x800
	PD_ENABLEPRINTHOOK            uint32 = 0x1000
	PD_ENABLESETUPHOOK            uint32 = 0x2000
	PD_ENABLEPRINTTEMPLATE        uint32 = 0x4000
	PD_ENABLESETUPTEMPLATE        uint32 = 0x8000
	PD_ENABLEPRINTTEMPLATEHANDLE  uint32 = 0x10000
	PD_ENABLESETUPTEMPLATEHANDLE  uint32 = 0x20000
	PD_USEDEVMODECOPIES           uint32 = 0x40000
	PD_USEDEVMODECOPIESANDCOLLATE uint32 = 0x40000
	PD_DISABLEPRINTTOFILE         uint32 = 0x80000
	PD_HIDEPRINTTOFILE            uint32 = 0x100000
	PD_NONETWORKBUTTON            uint32 = 0x200000
	PD_CURRENTPAGE                uint32 = 0x400000
	PD_NOCURRENTPAGE              uint32 = 0x800000
	PD_EXCLUSIONFLAGS             uint32 = 0x1000000
	PD_USELARGETEMPLATE           uint32 = 0x1000000
	PD_EXCL_COPIESANDCOLLATE      uint32 = DM_COPIES | DM_COLLATE
	START_PAGE_GENERAL            uint32 = 0xffffffff
)

type pdResult int

// Return values from PrintDlgExW.
const (
	PD_RESULT_CANCEL pdResult = 0
	PD_RESULT_PRINT  pdResult = 1
	PD_RESULT_APPLY  pdResult = 2
)

const (
	SRCCOPY uint32 = 0x00CC0020
)
