package main

import (
	"os"
	"strings"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/eldeal/collab/data"
	"github.com/nlopes/slack"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear("/tech|/technology|(?i)I know about(.*)").MessageHandler(addTech)
	//bot.Hear("/learn|(?i)I want to learn(.*)").MessageHandler(learn)
	//bot.Hear("/users|(?i)who knows(.*)").MessageHandler(AttachmentsHandler)
	bot.Run()
}

func addTech(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	list := strings.Split(evt.Text, " ")
	//trigger := list[0]
	s := strings.TrimSpace(list[1])
	if strings.Contains(s, " ") {
		list = strings.Split(s, " ")
	} else if strings.Contains(s, ",") {
		list = strings.Split(s, ",")
	} else {
		list = []string{s}
	}

	for _, tech := range list {
		doc := data.DB.FindTechnology(tech)
		if doc == nil {
			data.DB.NewTechnology(tech, evt.User, "tech")
			continue
		}

		//	switch msg.Trigger {
		//	case "tech:":
		doc.AddUser(evt.Name, tech)
		//		case "learn:":
		//		doc.AddLearner(msg.Name, tech)
		//		default:
		//			bot.Reply(evt, "You are already listed as a user of that technology.")
		//	}
		bot.Reply(evt, "Thanks for letting me know you use "+tech, slackbot.WithoutTyping)

	}
}
