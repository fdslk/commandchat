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
* How to use mock server to simulate
  * pre-preparing
    * docker
  * mock data is saved in the folder `stub`, if you want to add a new request, you can add a new file in its sub-folder `mappings`. And the response of a new request is saved in the sub-folder `__files`
  * modify the configuration `setting.json`, let the endpoint connect to the mocker server as follows:
    ```json
    {
      "modelName": "gpt-3.5-turbo",
      "apiUrl": "http://localhost:8080/v1/chat/completions"
    }
    ```
  * run `startstub.sh` to start the mocker server, you will see the following output in your console
    ```plain text
    2023-04-12 06:50:22.626 Verbose logging enabled
    2023-04-12 06:50:24.321 Verbose logging enabled
    /$$      /$$ /$$                     /$$      /$$                     /$$
    | $$  /$ | $$|__/                    | $$$    /$$$                    | $$
    | $$ /$$$| $$ /$$  /$$$$$$   /$$$$$$ | $$$$  /$$$$  /$$$$$$   /$$$$$$$| $$   /$$
    | $$/$$ $$ $$| $$ /$$__  $$ /$$__  $$| $$ $$/$$ $$ /$$__  $$ /$$_____/| $$  /$$/
    | $$$$_  $$$$| $$| $$  \__/| $$$$$$$$| $$  $$$| $$| $$  \ $$| $$      | $$$$$$/
    | $$$/ \  $$$| $$| $$      | $$_____/| $$\  $ | $$| $$  | $$| $$      | $$_  $$
    | $$/   \  $$| $$| $$      |  $$$$$$$| $$ \/  | $$|  $$$$$$/|  $$$$$$$| $$ \  $$
    |__/     \__/|__/|__/       \_______/|__/     |__/ \______/  \_______/|__/  \__/

    port:                         8080
    enable-browser-proxying:      false
    disable-banner:               false
    no-request-journal:           false
    verbose:                      true
    ```
  * Then you can run `go run . chat` to debug your app, you will see the verbose log in the console.
