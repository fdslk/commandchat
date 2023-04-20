/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"zqf.com/commandchat/cmdHelper"
)

// imageCmd represents the image command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "use this command to generator or edit picture",
	Long:  `use this command to generator or edit picture`,
	Run: func(cmd *cobra.Command, args []string) {
		generateImage()
	},
}

func init() {
	rootCmd.AddCommand(imageCmd)
}

func generateImage() {
	filepath := cmdHelper.CONFIGURATIONPATH + cmdHelper.FILE_NAME
	setting, err := cmdHelper.ReadFile(filepath)
	if err != nil {
		fmt.Printf("setting reading error %s", err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	quit := false

	for !quit {
		fmt.Print("Input your question (type `quit` to exit): ")

		if !scanner.Scan() {
			break
		}

		question := scanner.Text()

		switch question {
		case "quit":
			quit = true
		case "":
			continue
		default:
			cmdHelper.Display(question)
		}
	}
}
