package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {

	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}

}

func main() {
	os.Setenv("SLACK_BOT_TOKEM", "xoxb-5012960481395-5025759342113-GDhNwSTPnNMVlUXkpQ5xanwz")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A050CUU34FM-5020162617442-695e3e992eb59e3522a814357af7735baf824b4a0df920c2cabd9bc646f99e75")

	bot := slacker.NewClient("xoxb-5012960481395-5025759342113-GDhNwSTPnNMVlUXkpQ5xanwz", "xapp-1-A050CUU34FM-5020162617442-695e3e992eb59e3522a814357af7735baf824b4a0df920c2cabd9bc646f99e75")

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my year of birth is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				print("error")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}

}
