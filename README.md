# winprint
Repository for testing printing from Go with fyne on Windows 11

This repository contains code for determining how to print on Windows 11 from Go, possibly using the fyne GUI system.

There are three programs in this repository:
* onepage, which despite its name, prints 2 pages using GDI and GDI+. I could not get it to work using GDI only as I could not
get Bitblt to work.
* pdf, which creates a pdf document and attempts to print it. While the document is sent to the printer, it does not print.
* test, which attempts to print using raw text using the winspool API. This also sends the text to the printer, but it does not print.
The program is a slight modification to a [program](https://stackoverflow.com/questions/27917857/printing-via-winspool) found on StackOverflow.

## Warning

This is test code only. Once I have created an alpha version of fyne-print, I will be deleting this repository.


