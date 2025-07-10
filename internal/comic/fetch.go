package fetch


import (
	"fmt"
	"encoding/json"
	"net/http"
	"xkcd-cli/internal/models"
)

func fetchComic(comicId int) (models.XkcdJsonStruct, error) {
	var comic models.XkcdJsonStruct
	var url string = "https://xkcd.com/info.0.json"
	if comicId != 0 {
		url = fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicId)
	}

	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}

	if resp.StatusCode != 200 {
		return comic, fmt.Errorf("failed to fetch comic %d: %s", comicId, resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		return comic, err
	}

	return comic, nil
}
