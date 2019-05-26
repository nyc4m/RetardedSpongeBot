package main

import (
	"bytes"
	"fmt"
	"github.com/nyc4m/retarded-bob-generator/image"
	"gopkg.in/tucnak/telebot.v2"
	"image/png"
	"io/ioutil"
	"os"
	"time"
)

const (
	bobImagePath = "./res/bob_source.png"
	fontPath     = "./res/font/impact.ttf"
)

var Token = os.Getenv("TOKEN")

func main() {
	if Token == "" {
		fmt.Println("no Token, exiting")
		return
	}
	bobImageBytes, err := ioutil.ReadFile(bobImagePath)
	bobImage, err := png.Decode(bytes.NewBuffer(bobImageBytes))
	fontBytes, err := ioutil.ReadFile(fontPath)
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
