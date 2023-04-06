/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	commandchat "zqf.com/commandchat/commandchatChannel"
)

var currentChatHistory = make(map[string]string)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "plain text chat with AI, don't input any sensitive infornmation",
	Long: `When you use this command, you can ask any question to chatGpt. It will output you wanted answer, but keep
	it in you mind, don't upload any personal or sensitive data to it`,
	Run: func(cmd *cobra.Command, args []string) {
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
				newRequestBytes, err := commandchat.CreateCompletionsRequest(question)

				if err != nil {
					fmt.Println("error occurred:", err)
					return
				}

				rawResponse, err := commandchat.Chat(newRequestBytes)

				if err != nil {
					fmt.Println("error occurred:", err)
					return
				}

				response, err := commandchat.CreateCompletionsResponse(rawResponse)

				if err != nil {
					fmt.Println("error occurred:", err)
					return
				}

				AIOutPut(response.Choices[0].Text)
			}
		}
	},
}

func AIOutPut(answer string) {
	for _, c := range answer {
		fmt.Printf("%c", c)
		time.Sleep(time.Second / 5)
	}
	fmt.Println("")
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
