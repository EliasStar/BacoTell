package bacotell

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/EliasStar/BacoTell/internal/proxy"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

// connect starts the Discord bot client and connects to the Discord API.
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

// _onConnect gets called when successfully connected to Discord.
func _onConnect(*discordgo.Session, *discordgo.Connect) {
	discordLogger.Info("connected to discord gateway")
}

// _onDisconnect gets called when connection to Discord is lost.
func _onDisconnect(*discordgo.Session, *discordgo.Disconnect) {
	discordLogger.Info("disconnected from discord gateway")
}

// _onReady gets called when the ready event is received from Discord.
func _onReady(_ *discordgo.Session, evt *discordgo.Ready) {
	discordLogger.Info("logged in", "username", evt.User.Username, "tag", evt.User.Discriminator)
}

// _onGuildCreate gets called for every guild the bot is a member of.
//
// It registers all loaded commands to the guild which aren't already registered and updates the ones that are.
// It also deletes commands which are not loaded.
func _onGuildCreate(session *discordgo.Session, guild *discordgo.GuildCreate) {
	guildLogger := discordLogger.With("guild", guild.ID)
	guildLogger.Info("deploying commands")

	deployedCommandData, err := session.ApplicationCommands(session.State.User.ID, guild.ID)
	if err != nil {
		guildLogger.Warn("failed to retrieve deployed commands", "err", err)
		deployedCommandData = nil
	}

	for name, command := range commands {
		commandLogger := guildLogger.With("command", name)

		localData, err := command.Data()
		if err != nil {
			commandLogger.Warn("failed to get command data", "err", err)
			continue
		}

		index := -1
		for i, deployedData := range deployedCommandData {
			if deployedData.Name == localData.Name {
				index = i
				break
			}
		}

		if index != -1 {
			deployedData := deployedCommandData[index]
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

		commandLogger.Debug("deleted")
	}
}

// _onInteractionCreate gets called when a user uses a command, component or modal.
//
// It looks up the requested command, component or modal and calls the respective handler function.
func _onInteractionCreate(session *discordgo.Session, evt *discordgo.InteractionCreate) {
	interactionLogger := discordLogger.With("interaction", evt.ID)
	interactionLogger.Info("handling interaction", "guild", evt.GuildID, "type", evt.Type)

	switch evt.Type {
	case discordgo.InteractionApplicationCommand:
		proxy := proxy.NewExecuteProxy(session, evt.Interaction)

		name := evt.ApplicationCommandData().Name
		commandLogger := interactionLogger.With("command", name)

		cmd, ok := commands[name]
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

		cpt, ok := components[id]
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

		componentLogger.Debug("handled")

	case discordgo.InteractionApplicationCommandAutocomplete:
		proxy := proxy.NewAutocompleteProxy(session, evt.Interaction)

		name := evt.ApplicationCommandData().Name
		commandLogger := interactionLogger.With("command", name)

		cmd, ok := commands[name]
		if !ok {
			commandLogger.Warn("requested command not loaded")
			proxy.Respond()
			return
		}

		err := cmd.Autocomplete(proxy)
		if err != nil {
			commandLogger.Warn("command autocompletion failed", "err", err)
			return
		}

		commandLogger.Debug("autocompleted")

	case discordgo.InteractionModalSubmit:
		proxy := proxy.NewSubmitProxy(session, evt.Interaction)

		id := evt.ModalSubmitData().CustomID
		modalLogger := interactionLogger.With("modal", id)

		mod, ok := modals[id]
		if !ok {
			modalLogger.Warn("requested modal not loaded")
			proxy.Respond(common.Response{Content: "Modal not found: " + id}, true)
			return
		}

		err := mod.Submit(proxy)
		if err != nil {
			modalLogger.Warn("modal submission failed", "err", err)
			return
		}

		modalLogger.Debug("submitted")
	}
}
