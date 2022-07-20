package commander

import (
	"log"
	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Hostile359/homework-1/config"
	"gitlab.ozon.dev/Hostile359/homework-1/internal/commandhandler"
)

type CmdHandler func(string) string

type Commander struct {
	bot   *tgbotapi.BotAPI
	commandHandler *commandhandler.CommandHandler
}

func Init() (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	commandHandler := &commandhandler.CommandHandler{}
	commandHandler.Init()

	return &Commander{
		bot:   bot,
		commandHandler: commandHandler,
	}, nil
}

func (c *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := c.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		cmd := update.Message.Command()
		args := update.Message.CommandArguments()
		msg.Text = c.commandHandler.HandleCommand(cmd, args)
		
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.Wrap(err, "send tg message")
		}
	}
	return nil
}
