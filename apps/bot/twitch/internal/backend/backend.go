package backend

import (
	"context"

	"github.com/senchabot-dev/monorepo/apps/bot/twitch/internal/models"
)

type Backend interface {
	GetTwitchChannels(ctx context.Context) ([]*models.TwitchChannel, error)
	CreateTwitchChannel(ctx context.Context, channelId string, channelName string, userId *string) (bool, error)

	GetBotCommand(ctx context.Context, commandName string, twitchChannelId string) (*models.BotCommand, error)
	CreateBotCommand(ctx context.Context, commandName string, commandContent string, twitchChannelId string) (bool, error)
	UpdateBotCommand(ctx context.Context, commandName string, commandContent string, twitchChannelId string) error
	DeleteBotCommand(ctx context.Context, commandName string, twitchChannelId string) error
}