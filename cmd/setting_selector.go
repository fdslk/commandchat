package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"zqf.com/commandchat/cmdHelper"
)

// settingSelectorCmd represents the setting command by selection
var settingSelectorCmd = &cobra.Command{
	Use:   "settingSelector",
	Short: "modify gpt chat setting",
	Long:  `modify gpt chat setting like chat model, history count, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		displayCurrentSetting()
		updateModel()
		updateHistoryCount()
	},
}

func init() {
	rootCmd.AddCommand(settingSelectorCmd)
}

type AIModel struct {
	Name      string `json:"name"`
	ModelName string `json:"modelName"`
	ApiUrl    string `json:"apiUrl"`
}

var modelMap = map[int]AIModel{
	1: {
		Name:      "Prod",
		ModelName: "gpt-3.5-turbo",
		ApiUrl:    "https://api.openai.com/v1/chat/completions",
	},
	2: {
		Name:      "Prod",
		ModelName: "gpt-3.5-turbo",
		ApiUrl:    "http://localhost:8080/v1/chat/completions",
	},
}

func updateModel() {
	fmt.Println("Please select Model:")

	for key, value := range modelMap {
		fmt.Println(key, ") Model", value.Name, ":", value.ModelName, "API apiUrl:", value.ApiUrl)
	}

	var option int
	_, err := fmt.Scanln(&option)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	switch option {
	case 1:
		fmt.Println("You selected Option 1", modelMap[1].Name)
	case 2:
		fmt.Println("You selected Option 2", modelMap[2].Name)
	default:
		fmt.Println("Invalid input")
		return
	}

	fmt.Println(modelMap[option])
	cmdHelper.SaveFile(modelMap[option], cmdHelper.CONFIGURATIONPATH, cmdHelper.FILE_NAME)
	fmt.Println("Saved!")
}

func updateHistoryCount() {
	//TODO: save history count or size limiation into configuration
}

func displayCurrentSetting() error {
	settingLoaction := cmdHelper.CONFIGURATIONPATH + cmdHelper.FILE_NAME

	setting, err := cmdHelper.ReadFile(settingLoaction)
	if err != nil {
		return err
	}
	fmt.Printf("Current setting is %+v\n", setting)

	return nil
}
