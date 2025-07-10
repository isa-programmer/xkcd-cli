package comic

import (
	"github.com/charmbracelet/lipgloss"
	"fmt"
	"github.com/isa-programmer/xkcd-cli/internal/models"
)

func PrintDetails(comicData models.XkcdJsonStruct, config models.Config) {
	var output string
	output = fmt.Sprintf(`
	- comic Link: https://xkcd.com/%d
	- Title: %s
	- Date: %s/%s/%s
	Description: %s
	`, comicData.Num, comicData.Title, comicData.Day, comicData.Month, comicData.Year, comicData.Alt)

	style := lipgloss.NewStyle().
		Width(80).
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color(config.BorderColor)).
		Background(lipgloss.Color(config.BackgroundColor))

	fmt.Println(style.Render(output))
}