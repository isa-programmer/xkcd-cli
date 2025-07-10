package config

import (
	"path/filepath"
	"encoding/json"
	"xkcd-cli/internal/models"
)

func readConfigFile() (models.Config,error){
	var config models.Config
	home,_ := os.UserHomeDir()
	configPath := filepath.Join(home,".config","xkcd","xkcd.json")
	data,err := os.ReadFile(configPath)
	if err != nil{
		return config,err
	}
	err = json.Unmarshal(data,&config)
	return config, err
} 
