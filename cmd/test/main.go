/*
package main

import (

	"fmt"

	"github.com/alexbrainman/printer"

)

	func main() {
		defaultPrinterName, _ := printer.Default()
		fmt.Println(defaultPrinterName)
		p, err := printer.Open(defaultPrinterName)

		if err != nil {
			fmt.Printf("Open failed: %v\n", err)
		}

		defer p.Close()

		err = p.StartDocument("my document", "RAW")

		if err != nil {
			fmt.Printf("StartDocument failed: %v\n", err)
		}

		defer p.EndDocument()

		err = p.StartPage()

		if err != nil {
			fmt.Printf("StartPage failed: %v\n", err)
		}

		str := "testing 123"
		mySlice := []byte(str)

		_, err = p.Write(mySlice)

		if err != nil {
			fmt.Printf("Write failed: %v\n", err)
		}

		err = p.EndPage()

		if err != nil {
			fmt.Printf("EndPage failed: %v\n", err)
		}
	}
*/
package main

import (
	"fmt"
	"os"

	"github.com/Juicymo/printer"
)

func main() {
	defaultPrinterName, err := printer.Default()
	if err != nil {
		fmt.Printf("Failed to get default printer: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Default printer:", defaultPrinterName)

	p, err := printer.Open(defaultPrinterName)
	if err != nil {
		fmt.Printf("Open failed: %v\n", err)
		os.Exit(1)
	}
	defer p.Close()

	err = p.StartRawDocument("my document")
	if err != nil {
		fmt.Printf("StartDocument failed: %v\n", err)
		os.Exit(1)
	}
	defer p.EndDocument()

	err = p.StartPage()
	if err != nil {
		fmt.Printf("StartPage failed: %v\n", err)
		os.Exit(1)
	}
	defer p.EndPage()

	str := "testing 123\nThis is a test print\nAnother line of text\n" // Add some line breaks for better readability
	mySlice := []byte(str)

	_, err = p.Write(mySlice)
	if err != nil {
		fmt.Printf("Write failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Text sent to printer successfully.")
}
