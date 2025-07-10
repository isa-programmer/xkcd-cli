package comic

import (
	"net/http"
	"io"
	"os"
)

func DownloadImage(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create("/tmp/temp_file.png")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}