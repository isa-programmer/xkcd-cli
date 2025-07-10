package comic

import (
	"github.com/charmbracelet/lipgloss"
	"fmt"
	"github.com/isa-programmer/xkcd-cli/internal/models"
)

func PrintXkcdComic(comic models.XkcdJsonStruct, config models.Config) {
	var output string
	output = fmt.Sprintf(`
	- Comic Link: https://xkcd.com/%d
	- Title: %s
	- Date: %s/%s/%s
	Description: %s
	`, comic.Num, comic.Title, comic.Day, comic.Month, comic.Year, comic.Alt)

	style := lipgloss.NewStyle().
		Width(80).
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color(config.BorderColor)).
		Background(lipgloss.Color(config.BackgroundColor))

	fmt.Println(style.Render(output))
}