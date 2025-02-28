//go:build windows

package main

import (
	"bytes"
	"fmt"
	"os"

	"image/color"
	"io"

	//	"github.com/sunshineplan/imgconv"
	//	"github.com/sunshineplan/pdf"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/software"
	"github.com/jimorc/winprint/wingdi"
	"github.com/shahfarhadreza/go-gdiplus"
	"golang.org/x/image/bmp"
)

func main() {
	gdiIn := &gdiplus.GdiplusStartupInput{GdiplusVersion: 1}

	status := gdiplus.GdiplusStartup(gdiIn, nil)
	if status != 0 {
		fmt.Printf("Failed to initialize GDI+: %s\n", status.String())
		os.Exit(1)
	}
	fmt.Println("GDI+ initialized")
	defer gdiplus.GdiplusShutdown()

	a1 := app.New()

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

	pDC := wingdi.CreateDC("", defPrinter, pInfo2s[1].DevMode)
	fmt.Printf("pDC: %d, %d\n", pDC, uint32(pDC))
	docInfo := wingdi.NewDocInfo("Test Print Doc")
	printJob := wingdi.StartDoc(pDC, docInfo)
	//	printJob := wingdi.StartDocPrinter(handle, docInfo1)
	fmt.Printf("PrintJob: %t %d\n", int32(printJob) > 0, int(printJob))

	ok := wingdi.StartPage(pDC)
	fmt.Printf("StartPage status: %t\n", ok)

	var buf bytes.Buffer
	_ = io.Writer(&buf)

	circle := canvas.NewCircle(color.Gray16{Y: 10})
	line := canvas.NewLine(color.Black)
	circle.FillColor = color.Gray{Y: 0xC0}
	circle.StrokeColor = color.Black
	circle.StrokeWidth = 10
	circle.Resize(fyne.NewSize(50, 50))
	circle.Move(fyne.NewPos(100, 100))

	line.StrokeWidth = 10
	line.Position1 = fyne.NewPos(0, 0)
	line.Position2 = fyne.NewPos(120, 60)

	text := canvas.NewText("Test string", color.RGBA{255, 0, 0, 255})
	text.Move(fyne.NewPos(200, 200))

	c3 := container.New(wingdi.NewPrintPageLayout(), line, circle, text)
	c3.Resize(fyne.NewSize(1000, 1000))

	i := software.Render(c3, fyne.CurrentApp().Settings().Theme())
	img := canvas.NewImageFromImage(i)
	win := a1.NewWindow("Test")
	win.SetContent(img)
	win.Resize(fyne.NewSize(1000, 1000))
	win.ShowAndRun()

	bmp.Encode(&buf, i)
	bmp := buf.Bytes()
	fmt.Printf("bufLen = %d\n", len(bmp))
	nPlanes, bpp, offset := wingdi.DecodeGoBMP(&bmp)
	fmt.Printf("nPlanes: %d, bpp: %d, offset: %d\n", uint32(nPlanes), uint32(bpp), offset)

	g := gdiplus.NewGraphicsFromHDC(gdiplus.HDC(pDC))
	fmt.Printf("Graphics: %v\n", g)
	bitmap := gdiplus.NewBitmapEx(1000, 1000, -3000, gdiplus.PixelFormat24bppRGB, &bmp[len(bmp)-3-3000])
	fmt.Printf("NewBitmapFromHBITMAP: %v\n", bitmap)
	g.DrawImage(&bitmap.Image, 0, 0)

	ok = wingdi.EndPage(pDC)
	fmt.Printf("EndPage status: %t\n", ok)

	ok = wingdi.EndDoc(pDC)
	fmt.Printf("EndDoc Status: %t\n", ok)

	err := wingdi.ClosePrinter(handle)
	fmt.Printf("Result of ClosePrinter: %s\n", err.Error())
}
