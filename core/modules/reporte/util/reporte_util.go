package util

import (
	_r "acs/domain/repository"
	// "log"

	"github.com/xuri/excelize/v2"
)

type reporteUtil struct {
	Palette   _r.ColorPalette
	TextSizes _r.TextSizes
	Locale _r.Locale
}

func New(palette _r.ColorPalette, textSizes _r.TextSizes,locale _r.Locale) _r.ReporteUtil {
	return &reporteUtil{
		Palette:   palette,
		TextSizes: textSizes,
		Locale: locale,
	}
}

func (r *reporteUtil) SetUpReporteLayout(sheet string, f *excelize.File) (err error) {
	var styleId int
	if styleId, err = r.GetBlankStyle(f); err != nil {
		return
	}

	f.SetColStyle(sheet, "A", styleId)
	f.SetColStyle(sheet, "B", styleId)
	f.SetColStyle(sheet, "C", styleId)
	f.SetColStyle(sheet, "D", styleId)
	f.SetColStyle(sheet, "E", styleId)
	f.SetColStyle(sheet, "F", styleId)
	f.SetColStyle(sheet, "G", styleId)
	f.SetColStyle(sheet, "H", styleId)
	f.SetColStyle(sheet, "I", styleId)
	f.SetColStyle(sheet, "J", styleId)
	f.SetColStyle(sheet, "K", styleId)

	return
}

func (r *reporteUtil) GetBlankStyle(f *excelize.File) (styleId int, err error) {
	styleId, err = f.NewStyle(&excelize.Style{
		Border: []excelize.Border{{Type: "top", Style: 1, Color: r.Palette.ColorWhite}, {Type: "left", Style: 1, Color: r.Palette.ColorWhite},
			{Type: "bottom", Style: 1, Color: r.Palette.ColorWhite}, {Type: "right", Style: 1, Color: r.Palette.ColorWhite}},
		// Border:    []excelize.Border{{Type: "Bottom", Style: 2, Color: "1f7f3b"}},
	})
	return
}

func (r *reporteUtil) GetTitleStyle(f *excelize.File) (styleId int, err error) {
	styleId, err = f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Color: r.Palette.ColorWhite, Bold: true, Family: "Arial", Size: r.TextSizes.Small},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{r.Palette.PrimaryColor}, Pattern: 1},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
		Border: []excelize.Border{{Type: "top", Style: 1, Color: r.Palette.BorderColor}, {Type: "left", Style: 1, Color: r.Palette.BorderColor},
			{Type: "bottom", Style: 1, Color: r.Palette.BorderColor}, {Type: "right", Style: 1, Color: r.Palette.BorderColor}},
		// Border:    []excelize.Border{{Type: "Bottom", Style: 2, Color: "1f7f3b"}},
	})
	return
}

func (r *reporteUtil) GetCommonCellStyle(f *excelize.File) (styleId int, err error) {
	styleId, err = f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Family: "Arial", Size: r.TextSizes.Small},
		// Fill:      excelize.Fill{Type: "pattern", Color: []string{r.Palette.PrimaryColor}, Pattern: 1},
		// Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
		Border: []excelize.Border{{Type: "top", Style: 1, Color: r.Palette.BorderColor}, {Type: "left", Style: 1, Color: r.Palette.BorderColor},
			{Type: "bottom", Style: 1, Color: r.Palette.BorderColor}, {Type: "right", Style: 1, Color: r.Palette.BorderColor}},
		// Border:    []excelize.Border{{Type: "Bottom", Style: 2, Color: "1f7f3b"}},
	})
	return
}

func (r *reporteUtil) SetUpHeader(sheet string, f *excelize.File, d _r.ReportInfo, titleStyle, cellStyle int,lang string) (err error) {

	cell, err := excelize.CoordinatesToCellName(2, 2)
	if err != nil {
		return
	}
	if err = f.SetCellStyle(sheet, "B2", "B2", titleStyle); err != nil {
		return
	}
	if err = f.SetCellStyle(sheet, "C2", "C2", cellStyle); err != nil {
		return
	}
	f.SetSheetRow(sheet, cell, &[]string{r.Locale.MustLocalize("Name",lang) , d.EmployeName})

	cell, err = excelize.CoordinatesToCellName(2, 3)
	if err != nil {
		return
	}
	if err = f.SetCellStyle(sheet, "B3", "B3", titleStyle); err != nil {
		return
	}
	if err = f.SetCellStyle(sheet, "C3", "C3", cellStyle); err != nil {
		return
	}
	f.SetSheetRow(sheet, cell, &[]string{r.Locale.MustLocalize("Area",lang), d.GerenciaName})

	cell, err = excelize.CoordinatesToCellName(2, 4)
	if err != nil {
		return
	}
	if err = f.SetCellStyle(sheet, "B4", "B4", titleStyle); err != nil {
		return
	}
	if err = f.SetCellStyle(sheet, "C4", "C4", cellStyle); err != nil {
		return
	}
	f.SetSheetRow(sheet, cell, &[]string{r.Locale.MustLocalize("Place",lang), d.SitioName})

	cell, err = excelize.CoordinatesToCellName(5, 2)
	if err != nil {
		return
	}
	err = f.MergeCell(sheet, "E2", "F2")
	if err != nil {
		return
	}
	if err = f.SetCellStyle(sheet, "E2", "F2", titleStyle); err != nil {
		return
	}
	// log.Println(cell2,cell)
	f.SetSheetRow(sheet, cell, &[]string{r.Locale.MustLocalize("TimePeriod",lang)})

	// cell, err = excelize.CoordinatesToCellName(5, 2)
	// if err != nil {
	// 	return
	// }
	// err = f.MergeCell(sheet, "E2", "F2")
	// if err != nil {
	// 	return
	// }
	// if err = f.SetCellStyle(sheet, "E2", "F2", titleStyle); err != nil {
	// 	return
	// }
	// // log.Println(cell2,cell)
	// f.SetSheetRow(sheet, cell, &[]string{"FECHA INICIO"})

	return
}

// ffcc33
