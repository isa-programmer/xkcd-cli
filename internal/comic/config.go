package comic

import (
	"os"
	"path/filepath"
	"encoding/json"
	"github.com/isa-programmer/xkcd-cli/internal/models"
)




func CreateConfigFile() error {
	configFileContent := `
	{
		"border_color":"#07177f",
		"background_color":"#030b3f"
	}
	`
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".config", "xkcd.json")
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(configFileContent)
	if err != nil {
		return err
	}

	return nil
}


func ReadConfigFile() (models.Config,error){
	config := models.Config{}
	
	home,_ := os.UserHomeDir()
	configPath := filepath.Join(home,".config","xkcd.json")
	data,err := os.ReadFile(configPath)
	if err != nil {
		return config,err
	}
	err = json.Unmarshal(data,&config)
	return config, err
} 
