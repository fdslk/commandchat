package commandchat

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const FILE_NAME = "setting.json"

type ChatSetting struct {
	ModelName string `json:"modelName"`
	ApiUrl    string `json:"apiUrl"`
}

func ReadFile(filePath string) (ChatSetting, error) {
	var setting ChatSetting
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return setting, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&setting)
	if err != nil {
		log.Println(err)
		return setting, err
	}
	return setting, nil
}

func SaveFile(setting interface{}, filePath string) error {
	bytes, err := json.Marshal(setting)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath, bytes, 0644)
	if err != nil {
		panic(err)
	}
	return err
}

func (setting *ChatSetting) IsChatModel() bool {
	return setting.ModelName == "gpt-3.5-turbo"
}
