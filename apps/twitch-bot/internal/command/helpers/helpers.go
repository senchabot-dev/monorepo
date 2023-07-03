package helpers

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v3"
	"github.com/senchabot-dev/monorepo/apps/twitch-bot/internal/models"
	"github.com/senchabot-dev/monorepo/apps/twitch-bot/internal/service"
)

const (
	max = 70
	min = 18

	maxCommandNameLength    = 50
	maxCommandContentLength = 400

	maxAliasParamLength = 4
)

func FormatCommandContent(commandData *models.BotCommand, message twitch.PrivateMessage) string {
	msgContent := commandData.CommandContent

	userName := message.User.DisplayName
	dateTemplate := "02/01/2006"

	stringTemplates := map[string]string{
		"{user.name}":     userName,
		"{cmd.author}":    userName,
		"{random_number}": strconv.Itoa(rand.Intn(max-min) + min),
		"{date}":          message.Time.Format(dateTemplate),
		"{cmd.date}":      commandData.CreatedAt.Format(dateTemplate),
		"{channel.name}":  message.Channel,

		// we will keep these old string templates used in commands for a while for backward compatibility.
		"{user_name}": userName,
		"{cmd_date}":  commandData.CreatedAt.Format(dateTemplate),
	}

	for k, v := range stringTemplates {
		msgContent = strings.ReplaceAll(msgContent, k, v)
	}

	return msgContent
}

func ParseMessage(message string) (string, []string) {
	var splitMsg = strings.Split(message, " ")
	var cmdName = splitMsg[0]
	var params []string

	// Check if first word is a @mention
	if strings.HasPrefix(cmdName, "@") && len(params) > 2 {
		cmdName = splitMsg[1]
	} else {
		params = splitMsg[1:]
	}

	if !CheckIfCommand(cmdName) {
		return "", nil
	}

	cmdName = strings.TrimPrefix(cmdName, "!")

	return cmdName, params
}

func CheckIfCommand(param string) bool {
	return strings.HasPrefix(param, "!")
}

func CanExecuteCommand(context context.Context, service service.Service, message twitch.PrivateMessage) bool {
	// broadcaster can run the command
	if isBroadcaster(message.Tags["badges"]) {
		return true
	}

	// moderator can run the command
	if isModerator(message.Tags["badges"]) {
		check := service.CheckConfig(context, message.RoomID, "mods_manage_cmds_enabled", "1")
		return check
	}

	// everyone else can't run the command
	return false
}

func isBroadcaster(badgeTags string) bool {
	return strings.Contains(badgeTags, "broadcaster")
}

func isModerator(badgeTags string) bool {
	return strings.Contains(badgeTags, "moderator")
}

func MakeUniqueArray(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			entry = strings.TrimPrefix(entry, "!")
			list = append(list, entry)
		}
	}
	return list
}

func GetCommandCreateUpdateParams(params []string) (string, string, bool) {
	if check := ValidateCommandCreateParamsLength(params); !check {
		return "", "", false
	}

	var commandName = strings.ToLower(params[0])
	var commandContent = strings.Join(params[1:], " ")

	commandName = TrimExclamationPrefix(commandName)

	return commandName, commandContent, true
}

func GetAliasCommandCreateParams(params []string) (string, []string, bool) {
	if check := ValidateCommandCreateParamsLength(params); !check {
		return "", nil, false
	}

	command := strings.ToLower(params[0])
	params = params[1:]

	command = TrimExclamationPrefix(command)
	aliasCommands := MakeUniqueArray(params)

	return command, aliasCommands, true
}

func ValidateCommandCreateParamsLength(params []string) bool {
	fmt.Println("ValidateCommandCreateParamsLength", params)
	return len(params) >= 2
}

func ValidateAliasCommandsLength(aliasCommands []string) (string, bool) {
	if len(aliasCommands) > maxAliasParamLength {
		return fmt.Sprintf("Command Aliases length must be no more than %d", maxAliasParamLength), false
	}

	return "", true
}

func ValidateCommandCreateParams(commandName string, commandContent string) (string, bool) {
	if len(commandName) > maxCommandNameLength {
		return fmt.Sprintf("Command Name length must be no more than %d chars", maxCommandNameLength), false
	}
	if infoText, check := ValidateCommandContentLength(commandContent); !check {
		return infoText, check
	}

	return "", true
}

func ValidateCommandContentLength(commandContent string) (string, bool) {
	if len(commandContent) > maxCommandContentLength {
		return fmt.Sprintf("Command Content length must be no more than %d chars", maxCommandContentLength), false
	}

	return "", true
}

func ValidateCommandDeleteParamsLength(params []string) bool {
	return len(params) == 1
}

func IsCommandParamsLengthEqualToOne(params []string) bool {
	return len(params) == 1
}

func TrimExclamationPrefix(commandName string) string {
	return strings.TrimPrefix(commandName, "!")
}
