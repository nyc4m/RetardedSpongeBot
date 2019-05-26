package main

import (
	bobImage "github.com/nyc4m/retarded-bob-generator/image"
	"github.com/nyc4m/retarded-bob-generator/text"
	"golang.org/x/image/font"
	"gopkg.in/tucnak/telebot.v2"
	"image"
)

//SpongeBot Defines the bot itself
type SpongeBot struct {
	sourceImg image.Image
	font      font.Face
	bot       *telebot.Bot
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
	input := getInput(len("/retardedPic"), m.Text)
	retardedSentence := text.ToBobRetardedString(input)
	retardedBobBytes := bobImage.GenerateBobMeme(sponge.sourceImg, sponge.font, retardedSentence)
	photoToSend := telebot.Photo{
		File: telebot.FromReader(retardedBobBytes),
	}
	sponge.bot.Send(m.Sender, text.ToBobRetardedString("A few moments later..."))
	go sponge.bot.Send(m.Sender, &photoToSend)
}
