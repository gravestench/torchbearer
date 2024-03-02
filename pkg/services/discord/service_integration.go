package discord

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordAPI interface {
}

type HandlerApplicationCommandPermissionsUpdate interface {
	OnDiscordApplicationCommandPermissionsUpdate(session *discordgo.Session, e *discordgo.ApplicationCommandPermissionsUpdate)
}

type HandlerAutoModerationActionExecution interface {
	OnDiscordAutoModerationActionExecution(session *discordgo.Session, e *discordgo.AutoModerationActionExecution)
}

type HandlerAutoModerationRuleCreate interface {
	OnDiscordAutoModerationRuleCreate(session *discordgo.Session, e *discordgo.AutoModerationRuleCreate)
}

type HandlerAutoModerationRuleDelete interface {
	OnDiscordAutoModerationRuleDelete(session *discordgo.Session, e *discordgo.AutoModerationRuleDelete)
}

type HandlerAutoModerationRuleUpdate interface {
	OnDiscordAutoModerationRuleUpdate(session *discordgo.Session, e *discordgo.AutoModerationRuleUpdate)
}

type HandlerChannelCreate interface {
	OnDiscordChannelCreate(session *discordgo.Session, e *discordgo.ChannelCreate)
}

type HandlerChannelDelete interface {
	OnDiscordChannelDelete(session *discordgo.Session, e *discordgo.ChannelDelete)
}

type HandlerChannelPinsUpdate interface {
	OnDiscordChannelPinsUpdate(session *discordgo.Session, e *discordgo.ChannelPinsUpdate)
}

type HandlerChannelUpdate interface {
	OnDiscordChannelUpdate(session *discordgo.Session, e *discordgo.ChannelUpdate)
}

type HandlerConnect interface {
	OnDiscordConnect(session *discordgo.Session, e *discordgo.Connect)
}

type HandlerDisconnect interface {
	OnDiscordDisconnect(session *discordgo.Session, e *discordgo.Disconnect)
}

type HandlerEvent interface {
	OnDiscordEvent(session *discordgo.Session, e *discordgo.Event)
}

type HandlerGuildBanAdd interface {
	OnDiscordGuildBanAdd(session *discordgo.Session, e *discordgo.GuildBanAdd)
}

type HandlerGuildBanRemove interface {
	OnDiscordGuildBanRemove(session *discordgo.Session, e *discordgo.GuildBanRemove)
}

type HandlerGuildCreate interface {
	OnDiscordGuildCreate(session *discordgo.Session, e *discordgo.GuildCreate)
}

type HandlerGuildDelete interface {
	OnDiscordGuildDelete(session *discordgo.Session, e *discordgo.GuildDelete)
}

type HandlerGuildEmojisUpdate interface {
	OnDiscordGuildEmojisUpdate(session *discordgo.Session, e *discordgo.GuildEmojisUpdate)
}

type HandlerGuildIntegrationsUpdate interface {
	OnDiscordGuildIntegrationsUpdate(session *discordgo.Session, e *discordgo.GuildIntegrationsUpdate)
}

type HandlerGuildMemberAdd interface {
	OnDiscordGuildMemberAdd(session *discordgo.Session, e *discordgo.GuildMemberAdd)
}

type HandlerGuildMemberRemove interface {
	OnDiscordGuildMemberRemove(session *discordgo.Session, e *discordgo.GuildMemberRemove)
}

type HandlerGuildMemberUpdate interface {
	OnDiscordGuildMemberUpdate(session *discordgo.Session, e *discordgo.GuildMemberUpdate)
}

type HandlerGuildMembersChunk interface {
	OnDiscordGuildMembersChunk(session *discordgo.Session, e *discordgo.GuildMembersChunk)
}

type HandlerGuildRoleCreate interface {
	OnDiscordGuildRoleCreate(session *discordgo.Session, e *discordgo.GuildRoleCreate)
}

type HandlerGuildRoleDelete interface {
	OnDiscordGuildRoleDelete(session *discordgo.Session, e *discordgo.GuildRoleDelete)
}

type HandlerGuildRoleUpdate interface {
	OnDiscordGuildRoleUpdate(session *discordgo.Session, e *discordgo.GuildRoleUpdate)
}

type HandlerGuildScheduledEventCreate interface {
	OnDiscordGuildScheduledEventCreate(session *discordgo.Session, e *discordgo.GuildScheduledEventCreate)
}

type HandlerGuildScheduledEventDelete interface {
	OnDiscordGuildScheduledEventDelete(session *discordgo.Session, e *discordgo.GuildScheduledEventDelete)
}

type HandlerGuildScheduledEventUpdate interface {
	OnDiscordGuildScheduledEventUpdate(session *discordgo.Session, e *discordgo.GuildScheduledEventUpdate)
}

type HandlerGuildScheduledEventUserAdd interface {
	OnDiscordGuildScheduledEventUserAdd(session *discordgo.Session, e *discordgo.GuildScheduledEventUserAdd)
}

type HandlerGuildScheduledEventUserRemove interface {
	OnDiscordGuildScheduledEventUserRemove(session *discordgo.Session, e *discordgo.GuildScheduledEventUserRemove)
}

type HandlerGuildUpdate interface {
	OnDiscordGuildUpdate(session *discordgo.Session, e *discordgo.GuildUpdate)
}

type HandlerInteractionCreate interface {
	OnDiscordInteractionCreate(session *discordgo.Session)
}

type HandlerInviteCreate interface {
	OnDiscordInviteCreate(session *discordgo.Session, e *discordgo.InviteCreate)
}

type HandlerInviteDelete interface {
	OnDiscordInviteDelete(session *discordgo.Session, e *discordgo.InviteDelete)
}

type HandlerMessageCreate interface {
	OnDiscordMessageCreate(session *discordgo.Session, e *discordgo.MessageCreate)
}

type HandlerMessageDelete interface {
	OnDiscordMessageDelete(session *discordgo.Session, e *discordgo.MessageDelete)
}

type HandlerMessageDeleteBulk interface {
	OnDiscordMessageDeleteBulk(session *discordgo.Session, e *discordgo.MessageDeleteBulk)
}

type HandlerMessageReactionAdd interface {
	OnDiscordMessageReactionAdd(session *discordgo.Session, e *discordgo.MessageReactionAdd)
}

type HandlerMessageReactionRemove interface {
	OnDiscordMessageReactionRemove(session *discordgo.Session, e *discordgo.MessageReactionRemove)
}

type HandlerMessageReactionRemoveAll interface {
	OnDiscordMessageReactionRemoveAll(session *discordgo.Session, e *discordgo.MessageReactionRemoveAll)
}

type HandlerMessageUpdate interface {
	OnDiscordMessageUpdate(session *discordgo.Session, e *discordgo.MessageUpdate)
}

type HandlerPresenceUpdate interface {
	OnDiscordPresenceUpdate(session *discordgo.Session, e *discordgo.PresenceUpdate)
}

type HandlerPresencesReplace interface {
	OnDiscordPresencesReplace(session *discordgo.Session, e *discordgo.PresencesReplace)
}

type HandlerRateLimit interface {
	OnDiscordRateLimit(session *discordgo.Session, e *discordgo.RateLimit)
}

type HandlerReady interface {
	OnDiscordReady(session *discordgo.Session, e *discordgo.Ready)
}

type HandlerResumed interface {
	OnDiscordResumed(session *discordgo.Session, e *discordgo.Resumed)
}

type HandlerStageInstanceEventCreate interface {
	OnDiscordStageInstanceEventCreate(session *discordgo.Session, e *discordgo.StageInstanceEventCreate)
}

type HandlerStageInstanceEventDelete interface {
	OnDiscordStageInstanceEventDelete(session *discordgo.Session, e *discordgo.StageInstanceEventDelete)
}

type HandlerStageInstanceEventUpdate interface {
	OnDiscordStageInstanceEventUpdate(session *discordgo.Session, e *discordgo.StageInstanceEventUpdate)
}

type HandlerThreadCreate interface {
	OnDiscordThreadCreate(session *discordgo.Session, e *discordgo.ThreadCreate)
}

type HandlerThreadDelete interface {
	OnDiscordThreadDelete(session *discordgo.Session, e *discordgo.ThreadDelete)
}

type HandlerThreadListSync interface {
	OnDiscordThreadListSync(session *discordgo.Session, e *discordgo.ThreadListSync)
}

type HandlerThreadMemberUpdate interface {
	OnDiscordThreadMemberUpdate(session *discordgo.Session, e *discordgo.ThreadMemberUpdate)
}

type HandlerThreadMembersUpdate interface {
	OnDiscordThreadMembersUpdate(session *discordgo.Session, e *discordgo.ThreadMembersUpdate)
}

type HandlerThreadUpdate interface {
	OnDiscordThreadUpdate(session *discordgo.Session, e *discordgo.ThreadUpdate)
}

type HandlerTypingStart interface {
	OnDiscordTypingStart(session *discordgo.Session, e *discordgo.TypingStart)
}

type HandlerUserUpdate interface {
	OnDiscordUserUpdate(session *discordgo.Session, e *discordgo.UserUpdate)
}

type HandlerVoiceServerUpdate interface {
	OnDiscordVoiceServerUpdate(session *discordgo.Session, e *discordgo.VoiceServerUpdate)
}

type HandlerVoiceStateUpdate interface {
	OnDiscordVoiceStateUpdate(session *discordgo.Session, e *discordgo.VoiceStateUpdate)
}

type HandlerWebhooksUpdate interface {
	OnDiscordWebhooksUpdate(session *discordgo.Session, e *discordgo.WebhooksUpdate)
}
