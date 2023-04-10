/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
	commandchat "zqf.com/commandchat/commandchatChannel"
)

var currentChatHistory = make(map[string][]interface{})

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "plain text chat with AI, don't input any sensitive infornmation",
	Long: `When you use this command, you can ask any question to chatGpt. It will output you wanted answer, but keep
	it in you mind, don't upload any personal or sensitive data to it`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		quit := false
		// firstTalk := true
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
				// history := commandchat.Convert2HistoryMessage(currentChatHistory, firstTalk)

				newRequestBytes, err := commandchat.CreateCompletionsRequest(question, []commandchat.Message{})

				fmt.Println(string(newRequestBytes))

				if err != nil {
					fmt.Println("error occurred:", err)
					return
				}

				rawResponse, err := commandchat.Chat(newRequestBytes)

				if err != nil || rawResponse.Status != "200 OK" {
					body, _ := io.ReadAll(rawResponse.Body)
					fmt.Printf("error occurred:%s and the response is %s", err, string(body))
					return
				}

				response, err := commandchat.CreateCompletionsResponse(rawResponse)

				if err != nil {
					fmt.Println("error occurred:", err)
					return
				}

				responseText := response.Choices[0].Text
				AIOutPut(responseText)
				// firstTalk = false
				commandchat.UpdateMap(commandchat.USER, question, currentChatHistory)
				commandchat.UpdateMap(commandchat.ASSISTANT, responseText, currentChatHistory)
			}
		}
	},
}

func AIOutPut(answer string) {
	for _, c := range answer {
		fmt.Printf("%c", c)
		time.Sleep(time.Second / 20)
	}
	fmt.Println("")
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
