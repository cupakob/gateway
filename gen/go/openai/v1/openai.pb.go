// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: openai/v1/openai.proto

package openaiconnector

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

// Type of OpenAI Model
type OpenAIModel int32

const (
	// OpenAI Model is not Specified
	OpenAIModel_OPEN_AI_MODEL_UNSPECIFIED OpenAIModel = 0
	// OpenAI GPT3.5 Turbo 4K
	OpenAIModel_OPEN_AI_MODEL_GPT3_5_TURBO_4K OpenAIModel = 1
	// OpenAI GPT3.5 Turbo 16K
	OpenAIModel_OPEN_AI_MODEL_GPT3_5_TURBO_16K OpenAIModel = 2
	// OpenAI GPT4 8K
	OpenAIModel_OPEN_AI_MODEL_GPT4_8K OpenAIModel = 3
	// OpenAI GPT4 32K
	OpenAIModel_OPEN_AI_MODEL_GPT4_32K OpenAIModel = 4
)

// Enum value maps for OpenAIModel.
var (
	OpenAIModel_name = map[int32]string{
		0: "OPEN_AI_MODEL_UNSPECIFIED",
		1: "OPEN_AI_MODEL_GPT3_5_TURBO_4K",
		2: "OPEN_AI_MODEL_GPT3_5_TURBO_16K",
		3: "OPEN_AI_MODEL_GPT4_8K",
		4: "OPEN_AI_MODEL_GPT4_32K",
	}
	OpenAIModel_value = map[string]int32{
		"OPEN_AI_MODEL_UNSPECIFIED":      0,
		"OPEN_AI_MODEL_GPT3_5_TURBO_4K":  1,
		"OPEN_AI_MODEL_GPT3_5_TURBO_16K": 2,
		"OPEN_AI_MODEL_GPT4_8K":          3,
		"OPEN_AI_MODEL_GPT4_32K":         4,
	}
)

func (x OpenAIModel) Enum() *OpenAIModel {
	p := new(OpenAIModel)
	*p = x
	return p
}

func (x OpenAIModel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenAIModel) Descriptor() protoreflect.EnumDescriptor {
	return file_openai_v1_openai_proto_enumTypes[0].Descriptor()
}

func (OpenAIModel) Type() protoreflect.EnumType {
	return &file_openai_v1_openai_proto_enumTypes[0]
}

func (x OpenAIModel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenAIModel.Descriptor instead.
func (OpenAIModel) EnumDescriptor() ([]byte, []int) {
	return file_openai_v1_openai_proto_rawDescGZIP(), []int{0}
}

// Type of OpenAI Message
type OpenAIMessageRole int32

const (
	// OpenAI Message type is not Specified
	OpenAIMessageRole_OPEN_AI_MESSAGE_ROLE_UNSPECIFIED OpenAIMessageRole = 0
	// OpenAI System message
	OpenAIMessageRole_OPEN_AI_MESSAGE_ROLE_SYSTEM OpenAIMessageRole = 1
	// OpenAI User message
	OpenAIMessageRole_OPEN_AI_MESSAGE_ROLE_USER OpenAIMessageRole = 2
	// OpenAI Assistant message
	OpenAIMessageRole_OPEN_AI_MESSAGE_ROLE_ASSISTANT OpenAIMessageRole = 3
	// OpenAI Function message
	OpenAIMessageRole_OPEN_AI_MESSAGE_ROLE_FUNCTION OpenAIMessageRole = 4
)

// Enum value maps for OpenAIMessageRole.
var (
	OpenAIMessageRole_name = map[int32]string{
		0: "OPEN_AI_MESSAGE_ROLE_UNSPECIFIED",
		1: "OPEN_AI_MESSAGE_ROLE_SYSTEM",
		2: "OPEN_AI_MESSAGE_ROLE_USER",
		3: "OPEN_AI_MESSAGE_ROLE_ASSISTANT",
		4: "OPEN_AI_MESSAGE_ROLE_FUNCTION",
	}
	OpenAIMessageRole_value = map[string]int32{
		"OPEN_AI_MESSAGE_ROLE_UNSPECIFIED": 0,
		"OPEN_AI_MESSAGE_ROLE_SYSTEM":      1,
		"OPEN_AI_MESSAGE_ROLE_USER":        2,
		"OPEN_AI_MESSAGE_ROLE_ASSISTANT":   3,
		"OPEN_AI_MESSAGE_ROLE_FUNCTION":    4,
	}
)

func (x OpenAIMessageRole) Enum() *OpenAIMessageRole {
	p := new(OpenAIMessageRole)
	*p = x
	return p
}

func (x OpenAIMessageRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenAIMessageRole) Descriptor() protoreflect.EnumDescriptor {
	return file_openai_v1_openai_proto_enumTypes[1].Descriptor()
}

func (OpenAIMessageRole) Type() protoreflect.EnumType {
	return &file_openai_v1_openai_proto_enumTypes[1]
}

func (x OpenAIMessageRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenAIMessageRole.Descriptor instead.
func (OpenAIMessageRole) EnumDescriptor() ([]byte, []int) {
	return file_openai_v1_openai_proto_rawDescGZIP(), []int{1}
}

// An OpenAI function call
type OpenAIFunctionCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The signature of the function arguments
	Arguments string `protobuf:"bytes,1,opt,name=arguments,proto3" json:"arguments,omitempty"`
	// The function name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *OpenAIFunctionCall) Reset() {
	*x = OpenAIFunctionCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openai_v1_openai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAIFunctionCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAIFunctionCall) ProtoMessage() {}

func (x *OpenAIFunctionCall) ProtoReflect() protoreflect.Message {
	mi := &file_openai_v1_openai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAIFunctionCall.ProtoReflect.Descriptor instead.
func (*OpenAIFunctionCall) Descriptor() ([]byte, []int) {
	return file_openai_v1_openai_proto_rawDescGZIP(), []int{0}
}

func (x *OpenAIFunctionCall) GetArguments() string {
	if x != nil {
		return x.Arguments
	}
	return ""
}

func (x *OpenAIFunctionCall) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// An OpenAI Chat Message
type OpenAIMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The content of the message
	Content *string `protobuf:"bytes,1,opt,name=content,proto3,oneof" json:"content,omitempty"`
	// The role of the message
	Role OpenAIMessageRole `protobuf:"varint,2,opt,name=role,proto3,enum=openai.v1.OpenAIMessageRole" json:"role,omitempty"`
	// Name of the message author or function name
	Name *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// The signature function to invoke, if any
	FunctionCall *OpenAIFunctionCall `protobuf:"bytes,4,opt,name=function_call,json=functionCall,proto3,oneof" json:"function_call,omitempty"`
}

func (x *OpenAIMessage) Reset() {
	*x = OpenAIMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openai_v1_openai_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAIMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAIMessage) ProtoMessage() {}

func (x *OpenAIMessage) ProtoReflect() protoreflect.Message {
	mi := &file_openai_v1_openai_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAIMessage.ProtoReflect.Descriptor instead.
func (*OpenAIMessage) Descriptor() ([]byte, []int) {
	return file_openai_v1_openai_proto_rawDescGZIP(), []int{1}
}

func (x *OpenAIMessage) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *OpenAIMessage) GetRole() OpenAIMessageRole {
	if x != nil {
		return x.Role
	}
	return OpenAIMessageRole_OPEN_AI_MESSAGE_ROLE_UNSPECIFIED
}

func (x *OpenAIMessage) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *OpenAIMessage) GetFunctionCall() *OpenAIFunctionCall {
	if x != nil {
		return x.FunctionCall
	}
	return nil
}

// OpenAI API Request parameters
type OpenAIModelParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Temperature Sampling: https://platform.openai.com/docs/api-reference/chat/create#temperature
	Temperature *float32 `protobuf:"fixed32,1,opt,name=temperature,proto3,oneof" json:"temperature,omitempty"`
	// Nucleus Sampling: https://platform.openai.com/docs/api-reference/chat/create#top_p
	TopP *float32 `protobuf:"fixed32,2,opt,name=top_p,json=topP,proto3,oneof" json:"top_p,omitempty"`
	// Maximum Tokens after which the prompt will stop generating a response;
	MaxTokens *uint32 `protobuf:"varint,3,opt,name=max_tokens,json=maxTokens,proto3,oneof" json:"max_tokens,omitempty"`
	// Penalize New tokens: https://platform.openai.com/docs/api-reference/chat/create#presence_penalty
	PresencePenalty *float32 `protobuf:"fixed32,4,opt,name=presence_penalty,json=presencePenalty,proto3,oneof" json:"presence_penalty,omitempty"`
	// Penalize Repeated tokens: https://platform.openai.com/docs/api-reference/chat/create#frequency_penalty
	FrequencyPenalty *float32 `protobuf:"fixed32,5,opt,name=frequency_penalty,json=frequencyPenalty,proto3,oneof" json:"frequency_penalty,omitempty"`
}

func (x *OpenAIModelParameters) Reset() {
	*x = OpenAIModelParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openai_v1_openai_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAIModelParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAIModelParameters) ProtoMessage() {}

func (x *OpenAIModelParameters) ProtoReflect() protoreflect.Message {
	mi := &file_openai_v1_openai_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAIModelParameters.ProtoReflect.Descriptor instead.
func (*OpenAIModelParameters) Descriptor() ([]byte, []int) {
	return file_openai_v1_openai_proto_rawDescGZIP(), []int{2}
}

func (x *OpenAIModelParameters) GetTemperature() float32 {
	if x != nil && x.Temperature != nil {
		return *x.Temperature
	}
	return 0
}

func (x *OpenAIModelParameters) GetTopP() float32 {
	if x != nil && x.TopP != nil {
		return *x.TopP
	}
	return 0
}

func (x *OpenAIModelParameters) GetMaxTokens() uint32 {
	if x != nil && x.MaxTokens != nil {
		return *x.MaxTokens
	}
	return 0
}

func (x *OpenAIModelParameters) GetPresencePenalty() float32 {
	if x != nil && x.PresencePenalty != nil {
		return *x.PresencePenalty
	}
	return 0
}

func (x *OpenAIModelParameters) GetFrequencyPenalty() float32 {
	if x != nil && x.FrequencyPenalty != nil {
		return *x.FrequencyPenalty
	}
	return 0
}

// A Request for an OpenAI regular LLM Prompt
type OpenAIPromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// OpenAI Model identifier
	Model OpenAIModel `protobuf:"varint,1,opt,name=model,proto3,enum=openai.v1.OpenAIModel" json:"model,omitempty"`
	// Prompt Messages
	Messages []*OpenAIMessage `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	// OpenAI API Request parameters
	Parameters *OpenAIModelParameters `protobuf:"bytes,3,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// Unique application ID to keep track of conversations;
	ApplicationId *string `protobuf:"bytes,4,opt,name=application_id,json=applicationId,proto3,oneof" json:"application_id,omitempty"`
}

func (x *OpenAIPromptRequest) Reset() {
	*x = OpenAIPromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openai_v1_openai_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAIPromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAIPromptRequest) ProtoMessage() {}

func (x *OpenAIPromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_openai_v1_openai_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAIPromptRequest.ProtoReflect.Descriptor instead.
func (*OpenAIPromptRequest) Descriptor() ([]byte, []int) {
	return file_openai_v1_openai_proto_rawDescGZIP(), []int{3}
}

func (x *OpenAIPromptRequest) GetModel() OpenAIModel {
	if x != nil {
		return x.Model
	}
	return OpenAIModel_OPEN_AI_MODEL_UNSPECIFIED
}

func (x *OpenAIPromptRequest) GetMessages() []*OpenAIMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *OpenAIPromptRequest) GetParameters() *OpenAIModelParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

func (x *OpenAIPromptRequest) GetApplicationId() string {
	if x != nil && x.ApplicationId != nil {
		return *x.ApplicationId
	}
	return ""
}

// An OpenAI Prompt Response Message
type OpenAIPromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompt Content
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// Number of tokens used for the prompt
	PromptTokens uint32 `protobuf:"varint,2,opt,name=prompt_tokens,json=promptTokens,proto3" json:"prompt_tokens,omitempty"`
	// Number of tokens used to generate the completion
	CompletionTokens uint32 `protobuf:"varint,3,opt,name=completion_tokens,json=completionTokens,proto3" json:"completion_tokens,omitempty"`
	// Total number of tokens used to generate the response
	TotalTokens uint32 `protobuf:"varint,4,opt,name=total_tokens,json=totalTokens,proto3" json:"total_tokens,omitempty"`
}

func (x *OpenAIPromptResponse) Reset() {
	*x = OpenAIPromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openai_v1_openai_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAIPromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAIPromptResponse) ProtoMessage() {}

func (x *OpenAIPromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_openai_v1_openai_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAIPromptResponse.ProtoReflect.Descriptor instead.
func (*OpenAIPromptResponse) Descriptor() ([]byte, []int) {
	return file_openai_v1_openai_proto_rawDescGZIP(), []int{4}
}

func (x *OpenAIPromptResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *OpenAIPromptResponse) GetPromptTokens() uint32 {
	if x != nil {
		return x.PromptTokens
	}
	return 0
}

func (x *OpenAIPromptResponse) GetCompletionTokens() uint32 {
	if x != nil {
		return x.CompletionTokens
	}
	return 0
}

func (x *OpenAIPromptResponse) GetTotalTokens() uint32 {
	if x != nil {
		return x.TotalTokens
	}
	return 0
}

// An OpenAI Streaming Response Message
type OpenAIStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompt Content
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// Finish reason, if this is the last message
	FinishReason *string `protobuf:"bytes,2,opt,name=finish_reason,json=finishReason,proto3,oneof" json:"finish_reason,omitempty"`
}

func (x *OpenAIStreamResponse) Reset() {
	*x = OpenAIStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_openai_v1_openai_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAIStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAIStreamResponse) ProtoMessage() {}

func (x *OpenAIStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_openai_v1_openai_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAIStreamResponse.ProtoReflect.Descriptor instead.
func (*OpenAIStreamResponse) Descriptor() ([]byte, []int) {
	return file_openai_v1_openai_proto_rawDescGZIP(), []int{5}
}

func (x *OpenAIStreamResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *OpenAIStreamResponse) GetFinishReason() string {
	if x != nil && x.FinishReason != nil {
		return *x.FinishReason
	}
	return ""
}

var File_openai_v1_openai_proto protoreflect.FileDescriptor

var file_openai_v1_openai_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69,
	0x2e, 0x76, 0x31, 0x22, 0x46, 0x0a, 0x12, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xe9, 0x01, 0x0a, 0x0d,
	0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x47, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41,
	0x49, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x02, 0x52,
	0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x88, 0x01, 0x01,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x22, 0xb2, 0x02, 0x0a, 0x15, 0x4f, 0x70, 0x65, 0x6e,
	0x41, 0x49, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x25, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x50, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x48, 0x03, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x65, 0x6e, 0x61,
	0x6c, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x48, 0x04, 0x52, 0x10, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x65,
	0x6e, 0x61, 0x6c, 0x74, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x70,
	0x5f, 0x70, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x70,
	0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x70, 0x65, 0x6e, 0x61, 0x6c, 0x74, 0x79, 0x22, 0xfa, 0x01, 0x0a,
	0x13, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2a, 0x0a, 0x0e, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x14, 0x4f, 0x70,
	0x65, 0x6e, 0x41, 0x49, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x22, 0x6c, 0x0a, 0x14, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a,
	0xaa, 0x01, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x1d, 0x0a, 0x19, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x41, 0x49, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21,
	0x0a, 0x1d, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x41, 0x49, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f,
	0x47, 0x50, 0x54, 0x33, 0x5f, 0x35, 0x5f, 0x54, 0x55, 0x52, 0x42, 0x4f, 0x5f, 0x34, 0x4b, 0x10,
	0x01, 0x12, 0x22, 0x0a, 0x1e, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x41, 0x49, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x4c, 0x5f, 0x47, 0x50, 0x54, 0x33, 0x5f, 0x35, 0x5f, 0x54, 0x55, 0x52, 0x42, 0x4f, 0x5f,
	0x31, 0x36, 0x4b, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x41, 0x49,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x47, 0x50, 0x54, 0x34, 0x5f, 0x38, 0x4b, 0x10, 0x03,
	0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x41, 0x49, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x4c, 0x5f, 0x47, 0x50, 0x54, 0x34, 0x5f, 0x33, 0x32, 0x4b, 0x10, 0x04, 0x2a, 0xc0, 0x01, 0x0a,
	0x11, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x41, 0x49, 0x5f, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x50, 0x45, 0x4e,
	0x5f, 0x41, 0x49, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4f, 0x50, 0x45,
	0x4e, 0x5f, 0x41, 0x49, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x4f, 0x4c,
	0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x4f, 0x50, 0x45, 0x4e,
	0x5f, 0x41, 0x49, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45,
	0x5f, 0x41, 0x53, 0x53, 0x49, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d,
	0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x41, 0x49, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x32,
	0xb7, 0x01, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x51, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x12, 0x1e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x6e, 0x41, 0x49, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x6e, 0x41, 0x49, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x1e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x49, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x98, 0x01, 0x0a, 0x0d, 0x63, 0x6f,
	0x6d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x4f, 0x70, 0x65,
	0x6e, 0x61, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x03, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x2d, 0x61, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x69,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x15, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x61, 0x69,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_openai_v1_openai_proto_rawDescOnce sync.Once
	file_openai_v1_openai_proto_rawDescData = file_openai_v1_openai_proto_rawDesc
)

func file_openai_v1_openai_proto_rawDescGZIP() []byte {
	file_openai_v1_openai_proto_rawDescOnce.Do(func() {
		file_openai_v1_openai_proto_rawDescData = protoimpl.X.CompressGZIP(file_openai_v1_openai_proto_rawDescData)
	})
	return file_openai_v1_openai_proto_rawDescData
}

var file_openai_v1_openai_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_openai_v1_openai_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_openai_v1_openai_proto_goTypes = []interface{}{
	(OpenAIModel)(0),              // 0: openai.v1.OpenAIModel
	(OpenAIMessageRole)(0),        // 1: openai.v1.OpenAIMessageRole
	(*OpenAIFunctionCall)(nil),    // 2: openai.v1.OpenAIFunctionCall
	(*OpenAIMessage)(nil),         // 3: openai.v1.OpenAIMessage
	(*OpenAIModelParameters)(nil), // 4: openai.v1.OpenAIModelParameters
	(*OpenAIPromptRequest)(nil),   // 5: openai.v1.OpenAIPromptRequest
	(*OpenAIPromptResponse)(nil),  // 6: openai.v1.OpenAIPromptResponse
	(*OpenAIStreamResponse)(nil),  // 7: openai.v1.OpenAIStreamResponse
}
var file_openai_v1_openai_proto_depIdxs = []int32{
	1, // 0: openai.v1.OpenAIMessage.role:type_name -> openai.v1.OpenAIMessageRole
	2, // 1: openai.v1.OpenAIMessage.function_call:type_name -> openai.v1.OpenAIFunctionCall
	0, // 2: openai.v1.OpenAIPromptRequest.model:type_name -> openai.v1.OpenAIModel
	3, // 3: openai.v1.OpenAIPromptRequest.messages:type_name -> openai.v1.OpenAIMessage
	4, // 4: openai.v1.OpenAIPromptRequest.parameters:type_name -> openai.v1.OpenAIModelParameters
	5, // 5: openai.v1.OpenAIService.OpenAIPrompt:input_type -> openai.v1.OpenAIPromptRequest
	5, // 6: openai.v1.OpenAIService.OpenAIStream:input_type -> openai.v1.OpenAIPromptRequest
	6, // 7: openai.v1.OpenAIService.OpenAIPrompt:output_type -> openai.v1.OpenAIPromptResponse
	7, // 8: openai.v1.OpenAIService.OpenAIStream:output_type -> openai.v1.OpenAIStreamResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_openai_v1_openai_proto_init() }
func file_openai_v1_openai_proto_init() {
	if File_openai_v1_openai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_openai_v1_openai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAIFunctionCall); i {
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
		file_openai_v1_openai_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAIMessage); i {
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
		file_openai_v1_openai_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAIModelParameters); i {
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
		file_openai_v1_openai_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAIPromptRequest); i {
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
		file_openai_v1_openai_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAIPromptResponse); i {
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
		file_openai_v1_openai_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenAIStreamResponse); i {
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
	file_openai_v1_openai_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_openai_v1_openai_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_openai_v1_openai_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_openai_v1_openai_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_openai_v1_openai_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_openai_v1_openai_proto_goTypes,
		DependencyIndexes: file_openai_v1_openai_proto_depIdxs,
		EnumInfos:         file_openai_v1_openai_proto_enumTypes,
		MessageInfos:      file_openai_v1_openai_proto_msgTypes,
	}.Build()
	File_openai_v1_openai_proto = out.File
	file_openai_v1_openai_proto_rawDesc = nil
	file_openai_v1_openai_proto_goTypes = nil
	file_openai_v1_openai_proto_depIdxs = nil
}
