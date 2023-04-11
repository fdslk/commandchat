/*
Copyright © 2023 zengqiang <zqfangmaster@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// settingCmd represents the setting command
var settingCmd = &cobra.Command{
	Use:   "setting",
	Short: "modify gpt chat setting",
	Long:  `modify gpt chat setting like chat model, history count, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setting called")
	},
}

func init() {
	rootCmd.AddCommand(settingCmd)
}
