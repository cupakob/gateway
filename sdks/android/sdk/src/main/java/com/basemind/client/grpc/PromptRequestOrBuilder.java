// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: gateway/v1/gateway.proto

package com.basemind.client.grpc;

public interface PromptRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:gateway.v1.PromptRequest)
    com.google.protobuf.MessageLiteOrBuilder {

  /**
   * <pre>
   * The User prompt variables
   * This is a hash-map of variables that should have the same keys as those contained by the PromptConfigResponse
   * </pre>
   *
   * <code>map&lt;string, string&gt; template_variables = 1 [json_name = "templateVariables"];</code>
   */
  int getTemplateVariablesCount();
  /**
   * <pre>
   * The User prompt variables
   * This is a hash-map of variables that should have the same keys as those contained by the PromptConfigResponse
   * </pre>
   *
   * <code>map&lt;string, string&gt; template_variables = 1 [json_name = "templateVariables"];</code>
   */
  boolean containsTemplateVariables(
      java.lang.String key);
  /**
   * Use {@link #getTemplateVariablesMap()} instead.
   */
  @java.lang.Deprecated
  java.util.Map<java.lang.String, java.lang.String>
  getTemplateVariables();
  /**
   * <pre>
   * The User prompt variables
   * This is a hash-map of variables that should have the same keys as those contained by the PromptConfigResponse
   * </pre>
   *
   * <code>map&lt;string, string&gt; template_variables = 1 [json_name = "templateVariables"];</code>
   */
  java.util.Map<java.lang.String, java.lang.String>
  getTemplateVariablesMap();
  /**
   * <pre>
   * The User prompt variables
   * This is a hash-map of variables that should have the same keys as those contained by the PromptConfigResponse
   * </pre>
   *
   * <code>map&lt;string, string&gt; template_variables = 1 [json_name = "templateVariables"];</code>
   */

  /* nullable */
java.lang.String getTemplateVariablesOrDefault(
      java.lang.String key,
      /* nullable */
java.lang.String defaultValue);
  /**
   * <pre>
   * The User prompt variables
   * This is a hash-map of variables that should have the same keys as those contained by the PromptConfigResponse
   * </pre>
   *
   * <code>map&lt;string, string&gt; template_variables = 1 [json_name = "templateVariables"];</code>
   */

  java.lang.String getTemplateVariablesOrThrow(
      java.lang.String key);
}
