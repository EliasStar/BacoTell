// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: proto/discord_channel.proto

package discordpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Discriminator string `protobuf:"bytes,3,opt,name=discriminator,proto3" json:"discriminator,omitempty"`
	Email         string `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
	Avatar        string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Banner        string `protobuf:"bytes,8,opt,name=banner,proto3" json:"banner,omitempty"`
	AccentColor   uint32 `protobuf:"varint,9,opt,name=accent_color,json=accentColor,proto3" json:"accent_color,omitempty"`
	Locale        string `protobuf:"bytes,10,opt,name=locale,proto3" json:"locale,omitempty"`
	Bot           bool   `protobuf:"varint,5,opt,name=bot,proto3" json:"bot,omitempty"`
	System        bool   `protobuf:"varint,6,opt,name=system,proto3" json:"system,omitempty"`
	MfaEnabled    bool   `protobuf:"varint,7,opt,name=mfa_enabled,json=mfaEnabled,proto3" json:"mfa_enabled,omitempty"`
	Verified      bool   `protobuf:"varint,11,opt,name=verified,proto3" json:"verified,omitempty"`
	PremiumType   uint32 `protobuf:"varint,14,opt,name=premium_type,json=premiumType,proto3" json:"premium_type,omitempty"`
	Flags         uint64 `protobuf:"varint,13,opt,name=flags,proto3" json:"flags,omitempty"`
	PublicFlags   uint64 `protobuf:"varint,15,opt,name=public_flags,json=publicFlags,proto3" json:"public_flags,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discord_channel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_channel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_discord_channel_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetDiscriminator() string {
	if x != nil {
		return x.Discriminator
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetBanner() string {
	if x != nil {
		return x.Banner
	}
	return ""
}

func (x *User) GetAccentColor() uint32 {
	if x != nil {
		return x.AccentColor
	}
	return 0
}

func (x *User) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *User) GetBot() bool {
	if x != nil {
		return x.Bot
	}
	return false
}

func (x *User) GetSystem() bool {
	if x != nil {
		return x.System
	}
	return false
}

func (x *User) GetMfaEnabled() bool {
	if x != nil {
		return x.MfaEnabled
	}
	return false
}

func (x *User) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *User) GetPremiumType() uint32 {
	if x != nil {
		return x.PremiumType
	}
	return 0
}

func (x *User) GetFlags() uint64 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *User) GetPublicFlags() uint64 {
	if x != nil {
		return x.PublicFlags
	}
	return 0
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Color        uint32    `protobuf:"varint,3,opt,name=color,proto3" json:"color,omitempty"`
	UnicodeEmoji string    `protobuf:"bytes,6,opt,name=unicode_emoji,json=unicodeEmoji,proto3" json:"unicode_emoji,omitempty"`
	Icon         string    `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	Managed      bool      `protobuf:"varint,9,opt,name=managed,proto3" json:"managed,omitempty"`
	Mentionable  bool      `protobuf:"varint,10,opt,name=mentionable,proto3" json:"mentionable,omitempty"`
	Hoist        bool      `protobuf:"varint,4,opt,name=hoist,proto3" json:"hoist,omitempty"`
	Position     uint32    `protobuf:"varint,7,opt,name=position,proto3" json:"position,omitempty"`
	Permissions  int64     `protobuf:"varint,8,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Tags         *RoleTags `protobuf:"bytes,11,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discord_channel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_channel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_proto_discord_channel_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetColor() uint32 {
	if x != nil {
		return x.Color
	}
	return 0
}

func (x *Role) GetUnicodeEmoji() string {
	if x != nil {
		return x.UnicodeEmoji
	}
	return ""
}

func (x *Role) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Role) GetManaged() bool {
	if x != nil {
		return x.Managed
	}
	return false
}

func (x *Role) GetMentionable() bool {
	if x != nil {
		return x.Mentionable
	}
	return false
}

func (x *Role) GetHoist() bool {
	if x != nil {
		return x.Hoist
	}
	return false
}

func (x *Role) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Role) GetPermissions() int64 {
	if x != nil {
		return x.Permissions
	}
	return 0
}

func (x *Role) GetTags() *RoleTags {
	if x != nil {
		return x.Tags
	}
	return nil
}

type RoleTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotId                 string `protobuf:"bytes,1,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	IntegrationId         string `protobuf:"bytes,2,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	SubscriptionListingId string `protobuf:"bytes,4,opt,name=subscription_listing_id,json=subscriptionListingId,proto3" json:"subscription_listing_id,omitempty"`
	PremiumSubscriber     bool   `protobuf:"varint,3,opt,name=premium_subscriber,json=premiumSubscriber,proto3" json:"premium_subscriber,omitempty"`
	AvailableForPurchase  bool   `protobuf:"varint,5,opt,name=available_for_purchase,json=availableForPurchase,proto3" json:"available_for_purchase,omitempty"`
	GuildConnections      bool   `protobuf:"varint,6,opt,name=guild_connections,json=guildConnections,proto3" json:"guild_connections,omitempty"`
}

func (x *RoleTags) Reset() {
	*x = RoleTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discord_channel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleTags) ProtoMessage() {}

func (x *RoleTags) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_channel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleTags.ProtoReflect.Descriptor instead.
func (*RoleTags) Descriptor() ([]byte, []int) {
	return file_proto_discord_channel_proto_rawDescGZIP(), []int{2}
}

func (x *RoleTags) GetBotId() string {
	if x != nil {
		return x.BotId
	}
	return ""
}

func (x *RoleTags) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *RoleTags) GetSubscriptionListingId() string {
	if x != nil {
		return x.SubscriptionListingId
	}
	return ""
}

func (x *RoleTags) GetPremiumSubscriber() bool {
	if x != nil {
		return x.PremiumSubscriber
	}
	return false
}

func (x *RoleTags) GetAvailableForPurchase() bool {
	if x != nil {
		return x.AvailableForPurchase
	}
	return false
}

func (x *RoleTags) GetGuildConnections() bool {
	if x != nil {
		return x.GuildConnections
	}
	return false
}

type Channel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Channel) Reset() {
	*x = Channel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discord_channel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Channel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Channel) ProtoMessage() {}

func (x *Channel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_channel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Channel.ProtoReflect.Descriptor instead.
func (*Channel) Descriptor() ([]byte, []int) {
	return file_proto_discord_channel_proto_rawDescGZIP(), []int{3}
}

type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename     string  `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Url          string  `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	ProxyUrl     string  `protobuf:"bytes,7,opt,name=proxy_url,json=proxyUrl,proto3" json:"proxy_url,omitempty"`
	ContentType  string  `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Size         uint32  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Height       uint32  `protobuf:"varint,8,opt,name=height,proto3" json:"height,omitempty"`
	Width        uint32  `protobuf:"varint,9,opt,name=width,proto3" json:"width,omitempty"`
	Ephemeral    bool    `protobuf:"varint,10,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	DurationSecs float64 `protobuf:"fixed64,11,opt,name=duration_secs,json=durationSecs,proto3" json:"duration_secs,omitempty"`
	Waveform     string  `protobuf:"bytes,12,opt,name=waveform,proto3" json:"waveform,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discord_channel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_channel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_proto_discord_channel_proto_rawDescGZIP(), []int{4}
}

func (x *Attachment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attachment) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Attachment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Attachment) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Attachment) GetProxyUrl() string {
	if x != nil {
		return x.ProxyUrl
	}
	return ""
}

func (x *Attachment) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Attachment) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Attachment) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Attachment) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Attachment) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

func (x *Attachment) GetDurationSecs() float64 {
	if x != nil {
		return x.DurationSecs
	}
	return 0
}

func (x *Attachment) GetWaveform() string {
	if x != nil {
		return x.Waveform
	}
	return ""
}

var File_proto_discord_channel_proto protoreflect.FileDescriptor

var file_proto_discord_channel_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x9c, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x62, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x66, 0x61, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x6d, 0x66, 0x61, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x65,
	0x6d, 0x69, 0x75, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x22, 0xb0, 0x02, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x6e, 0x69, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x75, 0x6e, 0x69, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x68, 0x6f, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x68, 0x6f,
	0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x61,
	0x67, 0x73, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x92, 0x02, 0x0a, 0x08, 0x52, 0x6f, 0x6c,
	0x65, 0x54, 0x61, 0x67, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x70,
	0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x67, 0x75, 0x69,
	0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x09, 0x0a,
	0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xcd, 0x02, 0x0a, 0x0a, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x55, 0x72, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x70, 0x68,
	0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x70,
	0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x61, 0x76, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x77, 0x61, 0x76, 0x65, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6c, 0x69, 0x61, 0x73, 0x53, 0x74, 0x61, 0x72,
	0x2f, 0x42, 0x61, 0x63, 0x6f, 0x54, 0x65, 0x6c, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_discord_channel_proto_rawDescOnce sync.Once
	file_proto_discord_channel_proto_rawDescData = file_proto_discord_channel_proto_rawDesc
)

func file_proto_discord_channel_proto_rawDescGZIP() []byte {
	file_proto_discord_channel_proto_rawDescOnce.Do(func() {
		file_proto_discord_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_discord_channel_proto_rawDescData)
	})
	return file_proto_discord_channel_proto_rawDescData
}

var file_proto_discord_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_discord_channel_proto_goTypes = []interface{}{
	(*User)(nil),       // 0: discord.User
	(*Role)(nil),       // 1: discord.Role
	(*RoleTags)(nil),   // 2: discord.RoleTags
	(*Channel)(nil),    // 3: discord.Channel
	(*Attachment)(nil), // 4: discord.Attachment
}
var file_proto_discord_channel_proto_depIdxs = []int32{
	2, // 0: discord.Role.tags:type_name -> discord.RoleTags
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_discord_channel_proto_init() }
func file_proto_discord_channel_proto_init() {
	if File_proto_discord_channel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_discord_channel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_discord_channel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_discord_channel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleTags); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_discord_channel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Channel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_discord_channel_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_discord_channel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_discord_channel_proto_goTypes,
		DependencyIndexes: file_proto_discord_channel_proto_depIdxs,
		MessageInfos:      file_proto_discord_channel_proto_msgTypes,
	}.Build()
	File_proto_discord_channel_proto = out.File
	file_proto_discord_channel_proto_rawDesc = nil
	file_proto_discord_channel_proto_goTypes = nil
	file_proto_discord_channel_proto_depIdxs = nil
}
