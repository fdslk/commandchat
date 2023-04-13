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
	"zqf.com/commandchat/cmdHelper"
)

var currentChatHistory = make(map[string][]interface{})

type chatHistory struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

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
				createHistoryFile()
			case "":
				continue
			default:
				chat(question)
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

func chat(question string) {
	filepath := cmdHelper.CONFIGURATIONPATH + cmdHelper.FILE_NAME
	setting, err := cmdHelper.ReadFile(filepath)
	if err != nil {
		fmt.Printf("setting reading error %s", err.Error())
		os.Exit(1)
	}
	history := cmdHelper.Convert2HistoryMessage(currentChatHistory, setting)

	newRequestBytes, err := cmdHelper.CreateCompletionsRequest(question, history, setting)

	if err != nil {
		fmt.Println("error occurred:", err)
		return
	}

	rawResponse, err := cmdHelper.Chat(newRequestBytes, setting)

	if err != nil || rawResponse.Status != "200 OK" {
		body, _ := io.ReadAll(rawResponse.Body)
		fmt.Printf("error occurred:%s and the response is %s", err, string(body))
		return
	}

	response, err := cmdHelper.CreateCompletionsResponse(rawResponse)

	if err != nil {
		fmt.Println("error occurred:", err)
		return
	}

	var responseText string
	if setting.IsChatModel() {
		responseText = response.Choices[0].Message.Content
	} else {
		responseText = response.Choices[0].Text
	}
	AIOutPut(responseText)
	cmdHelper.UpdateMap(cmdHelper.USER, question, currentChatHistory)
	cmdHelper.UpdateMap(cmdHelper.ASSISTANT, responseText, currentChatHistory)
}

func createHistoryFile() {
	userHistory := currentChatHistory[cmdHelper.USER]
	assistantHistory := currentChatHistory[cmdHelper.ASSISTANT]
	if len(userHistory) > 0 {
		var historys []chatHistory
		for index, content := range userHistory {
			historys = append(historys, chatHistory{cmdHelper.USER, content.(string)})
			historys = append(historys, chatHistory{cmdHelper.ASSISTANT, assistantHistory[index].(string)})
		}

		fmt.Print("You can input your current chat history name, if no the first sentence in the current chat history\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		var fileName string
		inputName := scanner.Text()
		if inputName != "" {
			fileName = inputName + ".json"
		} else {
			fileName = userHistory[0].(string) + ".json"
		}
		cmdHelper.SaveFile(historys, cmdHelper.CHATHISTORYPATH, fileName)
	}
}

func init() {
	rootCmd.AddCommand(chatCmd)
}
