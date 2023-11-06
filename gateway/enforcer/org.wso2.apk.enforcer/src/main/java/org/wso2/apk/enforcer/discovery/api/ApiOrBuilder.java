// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/api/api.proto

package org.wso2.apk.enforcer.discovery.api;

public interface ApiOrBuilder extends
    // @@protoc_insertion_point(interface_extends:wso2.discovery.api.Api)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string id = 1;</code>
   * @return The id.
   */
  java.lang.String getId();
  /**
   * <code>string id = 1;</code>
   * @return The bytes for id.
   */
  com.google.protobuf.ByteString
      getIdBytes();

  /**
   * <code>string title = 2;</code>
   * @return The title.
   */
  java.lang.String getTitle();
  /**
   * <code>string title = 2;</code>
   * @return The bytes for title.
   */
  com.google.protobuf.ByteString
      getTitleBytes();

  /**
   * <code>string version = 3;</code>
   * @return The version.
   */
  java.lang.String getVersion();
  /**
   * <code>string version = 3;</code>
   * @return The bytes for version.
   */
  com.google.protobuf.ByteString
      getVersionBytes();

  /**
   * <code>string apiType = 4;</code>
   * @return The apiType.
   */
  java.lang.String getApiType();
  /**
   * <code>string apiType = 4;</code>
   * @return The bytes for apiType.
   */
  com.google.protobuf.ByteString
      getApiTypeBytes();

  /**
   * <code>bool disableAuthentications = 5;</code>
   * @return The disableAuthentications.
   */
  boolean getDisableAuthentications();

  /**
   * <code>bool disableScopes = 6;</code>
   * @return The disableScopes.
   */
  boolean getDisableScopes();

  /**
   * <code>string envType = 7;</code>
   * @return The envType.
   */
  java.lang.String getEnvType();
  /**
   * <code>string envType = 7;</code>
   * @return The bytes for envType.
   */
  com.google.protobuf.ByteString
      getEnvTypeBytes();

  /**
   * <code>repeated .wso2.discovery.api.Resource resources = 8;</code>
   */
  java.util.List<org.wso2.apk.enforcer.discovery.api.Resource> 
      getResourcesList();
  /**
   * <code>repeated .wso2.discovery.api.Resource resources = 8;</code>
   */
  org.wso2.apk.enforcer.discovery.api.Resource getResources(int index);
  /**
   * <code>repeated .wso2.discovery.api.Resource resources = 8;</code>
   */
  int getResourcesCount();
  /**
   * <code>repeated .wso2.discovery.api.Resource resources = 8;</code>
   */
  java.util.List<? extends org.wso2.apk.enforcer.discovery.api.ResourceOrBuilder> 
      getResourcesOrBuilderList();
  /**
   * <code>repeated .wso2.discovery.api.Resource resources = 8;</code>
   */
  org.wso2.apk.enforcer.discovery.api.ResourceOrBuilder getResourcesOrBuilder(
      int index);

  /**
   * <code>string basePath = 9;</code>
   * @return The basePath.
   */
  java.lang.String getBasePath();
  /**
   * <code>string basePath = 9;</code>
   * @return The bytes for basePath.
   */
  com.google.protobuf.ByteString
      getBasePathBytes();

  /**
   * <code>string tier = 10;</code>
   * @return The tier.
   */
  java.lang.String getTier();
  /**
   * <code>string tier = 10;</code>
   * @return The bytes for tier.
   */
  com.google.protobuf.ByteString
      getTierBytes();

  /**
   * <code>string apiLifeCycleState = 11;</code>
   * @return The apiLifeCycleState.
   */
  java.lang.String getApiLifeCycleState();
  /**
   * <code>string apiLifeCycleState = 11;</code>
   * @return The bytes for apiLifeCycleState.
   */
  com.google.protobuf.ByteString
      getApiLifeCycleStateBytes();

  /**
   * <code>string vhost = 12;</code>
   * @return The vhost.
   */
  java.lang.String getVhost();
  /**
   * <code>string vhost = 12;</code>
   * @return The bytes for vhost.
   */
  com.google.protobuf.ByteString
      getVhostBytes();

  /**
   * <code>string organizationId = 13;</code>
   * @return The organizationId.
   */
  java.lang.String getOrganizationId();
  /**
   * <code>string organizationId = 13;</code>
   * @return The bytes for organizationId.
   */
  com.google.protobuf.ByteString
      getOrganizationIdBytes();

  /**
   * <pre>
   * bool isMockedApi = 18;
   * </pre>
   *
   * <code>repeated .wso2.discovery.api.Certificate clientCertificates = 14;</code>
   */
  java.util.List<org.wso2.apk.enforcer.discovery.api.Certificate> 
      getClientCertificatesList();
  /**
   * <pre>
   * bool isMockedApi = 18;
   * </pre>
   *
   * <code>repeated .wso2.discovery.api.Certificate clientCertificates = 14;</code>
   */
  org.wso2.apk.enforcer.discovery.api.Certificate getClientCertificates(int index);
  /**
   * <pre>
   * bool isMockedApi = 18;
   * </pre>
   *
   * <code>repeated .wso2.discovery.api.Certificate clientCertificates = 14;</code>
   */
  int getClientCertificatesCount();
  /**
   * <pre>
   * bool isMockedApi = 18;
   * </pre>
   *
   * <code>repeated .wso2.discovery.api.Certificate clientCertificates = 14;</code>
   */
  java.util.List<? extends org.wso2.apk.enforcer.discovery.api.CertificateOrBuilder> 
      getClientCertificatesOrBuilderList();
  /**
   * <pre>
   * bool isMockedApi = 18;
   * </pre>
   *
   * <code>repeated .wso2.discovery.api.Certificate clientCertificates = 14;</code>
   */
  org.wso2.apk.enforcer.discovery.api.CertificateOrBuilder getClientCertificatesOrBuilder(
      int index);

  /**
   * <code>string mutualSSL = 15;</code>
   * @return The mutualSSL.
   */
  java.lang.String getMutualSSL();
  /**
   * <code>string mutualSSL = 15;</code>
   * @return The bytes for mutualSSL.
   */
  com.google.protobuf.ByteString
      getMutualSSLBytes();

  /**
   * <code>bool applicationSecurity = 16;</code>
   * @return The applicationSecurity.
   */
  boolean getApplicationSecurity();

  /**
   * <pre>
   *&#47; string graphQLSchema = 22;
   * repeated GraphqlComplexity graphqlComplexityInfo = 23;
   * </pre>
   *
   * <code>bool systemAPI = 24;</code>
   * @return The systemAPI.
   */
  boolean getSystemAPI();

  /**
   * <code>.wso2.discovery.api.BackendJWTTokenInfo backendJWTTokenInfo = 25;</code>
   * @return Whether the backendJWTTokenInfo field is set.
   */
  boolean hasBackendJWTTokenInfo();
  /**
   * <code>.wso2.discovery.api.BackendJWTTokenInfo backendJWTTokenInfo = 25;</code>
   * @return The backendJWTTokenInfo.
   */
  org.wso2.apk.enforcer.discovery.api.BackendJWTTokenInfo getBackendJWTTokenInfo();
  /**
   * <code>.wso2.discovery.api.BackendJWTTokenInfo backendJWTTokenInfo = 25;</code>
   */
  org.wso2.apk.enforcer.discovery.api.BackendJWTTokenInfoOrBuilder getBackendJWTTokenInfoOrBuilder();

  /**
   * <code>bytes apiDefinitionFile = 26;</code>
   * @return The apiDefinitionFile.
   */
  com.google.protobuf.ByteString getApiDefinitionFile();

  /**
   * <code>string environment = 27;</code>
   * @return The environment.
   */
  java.lang.String getEnvironment();
  /**
   * <code>string environment = 27;</code>
   * @return The bytes for environment.
   */
  com.google.protobuf.ByteString
      getEnvironmentBytes();

  /**
   * <code>bool subscriptionValidation = 28;</code>
   * @return The subscriptionValidation.
   */
  boolean getSubscriptionValidation();
}
