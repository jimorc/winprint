//go:build windows

package wingdi

import (
	"fmt"
	"unicode/utf8"
)

type Bitmap struct {
	bmType       uint32
	bmWidth      int32
	bmHeight     int32
	bmWidthBytes int32
	bmPlanes     uint16
	bmBitsPixel  uint16
	bmBits       *[]byte
}

func (bm *Bitmap) Print() {
	fmt.Println("Bitmap:")
	fmt.Printf("    bmType: %d\n", bm.bmType)
	fmt.Printf("    bmWidth: %d\n", bm.bmWidth)
	fmt.Printf("    bmHeight: %d\n", bm.bmHeight)
	fmt.Printf("    bmWidthBytes: %d\n", bm.bmWidthBytes)
	fmt.Printf("    bmPlanes: %d\n", bm.bmPlanes)
	fmt.Printf("    bmBitsPixel: %d\n", bm.bmBitsPixel)
}

func DecodeGoBMP(bmp *[]byte) (colorPlane uint16, bpp uint16, pixOffset uint32) {
	var sigBM string
	b := (*bmp)[0:2]
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		sigBM += string(r)

		b = b[size:]
	}
	fileSize := readUint32((*bmp)[2:6])
	pixOffset = readUint32((*bmp)[10:14])
	dibHeaderSize := readUint32((*bmp)[14:18])
	width := readUint32((*bmp)[18:22])
	height := readUint32((*bmp)[22:26])
	colorPlane = readUint16((*bmp)[26:28])
	bpp = readUint16((*bmp)[28:30])
	compression := readUint32((*bmp)[30:34])
	imageSize := readUint32((*bmp)[34:38])
	xPixelsPerMeter := readUint32((*bmp)[38:42])
	yPixelsPerMeter := readUint32((*bmp)[42:46])
	colorUse := readUint32((*bmp)[46:50])
	colorImportant := readUint32((*bmp)[50:54])

	fmt.Printf("sigBM: %s\n", sigBM)
	fmt.Printf("fileSize: %d\n", fileSize)
	fmt.Printf("pixOffset: %d\n", pixOffset)
	fmt.Printf("dibHeaderSize: %d\n", dibHeaderSize)
	fmt.Printf("width: %d\n", width)
	fmt.Printf("height: %d\n", height)
	fmt.Printf("colorPlane: %d\n", colorPlane)
	fmt.Printf("bpp: %d\n", bpp)
	fmt.Printf("compression: %d\n", compression)
	fmt.Printf("imageSize: %d\n", imageSize)
	fmt.Printf("xPixelsPerMeter: %d\n", xPixelsPerMeter)
	fmt.Printf("yPixelsPerMeter: %d\n", yPixelsPerMeter)
	fmt.Printf("colorUse: %d\n", colorUse)
	fmt.Printf("colorImportant: %d\n", colorImportant)

	return
}

// readUint16 and readUint32 are copied from image/bmp/reader.go
func readUint16(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

func readUint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}
