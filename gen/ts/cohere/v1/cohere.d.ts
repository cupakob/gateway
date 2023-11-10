// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies,long_type_string,output_javascript_es2020,server_grpc1,force_client_none
// @generated from protobuf file "cohere/v1/cohere.proto" (package "cohere.v1", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { Any } from "../../google/protobuf/any";
/**
 * Cohere RAG Connector
 *
 * @generated from protobuf message cohere.v1.CohereConnector
 */
export interface CohereConnector {
    /**
     * Cohere Connector type
     *
     * @generated from protobuf field: cohere.v1.CohereConnectorType id = 1;
     */
    id: CohereConnectorType;
    /**
     * Cohere Connector options
     *
     * @generated from protobuf field: map<string, google.protobuf.Any> options = 2;
     */
    options: {
        [key: string]: Any;
    };
}
/**
 * Cohere API Request parameters
 *
 * @generated from protobuf message cohere.v1.CohereModelParameters
 */
export interface CohereModelParameters {
    /**
     * Temperature Sampling.
     * Should be a non-negative float (0-1.0) that tunes the degree of randomness in generation. Lower temperatures mean
     * less random generations, and higher temperatures mean more random generations.
     *
     * @generated from protobuf field: optional float temperature = 1;
     */
    temperature?: number;
    /**
     * RAG connectors to use, when specified, the model's reply will be enriched with information found by querying each
     * of the connectors.
     *
     * @generated from protobuf field: repeated cohere.v1.CohereConnector connectors = 2;
     */
    connectors: CohereConnector[];
}
/**
 *  The CoherePromptRequest contains the data that will be sent to the Cohere API.
 *
 * @generated from protobuf message cohere.v1.CoherePromptRequest
 */
export interface CoherePromptRequest {
    /**
     * Cohere Model identifier
     *
     * @generated from protobuf field: cohere.v1.CohereModel model = 1;
     */
    model: CohereModel;
    /**
     * Prompt message
     *
     * @generated from protobuf field: string message = 2;
     */
    message: string;
    /**
     * Cohere API Request parameters
     *
     * @generated from protobuf field: cohere.v1.CohereModelParameters parameters = 3;
     */
    parameters?: CohereModelParameters;
    /**
     * An identifier for the conversation chain. Conversations can be resumed by providing the conversation's identifier.
     * The contents of `message` and the model's response will be stored as part of this conversation.
     * If a conversation with this id does not already exist, a new conversation will be created.
     *
     * @generated from protobuf field: optional string conversation_id = 4;
     */
    conversationId?: string;
}
/**
 * The CoherePromptResponse contains the data that is returned from the Cohere API.
 *
 * @generated from protobuf message cohere.v1.CoherePromptResponse
 */
export interface CoherePromptResponse {
    /**
     * Prompt Content
     *
     * @generated from protobuf field: optional string content = 1;
     */
    content?: string;
}
/**
 * The CohereStreamResponse contains the data that is streamed from the Cohere API.
 *
 * @generated from protobuf message cohere.v1.CohereStreamResponse
 */
export interface CohereStreamResponse {
    /**
     * Prompt Content
     *
     * @generated from protobuf field: optional string content = 1;
     */
    content?: string;
    /**
     * Finish reason, if this is the last message
     *
     * @generated from protobuf field: optional string finish_reason = 2;
     */
    finishReason?: string;
}
/**
 * Type of Cohere Model
 *
 * @generated from protobuf enum cohere.v1.CohereModel
 */
export declare enum CohereModel {
    /**
     * Cohere Model is not specified
     *
     * @generated from protobuf enum value: COHERE_MODEL_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * Command - the default Cohere model.
     *
     * @generated from protobuf enum value: COHERE_MODEL_COMMAND = 1;
     */
    COMMAND = 1,
    /**
     * Command Light - a faster but less accurate version of Command.
     *
     * @generated from protobuf enum value: COHERE_MODEL_COMMAND_LIGHT = 2;
     */
    COMMAND_LIGHT = 2,
    /**
     * Command Nightly - a nightly version of Command.
     *
     * @generated from protobuf enum value: COHERE_MODEL_COMMAND_NIGHTLY = 3;
     */
    COMMAND_NIGHTLY = 3,
    /**
     * Command Light Nightly - a nightly version of Command Light.
     *
     * @generated from protobuf enum value: COHERE_MODEL_COMMAND_LIGHT_NIGHTLY = 4;
     */
    COMMAND_LIGHT_NIGHTLY = 4
}
/**
 * Type of Cohere RAG Connector
 *
 * @generated from protobuf enum cohere.v1.CohereConnectorType
 */
export declare enum CohereConnectorType {
    /**
     * Cohere Connector is not specified
     *
     * @generated from protobuf enum value: COHERE_CONNECTOR_TYPE_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * Cohere Connector is a web search.
     *
     * @generated from protobuf enum value: COHERE_CONNECTOR_TYPE_WEB_SEARCH = 1;
     */
    WEB_SEARCH = 1,
    /**
     * Cohere Connector is a custom ID.
     *
     * @generated from protobuf enum value: COHERE_CONNECTOR_TYPE_ID = 2;
     */
    ID = 2
}
declare class CohereConnector$Type extends MessageType<CohereConnector> {
    constructor();
}
/**
 * @generated MessageType for protobuf message cohere.v1.CohereConnector
 */
export declare const CohereConnector: CohereConnector$Type;
declare class CohereModelParameters$Type extends MessageType<CohereModelParameters> {
    constructor();
}
/**
 * @generated MessageType for protobuf message cohere.v1.CohereModelParameters
 */
export declare const CohereModelParameters: CohereModelParameters$Type;
declare class CoherePromptRequest$Type extends MessageType<CoherePromptRequest> {
    constructor();
}
/**
 * @generated MessageType for protobuf message cohere.v1.CoherePromptRequest
 */
export declare const CoherePromptRequest: CoherePromptRequest$Type;
declare class CoherePromptResponse$Type extends MessageType<CoherePromptResponse> {
    constructor();
}
/**
 * @generated MessageType for protobuf message cohere.v1.CoherePromptResponse
 */
export declare const CoherePromptResponse: CoherePromptResponse$Type;
declare class CohereStreamResponse$Type extends MessageType<CohereStreamResponse> {
    constructor();
}
/**
 * @generated MessageType for protobuf message cohere.v1.CohereStreamResponse
 */
export declare const CohereStreamResponse: CohereStreamResponse$Type;
/**
 * @generated ServiceType for protobuf service cohere.v1.CohereService
 */
export declare const CohereService: any;
export {};
