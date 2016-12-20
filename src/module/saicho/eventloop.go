package saicho

import (
    "log"

    "github.com/nlopes/slack"
)

func EventLoop(api *slack.Client, rtm *slack.RTM) int {
	for {
        select {
        case msg := <-rtm.IncomingEvents:
            switch ev := msg.Data.(type) {
            case *slack.HelloEvent:
				Hello(api, rtm, ev)
					
			case *slack.UserTypingEvent:
				UserTyping(api, rtm, ev)

			case *slack.PresenceChangeEvent:
				PresenceChange(api, rtm, ev)

            case *slack.MessageEvent:
				Message(api, rtm, ev)

            case *slack.InvalidAuthEvent:
                log.Println("Invalid credentials")
                return 1
            }
        }
    }
	return 0
}

