// @generated by protobuf-ts 2.9.3 with parameter generate_dependencies,long_type_string,output_javascript_es2020,server_grpc1,force_client_none
// @generated from protobuf file "openai/v1/openai.proto" (package "openai.v1", syntax proto3)
// tslint:disable
import { OpenAIStreamResponse } from "./openai";
import { OpenAIPromptResponse } from "./openai";
import { OpenAIPromptRequest } from "./openai";
import type * as grpc from "@grpc/grpc-js";
/**
 * The OpenAIService service definition.
 *
 * @generated from protobuf service openai.v1.OpenAIService
 */
export interface IOpenAIService extends grpc.UntypedServiceImplementation {
    /**
     * Request a regular LLM prompt
     *
     * @generated from protobuf rpc: OpenAIPrompt(openai.v1.OpenAIPromptRequest) returns (openai.v1.OpenAIPromptResponse);
     */
    openAIPrompt: grpc.handleUnaryCall<OpenAIPromptRequest, OpenAIPromptResponse>;
    /**
     * Request a streaming LLM prompt
     *
     * @generated from protobuf rpc: OpenAIStream(openai.v1.OpenAIPromptRequest) returns (stream openai.v1.OpenAIStreamResponse);
     */
    openAIStream: grpc.handleServerStreamingCall<OpenAIPromptRequest, OpenAIStreamResponse>;
}
/**
 * @grpc/grpc-js definition for the protobuf service openai.v1.OpenAIService.
 *
 * Usage: Implement the interface IOpenAIService and add to a grpc server.
 *
 * ```typescript
 * const server = new grpc.Server();
 * const service: IOpenAIService = ...
 * server.addService(openAIServiceDefinition, service);
 * ```
 */
export declare const openAIServiceDefinition: grpc.ServiceDefinition<IOpenAIService>;
