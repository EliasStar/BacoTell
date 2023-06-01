package bacotell

import (
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/EliasStar/BacoTell/internal/proxy"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func connect() {
	session, err := discordgo.New("Bot " + viper.GetString(ConfigBotToken))
	if err != nil {
		discordLogger.Error("cannot connect to discord", "err", err)
		return
	}

	session.UserAgent = "BacoTell " + Version

	discordLogger.Info("registering handlers")
	session.AddHandler(_onConnect)
	session.AddHandler(_onDisconnect)
	session.AddHandler(_onReady)
	session.AddHandler(_onGuildCreate)
	session.AddHandler(_onInteractionCreate)

	discordLogger.Info("opening session")
	if session.Open() != nil {
		discordLogger.Error("failed to open session", "err", err)
		return
	}

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	<-signalChannel

	discordLogger.Info("closing session")
	if session.Close() != nil {
		discordLogger.Error("failed to close session", "err", err)
		return
	}
}

func _onConnect(*discordgo.Session, *discordgo.Connect) {
	discordLogger.Info("connected to discord gateway")
}

func _onDisconnect(*discordgo.Session, *discordgo.Disconnect) {
	discordLogger.Info("disconnected from discord gateway")
}

func _onReady(_ *discordgo.Session, evt *discordgo.Ready) {
	discordLogger.Info("logged in", "username", evt.User.Username, "tag", evt.User.Discriminator)
}

func _onGuildCreate(session *discordgo.Session, guild *discordgo.GuildCreate) {
	guildLogger := discordLogger.With("guild", guild.ID)
	guildLogger.Info("deploying commands")

	deployedCommandData, err := session.ApplicationCommands(session.State.User.ID, guild.ID)
	if err != nil {
		guildLogger.Warn("failed to retrieve deployed commands", "err", err)
		deployedCommandData = nil
	}

	for name, command := range applicationCommands {
		commandLogger := guildLogger.With("command", name)

		localData, err := command.CommandData()
		if err != nil {
			commandLogger.Warn("failed to get command data", "err", err)
			continue
		}

		localData.Name = name

		index := -1
		for i, deployedData := range deployedCommandData {
			if deployedData.Name == name {
				index = i
				break
			}
		}

		if index != -1 {
			deployedData := deployedCommandData[index]

			if _areCommandsEqual(&localData, deployedData) {
				commandLogger.Debug("skipped")
				continue
			}

			if _, err := session.ApplicationCommandEdit(session.State.User.ID, guild.ID, deployedData.ID, &localData); err != nil {
				commandLogger.Warn("failed to update", "err", err)
				continue
			}

			commandLogger.Debug("updated")
			deployedCommandData = append(deployedCommandData[:index], deployedCommandData[index+1:]...)
			continue
		}

		if _, err := session.ApplicationCommandCreate(session.State.User.ID, guild.ID, &localData); err != nil {
			commandLogger.Warn("failed to deploy", "err", err)
			continue
		}

		commandLogger.Debug("deployed")
	}

	for _, deployedData := range deployedCommandData {
		commandLogger := guildLogger.With("command", deployedData.Name)

		if err := session.ApplicationCommandDelete(session.State.User.ID, guild.ID, deployedData.ID); err != nil {
			commandLogger.Warn("failed to delete", "err", err)
			continue
		}

		commandLogger.Debug("deployed")
	}
}

func _onInteractionCreate(session *discordgo.Session, evt *discordgo.InteractionCreate) {
	interactionLogger := discordLogger.With("interaction", evt.ID)
	interactionLogger.Info("handling interaction", "guild", evt.GuildID, "type", evt.Type)

	switch evt.Type {
	case discordgo.InteractionApplicationCommand:
		proxy := proxy.NewExecuteProxy(session, evt.Interaction)

		name := evt.ApplicationCommandData().Name
		commandLogger := interactionLogger.With("command", name)

		cmd, ok := applicationCommands[name]
		if !ok {
			commandLogger.Warn("requested command not loaded")
			proxy.Respond(common.Response{Content: "Command not found: " + name}, true)
			return
		}

		err := cmd.Execute(proxy)
		if err != nil {
			commandLogger.Warn("command execution failed", "err", err)
			return
		}

		commandLogger.Debug("executed")

	case discordgo.InteractionMessageComponent:
		proxy := proxy.NewHandleProxy(session, evt.Interaction)

		id := evt.MessageComponentData().CustomID
		componentLogger := interactionLogger.With("component", id)

		cpt, ok := messageComponents[id]
		if !ok {
			componentLogger.Warn("requested component not loaded")
			proxy.Respond(common.Response{Content: "Component not found: " + id}, true)
			return
		}

		err := cpt.Handle(proxy)
		if err != nil {
			componentLogger.Warn("component handling failed", "err", err)
			return
		}

		componentLogger.Debug("executed")

	case discordgo.InteractionApplicationCommandAutocomplete:
		// TODO
	case discordgo.InteractionModalSubmit:
		// TODO
	}
}

func _areCommandsEqual(cmd1, cmd2 *discordgo.ApplicationCommand) bool {
	equal := true

	equal = equal && reflect.DeepEqual(cmd1.Type, cmd2.Type)
	equal = equal && reflect.DeepEqual(cmd1.Name, cmd2.Name)
	equal = equal && reflect.DeepEqual(cmd1.Description, cmd2.Description)

	equal = equal && reflect.DeepEqual(cmd1.NameLocalizations, cmd2.NameLocalizations)
	equal = equal && reflect.DeepEqual(cmd1.DescriptionLocalizations, cmd2.DescriptionLocalizations)

	equal = equal && reflect.DeepEqual(cmd1.DefaultPermission, cmd2.DefaultPermission)
	equal = equal && reflect.DeepEqual(cmd1.DefaultMemberPermissions, cmd2.DefaultMemberPermissions)
	equal = equal && reflect.DeepEqual(cmd1.DMPermission, cmd2.DMPermission)
	equal = equal && reflect.DeepEqual(cmd1.NSFW, cmd2.NSFW)

	equal = equal && reflect.DeepEqual(cmd1.Options, cmd2.Options)

	return equal
}
