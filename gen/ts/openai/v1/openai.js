// @generated by protobuf-ts 2.9.3 with parameter generate_dependencies,long_type_string,output_javascript_es2020,server_grpc1,force_client_none
// @generated from protobuf file "openai/v1/openai.proto" (package "openai.v1", syntax proto3)
// tslint:disable
// @generated by protobuf-ts 2.9.3 with parameter generate_dependencies,long_type_string,output_javascript_es2020,server_grpc1,force_client_none
// @generated from protobuf file "openai/v1/openai.proto" (package "openai.v1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * Type of OpenAI Model
 *
 * @generated from protobuf enum openai.v1.OpenAIModel
 */
export var OpenAIModel;
(function (OpenAIModel) {
    /**
     * OpenAI Model is not Specified
     *
     * @generated from protobuf enum value: OPEN_AI_MODEL_UNSPECIFIED = 0;
     */
    OpenAIModel[OpenAIModel["OPEN_AI_MODEL_UNSPECIFIED"] = 0] = "OPEN_AI_MODEL_UNSPECIFIED";
    /**
     * OpenAI GPT3.5 Turbo 4K
     *
     * @generated from protobuf enum value: OPEN_AI_MODEL_GPT3_5_TURBO_4K = 1;
     */
    OpenAIModel[OpenAIModel["OPEN_AI_MODEL_GPT3_5_TURBO_4K"] = 1] = "OPEN_AI_MODEL_GPT3_5_TURBO_4K";
    /**
     * OpenAI GPT3.5 Turbo 16K
     *
     * @generated from protobuf enum value: OPEN_AI_MODEL_GPT3_5_TURBO_16K = 2;
     */
    OpenAIModel[OpenAIModel["OPEN_AI_MODEL_GPT3_5_TURBO_16K"] = 2] = "OPEN_AI_MODEL_GPT3_5_TURBO_16K";
    /**
     * OpenAI GPT4 8K
     *
     * @generated from protobuf enum value: OPEN_AI_MODEL_GPT4_8K = 3;
     */
    OpenAIModel[OpenAIModel["OPEN_AI_MODEL_GPT4_8K"] = 3] = "OPEN_AI_MODEL_GPT4_8K";
    /**
     * OpenAI GPT4 32K
     *
     * @generated from protobuf enum value: OPEN_AI_MODEL_GPT4_32K = 4;
     */
    OpenAIModel[OpenAIModel["OPEN_AI_MODEL_GPT4_32K"] = 4] = "OPEN_AI_MODEL_GPT4_32K";
})(OpenAIModel || (OpenAIModel = {}));
/**
 * Type of OpenAI Message
 *
 * @generated from protobuf enum openai.v1.OpenAIMessageRole
 */
export var OpenAIMessageRole;
(function (OpenAIMessageRole) {
    /**
     * OpenAI Message type is not Specified
     *
     * @generated from protobuf enum value: OPEN_AI_MESSAGE_ROLE_UNSPECIFIED = 0;
     */
    OpenAIMessageRole[OpenAIMessageRole["OPEN_AI_MESSAGE_ROLE_UNSPECIFIED"] = 0] = "OPEN_AI_MESSAGE_ROLE_UNSPECIFIED";
    /**
     * OpenAI System message
     *
     * @generated from protobuf enum value: OPEN_AI_MESSAGE_ROLE_SYSTEM = 1;
     */
    OpenAIMessageRole[OpenAIMessageRole["OPEN_AI_MESSAGE_ROLE_SYSTEM"] = 1] = "OPEN_AI_MESSAGE_ROLE_SYSTEM";
    /**
     * OpenAI User message
     *
     * @generated from protobuf enum value: OPEN_AI_MESSAGE_ROLE_USER = 2;
     */
    OpenAIMessageRole[OpenAIMessageRole["OPEN_AI_MESSAGE_ROLE_USER"] = 2] = "OPEN_AI_MESSAGE_ROLE_USER";
    /**
     * OpenAI Assistant message
     *
     * @generated from protobuf enum value: OPEN_AI_MESSAGE_ROLE_ASSISTANT = 3;
     */
    OpenAIMessageRole[OpenAIMessageRole["OPEN_AI_MESSAGE_ROLE_ASSISTANT"] = 3] = "OPEN_AI_MESSAGE_ROLE_ASSISTANT";
    /**
     * OpenAI Function message
     *
     * @generated from protobuf enum value: OPEN_AI_MESSAGE_ROLE_FUNCTION = 4;
     */
    OpenAIMessageRole[OpenAIMessageRole["OPEN_AI_MESSAGE_ROLE_FUNCTION"] = 4] = "OPEN_AI_MESSAGE_ROLE_FUNCTION";
})(OpenAIMessageRole || (OpenAIMessageRole = {}));
// @generated message type with reflection information, may provide speed optimized methods
class OpenAIFunctionCall$Type extends MessageType {
    constructor() {
        super("openai.v1.OpenAIFunctionCall", [
            { no: 1, name: "arguments", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message openai.v1.OpenAIFunctionCall
 */
export const OpenAIFunctionCall = new OpenAIFunctionCall$Type();
// @generated message type with reflection information, may provide speed optimized methods
class OpenAIMessage$Type extends MessageType {
    constructor() {
        super("openai.v1.OpenAIMessage", [
            { no: 1, name: "content", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "role", kind: "enum", T: () => ["openai.v1.OpenAIMessageRole", OpenAIMessageRole] },
            { no: 3, name: "name", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "function_call", kind: "message", T: () => OpenAIFunctionCall }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message openai.v1.OpenAIMessage
 */
export const OpenAIMessage = new OpenAIMessage$Type();
// @generated message type with reflection information, may provide speed optimized methods
class OpenAIModelParameters$Type extends MessageType {
    constructor() {
        super("openai.v1.OpenAIModelParameters", [
            { no: 1, name: "temperature", kind: "scalar", opt: true, T: 2 /*ScalarType.FLOAT*/ },
            { no: 2, name: "top_p", kind: "scalar", opt: true, T: 2 /*ScalarType.FLOAT*/ },
            { no: 3, name: "max_tokens", kind: "scalar", opt: true, T: 13 /*ScalarType.UINT32*/ },
            { no: 4, name: "presence_penalty", kind: "scalar", opt: true, T: 2 /*ScalarType.FLOAT*/ },
            { no: 5, name: "frequency_penalty", kind: "scalar", opt: true, T: 2 /*ScalarType.FLOAT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message openai.v1.OpenAIModelParameters
 */
export const OpenAIModelParameters = new OpenAIModelParameters$Type();
// @generated message type with reflection information, may provide speed optimized methods
class OpenAIPromptRequest$Type extends MessageType {
    constructor() {
        super("openai.v1.OpenAIPromptRequest", [
            { no: 1, name: "model", kind: "enum", T: () => ["openai.v1.OpenAIModel", OpenAIModel] },
            { no: 2, name: "messages", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => OpenAIMessage },
            { no: 3, name: "parameters", kind: "message", T: () => OpenAIModelParameters },
            { no: 4, name: "application_id", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message openai.v1.OpenAIPromptRequest
 */
export const OpenAIPromptRequest = new OpenAIPromptRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class OpenAIPromptResponse$Type extends MessageType {
    constructor() {
        super("openai.v1.OpenAIPromptResponse", [
            { no: 1, name: "content", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "prompt_tokens", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 3, name: "completion_tokens", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 4, name: "total_tokens", kind: "scalar", T: 13 /*ScalarType.UINT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message openai.v1.OpenAIPromptResponse
 */
export const OpenAIPromptResponse = new OpenAIPromptResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class OpenAIStreamResponse$Type extends MessageType {
    constructor() {
        super("openai.v1.OpenAIStreamResponse", [
            { no: 1, name: "content", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "finish_reason", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message openai.v1.OpenAIStreamResponse
 */
export const OpenAIStreamResponse = new OpenAIStreamResponse$Type();
/**
 * @generated ServiceType for protobuf service openai.v1.OpenAIService
 */
export const OpenAIService = new ServiceType("openai.v1.OpenAIService", [
    { name: "OpenAIPrompt", options: {}, I: OpenAIPromptRequest, O: OpenAIPromptResponse },
    { name: "OpenAIStream", serverStreaming: true, options: {}, I: OpenAIPromptRequest, O: OpenAIStreamResponse }
]);
