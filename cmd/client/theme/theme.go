// Package theme provides custom colors and a theme for the app.
package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/stnokott/spacetrader/res"
)

var (
	// ColorBg is the background color.
	ColorBg = color.RGBA{35, 49, 66, 255}
	// ColorFg is the foreground color, mostly used for text.
	ColorFg = color.RGBA{227, 227, 227, 255}
	// ColorPrimary is the accent color.
	ColorPrimary = color.RGBA{249, 89, 89, 255}
	// ColorCredits is the color for displaying monetary values.
	ColorCredits = color.RGBA{234, 250, 90, 255}
	// ColorSuccess is used for displaying messages related to success.
	ColorSuccess = color.RGBA{59, 201, 49, 255}
	// ColorWarning is used for displaying messages related to warnings.
	ColorWarning = color.RGBA{250, 159, 90, 255}
	// ColorError is used for displaying messages related to errors.
	ColorError = ColorPrimary
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
			return ColorBg
		case theme.ColorNameForeground:
			return ColorFg
		case theme.ColorNamePrimary:
			return ColorPrimary
		case theme.ColorNameSuccess:
			return ColorSuccess
		case theme.ColorNameWarning:
			return ColorWarning
		case theme.ColorNameError:
			return ColorError
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
