// Package styles implements the theming logic for slides
package styles

import (
	_ "embed"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const (
	salmon = lipgloss.Color("#E8B4BC")
)

var (
	Author = lipgloss.NewStyle().Foreground(salmon).Align(lipgloss.Left).MarginLeft(2)
	Date   = lipgloss.NewStyle().Faint(true).Align(lipgloss.Left).Margin(0, 1)
	Page   = lipgloss.NewStyle().Foreground(salmon).Align(lipgloss.Right).MarginRight(3)
	Slide  = lipgloss.NewStyle().Padding(1)
	Status = lipgloss.NewStyle().Padding(1)
)

var (
	//go:embed theme.json
	DefaultTheme []byte
)

func JoinHorizontal(left, right string, width int) string {
	length := lipgloss.Width(left + right)
	if width < length {
		return left + " " + right
	}
	padding := strings.Repeat(" ", width-length)
	return left + padding + right
}

func JoinVertical(top, bottom string, height int) string {
	h := lipgloss.Height(top) + lipgloss.Height(bottom)
	if height < h {
		return top + "\n" + bottom
	}
	fill := strings.Repeat("\n", height-h)
	return top + fill + bottom
}

// SelectTheme picks a glamour style config based
// on the theme provided in the markdown header
func SelectTheme(theme string) glamour.TermRendererOption {
	switch theme {
	case "ascii":
		return glamour.WithStyles(glamour.ASCIIStyleConfig)
	case "light":
		return glamour.WithStyles(glamour.LightStyleConfig)
	case "dark":
		return glamour.WithStyles(glamour.DarkStyleConfig)
	case "notty":
		return glamour.WithStyles(glamour.NoTTYStyleConfig)
	default:
		bytes, err := os.ReadFile(theme)
		if err == nil {
			return glamour.WithStylesFromJSONBytes(bytes)
		}
		// Should log a warning so the user knows we failed to read their theme file

		if termenv.EnvNoColor() {
			return glamour.WithStyles(glamour.NoTTYStyleConfig)
		}

		if !termenv.HasDarkBackground() {
			return glamour.WithStyles(glamour.LightStyleConfig)
		}

		return glamour.WithStylesFromJSONBytes(DefaultTheme)
	}
}
