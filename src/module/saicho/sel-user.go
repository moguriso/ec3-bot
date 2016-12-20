package saicho

import (
//  "log"
	"strings"

    "github.com/nlopes/slack"
)

func selectUser(api *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent) {

//	user, err := api.GetUserInfo(ev.User)
//	if err != nil {
//		log.Printf("%s\n", err)
//		return
//	}

	if strings.Index(ev.Text, "@XXXXXXXX") != -1 {		  /* FIXME		  */
		replyToNormal(api, rtm, ev)						  /* edit user ID */
	}
}

