package main

import (
	"image"

	bobImage "github.com/nyc4m/retarded-bob-generator/image"
	"github.com/nyc4m/retarded-bob-generator/text"
	"golang.org/x/image/font"
	"gopkg.in/tucnak/telebot.v2"
)

//SpongeBot Defines the bot itself
type SpongeBot struct {
	sourceImg   image.Image
	thumbsUpBob image.Image
	font        font.Face
	bot         *telebot.Bot
}

//getInput returns the argument of the command.
//offset is the emplacement of the first arg to read
//so basically it should be the length of the route name
func getInput(offset int, rawContent string) (arg string) {
	arg = rawContent[offset:]
	return
}

//RetardedText defines the handler to get some retarded text
func (sponge *SpongeBot) RetardedText(m *telebot.Message) {
	input := getInput(len("/retarded"), m.Text)
	output := text.ToBobRetardedString(input)
	sponge.bot.Send(m.Sender, output)
}

//RetardedPic defines the handler to send
//the retarded sponge bot meme
func (sponge *SpongeBot) RetardedPic(m *telebot.Message) {
	input := m.Text
	retardedSentence := text.ToBobRetardedString(input)
	retardedBobBytes := bobImage.GenerateBobMeme(sponge.sourceImg, sponge.font, retardedSentence)
	photoToSend := telebot.Photo{
		File: telebot.FromReader(retardedBobBytes),
	}
	sponge.bot.Send(m.Sender, text.ToBobRetardedString("A few moments later..."))
	go sponge.bot.Send(m.Sender, &photoToSend)
}

//Help returns a helpful message from the bot
//in order for the user to use to bot
func (sponge *SpongeBot) Help(m *telebot.Message) {
	helpMessage := `
	Hey ! if you need any help, that's how it works : 
	- /retarded will convert a sentence to a retarded one : Hello world => HeLlO wOrLd
	- if you type any text, it will generate the meme from sponge bob (image + retarded text)
	`
	photoToSend := telebot.Photo{
		File: telebot.FromReader(bobImage.PngToBytes(sponge.thumbsUpBob)),
	}
	sponge.bot.Send(m.Sender, "Gotcha")
	sponge.bot.Send(m.Sender, &photoToSend)
	sponge.bot.Send(m.Sender, helpMessage)
}
