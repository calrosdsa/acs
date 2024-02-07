package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	styleId, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
	})
	sheet := "Sheet1"

	// Insert a picture.
	f.SetColWidth(sheet, "B", "B", 13)
	f.SetColWidth(sheet, "C", "C", 25)
	if err := f.SetRowHeight(sheet, 1, 36); err != nil {
		fmt.Println(err)
		return
	}

	f.MergeCell(sheet, "B1", "C1")

	if err = f.SetCellStyle(sheet, "B1", "C1", styleId); err != nil {
		return
	}

	if err := f.AddPicture(sheet, "B1", "./app/media/logo.png", &excelize.GraphicOptions{
		AutoFit: true,
		OffsetY: 10,
		ScaleX:  1,
		ScaleY:  1,
	}); err != nil {
		fmt.Println(err)
		return
	}
	// Insert a picture scaling in the cell with location hyperlink.
	if err := f.SaveAs("./app/media/template.xlsx"); err != nil {
		fmt.Println(err)
	}
}
