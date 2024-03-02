package discord

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"torchbearer/pkg/services/config"
)

type Service struct {
	logger       *slog.Logger
	cfgManager   config.Dependency
	session      *discordgo.Session
	alreadyBound map[string]bool
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) ConfigFileName() string {
	return "discord.json"
}

func (s *Service) DefaultConfig() (cfg config.Config) {
	cfg.Group("Discord").Set("api key", "change me")

	return
}

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		switch candidate := service.(type) {
		case config.Dependency:
			s.cfgManager = candidate
		}
	}
}

func (s *Service) Name() string {
	return "Discord"
}

func (s *Service) Init(mesh servicemesh.Mesh) {
	s.alreadyBound = make(map[string]bool)
	s.initSession()

	for _, service := range services {
		s.bindHandlers(service)
	}

	// Wait here until CTRL-C or other term signal is received.
	s.logger.Info().Msg("Bot is now running. Press CTRL-C to exit.")
}

func (s *Service) initSession() {
	cfg, err := s.cfgManager.GetConfigByFileName(s.ConfigFileName())
	if err != nil {
		s.logger.Fatal().Msgf("could not load config file: %v", err)
	}

	key := cfg.Group("Discord").GetString("api key")
	if key == "" {
		fp := s.cfgManager.GetFilePath(s.ConfigFileName())
		s.cfgManager.SaveConfigWithFileName(s.ConfigFileName())
		s.logger.Fatal().Msgf("You need to set your discord API key in %q", fp)
	}

	session, err := discordgo.New(fmt.Sprintf("Bot %s", key))
	if err != nil {
		s.logger.Fatal().Msgf("could not init client: %v", err)
	}

	// Open a websocket connection to Discord and begin listening.
	if err := session.Open(); err != nil {
		s.logger.Fatal().Msgf("error opening connection,", err)
	}

	s.session = session
}

func (s *Service) OnServiceAdded(args ...interface{}) {
	for _, candidate := range args {
		if service, ok := candidate.(servicemesh.Service); ok {
			s.bindHandlers(service)
		}
	}
}

func (s *Service) bindHandlers(service servicemesh.Service) {
	for s.session == nil {
		time.Sleep(time.Millisecond * 10)
	}

	if s.alreadyBound == nil {
		s.alreadyBound = make(map[string]bool)
	}

	if _, bound := s.alreadyBound[service.Name()]; bound {
		return
	}

	s.alreadyBound[service.Name()] = true

	s.session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	if candidate, ok := service.(HandlerApplicationCommandPermissionsUpdate); ok {
		s.logger.Info().Msgf("binding 'ApplicationCommandPermissionsUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordApplicationCommandPermissionsUpdate)
	}

	if candidate, ok := service.(HandlerAutoModerationActionExecution); ok {
		s.logger.Info().Msgf("binding 'AutoModerationActionExecution' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordAutoModerationActionExecution)
	}

	if candidate, ok := service.(HandlerAutoModerationRuleCreate); ok {
		s.logger.Info().Msgf("binding 'AutoModerationRuleCreate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordAutoModerationRuleCreate)
	}

	if candidate, ok := service.(HandlerAutoModerationRuleDelete); ok {
		s.logger.Info().Msgf("binding 'AutoModerationRuleDelete' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordAutoModerationRuleDelete)
	}

	if candidate, ok := service.(HandlerAutoModerationRuleUpdate); ok {
		s.logger.Info().Msgf("binding 'AutoModerationRuleUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordAutoModerationRuleUpdate)
	}

	if candidate, ok := service.(HandlerChannelCreate); ok {
		s.logger.Info().Msgf("binding 'ChannelCreate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordChannelCreate)
	}

	if candidate, ok := service.(HandlerChannelDelete); ok {
		s.logger.Info().Msgf("binding 'ChannelDelete' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordChannelDelete)
	}

	if candidate, ok := service.(HandlerChannelPinsUpdate); ok {
		s.logger.Info().Msgf("binding 'ChannelPinsUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordChannelPinsUpdate)
	}

	if candidate, ok := service.(HandlerChannelUpdate); ok {
		s.logger.Info().Msgf("binding 'ChannelUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordChannelUpdate)
	}

	if candidate, ok := service.(HandlerConnect); ok {
		s.logger.Info().Msgf("binding 'Connect' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordConnect)
	}

	if candidate, ok := service.(HandlerDisconnect); ok {
		s.logger.Info().Msgf("binding 'Disconnect' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordDisconnect)
	}

	if candidate, ok := service.(HandlerEvent); ok {
		s.logger.Info().Msgf("binding 'Event' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordEvent)
	}

	if candidate, ok := service.(HandlerGuildBanAdd); ok {
		s.logger.Info().Msgf("binding 'GuildBanAdd' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildBanAdd)
	}

	if candidate, ok := service.(HandlerGuildBanRemove); ok {
		s.logger.Info().Msgf("binding 'GuildBanRemove' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildBanRemove)
	}

	if candidate, ok := service.(HandlerGuildCreate); ok {
		s.logger.Info().Msgf("binding 'GuildCreate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildCreate)
	}

	if candidate, ok := service.(HandlerGuildDelete); ok {
		s.logger.Info().Msgf("binding 'GuildDelete' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildDelete)
	}

	if candidate, ok := service.(HandlerGuildEmojisUpdate); ok {
		s.logger.Info().Msgf("binding 'GuildEmojisUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildEmojisUpdate)
	}

	if candidate, ok := service.(HandlerGuildIntegrationsUpdate); ok {
		s.logger.Info().Msgf("binding 'GuildIntegrationsUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildIntegrationsUpdate)
	}

	if candidate, ok := service.(HandlerGuildMemberAdd); ok {
		s.logger.Info().Msgf("binding 'GuildMemberAdd' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildMemberAdd)
	}

	if candidate, ok := service.(HandlerGuildMemberRemove); ok {
		s.logger.Info().Msgf("binding 'GuildMemberRemove' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildMemberRemove)
	}

	if candidate, ok := service.(HandlerGuildMemberUpdate); ok {
		s.logger.Info().Msgf("binding 'GuildMemberUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildMemberUpdate)
	}

	if candidate, ok := service.(HandlerGuildMembersChunk); ok {
		s.logger.Info().Msgf("binding 'GuildMembersChunk' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildMembersChunk)
	}

	if candidate, ok := service.(HandlerGuildRoleCreate); ok {
		s.logger.Info().Msgf("binding 'GuildRoleCreate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildRoleCreate)
	}

	if candidate, ok := service.(HandlerGuildRoleDelete); ok {
		s.logger.Info().Msgf("binding 'GuildRoleDelete' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildRoleDelete)
	}

	if candidate, ok := service.(HandlerGuildRoleUpdate); ok {
		s.logger.Info().Msgf("binding 'GuildRoleUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildRoleUpdate)
	}

	if candidate, ok := service.(HandlerGuildScheduledEventCreate); ok {
		s.logger.Info().Msgf("binding 'GuildScheduledEventCreate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildScheduledEventCreate)
	}

	if candidate, ok := service.(HandlerGuildScheduledEventDelete); ok {
		s.logger.Info().Msgf("binding 'GuildScheduledEventDelete' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildScheduledEventDelete)
	}

	if candidate, ok := service.(HandlerGuildScheduledEventUpdate); ok {
		s.logger.Info().Msgf("binding 'GuildScheduledEventUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildScheduledEventUpdate)
	}

	if candidate, ok := service.(HandlerGuildScheduledEventUserAdd); ok {
		s.logger.Info().Msgf("binding 'GuildScheduledEventUserAdd' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildScheduledEventUserAdd)
	}

	if candidate, ok := service.(HandlerGuildScheduledEventUserRemove); ok {
		s.logger.Info().Msgf("binding 'GuildScheduledEventUserRemove' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildScheduledEventUserRemove)
	}

	if candidate, ok := service.(HandlerGuildUpdate); ok {
		s.logger.Info().Msgf("binding 'GuildUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordGuildUpdate)
	}

	if candidate, ok := service.(HandlerInteractionCreate); ok {
		s.logger.Info().Msgf("binding 'InteractionCreate' handler for service %q", service.Name())
		candidate.OnDiscordInteractionCreate(s.session)
	}

	if candidate, ok := service.(HandlerInviteCreate); ok {
		s.logger.Info().Msgf("binding 'InviteCreate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordInviteCreate)
	}

	if candidate, ok := service.(HandlerInviteDelete); ok {
		s.logger.Info().Msgf("binding 'InviteDelete' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordInviteDelete)
	}

	if candidate, ok := service.(HandlerMessageCreate); ok {
		s.logger.Info().Msgf("binding 'MessageCreate' handler for service %q", service.Name())
		s.session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent
		s.session.AddHandler(candidate.OnDiscordMessageCreate)
	}
	if candidate, ok := service.(HandlerMessageDelete); ok {
		s.logger.Info().Msgf("binding 'MessageDelete' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordMessageDelete)
	}

	if candidate, ok := service.(HandlerMessageDeleteBulk); ok {
		s.logger.Info().Msgf("binding 'MessageDeleteBulk' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordMessageDeleteBulk)
	}

	if candidate, ok := service.(HandlerMessageReactionAdd); ok {
		s.logger.Info().Msgf("binding 'MessageReactionAdd' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordMessageReactionAdd)
	}

	if candidate, ok := service.(HandlerMessageReactionRemove); ok {
		s.logger.Info().Msgf("binding 'MessageReactionRemove' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordMessageReactionRemove)
	}

	if candidate, ok := service.(HandlerMessageReactionRemoveAll); ok {
		s.logger.Info().Msgf("binding 'MessageReactionRemoveAll' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordMessageReactionRemoveAll)
	}

	if candidate, ok := service.(HandlerMessageUpdate); ok {
		s.logger.Info().Msgf("binding 'MessageUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordMessageUpdate)
	}

	if candidate, ok := service.(HandlerPresenceUpdate); ok {
		s.logger.Info().Msgf("binding 'PresenceUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordPresenceUpdate)
	}

	if candidate, ok := service.(HandlerPresencesReplace); ok {
		s.logger.Info().Msgf("binding 'PresencesReplace' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordPresencesReplace)
	}

	if candidate, ok := service.(HandlerRateLimit); ok {
		s.logger.Info().Msgf("binding 'RateLimit' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordRateLimit)
	}

	if candidate, ok := service.(HandlerReady); ok {
		s.logger.Info().Msgf("binding 'Ready' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordReady)
	}

	if candidate, ok := service.(HandlerResumed); ok {
		s.logger.Info().Msgf("binding 'Resumed' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordResumed)
	}

	if candidate, ok := service.(HandlerStageInstanceEventCreate); ok {
		s.logger.Info().Msgf("binding 'StageInstanceEventCreate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordStageInstanceEventCreate)
	}

	if candidate, ok := service.(HandlerStageInstanceEventDelete); ok {
		s.logger.Info().Msgf("binding 'StageInstanceEventDelete' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordStageInstanceEventDelete)
	}

	if candidate, ok := service.(HandlerStageInstanceEventUpdate); ok {
		s.logger.Info().Msgf("binding 'StageInstanceEventUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordStageInstanceEventUpdate)
	}

	if candidate, ok := service.(HandlerThreadCreate); ok {
		s.logger.Info().Msgf("binding 'ThreadCreate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordThreadCreate)
	}

	if candidate, ok := service.(HandlerThreadDelete); ok {
		s.logger.Info().Msgf("binding 'ThreadDelete' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordThreadDelete)
	}

	if candidate, ok := service.(HandlerThreadListSync); ok {
		s.logger.Info().Msgf("binding 'ThreadListSync' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordThreadListSync)
	}

	if candidate, ok := service.(HandlerThreadMemberUpdate); ok {
		s.logger.Info().Msgf("binding 'ThreadMemberUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordThreadMemberUpdate)
	}

	if candidate, ok := service.(HandlerThreadMembersUpdate); ok {
		s.logger.Info().Msgf("binding 'ThreadMembersUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordThreadMembersUpdate)
	}

	if candidate, ok := service.(HandlerThreadUpdate); ok {
		s.logger.Info().Msgf("binding 'ThreadUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordThreadUpdate)
	}

	if candidate, ok := service.(HandlerTypingStart); ok {
		s.logger.Info().Msgf("binding 'TypingStart' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordTypingStart)
	}

	if candidate, ok := service.(HandlerUserUpdate); ok {
		s.logger.Info().Msgf("binding 'UserUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordUserUpdate)
	}

	if candidate, ok := service.(HandlerVoiceServerUpdate); ok {
		s.logger.Info().Msgf("binding 'VoiceServerUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordVoiceServerUpdate)
	}

	if candidate, ok := service.(HandlerVoiceStateUpdate); ok {
		s.logger.Info().Msgf("binding 'VoiceStateUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordVoiceStateUpdate)
	}

	if candidate, ok := service.(HandlerWebhooksUpdate); ok {
		s.logger.Info().Msgf("binding 'WebhooksUpdate' handler for service %q", service.Name())
		s.session.AddHandler(candidate.OnDiscordWebhooksUpdate)
	}
}

func (s *Service) OnDiscordMessageCreate(session *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == session.State.User.ID {
		return
	}

	// In this example, we only care about messages that are "ping".
	if m.Content != "ping" {
		return
	}

	// We create the private channel with the user who sent the message.
	channel, err := session.UserChannelCreate(m.Author.ID)
	if err != nil {
		// If an error occurred, we failed to create the channel.
		//
		// Some common causes are:
		// 1. We don't share a server with the user (not possible here).
		// 2. We opened enough DM channels quickly enough for Discord to
		//    label us as abusing the endpoint, blocking us from opening
		//    new ones.
		fmt.Println("error creating channel:", err)
		session.ChannelMessageSend(
			m.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}
	// Then we send the message through the channel we created.
	_, err = session.ChannelMessageSend(channel.ID, "Pong!")
	if err != nil {
		// If an error occurred, we failed to send the message.
		//
		// It may occur either when we do not share a server with the
		// user (highly unlikely as we just received a message) or
		// the user disabled DM in their settings (more likely).
		fmt.Println("error sending DM message:", err)
		session.ChannelMessageSend(
			m.ChannelID,
			"Failed to send you a DM. "+
				"Did you disable DM in your privacy settings?",
		)
	}
}
