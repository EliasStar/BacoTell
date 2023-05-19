package discord

import (
	"reflect"

	"github.com/EliasStar/BacoTell/internal/loader"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/bwmarrin/discordgo"
)

var (
	commandDataCache []discordgo.ApplicationCommand
	commandCache     = make(map[string]provider.Command)
	componentCache   = make(map[string]provider.Component)
)

func interactionOnConnect(*discordgo.Session, *discordgo.Connect) {
	for prefix, cmds := range loader.GetApplicationCommands() {
		for _, cmd := range cmds {
			data, err := cmd.CommandData()
			if err != nil {
				logger.Warn("cannot get command data", "prefix", prefix, "err", err)
				continue
			}

			data.Name = prefix + "-" + data.Name
			commandDataCache = append(commandDataCache, data)
			commandCache[data.Name] = cmd
		}
	}

	for prefix, cpts := range loader.GetMessageComponents() {
		for _, cpt := range cpts {
			id, err := cpt.CustomId()
			if err != nil {
				logger.Warn("cannot get custom id", "prefix", prefix, "err", err)
				continue
			}

			componentCache[prefix+"-"+id] = cpt
		}
	}
}

func interactionOnGuildCreate(session *discordgo.Session, guild *discordgo.GuildCreate) {
	logger.Info("deploying commands", "guild", guild.ID)

	deployedCommandData, err := session.ApplicationCommands(session.State.User.ID, guild.ID)
	if err != nil {
		logger.Warn("failed to retrieve deployed commands", "err", err)
		deployedCommandData = nil
	}

	for _, localData := range commandDataCache {
		index := -1
		for i, deployedData := range deployedCommandData {
			if localData.Name == deployedData.Name {
				index = i
				break
			}
		}

		if index != -1 {
			deployedData := deployedCommandData[index]

			if isCommandDataEqual(&localData, deployedData) {
				logger.Info("skipped", "command", localData.Name)
				continue
			}

			if _, err := session.ApplicationCommandEdit(session.State.User.ID, guild.ID, deployedData.ID, &localData); err != nil {
				logger.Warn("failed to update", "command", localData.Name, "err", err)
				continue
			}

			deployedCommandData = append(deployedCommandData[:index], deployedCommandData[index+1:]...)
			continue
		}

		if _, err := session.ApplicationCommandCreate(session.State.User.ID, guild.ID, &localData); err != nil {
			logger.Warn("failed to deploy", "command", localData.Name, "err", err)
		}
	}

	for _, deployedData := range deployedCommandData {
		if err := session.ApplicationCommandDelete(session.State.User.ID, guild.ID, deployedData.ID); err != nil {
			logger.Warn("failed to delete:", "command", deployedData.Name, "err", err)
		}
	}
}

func interactionOnInteractionCreate(session *discordgo.Session, evt *discordgo.InteractionCreate) {
	logger.Info("handling interaction", "type", evt.Type, "guild", evt.GuildID)

	switch evt.Type {
	case discordgo.InteractionApplicationCommand:
		name := evt.ApplicationCommandData().Name
		cmd, ok := commandCache[name]
		if !ok {
			logger.Warn("cannot execute nonexistent command", "command", name)
			RespondEphemeral(session, evt.Interaction, "Command not found: "+name)
			return
		}

		err := cmd.Execute(nil)
		if err != nil {
			logger.Warn("command execution failed", "command", name, "err", err)
		}

	case discordgo.InteractionMessageComponent:
		id := evt.MessageComponentData().CustomID
		cpt, ok := componentCache[id]
		if !ok {
			logger.Warn("cannot handle nonexistent component", "component", id)
			RespondEphemeral(session, evt.Interaction, "Component not found: "+id)
			return
		}

		err := cpt.Handle(nil)
		if err != nil {
			logger.Warn("component handling failed", "component", id, "err", err)
		}

	case discordgo.InteractionApplicationCommandAutocomplete:
		// TODO
	case discordgo.InteractionModalSubmit:
		// TODO
	}
}

func isCommandDataEqual(cmd1, cmd2 *discordgo.ApplicationCommand) bool {
	equal := cmd1.Type == cmd2.Type && cmd1.Name == cmd2.Name && cmd1.Description == cmd2.Description

	equal = equal && reflect.DeepEqual(cmd1.NameLocalizations, cmd2.NameLocalizations)
	equal = equal && reflect.DeepEqual(cmd1.DescriptionLocalizations, cmd2.DescriptionLocalizations)

	equal = equal && reflect.DeepEqual(cmd1.DefaultMemberPermissions, cmd2.DefaultMemberPermissions)
	equal = equal && reflect.DeepEqual(cmd1.DMPermission, cmd2.DMPermission)

	return equal && reflect.DeepEqual(cmd1.Options, cmd2.Options)
}
