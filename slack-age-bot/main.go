package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
)

func printCommandEvents(commandEvents chan *slacker.CommandEvent) {
	for commandEvent := range commandEvents {
		fmt.Println(commandEvent)

	}
}

func main() {
	os.Setenv("Slack_Bot_Token", "xoxb-5950170140819-5947378409013-u8fhCm201y9VCv1AyA2iUdtk")
	os.Setenv("Slack_App_Token", "xapp-1-A05UM82JSN4-5950188269731-b56c7f50b079d7990ed8b755faa3545dbe889bab45cea334d3e8fd153e00c86c")
	bot := slacker.NewClient(os.Getenv("Slack_Bot_Token"), os.Getenv("Slack_App_Token"))
	go printCommandEvents(bot.CommandEvents())
	bot.Command("hello", &slacker.CommandDefinition{
		Description: "Say Hello",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Hello")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
