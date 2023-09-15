// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: gateway/v1/gateway.proto

package gateway

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

// A request for a prompt configuration - retrieving the expected prompt variables
type PromptConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The application ID, this value represents the APP ID as configured in our db.
	ApplicationId string `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (x *PromptConfigRequest) Reset() {
	*x = PromptConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptConfigRequest) ProtoMessage() {}

func (x *PromptConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptConfigRequest.ProtoReflect.Descriptor instead.
func (*PromptConfigRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *PromptConfigRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

// A response for a prompt configuration - retrieving the expected prompt variables
type PromptConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The expected prompt variables
	ExpectedPromptVariables map[string]string `protobuf:"bytes,1,rep,name=expected_prompt_variables,json=expectedPromptVariables,proto3" json:"expected_prompt_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PromptConfigResponse) Reset() {
	*x = PromptConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptConfigResponse) ProtoMessage() {}

func (x *PromptConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptConfigResponse.ProtoReflect.Descriptor instead.
func (*PromptConfigResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *PromptConfigResponse) GetExpectedPromptVariables() map[string]string {
	if x != nil {
		return x.ExpectedPromptVariables
	}
	return nil
}

// A request for a prompt - sending user input to the server.
type PromptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The application ID, this value represents the APP ID as configured in our db.
	ApplicationId string `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	// The User prompt variables
	// This is a hash-map of variables that should have the same keys as those contained by the PromptConfigResponse
	TemplateVariables map[string]string `protobuf:"bytes,2,rep,name=template_variables,json=templateVariables,proto3" json:"template_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PromptRequest) Reset() {
	*x = PromptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptRequest) ProtoMessage() {}

func (x *PromptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptRequest.ProtoReflect.Descriptor instead.
func (*PromptRequest) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *PromptRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *PromptRequest) GetTemplateVariables() map[string]string {
	if x != nil {
		return x.TemplateVariables
	}
	return nil
}

// A Prompt Response Message
type PromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompt Content
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// Number of tokens used for the prompt
	PromptTokens uint32 `protobuf:"varint,2,opt,name=prompt_tokens,json=promptTokens,proto3" json:"prompt_tokens,omitempty"`
}

func (x *PromptResponse) Reset() {
	*x = PromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptResponse) ProtoMessage() {}

func (x *PromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptResponse.ProtoReflect.Descriptor instead.
func (*PromptResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *PromptResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PromptResponse) GetPromptTokens() uint32 {
	if x != nil {
		return x.PromptTokens
	}
	return 0
}

// An Streaming Prompt Response Message
type StreamingPromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompt Content
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// Finish reason, if this is the last message
	FinishReason *string `protobuf:"bytes,2,opt,name=finish_reason,json=finishReason,proto3,oneof" json:"finish_reason,omitempty"`
	// Number of tokens used for the prompt, given when the stream ends
	PromptTokens *uint32 `protobuf:"varint,3,opt,name=prompt_tokens,json=promptTokens,proto3,oneof" json:"prompt_tokens,omitempty"`
}

func (x *StreamingPromptResponse) Reset() {
	*x = StreamingPromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_v1_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingPromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingPromptResponse) ProtoMessage() {}

func (x *StreamingPromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_v1_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingPromptResponse.ProtoReflect.Descriptor instead.
func (*StreamingPromptResponse) Descriptor() ([]byte, []int) {
	return file_gateway_v1_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *StreamingPromptResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *StreamingPromptResponse) GetFinishReason() string {
	if x != nil && x.FinishReason != nil {
		return *x.FinishReason
	}
	return ""
}

func (x *StreamingPromptResponse) GetPromptTokens() uint32 {
	if x != nil && x.PromptTokens != nil {
		return *x.PromptTokens
	}
	return 0
}

var File_gateway_v1_gateway_proto protoreflect.FileDescriptor

var file_gateway_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x3c, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a,
	0x19, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x17, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x4a, 0x0a, 0x1c, 0x45, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xdd, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x5f, 0x0a,
	0x12, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x44,
	0x0a, 0x16, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x4f, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x17, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x32, 0x97, 0x02, 0x0a, 0x11, 0x41, 0x50, 0x49, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5c, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x2d, 0x61, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70, 0x6f,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_v1_gateway_proto_rawDescOnce sync.Once
	file_gateway_v1_gateway_proto_rawDescData = file_gateway_v1_gateway_proto_rawDesc
)

func file_gateway_v1_gateway_proto_rawDescGZIP() []byte {
	file_gateway_v1_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_v1_gateway_proto_rawDescData)
	})
	return file_gateway_v1_gateway_proto_rawDescData
}

var file_gateway_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_gateway_v1_gateway_proto_goTypes = []interface{}{
	(*PromptConfigRequest)(nil),     // 0: gateway.v1.PromptConfigRequest
	(*PromptConfigResponse)(nil),    // 1: gateway.v1.PromptConfigResponse
	(*PromptRequest)(nil),           // 2: gateway.v1.PromptRequest
	(*PromptResponse)(nil),          // 3: gateway.v1.PromptResponse
	(*StreamingPromptResponse)(nil), // 4: gateway.v1.StreamingPromptResponse
	nil,                             // 5: gateway.v1.PromptConfigResponse.ExpectedPromptVariablesEntry
	nil,                             // 6: gateway.v1.PromptRequest.TemplateVariablesEntry
}
var file_gateway_v1_gateway_proto_depIdxs = []int32{
	5, // 0: gateway.v1.PromptConfigResponse.expected_prompt_variables:type_name -> gateway.v1.PromptConfigResponse.ExpectedPromptVariablesEntry
	6, // 1: gateway.v1.PromptRequest.template_variables:type_name -> gateway.v1.PromptRequest.TemplateVariablesEntry
	0, // 2: gateway.v1.APIGatewayService.RequestPromptConfig:input_type -> gateway.v1.PromptConfigRequest
	2, // 3: gateway.v1.APIGatewayService.RequestPrompt:input_type -> gateway.v1.PromptRequest
	2, // 4: gateway.v1.APIGatewayService.RequestStreamingPrompt:input_type -> gateway.v1.PromptRequest
	1, // 5: gateway.v1.APIGatewayService.RequestPromptConfig:output_type -> gateway.v1.PromptConfigResponse
	3, // 6: gateway.v1.APIGatewayService.RequestPrompt:output_type -> gateway.v1.PromptResponse
	4, // 7: gateway.v1.APIGatewayService.RequestStreamingPrompt:output_type -> gateway.v1.StreamingPromptResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gateway_v1_gateway_proto_init() }
func file_gateway_v1_gateway_proto_init() {
	if File_gateway_v1_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_v1_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptConfigRequest); i {
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
		file_gateway_v1_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptConfigResponse); i {
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
		file_gateway_v1_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptRequest); i {
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
		file_gateway_v1_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptResponse); i {
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
		file_gateway_v1_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingPromptResponse); i {
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
	file_gateway_v1_gateway_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_v1_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_v1_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_v1_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_v1_gateway_proto_msgTypes,
	}.Build()
	File_gateway_v1_gateway_proto = out.File
	file_gateway_v1_gateway_proto_rawDesc = nil
	file_gateway_v1_gateway_proto_goTypes = nil
	file_gateway_v1_gateway_proto_depIdxs = nil
}
