package command

import (
	"context"

	"github.com/gempir/go-twitch-irc/v3"
)

func (s *commands) AstraCommand(context context.Context, message twitch.PrivateMessage, commandName string, params []string) {
	s.client.Twitch.Say(message.Channel, "[We did it!] Astra UI Kit: https://docs.astraui.com")
}
