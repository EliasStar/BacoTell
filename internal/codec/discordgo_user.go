package codec

import (
	"time"

	"github.com/EliasStar/BacoTell/internal/proto/discordgopb"
	util "github.com/EliasStar/BacoTell/pkg/bacotell_util"
	"github.com/bwmarrin/discordgo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func encodeUser(user *discordgo.User) *discordgopb.User {
	if user == nil {
		return nil
	}

	return &discordgopb.User{
		Id:       user.ID,
		Email:    user.Email,
		Username: user.Username,

		Avatar:        user.Avatar,
		Locale:        user.Locale,
		Discriminator: user.Discriminator,
		Token:         user.Token,

		Verified:   user.Verified,
		MfaEnabled: user.MFAEnabled,

		Banner:      user.Banner,
		AccentColor: int32(user.AccentColor),

		Bot:         user.Bot,
		PublicFlags: int32(user.PublicFlags),
		PremiumType: int32(user.PremiumType),
		System:      user.System,
		Flags:       int32(user.Flags),
	}
}

func encodeUsers(users []*discordgo.User) []*discordgopb.User {
	result := make([]*discordgopb.User, len(users))

	for i, user := range users {
		result[i] = encodeUser(user)
	}

	return result
}

func decodeUser(user *discordgopb.User) *discordgo.User {
	if user == nil {
		return nil
	}

	return &discordgo.User{
		ID:       user.Id,
		Email:    user.Email,
		Username: user.Username,

		Avatar:        user.Avatar,
		Locale:        user.Locale,
		Discriminator: user.Discriminator,
		Token:         user.Token,

		Verified:   user.Verified,
		MFAEnabled: user.MfaEnabled,

		Banner:      user.Banner,
		AccentColor: int(user.AccentColor),

		Bot:         user.Bot,
		PublicFlags: discordgo.UserFlags(user.PublicFlags),
		PremiumType: int(user.PremiumType),
		System:      user.System,
		Flags:       int(user.Flags),
	}
}

func decodeUsers(users []*discordgopb.User) []*discordgo.User {
	result := make([]*discordgo.User, len(users))

	for i, user := range users {
		result[i] = decodeUser(user)
	}

	return result
}

func encodeMember(member *discordgo.Member) *discordgopb.Member {
	if member == nil {
		return nil
	}

	var premiumSince *timestamppb.Timestamp
	if member.PremiumSince != nil {
		premiumSince = timestamppb.New(*member.PremiumSince)
	}

	var communicationDisabledUntil *timestamppb.Timestamp
	if member.CommunicationDisabledUntil != nil {
		communicationDisabledUntil = timestamppb.New(*member.CommunicationDisabledUntil)
	}

	return &discordgopb.Member{
		GuildId:  member.GuildID,
		JoinedAt: timestamppb.New(member.JoinedAt),
		Nick:     member.Nick,

		Deaf: member.Deaf,
		Mute: member.Mute,

		Avatar:       member.Avatar,
		User:         encodeUser(member.User),
		Roles:        member.Roles,
		PremiumSince: premiumSince,

		Pending:                    member.Pending,
		Permissions:                member.Permissions,
		CommunicationDisabledUntil: communicationDisabledUntil,
	}
}

func decodeMember(member *discordgopb.Member) *discordgo.Member {
	if member == nil {
		return nil
	}

	var premiumSince *time.Time
	if member.PremiumSince != nil {
		premiumSince = util.Ptr(member.PremiumSince.AsTime())
	}

	var communicationDisabledUntil *time.Time
	if member.CommunicationDisabledUntil != nil {
		communicationDisabledUntil = util.Ptr(member.CommunicationDisabledUntil.AsTime())
	}

	return &discordgo.Member{
		GuildID:  member.GuildId,
		JoinedAt: member.JoinedAt.AsTime(),
		Nick:     member.Nick,

		Deaf: member.Deaf,
		Mute: member.Mute,

		Avatar:       member.Avatar,
		User:         decodeUser(member.User),
		Roles:        member.Roles,
		PremiumSince: premiumSince,

		Pending:                    member.Pending,
		Permissions:                member.Permissions,
		CommunicationDisabledUntil: communicationDisabledUntil,
	}
}

func encodeRole(role *discordgo.Role) *discordgopb.Role {
	if role == nil {
		return nil
	}

	return &discordgopb.Role{
		Id:   role.ID,
		Name: role.Name,

		Managed:     role.Managed,
		Mentionable: role.Mentionable,
		Hoist:       role.Hoist,

		Color: int32(role.Color),

		Position:    int32(role.Position),
		Permissions: role.Permissions,
	}
}

func decodeRole(role *discordgopb.Role) *discordgo.Role {
	if role == nil {
		return nil
	}

	return &discordgo.Role{
		ID:   role.Id,
		Name: role.Name,

		Managed:     role.Managed,
		Mentionable: role.Mentionable,
		Hoist:       role.Hoist,

		Color: int(role.Color),

		Position:    int(role.Position),
		Permissions: role.Permissions,
	}
}
