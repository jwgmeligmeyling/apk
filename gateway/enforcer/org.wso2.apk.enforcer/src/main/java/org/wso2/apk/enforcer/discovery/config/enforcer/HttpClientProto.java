// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/config/enforcer/client.proto

package org.wso2.apk.enforcer.discovery.config.enforcer;

public final class HttpClientProto {
  private HttpClientProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_wso2_discovery_config_enforcer_HttpClient_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_wso2_discovery_config_enforcer_HttpClient_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n+wso2/discovery/config/enforcer/client." +
      "proto\022\036wso2.discovery.config.enforcer\"\243\001" +
      "\n\nHttpClient\022\017\n\007skipSSl\030\001 \001(\010\022\030\n\020hostnam" +
      "eVerifier\030\002 \001(\t\022\033\n\023maxTotalConnections\030\003" +
      " \001(\005\022\036\n\026maxConnectionsPerRoute\030\004 \001(\005\022\026\n\016" +
      "connectTimeout\030\005 \001(\005\022\025\n\rsocketTimeout\030\006 " +
      "\001(\005B\224\001\n/org.wso2.apk.enforcer.discovery." +
      "config.enforcerB\017HttpClientProtoP\001ZNgith" +
      "ub.com/envoyproxy/go-control-plane/wso2/" +
      "discovery/config/enforcer;enforcerb\006prot" +
      "o3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_wso2_discovery_config_enforcer_HttpClient_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_wso2_discovery_config_enforcer_HttpClient_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_wso2_discovery_config_enforcer_HttpClient_descriptor,
        new java.lang.String[] { "SkipSSl", "HostnameVerifier", "MaxTotalConnections", "MaxConnectionsPerRoute", "ConnectTimeout", "SocketTimeout", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
