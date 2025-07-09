package main

// LICENSE: GPLv3
// Author: İsa
// URL: https://github.com/isa-programmer/xkcd-cli
// Description: A command-line interface for xkcd.com comics
// Usage:
//		xkcd-cli # A random comic
// 		xkcd-cli 927 # You can choose comic ID yourself

import (
	"net/http"
	"os/exec"
	"math/rand"
	"github.com/charmbracelet/lipgloss"
	"time"
	"strconv"
	"encoding/json"
	"fmt"
	"os"
	"io"
)



func printXkcdComic(comic XkcdJsonStruct,color string) {
	var output string
	output = fmt.Sprintf(`
	- Comic Link: https://xkcd.com/%d
	- Title: %s
	- Date: %s/%s/%s
	Description: %s
	`,comic.Num,comic.Title,comic.Day,comic.Month,comic.Year,comic.Alt)
	
	style := lipgloss.NewStyle().
				Width(80).
				Border(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.Color(color)).
				Background(lipgloss.Color(color))

	fmt.Println(style.Render(output))
}

type XkcdJsonStruct struct{
	Year 	string	`json:"year"`
	Month 	string	`json:"month"`
	Day 	string	`json:"day"`
	Title 	string	`json:"title"`
	Alt 	string	`json:"alt"`
	Img 	string	`json:"img"`
	Num 	int		`json:"num"`
}

func fetchComic(comicId int) (XkcdJsonStruct, error){
	var comic XkcdJsonStruct
	var url string = "https://xkcd.com/info.0.json"
	if comicId != 0{
		url = fmt.Sprintf("https://xkcd.com/%d/info.0.json",comicId)
	}
	
	
	resp, err := http.Get(url)
	if err != nil{
		return comic,err
	}

	if resp.StatusCode != 200 {
		return comic, fmt.Errorf("failed to fetch comic %d: %s", comicId, resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil{
		return comic,err
	}

	return comic,nil
}

func generateRandomComicNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3112) + 1
}


func downloadImage(url string) error{
	resp, err := http.Get(url)
	if err != nil{
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create("/tmp/temp_file.png")
	if err != nil{
		return err
	}
	defer file.Close()

	_, err = io.Copy(file,resp.Body)
	if err != nil{
		return err
	}
	return nil
}

func main(){
	var comicId int = 0
	var err error
	if len(os.Args) > 1 {
		if os.Args[1] == "random"{
			comicId = generateRandomComicNumber()
		} else {
			comicId, err = strconv.Atoi(os.Args[1])
			if err != nil{
				fmt.Printf("Comic ID should be an integer! Not %s",os.Args[1])
				return
			}
		}
	}
	
	comic,err := fetchComic(comicId)
	if err != nil{
		fmt.Println(err)
		return
	}
	
	printXkcdComic(comic,"#820173")

	err = downloadImage(comic.Img)
	if err != nil{
		fmt.Println(err)
	}

	cmd := exec.Command("xdg-open","/tmp/temp_file.png")

	err = cmd.Run()
	if err != nil{
		fmt.Println(err)
	}
	
}
