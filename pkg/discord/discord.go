package discord

import "github.com/bwmarrin/discordgo"

type Bot struct {
	token   string
	Session *discordgo.Session
}

func NewDiscord(apiToken string) (*Bot, error) {
	var (
		err error
		bot = &Bot{}
	)

	bot.token = apiToken
	bot.Session, err = discordgo.New("Bot " + bot.token)
	if err != nil {
		return nil, err
	}

	return bot, nil
}
