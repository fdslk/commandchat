# chatGpt golang command line tool

![build status](https://github.com/Fdslk/commandchat/actions/workflows/go.yml/badge.svg)

## What is it?

* It is a command line tool that can be used to chat with GPT, which will send messages with the latest two conversation record
* It is based on Golang

## How to run it in local

* run `go run . chat`

## How to use it

* run `export API_KEY="[your api key]"`
* run `commandchat chat"` to chat with Chatgpt and the default chat model is `gpt-3.5-turbo`
* run `commandchat setting` to modify the chatting model, **FYI**, the input should be formatted as `json`
  * eg: switch to model `text-davinci-003`,

  ```json
  {
    "modelName": "text-davinci-003",
    "apiUrl": "https://api.openai.com/v1/completions"
  }
  ```

![output](https://user-images.githubusercontent.com/6279298/231410956-5555a391-1557-406f-a088-a15e9accc25c.gif)

## Development tip

* In case OpenAi locks down your APIKEY or account, you can use the Wiremock as your mock server to develop or debug your code. You can run `./startstub.sh`
* check your Wiremock works, input Url `http://localhost:8080/__admin` in your browser
