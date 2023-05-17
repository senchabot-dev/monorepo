package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/senchabot-dev/monorepo/apps/twitch-bot/client"
	"github.com/senchabot-dev/monorepo/apps/twitch-bot/internal/command/helpers"
	"github.com/senchabot-dev/monorepo/apps/twitch-bot/server"
)

const DELETE_COMMAND_INFO = "For example: !dcmd [command_name]"

func DeleteCommandCommand(client *client.Clients, server *server.SenchabotAPIServer, message twitch.PrivateMessage, commandName string, params []string) {
	if !helpers.CanExecuteCommand(context.Background(), server, message) {
		return
	}
	if len(params) < 1 {
		client.Twitch.Say(message.Channel, DELETE_COMMAND_INFO)
		return
	}
	var command_name = strings.ToLower(params[0])
	if commandName != "" {
		err := server.DeleteBotCommand(context.Background(), command_name, message.RoomID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("COMMAND_DELETE: command_name:", command_name)

		client.Twitch.Say(message.Channel, "Command Deleted: "+command_name)
	}
}
