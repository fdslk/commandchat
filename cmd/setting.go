/*
Copyright © 2023 zengqiang <zqfangmaster@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"zqf.com/commandchat/cmdHelper"
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
	settingLoaction := cmdHelper.CONFIGURATIONPATH + cmdHelper.FILE_NAME
	save := false
	var newSetting string
	setting, err := cmdHelper.ReadFile(settingLoaction)
	if err != nil {
		return err
	}
	fmt.Printf("Current setting is %+v\n", setting)

	for !save {
		newSetting, err = cmdHelper.ReadJson()
		if err != nil {
			fmt.Printf("Your current input is invalidate and err is `%s`, please input again\n", err.Error())
			continue
		}

		switch newSetting {
		case "save":
			cmdHelper.SaveFile(setting, cmdHelper.CONFIGURATIONPATH, cmdHelper.FILE_NAME)
			save = true
		case "":
			continue
		default:
			fmt.Print(newSetting)
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
