package main

import (
	"bytes"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/nyc4m/retarded-bob-generator/image"
	"gopkg.in/tucnak/telebot.v2"
	"image/png"
	"os"
	"time"
)

const (
	bobImagePath = "/bob_source.png"
	fontPath     = "/font/impact.ttf"
)

var Token = os.Getenv("TOKEN")

func main() {
	if Token == "" {
		fmt.Println("no Token, exiting")
		return
	}
	box := rice.MustFindBox("./res")
	bobImageBytes := box.MustBytes(bobImagePath)
	bobImage, err := png.Decode(bytes.NewBuffer(bobImageBytes))
	fontBytes := box.MustBytes(fontPath)
	font, err := image.LoadFontFromBytes(fontBytes, 30)
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	spongeBot := SpongeBot{bot: bot, sourceImg: bobImage, font: font}
	if err != nil {
		fmt.Println("Bot is not working")
		return
	}

	bot.Handle("/retarded", spongeBot.RetardedText)

	bot.Handle("/retardedpic", spongeBot.RetardedPic)

	bot.Start()
	fmt.Println("salut")
}
