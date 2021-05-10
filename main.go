package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

var rest = resty.New().SetRetryCount(3).
	SetRetryWaitTime(5 * time.Second).
	SetRetryMaxWaitTime(60 * time.Second).
	SetHostURL("https://api.letserver.run")

// args:
// 1. user token (required)
// 2. status (required, auto input by action)
// 3. success
// 4. failed
// 5. command
// 6. source
func main() {
	if len(os.Args) < 7 {
		fmt.Println("please run in github actions.")
		os.Exit(1)
	}
	// inputs
	var token = os.Args[1]
	var status = os.Args[2]
	var success = os.Args[3]
	var failed = os.Args[4]
	var command = os.Args[5]
	var source = os.Args[6]
	if source == "" {
		source = "Github Action"
	}
	if token == "" {
		fmt.Println("skadi user token can not be empty.")
		os.Exit(1)
	}
	rest = rest.SetAuthToken(token)
	// run
	if status == "success" {
		sendNotify(source, success)
		if command != "" {
			sendCommand(command)
		}
	} else {
		sendError(source, failed)
	}
}

func sendNotify(source, msg string) {
	if msg == "" {
		msg = "Action Success"
	}
	_, err := rest.R().SetBody(map[string]string{
		"message": msg,
		"source":  source,
	}).Post("/message/info")
	if err != nil {
		fmt.Println("send skadi notify error: ", err)
		os.Exit(1)
	}
	fmt.Println("send skadi notify successful.")
}

func sendError(source, msg string) {
	_, err := rest.R().SetBody(map[string]string{
		"message": msg,
		"source":  source,
	}).Post("/message/warning")
	if err != nil {
		fmt.Println("send skadi warning error: ", err)
		os.Exit(1)
	}
	fmt.Println("send skadi warning successful.")
}

func sendCommand(msg string) {
	_, err := rest.R().SetBody(map[string]string{
		"message": msg,
	}).Post("/job/add")
	if err != nil {
		fmt.Println("send skadi command error: ", err)
		os.Exit(1)
	}
	fmt.Println("send skadi command successful.")
}
