package cmd

import "github.com/charmbracelet/lipgloss"

var (
	width     = 80
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	success   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	successStyle = lipgloss.NewStyle().Foreground(success)

	paraStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
		//Foreground(lipgloss.Color("#FAFAFA")).
		//Background(highlight).
		Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true) //.
		//Foreground(highlight)

	hiliteStyle = lipgloss.NewStyle().
		//Foreground(lipgloss.Color("#FFFDF5")).
		//Background(lipgloss.Color("#FF5F87")).
		Padding(0, 1)
)

func helpParagraph(para string) string {
	return lipgloss.JoinVertical(lipgloss.Center,
		paraStyle.Width(width).Margin(1, 0).Render(para),
		lipgloss.PlaceHorizontal(width, lipgloss.Center,
			hiliteStyle.Margin(0, 1).Render(`Tlön, Uqbar, Orbis Tertius`),
			lipgloss.WithWhitespaceChars("="),
			lipgloss.WithWhitespaceForeground(subtle),
		))
}
