package main

import (
    "os"

	"module/saicho"
    "github.com/nlopes/slack"
)

func run(api *slack.Client) int {
    rtm := api.NewRTM()
    go rtm.ManageConnection()

	return saicho.EventLoop(api, rtm)
}

func main() {
	api := slack.New("")
    os.Exit(run(api))
}

