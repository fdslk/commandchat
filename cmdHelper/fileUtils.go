package cmdHelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	FILE_NAME         = "setting.json"
	CONFIGURATIONPATH = "Configuration/"
	CHATHISTORYPATH   = "historys/"
)

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

func SaveFile(data interface{}, filePath string, fileName string) error {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		// 创建目录
		err = os.MkdirAll(filePath, 0755)
		if err != nil {
			fmt.Println(err)
			return err
		}

		// 创建文件
		file, err := os.Create(filePath + fileName)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer file.Close()

		fmt.Println("File created:", filePath+fileName)
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filePath+fileName, bytes, 0644)
	if err != nil {
		panic(err)
	}
	return err
}

func (setting *ChatSetting) IsChatModel() bool {
	return setting.ModelName == "gpt-3.5-turbo"
}
