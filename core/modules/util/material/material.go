package material

import (
	_r "acs/domain/repository"
)

func NewPaletteColor() _r.ColorPalette {
	return _r.ColorPalette{
		PrimaryColor:   "008542",
		ColorWhite:     "ffffff",
		SecondaryColor: "ffcc33",
		BorderColor:    "6e6e6e",
	}
}

func NewTextSize() _r.TextSizes {
	return _r.TextSizes{
		Small:  8.5,
		Medium: 12,
		Large:  16,
	}
}
