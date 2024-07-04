package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/stnokott/spacetrader/res"
)

var (
	_colorBg      = color.RGBA{35, 49, 66, 255}
	_colorFg      = color.RGBA{227, 227, 227, 255}
	_colorPrimary = color.RGBA{249, 89, 89, 255}
	_colorCredits = color.RGBA{234, 250, 90, 255}
	_colorSuccess = color.RGBA{59, 201, 49, 255}
	_colorWarning = color.RGBA{250, 159, 90, 255}
	_colorError   = _colorPrimary
)

// Theme implements fyne.Theme.
type Theme struct {
}

var _ fyne.Theme = (*Theme)(nil)

// Color implements fyne.Theme.
func (Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if variant == theme.VariantDark {
		switch name {
		case theme.ColorNameBackground:
			return _colorBg
		case theme.ColorNameForeground:
			return _colorFg
		case theme.ColorNamePrimary:
			return _colorPrimary
		case theme.ColorNameSuccess:
			return _colorSuccess
		case theme.ColorNameWarning:
			return _colorWarning
		case theme.ColorNameError:
			return _colorError
		case theme.ColorNameDisabled:
			return color.RGBA{69, 93, 122, 255} // #455d7a
		case theme.ColorNameFocus:
			return color.RGBA{97, 129, 168, 255} // #6181A8
		}
	}
	return theme.DefaultTheme().Color(name, variant)
}

var (
	fontRegular = fyne.NewStaticResource("font-regular", res.MonoFontRegular)
	fontBold    = fyne.NewStaticResource("font-bold", res.MonoFontBold)
)

// Font implements fyne.Theme.
func (Theme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		return fontBold
	}
	return fontRegular
}

// Icon implements fyne.Theme.
func (Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name) // TODO
}

// Size implements fyne.Theme.
func (Theme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNamePadding:
		return 2
	case theme.SizeNameInnerPadding:
		return 0
	}
	return theme.DefaultTheme().Size(name)
}
