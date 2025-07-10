package main

// LICENSE: GPLv3
// Author: İsa
// URL: https://github.com/isa-programmer/xkcd-cli
// Description: A command-line interface for xkcd.com comics
// Usage:
//		xkcd-cli # A random comic
// 		xkcd-cli 927 # You can choose comic ID yourself

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
	"xkcd-cli/internal/comic/printcomic"
	"xkcd-cli/internal/comic/fetch"
	"xkcd-cli/internal/comic/download"
	"xkcd-cli/internal/comic/config"
	"xkcd-cli/internal/models"
)

func generateRandomComicNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3112) + 1
}

func main() {
	var comicId int = 0
	var err error

	userConfig,err := config.readConfigFile()
	if err != nil{
		fmt.Println("Error while reading config file:",err)
		return
	}
	
	if len(os.Args) > 1 {
		if os.Args[1] == "random" {
			comicId = generateRandomComicNumber()
		} else {
			comicId, err = strconv.Atoi(os.Args[1])
			if err != nil {
				fmt.Printf("Comic ID should be an integer! Not %s", os.Args[1])
				return
			}
		}
	}

	comic, err := fetchComic(comicId)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	printXkcdComic(comic, userConfig)

	err = download.downloadImage(comic.Img)
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("xdg-open", "/tmp/temp_file.png")

	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}
