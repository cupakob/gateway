// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: cohere/v1/cohere.proto

package cohereconnector

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Type of Cohere Model
type CohereModel int32

const (
	// Cohere Model is not specified
	CohereModel_COHERE_MODEL_UNSPECIFIED CohereModel = 0
	// Command - the default Cohere model.
	CohereModel_COHERE_MODEL_COMMAND CohereModel = 1
	// Command Light - a faster but less accurate version of Command.
	CohereModel_COHERE_MODEL_COMMAND_LIGHT CohereModel = 2
	// Command Nightly - a nightly version of Command.
	CohereModel_COHERE_MODEL_COMMAND_NIGHTLY CohereModel = 3
	// Command Light Nightly - a nightly version of Command Light.
	CohereModel_COHERE_MODEL_COMMAND_LIGHT_NIGHTLY CohereModel = 4
)

// Enum value maps for CohereModel.
var (
	CohereModel_name = map[int32]string{
		0: "COHERE_MODEL_UNSPECIFIED",
		1: "COHERE_MODEL_COMMAND",
		2: "COHERE_MODEL_COMMAND_LIGHT",
		3: "COHERE_MODEL_COMMAND_NIGHTLY",
		4: "COHERE_MODEL_COMMAND_LIGHT_NIGHTLY",
	}
	CohereModel_value = map[string]int32{
		"COHERE_MODEL_UNSPECIFIED":           0,
		"COHERE_MODEL_COMMAND":               1,
		"COHERE_MODEL_COMMAND_LIGHT":         2,
		"COHERE_MODEL_COMMAND_NIGHTLY":       3,
		"COHERE_MODEL_COMMAND_LIGHT_NIGHTLY": 4,
	}
)

func (x CohereModel) Enum() *CohereModel {
	p := new(CohereModel)
	*p = x
	return p
}

func (x CohereModel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CohereModel) Descriptor() protoreflect.EnumDescriptor {
	return file_cohere_v1_cohere_proto_enumTypes[0].Descriptor()
}

func (CohereModel) Type() protoreflect.EnumType {
	return &file_cohere_v1_cohere_proto_enumTypes[0]
}

func (x CohereModel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CohereModel.Descriptor instead.
func (CohereModel) EnumDescriptor() ([]byte, []int) {
	return file_cohere_v1_cohere_proto_rawDescGZIP(), []int{0}
}

// Type of Cohere RAG Connector
type CohereConnectorType int32

const (
	// Cohere Connector is not specified
	CohereConnectorType_COHERE_CONNECTOR_TYPE_UNSPECIFIED CohereConnectorType = 0
	// Cohere Connector is a web search.
	CohereConnectorType_COHERE_CONNECTOR_TYPE_WEB_SEARCH CohereConnectorType = 1
	// Cohere Connector is a custom ID.
	CohereConnectorType_COHERE_CONNECTOR_TYPE_ID CohereConnectorType = 2
)

// Enum value maps for CohereConnectorType.
var (
	CohereConnectorType_name = map[int32]string{
		0: "COHERE_CONNECTOR_TYPE_UNSPECIFIED",
		1: "COHERE_CONNECTOR_TYPE_WEB_SEARCH",
		2: "COHERE_CONNECTOR_TYPE_ID",
	}
	CohereConnectorType_value = map[string]int32{
		"COHERE_CONNECTOR_TYPE_UNSPECIFIED": 0,
		"COHERE_CONNECTOR_TYPE_WEB_SEARCH":  1,
		"COHERE_CONNECTOR_TYPE_ID":          2,
	}
)

func (x CohereConnectorType) Enum() *CohereConnectorType {
	p := new(CohereConnectorType)
	*p = x
	return p
}

func (x CohereConnectorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CohereConnectorType) Descriptor() protoreflect.EnumDescriptor {
	return file_cohere_v1_cohere_proto_enumTypes[1].Descriptor()
}

func (CohereConnectorType) Type() protoreflect.EnumType {
	return &file_cohere_v1_cohere_proto_enumTypes[1]
}

func (x CohereConnectorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CohereConnectorType.Descriptor instead.
func (CohereConnectorType) EnumDescriptor() ([]byte, []int) {
	return file_cohere_v1_cohere_proto_rawDescGZIP(), []int{1}
}

// Cohere RAG Connector
type CohereConnector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cohere Connector type
	Id CohereConnectorType `protobuf:"varint,1,opt,name=id,proto3,enum=cohere.v1.CohereConnectorType" json:"id,omitempty"`
	// Cohere Connector options
	Options map[string]*anypb.Any `protobuf:"bytes,2,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CohereConnector) Reset() {
	*x = CohereConnector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cohere_v1_cohere_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CohereConnector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CohereConnector) ProtoMessage() {}

func (x *CohereConnector) ProtoReflect() protoreflect.Message {
	mi := &file_cohere_v1_cohere_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CohereConnector.ProtoReflect.Descriptor instead.
func (*CohereConnector) Descriptor() ([]byte, []int) {
	return file_cohere_v1_cohere_proto_rawDescGZIP(), []int{0}
}

func (x *CohereConnector) GetId() CohereConnectorType {
	if x != nil {
		return x.Id
	}
	return CohereConnectorType_COHERE_CONNECTOR_TYPE_UNSPECIFIED
}

func (x *CohereConnector) GetOptions() map[string]*anypb.Any {
	if x != nil {
		return x.Options
	}
	return nil
}

// Cohere API Request parameters
type CohereModelParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Temperature Sampling.
	// Should be a non-negative float (0-1.0) that tunes the degree of randomness in generation. Lower temperatures mean
	// less random generations, and higher temperatures mean more random generations.
	Temperature *float32 `protobuf:"fixed32,1,opt,name=temperature,proto3,oneof" json:"temperature,omitempty"`
	// RAG connectors to use, when specified, the model's reply will be enriched with information found by querying each
	// of the connectors.
	Connectors []*CohereConnector `protobuf:"bytes,2,rep,name=connectors,proto3" json:"connectors,omitempty"`
}

func (x *CohereModelParameters) Reset() {
	*x = CohereModelParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cohere_v1_cohere_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CohereModelParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CohereModelParameters) ProtoMessage() {}

func (x *CohereModelParameters) ProtoReflect() protoreflect.Message {
	mi := &file_cohere_v1_cohere_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CohereModelParameters.ProtoReflect.Descriptor instead.
func (*CohereModelParameters) Descriptor() ([]byte, []int) {
	return file_cohere_v1_cohere_proto_rawDescGZIP(), []int{1}
}

func (x *CohereModelParameters) GetTemperature() float32 {
	if x != nil && x.Temperature != nil {
		return *x.Temperature
	}
	return 0
}

func (x *CohereModelParameters) GetConnectors() []*CohereConnector {
	if x != nil {
		return x.Connectors
	}
	return nil
}

// The CoherePromptRequest contains the data that will be sent to the Cohere API.
type CoherePromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cohere Model identifier
	Model CohereModel `protobuf:"varint,1,opt,name=model,proto3,enum=cohere.v1.CohereModel" json:"model,omitempty"`
	// Prompt message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Cohere API Request parameters
	Parameters *CohereModelParameters `protobuf:"bytes,3,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// An identifier for the conversation chain. Conversations can be resumed by providing the conversation's identifier.
	// The contents of `message` and the model's response will be stored as part of this conversation.
	// If a conversation with this id does not already exist, a new conversation will be created.
	ConversationId *string `protobuf:"bytes,4,opt,name=conversation_id,json=conversationId,proto3,oneof" json:"conversation_id,omitempty"`
}

func (x *CoherePromptRequest) Reset() {
	*x = CoherePromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cohere_v1_cohere_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoherePromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoherePromptRequest) ProtoMessage() {}

func (x *CoherePromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cohere_v1_cohere_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoherePromptRequest.ProtoReflect.Descriptor instead.
func (*CoherePromptRequest) Descriptor() ([]byte, []int) {
	return file_cohere_v1_cohere_proto_rawDescGZIP(), []int{2}
}

func (x *CoherePromptRequest) GetModel() CohereModel {
	if x != nil {
		return x.Model
	}
	return CohereModel_COHERE_MODEL_UNSPECIFIED
}

func (x *CoherePromptRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CoherePromptRequest) GetParameters() *CohereModelParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *CoherePromptRequest) GetConversationId() string {
	if x != nil && x.ConversationId != nil {
		return *x.ConversationId
	}
	return ""
}

// The CoherePromptResponse contains the data that is returned from the Cohere API.
type CoherePromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompt Content
	Content *string `protobuf:"bytes,1,opt,name=content,proto3,oneof" json:"content,omitempty"`
}

func (x *CoherePromptResponse) Reset() {
	*x = CoherePromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cohere_v1_cohere_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoherePromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoherePromptResponse) ProtoMessage() {}

func (x *CoherePromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cohere_v1_cohere_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoherePromptResponse.ProtoReflect.Descriptor instead.
func (*CoherePromptResponse) Descriptor() ([]byte, []int) {
	return file_cohere_v1_cohere_proto_rawDescGZIP(), []int{3}
}

func (x *CoherePromptResponse) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

// The CohereStreamResponse contains the data that is streamed from the Cohere API.
type CohereStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompt Content
	Content *string `protobuf:"bytes,1,opt,name=content,proto3,oneof" json:"content,omitempty"`
	// Finish reason, if this is the last message
	FinishReason *string `protobuf:"bytes,2,opt,name=finish_reason,json=finishReason,proto3,oneof" json:"finish_reason,omitempty"`
}

func (x *CohereStreamResponse) Reset() {
	*x = CohereStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cohere_v1_cohere_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CohereStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CohereStreamResponse) ProtoMessage() {}

func (x *CohereStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cohere_v1_cohere_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CohereStreamResponse.ProtoReflect.Descriptor instead.
func (*CohereStreamResponse) Descriptor() ([]byte, []int) {
	return file_cohere_v1_cohere_proto_rawDescGZIP(), []int{4}
}

func (x *CohereStreamResponse) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *CohereStreamResponse) GetFinishReason() string {
	if x != nil && x.FinishReason != nil {
		return *x.FinishReason
	}
	return ""
}

var File_cohere_v1_cohere_proto protoreflect.FileDescriptor

var file_cohere_v1_cohere_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x68, 0x65,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6,
	0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e,
	0x2e, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x68, 0x65, 0x72,
	0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x41, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8a, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x68, 0x65,
	0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x6f, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x6f,
	0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x68, 0x65, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x14, 0x43, 0x6f, 0x68, 0x65,
	0x72, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7d, 0x0a, 0x14, 0x43,
	0x6f, 0x68, 0x65, 0x72, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0xaf, 0x01, 0x0a, 0x0b, 0x43,
	0x6f, 0x68, 0x65, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f,
	0x48, 0x45, 0x52, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x48, 0x45,
	0x52, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x4c, 0x49, 0x47, 0x48, 0x54,
	0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x4e, 0x49, 0x47, 0x48, 0x54,
	0x4c, 0x59, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x4f, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x4c, 0x49, 0x47,
	0x48, 0x54, 0x5f, 0x4e, 0x49, 0x47, 0x48, 0x54, 0x4c, 0x59, 0x10, 0x04, 0x2a, 0x80, 0x01, 0x0a,
	0x13, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x4f, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x43,
	0x4f, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10,
	0x01, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x48, 0x45, 0x52, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x4e,
	0x45, 0x43, 0x54, 0x4f, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x02, 0x32,
	0xb7, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x51, 0x0a, 0x0c, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x68, 0x65, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x68, 0x65, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x98, 0x01, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43, 0x6f, 0x68,
	0x65, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x03, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x2d, 0x61, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x63, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x15, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x43, 0x6f, 0x68, 0x65, 0x72, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cohere_v1_cohere_proto_rawDescOnce sync.Once
	file_cohere_v1_cohere_proto_rawDescData = file_cohere_v1_cohere_proto_rawDesc
)

func file_cohere_v1_cohere_proto_rawDescGZIP() []byte {
	file_cohere_v1_cohere_proto_rawDescOnce.Do(func() {
		file_cohere_v1_cohere_proto_rawDescData = protoimpl.X.CompressGZIP(file_cohere_v1_cohere_proto_rawDescData)
	})
	return file_cohere_v1_cohere_proto_rawDescData
}

var file_cohere_v1_cohere_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cohere_v1_cohere_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cohere_v1_cohere_proto_goTypes = []interface{}{
	(CohereModel)(0),              // 0: cohere.v1.CohereModel
	(CohereConnectorType)(0),      // 1: cohere.v1.CohereConnectorType
	(*CohereConnector)(nil),       // 2: cohere.v1.CohereConnector
	(*CohereModelParameters)(nil), // 3: cohere.v1.CohereModelParameters
	(*CoherePromptRequest)(nil),   // 4: cohere.v1.CoherePromptRequest
	(*CoherePromptResponse)(nil),  // 5: cohere.v1.CoherePromptResponse
	(*CohereStreamResponse)(nil),  // 6: cohere.v1.CohereStreamResponse
	nil,                           // 7: cohere.v1.CohereConnector.OptionsEntry
	(*anypb.Any)(nil),             // 8: google.protobuf.Any
}
var file_cohere_v1_cohere_proto_depIdxs = []int32{
	1, // 0: cohere.v1.CohereConnector.id:type_name -> cohere.v1.CohereConnectorType
	7, // 1: cohere.v1.CohereConnector.options:type_name -> cohere.v1.CohereConnector.OptionsEntry
	2, // 2: cohere.v1.CohereModelParameters.connectors:type_name -> cohere.v1.CohereConnector
	0, // 3: cohere.v1.CoherePromptRequest.model:type_name -> cohere.v1.CohereModel
	3, // 4: cohere.v1.CoherePromptRequest.parameters:type_name -> cohere.v1.CohereModelParameters
	8, // 5: cohere.v1.CohereConnector.OptionsEntry.value:type_name -> google.protobuf.Any
	4, // 6: cohere.v1.CohereService.CoherePrompt:input_type -> cohere.v1.CoherePromptRequest
	4, // 7: cohere.v1.CohereService.CohereStream:input_type -> cohere.v1.CoherePromptRequest
	5, // 8: cohere.v1.CohereService.CoherePrompt:output_type -> cohere.v1.CoherePromptResponse
	6, // 9: cohere.v1.CohereService.CohereStream:output_type -> cohere.v1.CohereStreamResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cohere_v1_cohere_proto_init() }
func file_cohere_v1_cohere_proto_init() {
	if File_cohere_v1_cohere_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cohere_v1_cohere_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CohereConnector); i {
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
		file_cohere_v1_cohere_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CohereModelParameters); i {
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
		file_cohere_v1_cohere_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoherePromptRequest); i {
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
		file_cohere_v1_cohere_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoherePromptResponse); i {
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
		file_cohere_v1_cohere_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CohereStreamResponse); i {
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
	file_cohere_v1_cohere_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_cohere_v1_cohere_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_cohere_v1_cohere_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_cohere_v1_cohere_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cohere_v1_cohere_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cohere_v1_cohere_proto_goTypes,
		DependencyIndexes: file_cohere_v1_cohere_proto_depIdxs,
		EnumInfos:         file_cohere_v1_cohere_proto_enumTypes,
		MessageInfos:      file_cohere_v1_cohere_proto_msgTypes,
	}.Build()
	File_cohere_v1_cohere_proto = out.File
	file_cohere_v1_cohere_proto_rawDesc = nil
	file_cohere_v1_cohere_proto_goTypes = nil
	file_cohere_v1_cohere_proto_depIdxs = nil
}
