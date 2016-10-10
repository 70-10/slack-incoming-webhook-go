# slack-incoming-webhook-go

Slack Incoming Webhook SDK for Go

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/70-10/slack-incoming-webhook-go"
)

func main() {
	c := siw.NewClient("<YOUR INCOMING WEBHOOK URL>")

	payload := siw.Payload{
		Channel:  "#general",
		UserName: "webhookbot",
		Text:     "This is posted to #general and comes from a bot named webhookbot.",
		IconURL:  "https://slack.com/img/icons/app-57.png",
	}

	body, err := c.PostPayload(payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)
}
```
