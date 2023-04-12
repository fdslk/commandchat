/*
Copyright © 2023 zengqiang <zqfangmaster@gmail.com>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	commandchat "zqf.com/commandchat/commandchatChannel"
)

// settingCmd represents the setting command
var settingCmd = &cobra.Command{
	Use:   "setting",
	Short: "modify gpt chat setting",
	Long:  `modify gpt chat setting like chat model, history count, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		editSetting()
	},
}

func init() {
	rootCmd.AddCommand(settingCmd)
}

func editSetting() error {
	settingLoaction := "Configuration/" + commandchat.FILE_NAME
	scanner := bufio.NewScanner(os.Stdin)
	save := false
	setting, err := commandchat.ReadFile(settingLoaction)
	if err != nil {
		return err
	}
	fmt.Printf("Current setting is %+v\n", setting)
	for !save {
		fmt.Println("Please Input the setting you want to set (type `save` to save setting): ")
		if !scanner.Scan() {
			break
		}
		newSetting := scanner.Text()

		switch newSetting {
		case "save":
			commandchat.SaveFile(setting, settingLoaction)
			save = true
		case "":
			continue
		default:
			err = json.Unmarshal([]byte(newSetting), &setting)
			if err != nil {
				fmt.Printf("Your current input is invalidate and err is `%s`, please input again\n", err.Error())
				continue
			}
			fmt.Printf("Current modification setting is %+v\n", setting)
		}
	}

	return nil
}
