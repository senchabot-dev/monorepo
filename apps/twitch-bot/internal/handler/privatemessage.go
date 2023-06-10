package handler

import (
	"context"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/senchabot-dev/monorepo/apps/twitch-bot/client"
	"github.com/senchabot-dev/monorepo/apps/twitch-bot/internal/command"
	"github.com/senchabot-dev/monorepo/apps/twitch-bot/internal/command/helpers"

	"github.com/senchabot-dev/monorepo/apps/twitch-bot/internal/service"
)

func PrivateMessage(client *client.Clients, service service.Service) {
	commands := command.NewCommands(client, service)
	ctx := context.Background()

	client.Twitch.OnPrivateMessage(func(message twitch.PrivateMessage) {
		cmdName, params := helpers.SplitMessage(message.Message)
		if cmdName == "" {
			return
		}
		cmds := commands.GetCommands()

		if cmd, ok := cmds[cmdName]; ok {
			cmd(ctx, message, cmdName, params)
			service.SaveBotCommandActivity(ctx, cmdName, message.RoomID, message.User.DisplayName)
			return
		}

		commands.RunCommand(ctx, cmdName, message)
	})
}
