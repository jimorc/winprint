//go:build windows

package main

import (
	"bytes"
	"fmt"
	"os"

	"image/color"

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

	printerDefs := wingdi.NewPrinterDefaults("RAW", pInfo2s[2].DevMode, wingdi.PRINTER_ACCESS_USE)
	//	prNames, _ := printer.ReadNames()
	//	fmt.Printf("Printers: %v\n", prNames)
	pr, err := wingdi.OpenPrinter(defPrinter, printerDefs)
	fmt.Printf("Default printer:: %v\n", defPrinter)
	if pr == 0 {
		fmt.Printf("Open err: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Printer handle: 0x%x\n", pr)
	//	defer pr.Close()
	dInfo := wingdi.NewDocInfo(defPrinter)
	printJob := int32(wingdi.StartDoc(pr, dInfo))
	if printJob <= 0 {
		fmt.Println("StartDoc failed")
		os.Exit(1)
	}
	fmt.Printf("PrintJob is %d\n", printJob)
	//	defer pr.EndDocument()

	pDC := wingdi.CreateDC("", defPrinter, pInfo2s[1].DevMode)
	fmt.Printf("pDC: %d, %d\n", pDC, uint32(pDC))
	//	docInfo1 := wingdi.NewDocInfo1("Test Print Doc")
	//	printJob := wingdi.StartDocPrinter(pDC, docInfo1)
	/*	printJob, err := handle.StartDoc("Test Print Doc")
		if err != nil {
			fmt.Printf("StartDoc err: %s\n", err.Error())
		}
		fmt.Printf("PrintJob: %d\n", int(printJob))
	*/
	//	ok := wingdi.StartPagePrinter(handle)
	//	fmt.Printf("StartPagePrinter status: %t\n", ok)

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

	var buf bytes.Buffer
	bmp.Encode(&buf, i)
	bmp := buf.Bytes()
	fmt.Printf("bmpLen = %d\n", len(bmp))
	/*
		pdf := fpdf.New("P", "in", "Letter", "")
		pdf.AddPage()
		opts := fpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}
		pdf.RegisterImageOptionsReader("page1", opts, &buf)
		pdf.ImageOptions("page1", 0, 0, 0, 0, false, opts, 0, "")
		var out bytes.Buffer
		//	out := outBuf.Bytes()
		//	pdf.OutputFileAndClose("test.pdf")
		//	_ = io.Writer(&out)
		err = pdf.Output(&out)
		if err != nil {
			fmt.Printf("pdf.Output err: %s\n", err.Error())
		}*/
	nPlanes, bpp, offset := wingdi.DecodeGoBMP(&bmp)
	fmt.Printf("nPlanes: %d, bpp: %d, offset: %d\n", uint32(nPlanes), uint32(bpp), offset)

	wingdi.StartPage(pDC)
	g := gdiplus.NewGraphicsFromHDC(gdiplus.HDC(pDC))
	fmt.Printf("Graphics: %v\n", g)
	bitmap := gdiplus.NewBitmapEx(1000, 1000, -3000, gdiplus.PixelFormat24bppRGB, &bmp[len(bmp)-3-3000])
	fmt.Printf("NewBitmapFromHBITMAP: %v\n", bitmap)
	g.DrawImage(&bitmap.Image, 0, 0)
	wingdi.EndPage(pDC)
	wingdi.StartPage(pDC)
	g.DrawImage(&bitmap.Image, 100, 100)
	wingdi.EndPage(pDC)

	/*
		outBuf := out.Bytes()
		len := len(outBuf)
		c, ok := wingdi.WritePrinter(pr, uintptr(unsafe.Pointer(&outBuf[0])), len)
		if !ok {
			fmt.Println("Write failed")
		}
		fmt.Printf("%d bytes written to printer\n", c)
	*/
	//	count, ok := wingdi.WritePrinter(handle, uintptr(unsafe.Pointer(&outBuf)), len(outBuf))
	/*	err = pr.StartPage()
			if err != nil {
				fmt.Printf("StartPage err: %s\n", err.Error())
			}
			str := "This is a test"
			bstr := []byte(str)
			//	bStr, _ := syscall.UTF16FromString(str)
			count, err := fmt.Fprintln(pr, bstr)
			if err != nil {
				fmt.Printf("Write err: %s\n", err.Error())
			}
			fmt.Printf("Write count: %d\n", count)
			err = pr.EndPage()
			if err != nil {
				fmt.Printf("EndPage err: %s\n", err.Error())
			}
		//	ok = wingdi.EndPagePrinter(handle)
		//	fmt.Printf("EndPagePrinter status: %t\n", ok)*/
	ok := wingdi.EndDoc(pr)
	if !ok {
		fmt.Println("EndDoc failed")
	}

	err = wingdi.ClosePrinter(pr)
	if err != nil {
		fmt.Printf("Result of ClosePrinter: %s\n", err.Error())
	}
}
