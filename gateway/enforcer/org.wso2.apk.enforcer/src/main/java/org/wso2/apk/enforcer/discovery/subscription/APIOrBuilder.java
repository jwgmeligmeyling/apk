// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/subscription/subscription.proto

package org.wso2.apk.enforcer.discovery.subscription;

public interface APIOrBuilder extends
    // @@protoc_insertion_point(interface_extends:wso2.discovery.subscription.API)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string name = 1;</code>
   * @return The name.
   */
  java.lang.String getName();
  /**
   * <code>string name = 1;</code>
   * @return The bytes for name.
   */
  com.google.protobuf.ByteString
      getNameBytes();

  /**
   * <code>repeated string versions = 2;</code>
   * @return A list containing the versions.
   */
  java.util.List<java.lang.String>
      getVersionsList();
  /**
   * <code>repeated string versions = 2;</code>
   * @return The count of versions.
   */
  int getVersionsCount();
  /**
   * <code>repeated string versions = 2;</code>
   * @param index The index of the element to return.
   * @return The versions at the given index.
   */
  java.lang.String getVersions(int index);
  /**
   * <code>repeated string versions = 2;</code>
   * @param index The index of the value to return.
   * @return The bytes of the versions at the given index.
   */
  com.google.protobuf.ByteString
      getVersionsBytes(int index);
}
