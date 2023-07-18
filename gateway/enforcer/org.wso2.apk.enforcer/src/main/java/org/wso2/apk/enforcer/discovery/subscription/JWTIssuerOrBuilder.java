// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/subscription/jwtIssuer.proto

package org.wso2.apk.enforcer.discovery.subscription;

public interface JWTIssuerOrBuilder extends
    // @@protoc_insertion_point(interface_extends:wso2.discovery.subscription.JWTIssuer)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string eventId = 1;</code>
   * @return The eventId.
   */
  java.lang.String getEventId();
  /**
   * <code>string eventId = 1;</code>
   * @return The bytes for eventId.
   */
  com.google.protobuf.ByteString
      getEventIdBytes();

  /**
   * <code>string name = 2;</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 2;</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>string organization = 3;</code>
   * @return The organization.
   */
  java.lang.String getOrganization();
  /**
   * <code>string organization = 3;</code>
   * @return The bytes for organization.
   */
  com.google.protobuf.ByteString
      getOrganizationBytes();

  /**
   * <code>string issuer = 4;</code>
   * @return The issuer.
   */
  java.lang.String getIssuer();
  /**
   * <code>string issuer = 4;</code>
   * @return The bytes for issuer.
   */
  com.google.protobuf.ByteString
      getIssuerBytes();

  /**
   * <code>.wso2.discovery.subscription.Certificate certificate = 5;</code>
   * @return Whether the certificate field is set.
   */
  boolean hasCertificate();
  /**
   * <code>.wso2.discovery.subscription.Certificate certificate = 5;</code>
   * @return The certificate.
   */
  org.wso2.apk.enforcer.discovery.subscription.Certificate getCertificate();
  /**
   * <code>.wso2.discovery.subscription.Certificate certificate = 5;</code>
   */
  org.wso2.apk.enforcer.discovery.subscription.CertificateOrBuilder getCertificateOrBuilder();

  /**
   * <code>string consumerKeyClaim = 6;</code>
   * @return The consumerKeyClaim.
   */
  java.lang.String getConsumerKeyClaim();
  /**
   * <code>string consumerKeyClaim = 6;</code>
   * @return The bytes for consumerKeyClaim.
   */
  com.google.protobuf.ByteString
      getConsumerKeyClaimBytes();

  /**
   * <code>string scopesClaim = 7;</code>
   * @return The scopesClaim.
   */
  java.lang.String getScopesClaim();
  /**
   * <code>string scopesClaim = 7;</code>
   * @return The bytes for scopesClaim.
   */
  com.google.protobuf.ByteString
      getScopesClaimBytes();

  /**
   * <code>map&lt;string, string&gt; claimMapping = 8;</code>
   */
  int getClaimMappingCount();
  /**
   * <code>map&lt;string, string&gt; claimMapping = 8;</code>
   */
  boolean containsClaimMapping(
      java.lang.String key);
  /**
   * Use {@link #getClaimMappingMap()} instead.
   */
  @java.lang.Deprecated
  java.util.Map<java.lang.String, java.lang.String>
  getClaimMapping();
  /**
   * <code>map&lt;string, string&gt; claimMapping = 8;</code>
   */
  java.util.Map<java.lang.String, java.lang.String>
  getClaimMappingMap();
  /**
   * <code>map&lt;string, string&gt; claimMapping = 8;</code>
   */

  java.lang.String getClaimMappingOrDefault(
      java.lang.String key,
      java.lang.String defaultValue);
  /**
   * <code>map&lt;string, string&gt; claimMapping = 8;</code>
   */

  java.lang.String getClaimMappingOrThrow(
      java.lang.String key);
}
