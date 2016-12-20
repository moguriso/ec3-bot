package saicho

import (
    "log"

    "github.com/nlopes/slack"
)

/* bot login */
func Hello(api *slack.Client, rtm *slack.RTM, ev *slack.HelloEvent) {
	// log.Println("Hello Event")
	initRandomReplyforXXXX()
}

func UserTyping(api *slack.Client, rtm *slack.RTM, ev *slack.UserTypingEvent) {
	// log.Println("UserTyping Event")
	user, err := api.GetUserInfo(ev.User)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
	log.Println(user.Name+ "さんが文字打ち中")
}

func PresenceChange(api *slack.Client, rtm *slack.RTM, ev *slack.PresenceChangeEvent){
	// log.Println("PresenceChange Event")
	isOnline := "オンライン"
	if ev.Presence == "away" {
		isOnline = "オフライン"
	}

	user, err := api.GetUserInfo(ev.User)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	log.Println(user.Name+ "さんが" + isOnline + "になりました")
}

func Message(api *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent) {
	log.Printf("Message: %v\n", ev)
		
	selectUser(api, rtm, ev)
}

