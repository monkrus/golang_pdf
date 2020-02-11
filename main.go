package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	pdf.SetFont("arial", "B", 38)
	_, lineHt := pdf.GetFontSize()
	
	pdf.SetTextColor(255, 0, 0)
	pdf.Text(0, lineHt, "PPPP")


	pdf.SetFont("times", "", 16)
	pdf.MoveTo(0, lineHt*2)


	err := pdf.OutputFileAndClose("sergei.pdf")
	if err != nil {
		panic(err)
	}
}
