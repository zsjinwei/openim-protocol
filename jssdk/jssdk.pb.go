// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.2
// source: jssdk/jssdk.proto

package jssdk

import (
	conversation "github.com/zsjinwei/openim-protocol/conversation"
	relation "github.com/zsjinwei/openim-protocol/relation"
	sdkws "github.com/zsjinwei/openim-protocol/sdkws"
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

type ConversationMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conversation *conversation.Conversation `protobuf:"bytes,1,opt,name=conversation,proto3" json:"conversation"`
	LastMsg      *sdkws.MsgData             `protobuf:"bytes,2,opt,name=lastMsg,proto3" json:"lastMsg"`
	User         *sdkws.UserInfo            `protobuf:"bytes,3,opt,name=user,proto3" json:"user"`
	Friend       *relation.FriendInfoOnly   `protobuf:"bytes,4,opt,name=friend,proto3" json:"friend"`
	Group        *sdkws.GroupInfo           `protobuf:"bytes,5,opt,name=group,proto3" json:"group"`
	MaxSeq       int64                      `protobuf:"varint,6,opt,name=maxSeq,proto3" json:"maxSeq"`
	ReadSeq      int64                      `protobuf:"varint,7,opt,name=readSeq,proto3" json:"readSeq"`
}

func (x *ConversationMsg) Reset() {
	*x = ConversationMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jssdk_jssdk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversationMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversationMsg) ProtoMessage() {}

func (x *ConversationMsg) ProtoReflect() protoreflect.Message {
	mi := &file_jssdk_jssdk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversationMsg.ProtoReflect.Descriptor instead.
func (*ConversationMsg) Descriptor() ([]byte, []int) {
	return file_jssdk_jssdk_proto_rawDescGZIP(), []int{0}
}

func (x *ConversationMsg) GetConversation() *conversation.Conversation {
	if x != nil {
		return x.Conversation
	}
	return nil
}

func (x *ConversationMsg) GetLastMsg() *sdkws.MsgData {
	if x != nil {
		return x.LastMsg
	}
	return nil
}

func (x *ConversationMsg) GetUser() *sdkws.UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ConversationMsg) GetFriend() *relation.FriendInfoOnly {
	if x != nil {
		return x.Friend
	}
	return nil
}

func (x *ConversationMsg) GetGroup() *sdkws.GroupInfo {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *ConversationMsg) GetMaxSeq() int64 {
	if x != nil {
		return x.MaxSeq
	}
	return 0
}

func (x *ConversationMsg) GetReadSeq() int64 {
	if x != nil {
		return x.ReadSeq
	}
	return 0
}

type GetActiveConversationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerUserID string `protobuf:"bytes,1,opt,name=ownerUserID,proto3" json:"ownerUserID"`
	Count       int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count"`
}

func (x *GetActiveConversationsReq) Reset() {
	*x = GetActiveConversationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jssdk_jssdk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActiveConversationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActiveConversationsReq) ProtoMessage() {}

func (x *GetActiveConversationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_jssdk_jssdk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActiveConversationsReq.ProtoReflect.Descriptor instead.
func (*GetActiveConversationsReq) Descriptor() ([]byte, []int) {
	return file_jssdk_jssdk_proto_rawDescGZIP(), []int{1}
}

func (x *GetActiveConversationsReq) GetOwnerUserID() string {
	if x != nil {
		return x.OwnerUserID
	}
	return ""
}

func (x *GetActiveConversationsReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetActiveConversationsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnreadCount   int64              `protobuf:"varint,1,opt,name=unreadCount,proto3" json:"unreadCount"`
	Conversations []*ConversationMsg `protobuf:"bytes,2,rep,name=conversations,proto3" json:"conversations"`
}

func (x *GetActiveConversationsResp) Reset() {
	*x = GetActiveConversationsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jssdk_jssdk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetActiveConversationsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActiveConversationsResp) ProtoMessage() {}

func (x *GetActiveConversationsResp) ProtoReflect() protoreflect.Message {
	mi := &file_jssdk_jssdk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActiveConversationsResp.ProtoReflect.Descriptor instead.
func (*GetActiveConversationsResp) Descriptor() ([]byte, []int) {
	return file_jssdk_jssdk_proto_rawDescGZIP(), []int{2}
}

func (x *GetActiveConversationsResp) GetUnreadCount() int64 {
	if x != nil {
		return x.UnreadCount
	}
	return 0
}

func (x *GetActiveConversationsResp) GetConversations() []*ConversationMsg {
	if x != nil {
		return x.Conversations
	}
	return nil
}

type GetConversationsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerUserID     string   `protobuf:"bytes,1,opt,name=ownerUserID,proto3" json:"ownerUserID"`
	ConversationIDs []string `protobuf:"bytes,2,rep,name=conversationIDs,proto3" json:"conversationIDs"`
}

func (x *GetConversationsReq) Reset() {
	*x = GetConversationsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jssdk_jssdk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConversationsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConversationsReq) ProtoMessage() {}

func (x *GetConversationsReq) ProtoReflect() protoreflect.Message {
	mi := &file_jssdk_jssdk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConversationsReq.ProtoReflect.Descriptor instead.
func (*GetConversationsReq) Descriptor() ([]byte, []int) {
	return file_jssdk_jssdk_proto_rawDescGZIP(), []int{3}
}

func (x *GetConversationsReq) GetOwnerUserID() string {
	if x != nil {
		return x.OwnerUserID
	}
	return ""
}

func (x *GetConversationsReq) GetConversationIDs() []string {
	if x != nil {
		return x.ConversationIDs
	}
	return nil
}

type GetConversationsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnreadCount   int64              `protobuf:"varint,1,opt,name=unreadCount,proto3" json:"unreadCount"`
	Conversations []*ConversationMsg `protobuf:"bytes,2,rep,name=conversations,proto3" json:"conversations"`
}

func (x *GetConversationsResp) Reset() {
	*x = GetConversationsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jssdk_jssdk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConversationsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConversationsResp) ProtoMessage() {}

func (x *GetConversationsResp) ProtoReflect() protoreflect.Message {
	mi := &file_jssdk_jssdk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConversationsResp.ProtoReflect.Descriptor instead.
func (*GetConversationsResp) Descriptor() ([]byte, []int) {
	return file_jssdk_jssdk_proto_rawDescGZIP(), []int{4}
}

func (x *GetConversationsResp) GetUnreadCount() int64 {
	if x != nil {
		return x.UnreadCount
	}
	return 0
}

func (x *GetConversationsResp) GetConversations() []*ConversationMsg {
	if x != nil {
		return x.Conversations
	}
	return nil
}

var File_jssdk_jssdk_proto protoreflect.FileDescriptor

var file_jssdk_jssdk_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6a, 0x73, 0x73, 0x64, 0x6b, 0x2f, 0x6a, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x6a, 0x73, 0x73, 0x64,
	0x6b, 0x1a, 0x11, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf,
	0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x73, 0x67, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x07, 0x6c, 0x61, 0x73,
	0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x69, 0x6d, 0x2e, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69,
	0x6d, 0x2e, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x06, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x06, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12,
	0x2d, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x73, 0x64, 0x6b, 0x77, 0x73, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6d, 0x61, 0x78, 0x53, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x53, 0x65,
	0x71, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64, 0x53, 0x65, 0x71,
	0x22, 0x53, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a,
	0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x6a, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x61, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x73, 0x22, 0x7d,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x6e, 0x72,
	0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2e, 0x6a, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x2b, 0x5a,
	0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x73, 0x6a, 0x69,
	0x6e, 0x77, 0x65, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x69, 0x6d, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6a, 0x73, 0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_jssdk_jssdk_proto_rawDescOnce sync.Once
	file_jssdk_jssdk_proto_rawDescData = file_jssdk_jssdk_proto_rawDesc
)

func file_jssdk_jssdk_proto_rawDescGZIP() []byte {
	file_jssdk_jssdk_proto_rawDescOnce.Do(func() {
		file_jssdk_jssdk_proto_rawDescData = protoimpl.X.CompressGZIP(file_jssdk_jssdk_proto_rawDescData)
	})
	return file_jssdk_jssdk_proto_rawDescData
}

var file_jssdk_jssdk_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_jssdk_jssdk_proto_goTypes = []interface{}{
	(*ConversationMsg)(nil),            // 0: openim.jssdk.ConversationMsg
	(*GetActiveConversationsReq)(nil),  // 1: openim.jssdk.GetActiveConversationsReq
	(*GetActiveConversationsResp)(nil), // 2: openim.jssdk.GetActiveConversationsResp
	(*GetConversationsReq)(nil),        // 3: openim.jssdk.GetConversationsReq
	(*GetConversationsResp)(nil),       // 4: openim.jssdk.GetConversationsResp
	(*conversation.Conversation)(nil),  // 5: openim.conversation.Conversation
	(*sdkws.MsgData)(nil),              // 6: openim.sdkws.MsgData
	(*sdkws.UserInfo)(nil),             // 7: openim.sdkws.UserInfo
	(*relation.FriendInfoOnly)(nil),    // 8: openim.relation.FriendInfoOnly
	(*sdkws.GroupInfo)(nil),            // 9: openim.sdkws.GroupInfo
}
var file_jssdk_jssdk_proto_depIdxs = []int32{
	5, // 0: openim.jssdk.ConversationMsg.conversation:type_name -> openim.conversation.Conversation
	6, // 1: openim.jssdk.ConversationMsg.lastMsg:type_name -> openim.sdkws.MsgData
	7, // 2: openim.jssdk.ConversationMsg.user:type_name -> openim.sdkws.UserInfo
	8, // 3: openim.jssdk.ConversationMsg.friend:type_name -> openim.relation.FriendInfoOnly
	9, // 4: openim.jssdk.ConversationMsg.group:type_name -> openim.sdkws.GroupInfo
	0, // 5: openim.jssdk.GetActiveConversationsResp.conversations:type_name -> openim.jssdk.ConversationMsg
	0, // 6: openim.jssdk.GetConversationsResp.conversations:type_name -> openim.jssdk.ConversationMsg
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_jssdk_jssdk_proto_init() }
func file_jssdk_jssdk_proto_init() {
	if File_jssdk_jssdk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jssdk_jssdk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversationMsg); i {
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
		file_jssdk_jssdk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActiveConversationsReq); i {
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
		file_jssdk_jssdk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetActiveConversationsResp); i {
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
		file_jssdk_jssdk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConversationsReq); i {
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
		file_jssdk_jssdk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConversationsResp); i {
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
			RawDescriptor: file_jssdk_jssdk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jssdk_jssdk_proto_goTypes,
		DependencyIndexes: file_jssdk_jssdk_proto_depIdxs,
		MessageInfos:      file_jssdk_jssdk_proto_msgTypes,
	}.Build()
	File_jssdk_jssdk_proto = out.File
	file_jssdk_jssdk_proto_rawDesc = nil
	file_jssdk_jssdk_proto_goTypes = nil
	file_jssdk_jssdk_proto_depIdxs = nil
}
