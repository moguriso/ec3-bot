package saicho

import (
    "log"
	"strings"
	"math/rand"
	"time"

    "github.com/nlopes/slack"
)

func checkNormalMorning(text string) string {
	r_text := ""
	if (strings.Index(text, "おはよう") != -1) || (strings.Index(text, "オハヨウ") != -1) {
		r_text = "Good morning!"
	}
	return r_text
}

func initRandomReplyforXXXX() {
	rand.Seed(time.Now().UnixNano())
}

func randamReply() string {
	rnd := rand.Intn(10000)
	r_text := ""
	log.Printf("randome = %d\n", rnd)
	if rnd > 6000 && rnd < 8000 {
		if (rnd % 29) == 0 {
			/* FIXME: edit some reply text */
			r_text = ""
		} else {
			log.Printf("don't show message\n")
		}
	}
	return r_text
}

func checkNormalYoukai(text string) string {
	r_text := ""
	if strings.Index(text, "妖怪人間") != -1 {
		r_text = "早く人間になりたぁい！"
	} 
	return r_text
}

func checkNormalWhatauDoing(text string) string {
	rnd := rand.Intn(1) /* FIXME: endter value of Intn */
	r_text := ""
	/* FIXME: edit some word */
	if strings.Index(text, "") != -1 {
		switch rnd {
			/* FIXME: edit some reply text */
			default:
				r_text = ""
		}
	} 
	return r_text
}

func replyToNormal(api *slack.Client, rtm *slack.RTM, ev *slack.MessageEvent) {
	if rep_text := randamReply() ; rep_text != "" {
		rtm.SendMessage(rtm.NewOutgoingMessage(rep_text, ev.Channel))
	} else if rep_text := checkNormalWhatauDoing(ev.Text) ; rep_text != "" {
		rtm.SendMessage(rtm.NewOutgoingMessage(rep_text, ev.Channel))
	} else if rep_text := checkNormalMorning(ev.Text) ; rep_text != "" {
		rtm.SendMessage(rtm.NewOutgoingMessage(rep_text, ev.Channel))
	} else if rep_text := checkNormalYoukai(ev.Text) ; rep_text != "" {
		rtm.SendMessage(rtm.NewOutgoingMessage(rep_text, ev.Channel))
	}
}

