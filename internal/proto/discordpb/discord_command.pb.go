// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: proto/discord_command.proto

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

type ApplicationCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                     uint32                      `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                     string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NameLocalizations        map[string]string           `protobuf:"bytes,3,rep,name=name_localizations,json=nameLocalizations,proto3" json:"name_localizations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description              string                      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DescriptionLocalizations map[string]string           `protobuf:"bytes,5,rep,name=description_localizations,json=descriptionLocalizations,proto3" json:"description_localizations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Options                  []*ApplicationCommandOption `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty"`
	DefaultMemberPermissions int64                       `protobuf:"varint,7,opt,name=default_member_permissions,json=defaultMemberPermissions,proto3" json:"default_member_permissions,omitempty"`
	Nsfw                     bool                        `protobuf:"varint,8,opt,name=nsfw,proto3" json:"nsfw,omitempty"`
}

func (x *ApplicationCommand) Reset() {
	*x = ApplicationCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discord_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationCommand) ProtoMessage() {}

func (x *ApplicationCommand) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationCommand.ProtoReflect.Descriptor instead.
func (*ApplicationCommand) Descriptor() ([]byte, []int) {
	return file_proto_discord_command_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationCommand) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ApplicationCommand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationCommand) GetNameLocalizations() map[string]string {
	if x != nil {
		return x.NameLocalizations
	}
	return nil
}

func (x *ApplicationCommand) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApplicationCommand) GetDescriptionLocalizations() map[string]string {
	if x != nil {
		return x.DescriptionLocalizations
	}
	return nil
}

func (x *ApplicationCommand) GetOptions() []*ApplicationCommandOption {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ApplicationCommand) GetDefaultMemberPermissions() int64 {
	if x != nil {
		return x.DefaultMemberPermissions
	}
	return 0
}

func (x *ApplicationCommand) GetNsfw() bool {
	if x != nil {
		return x.Nsfw
	}
	return false
}

type ApplicationCommandOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                     uint32                            `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                     string                            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NameLocalizations        map[string]string                 `protobuf:"bytes,3,rep,name=name_localizations,json=nameLocalizations,proto3" json:"name_localizations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description              string                            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DescriptionLocalizations map[string]string                 `protobuf:"bytes,5,rep,name=description_localizations,json=descriptionLocalizations,proto3" json:"description_localizations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Required                 bool                              `protobuf:"varint,6,opt,name=required,proto3" json:"required,omitempty"`
	Choices                  []*ApplicationCommandOptionChoice `protobuf:"bytes,7,rep,name=choices,proto3" json:"choices,omitempty"`
	Options                  []*ApplicationCommandOption       `protobuf:"bytes,8,rep,name=options,proto3" json:"options,omitempty"`
	ChannelTypes             []uint32                          `protobuf:"varint,9,rep,packed,name=channel_types,json=channelTypes,proto3" json:"channel_types,omitempty"`
	MinValue                 float64                           `protobuf:"fixed64,10,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	MaxValue                 float64                           `protobuf:"fixed64,11,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	MinLength                uint32                            `protobuf:"varint,12,opt,name=min_length,json=minLength,proto3" json:"min_length,omitempty"`
	MaxLength                uint32                            `protobuf:"varint,13,opt,name=max_length,json=maxLength,proto3" json:"max_length,omitempty"`
	Autocomplete             bool                              `protobuf:"varint,14,opt,name=autocomplete,proto3" json:"autocomplete,omitempty"`
}

func (x *ApplicationCommandOption) Reset() {
	*x = ApplicationCommandOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discord_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationCommandOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationCommandOption) ProtoMessage() {}

func (x *ApplicationCommandOption) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationCommandOption.ProtoReflect.Descriptor instead.
func (*ApplicationCommandOption) Descriptor() ([]byte, []int) {
	return file_proto_discord_command_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationCommandOption) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ApplicationCommandOption) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationCommandOption) GetNameLocalizations() map[string]string {
	if x != nil {
		return x.NameLocalizations
	}
	return nil
}

func (x *ApplicationCommandOption) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApplicationCommandOption) GetDescriptionLocalizations() map[string]string {
	if x != nil {
		return x.DescriptionLocalizations
	}
	return nil
}

func (x *ApplicationCommandOption) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *ApplicationCommandOption) GetChoices() []*ApplicationCommandOptionChoice {
	if x != nil {
		return x.Choices
	}
	return nil
}

func (x *ApplicationCommandOption) GetOptions() []*ApplicationCommandOption {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ApplicationCommandOption) GetChannelTypes() []uint32 {
	if x != nil {
		return x.ChannelTypes
	}
	return nil
}

func (x *ApplicationCommandOption) GetMinValue() float64 {
	if x != nil {
		return x.MinValue
	}
	return 0
}

func (x *ApplicationCommandOption) GetMaxValue() float64 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

func (x *ApplicationCommandOption) GetMinLength() uint32 {
	if x != nil {
		return x.MinLength
	}
	return 0
}

func (x *ApplicationCommandOption) GetMaxLength() uint32 {
	if x != nil {
		return x.MaxLength
	}
	return 0
}

func (x *ApplicationCommandOption) GetAutocomplete() bool {
	if x != nil {
		return x.Autocomplete
	}
	return false
}

type ApplicationCommandOptionChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NameLocalizations map[string]string `protobuf:"bytes,2,rep,name=name_localizations,json=nameLocalizations,proto3" json:"name_localizations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Value             []byte            `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ApplicationCommandOptionChoice) Reset() {
	*x = ApplicationCommandOptionChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_discord_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationCommandOptionChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationCommandOptionChoice) ProtoMessage() {}

func (x *ApplicationCommandOptionChoice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discord_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationCommandOptionChoice.ProtoReflect.Descriptor instead.
func (*ApplicationCommandOptionChoice) Descriptor() ([]byte, []int) {
	return file_proto_discord_command_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationCommandOptionChoice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationCommandOptionChoice) GetNameLocalizations() map[string]string {
	if x != nil {
		return x.NameLocalizations
	}
	return nil
}

func (x *ApplicationCommandOptionChoice) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_proto_discord_command_proto protoreflect.FileDescriptor

var file_proto_discord_command_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x22, 0xdb, 0x04, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x61, 0x0a, 0x12, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x76, 0x0a, 0x19, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x3c, 0x0a, 0x1a, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x18, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x73, 0x66, 0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x73, 0x66,
	0x77, 0x1a, 0x44, 0x0a, 0x16, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x1d, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xbb, 0x06, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x67, 0x0a, 0x12, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x11, 0x6e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7c, 0x0a, 0x19, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x41,
	0x0a, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x07, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x69, 0x6e, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x6d, 0x69, 0x6e, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x61,
	0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x1a,
	0x44, 0x0a, 0x16, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4b, 0x0a, 0x1d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xff, 0x01, 0x0a, 0x1e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6d, 0x0a, 0x12, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x44,
	0x0a, 0x16, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x45, 0x6c, 0x69, 0x61, 0x73, 0x53, 0x74, 0x61, 0x72, 0x2f, 0x42, 0x61, 0x63,
	0x6f, 0x54, 0x65, 0x6c, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_discord_command_proto_rawDescOnce sync.Once
	file_proto_discord_command_proto_rawDescData = file_proto_discord_command_proto_rawDesc
)

func file_proto_discord_command_proto_rawDescGZIP() []byte {
	file_proto_discord_command_proto_rawDescOnce.Do(func() {
		file_proto_discord_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_discord_command_proto_rawDescData)
	})
	return file_proto_discord_command_proto_rawDescData
}

var file_proto_discord_command_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_discord_command_proto_goTypes = []interface{}{
	(*ApplicationCommand)(nil),             // 0: discord.ApplicationCommand
	(*ApplicationCommandOption)(nil),       // 1: discord.ApplicationCommandOption
	(*ApplicationCommandOptionChoice)(nil), // 2: discord.ApplicationCommandOptionChoice
	nil,                                    // 3: discord.ApplicationCommand.NameLocalizationsEntry
	nil,                                    // 4: discord.ApplicationCommand.DescriptionLocalizationsEntry
	nil,                                    // 5: discord.ApplicationCommandOption.NameLocalizationsEntry
	nil,                                    // 6: discord.ApplicationCommandOption.DescriptionLocalizationsEntry
	nil,                                    // 7: discord.ApplicationCommandOptionChoice.NameLocalizationsEntry
}
var file_proto_discord_command_proto_depIdxs = []int32{
	3, // 0: discord.ApplicationCommand.name_localizations:type_name -> discord.ApplicationCommand.NameLocalizationsEntry
	4, // 1: discord.ApplicationCommand.description_localizations:type_name -> discord.ApplicationCommand.DescriptionLocalizationsEntry
	1, // 2: discord.ApplicationCommand.options:type_name -> discord.ApplicationCommandOption
	5, // 3: discord.ApplicationCommandOption.name_localizations:type_name -> discord.ApplicationCommandOption.NameLocalizationsEntry
	6, // 4: discord.ApplicationCommandOption.description_localizations:type_name -> discord.ApplicationCommandOption.DescriptionLocalizationsEntry
	2, // 5: discord.ApplicationCommandOption.choices:type_name -> discord.ApplicationCommandOptionChoice
	1, // 6: discord.ApplicationCommandOption.options:type_name -> discord.ApplicationCommandOption
	7, // 7: discord.ApplicationCommandOptionChoice.name_localizations:type_name -> discord.ApplicationCommandOptionChoice.NameLocalizationsEntry
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_discord_command_proto_init() }
func file_proto_discord_command_proto_init() {
	if File_proto_discord_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_discord_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationCommand); i {
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
		file_proto_discord_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationCommandOption); i {
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
		file_proto_discord_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationCommandOptionChoice); i {
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
			RawDescriptor: file_proto_discord_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_discord_command_proto_goTypes,
		DependencyIndexes: file_proto_discord_command_proto_depIdxs,
		MessageInfos:      file_proto_discord_command_proto_msgTypes,
	}.Build()
	File_proto_discord_command_proto = out.File
	file_proto_discord_command_proto_rawDesc = nil
	file_proto_discord_command_proto_goTypes = nil
	file_proto_discord_command_proto_depIdxs = nil
}
