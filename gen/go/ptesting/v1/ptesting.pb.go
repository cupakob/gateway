// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: ptesting/v1/ptesting.proto

package ptesting

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

// A request for a prompt - sending user input to the server.
type PromptTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The project ID
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The application ID
	ApplicationId string `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	// The project ID
	PromptConfigId string `protobuf:"bytes,3,opt,name=prompt_config_id,json=promptConfigId,proto3" json:"prompt_config_id,omitempty"`
	// The model vendor, for example "OPEN_AI"
	ModelVendor string `protobuf:"bytes,4,opt,name=model_vendor,json=modelVendor,proto3" json:"model_vendor,omitempty"`
	// The model type to use, for example "gpt-3.5-turbo"
	ModelType string `protobuf:"bytes,5,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	// A serialized JSON object containing the model parameters
	ModelParameters []byte `protobuf:"bytes,6,opt,name=model_parameters,json=modelParameters,proto3" json:"model_parameters,omitempty"`
	// A serialized JSON array of provider message objects
	ProviderPromptMessages []byte `protobuf:"bytes,7,opt,name=provider_prompt_messages,json=providerPromptMessages,proto3" json:"provider_prompt_messages,omitempty"`
	// The User prompt variables
	TemplateVariables map[string]string `protobuf:"bytes,8,rep,name=template_variables,json=templateVariables,proto3" json:"template_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of expected prompt variables
	ExpectedTemplateVariables []string `protobuf:"bytes,9,rep,name=expected_template_variables,json=expectedTemplateVariables,proto3" json:"expected_template_variables,omitempty"`
}

func (x *PromptTestRequest) Reset() {
	*x = PromptTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptesting_v1_ptesting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptTestRequest) ProtoMessage() {}

func (x *PromptTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ptesting_v1_ptesting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptTestRequest.ProtoReflect.Descriptor instead.
func (*PromptTestRequest) Descriptor() ([]byte, []int) {
	return file_ptesting_v1_ptesting_proto_rawDescGZIP(), []int{0}
}

func (x *PromptTestRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *PromptTestRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *PromptTestRequest) GetPromptConfigId() string {
	if x != nil {
		return x.PromptConfigId
	}
	return ""
}

func (x *PromptTestRequest) GetModelVendor() string {
	if x != nil {
		return x.ModelVendor
	}
	return ""
}

func (x *PromptTestRequest) GetModelType() string {
	if x != nil {
		return x.ModelType
	}
	return ""
}

func (x *PromptTestRequest) GetModelParameters() []byte {
	if x != nil {
		return x.ModelParameters
	}
	return nil
}

func (x *PromptTestRequest) GetProviderPromptMessages() []byte {
	if x != nil {
		return x.ProviderPromptMessages
	}
	return nil
}

func (x *PromptTestRequest) GetTemplateVariables() map[string]string {
	if x != nil {
		return x.TemplateVariables
	}
	return nil
}

func (x *PromptTestRequest) GetExpectedTemplateVariables() []string {
	if x != nil {
		return x.ExpectedTemplateVariables
	}
	return nil
}

// An Streaming Prompt Response Message
type PromptTestingStreamingPromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prompt Content
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// Finish reason, given when the stream ends
	FinishReason *string `protobuf:"bytes,2,opt,name=finish_reason,json=finishReason,proto3,oneof" json:"finish_reason,omitempty"`
	// The prompt request record ID, given when the stream ends
	PromptRequestRecordId *string `protobuf:"bytes,3,opt,name=prompt_request_record_id,json=promptRequestRecordId,proto3,oneof" json:"prompt_request_record_id,omitempty"`
}

func (x *PromptTestingStreamingPromptResponse) Reset() {
	*x = PromptTestingStreamingPromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptesting_v1_ptesting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromptTestingStreamingPromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromptTestingStreamingPromptResponse) ProtoMessage() {}

func (x *PromptTestingStreamingPromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ptesting_v1_ptesting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromptTestingStreamingPromptResponse.ProtoReflect.Descriptor instead.
func (*PromptTestingStreamingPromptResponse) Descriptor() ([]byte, []int) {
	return file_ptesting_v1_ptesting_proto_rawDescGZIP(), []int{1}
}

func (x *PromptTestingStreamingPromptResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PromptTestingStreamingPromptResponse) GetFinishReason() string {
	if x != nil && x.FinishReason != nil {
		return *x.FinishReason
	}
	return ""
}

func (x *PromptTestingStreamingPromptResponse) GetPromptRequestRecordId() string {
	if x != nil && x.PromptRequestRecordId != nil {
		return *x.PromptRequestRecordId
	}
	return ""
}

var File_ptesting_v1_ptesting_proto protoreflect.FileDescriptor

var file_ptesting_v1_ptesting_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x96, 0x04, 0x0a, 0x11, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x18,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x16,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x64, 0x0a, 0x12, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x1b,
	0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x19, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x44, 0x0a, 0x16,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xd7, 0x01, 0x0a, 0x24, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x3c, 0x0a, 0x18, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42,
	0x1b, 0x0a, 0x19, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x32, 0x7b, 0x0a, 0x14,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x9d, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x70, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x50,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x03, 0x50, 0x01,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x6d, 0x69, 0x6e, 0x64, 0x2d, 0x61, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x72, 0x65, 0x70,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0xa2, 0x02,
	0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x50, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x0b, 0x50, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x17, 0x50, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x50, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ptesting_v1_ptesting_proto_rawDescOnce sync.Once
	file_ptesting_v1_ptesting_proto_rawDescData = file_ptesting_v1_ptesting_proto_rawDesc
)

func file_ptesting_v1_ptesting_proto_rawDescGZIP() []byte {
	file_ptesting_v1_ptesting_proto_rawDescOnce.Do(func() {
		file_ptesting_v1_ptesting_proto_rawDescData = protoimpl.X.CompressGZIP(file_ptesting_v1_ptesting_proto_rawDescData)
	})
	return file_ptesting_v1_ptesting_proto_rawDescData
}

var file_ptesting_v1_ptesting_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ptesting_v1_ptesting_proto_goTypes = []interface{}{
	(*PromptTestRequest)(nil),                    // 0: ptesting.v1.PromptTestRequest
	(*PromptTestingStreamingPromptResponse)(nil), // 1: ptesting.v1.PromptTestingStreamingPromptResponse
	nil, // 2: ptesting.v1.PromptTestRequest.TemplateVariablesEntry
}
var file_ptesting_v1_ptesting_proto_depIdxs = []int32{
	2, // 0: ptesting.v1.PromptTestRequest.template_variables:type_name -> ptesting.v1.PromptTestRequest.TemplateVariablesEntry
	0, // 1: ptesting.v1.PromptTestingService.TestPrompt:input_type -> ptesting.v1.PromptTestRequest
	1, // 2: ptesting.v1.PromptTestingService.TestPrompt:output_type -> ptesting.v1.PromptTestingStreamingPromptResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ptesting_v1_ptesting_proto_init() }
func file_ptesting_v1_ptesting_proto_init() {
	if File_ptesting_v1_ptesting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ptesting_v1_ptesting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptTestRequest); i {
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
		file_ptesting_v1_ptesting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromptTestingStreamingPromptResponse); i {
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
	file_ptesting_v1_ptesting_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ptesting_v1_ptesting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ptesting_v1_ptesting_proto_goTypes,
		DependencyIndexes: file_ptesting_v1_ptesting_proto_depIdxs,
		MessageInfos:      file_ptesting_v1_ptesting_proto_msgTypes,
	}.Build()
	File_ptesting_v1_ptesting_proto = out.File
	file_ptesting_v1_ptesting_proto_rawDesc = nil
	file_ptesting_v1_ptesting_proto_goTypes = nil
	file_ptesting_v1_ptesting_proto_depIdxs = nil
}
