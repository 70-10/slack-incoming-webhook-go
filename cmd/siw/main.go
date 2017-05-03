package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/70-10/slack-incoming-webhook-go"
)

var (
	payloadFilePath = flag.String("payload", defaultPayloadPath(), "payload json file")
	slackURL        = flag.String("url", "", "Slack Incoming Webhook URL")
)

func main() {
	flag.Parse()

	os.Exit(Run(*slackURL, *payloadFilePath))
}

func defaultPayloadPath() string {
	wd, _ := os.Getwd()
	return filepath.Join(wd, "payload.json")
}

func Run(slackURL, payloadFilePath string) int {
	bufPayload, err := ioutil.ReadFile(payloadFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	payload := &siw.Payload{}
	err = json.Unmarshal(bufPayload, payload)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	c := siw.NewClient(slackURL)
	body, err := c.PostPayload(*payload)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Fprintln(os.Stdout, body)

	return 0
}
