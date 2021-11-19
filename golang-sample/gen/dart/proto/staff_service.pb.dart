///
//  Generated code. Do not modify.
//  source: staff_service.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'staff.pb.dart' as $0;

class ListStaffV1Request extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ListStaffV1Request', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'offset', $pb.PbFieldType.O3)
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'limit', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  ListStaffV1Request._() : super();
  factory ListStaffV1Request() => create();
  factory ListStaffV1Request.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListStaffV1Request.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListStaffV1Request clone() => ListStaffV1Request()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListStaffV1Request copyWith(void Function(ListStaffV1Request) updates) => super.copyWith((message) => updates(message as ListStaffV1Request)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListStaffV1Request create() => ListStaffV1Request._();
  ListStaffV1Request createEmptyInstance() => create();
  static $pb.PbList<ListStaffV1Request> createRepeated() => $pb.PbList<ListStaffV1Request>();
  @$core.pragma('dart2js:noInline')
  static ListStaffV1Request getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListStaffV1Request>(create);
  static ListStaffV1Request _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get offset => $_getIZ(0);
  @$pb.TagNumber(1)
  set offset($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasOffset() => $_has(0);
  @$pb.TagNumber(1)
  void clearOffset() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get limit => $_getIZ(1);
  @$pb.TagNumber(2)
  set limit($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLimit() => $_has(1);
  @$pb.TagNumber(2)
  void clearLimit() => clearField(2);
}

class ListStaffV1Response extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'ListStaffV1Response', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'subCode', $pb.PbFieldType.O3)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errorMessage')
    ..a<$core.int>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'total', $pb.PbFieldType.O3)
    ..pc<$0.StaffV1>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'staff', $pb.PbFieldType.PM, subBuilder: $0.StaffV1.create)
    ..hasRequiredFields = false
  ;

  ListStaffV1Response._() : super();
  factory ListStaffV1Response() => create();
  factory ListStaffV1Response.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ListStaffV1Response.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  ListStaffV1Response clone() => ListStaffV1Response()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  ListStaffV1Response copyWith(void Function(ListStaffV1Response) updates) => super.copyWith((message) => updates(message as ListStaffV1Response)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ListStaffV1Response create() => ListStaffV1Response._();
  ListStaffV1Response createEmptyInstance() => create();
  static $pb.PbList<ListStaffV1Response> createRepeated() => $pb.PbList<ListStaffV1Response>();
  @$core.pragma('dart2js:noInline')
  static ListStaffV1Response getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ListStaffV1Response>(create);
  static ListStaffV1Response _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get code => $_getIZ(0);
  @$pb.TagNumber(1)
  set code($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get subCode => $_getIZ(1);
  @$pb.TagNumber(2)
  set subCode($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSubCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearSubCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get errorMessage => $_getSZ(2);
  @$pb.TagNumber(3)
  set errorMessage($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasErrorMessage() => $_has(2);
  @$pb.TagNumber(3)
  void clearErrorMessage() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get total => $_getIZ(3);
  @$pb.TagNumber(4)
  set total($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTotal() => $_has(3);
  @$pb.TagNumber(4)
  void clearTotal() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$0.StaffV1> get staff => $_getList(4);
}

class CreateStaffV1Request extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CreateStaffV1Request', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..aOM<$0.StaffV1>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'staff', subBuilder: $0.StaffV1.create)
    ..hasRequiredFields = false
  ;

  CreateStaffV1Request._() : super();
  factory CreateStaffV1Request() => create();
  factory CreateStaffV1Request.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateStaffV1Request.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateStaffV1Request clone() => CreateStaffV1Request()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateStaffV1Request copyWith(void Function(CreateStaffV1Request) updates) => super.copyWith((message) => updates(message as CreateStaffV1Request)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CreateStaffV1Request create() => CreateStaffV1Request._();
  CreateStaffV1Request createEmptyInstance() => create();
  static $pb.PbList<CreateStaffV1Request> createRepeated() => $pb.PbList<CreateStaffV1Request>();
  @$core.pragma('dart2js:noInline')
  static CreateStaffV1Request getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateStaffV1Request>(create);
  static CreateStaffV1Request _defaultInstance;

  @$pb.TagNumber(1)
  $0.StaffV1 get staff => $_getN(0);
  @$pb.TagNumber(1)
  set staff($0.StaffV1 v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStaff() => $_has(0);
  @$pb.TagNumber(1)
  void clearStaff() => clearField(1);
  @$pb.TagNumber(1)
  $0.StaffV1 ensureStaff() => $_ensure(0);
}

class CreateStaffV1Response extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CreateStaffV1Response', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'subCode', $pb.PbFieldType.O3)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errorMessage')
    ..aOM<$0.StaffV1>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'staff', subBuilder: $0.StaffV1.create)
    ..hasRequiredFields = false
  ;

  CreateStaffV1Response._() : super();
  factory CreateStaffV1Response() => create();
  factory CreateStaffV1Response.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateStaffV1Response.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateStaffV1Response clone() => CreateStaffV1Response()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateStaffV1Response copyWith(void Function(CreateStaffV1Response) updates) => super.copyWith((message) => updates(message as CreateStaffV1Response)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CreateStaffV1Response create() => CreateStaffV1Response._();
  CreateStaffV1Response createEmptyInstance() => create();
  static $pb.PbList<CreateStaffV1Response> createRepeated() => $pb.PbList<CreateStaffV1Response>();
  @$core.pragma('dart2js:noInline')
  static CreateStaffV1Response getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateStaffV1Response>(create);
  static CreateStaffV1Response _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get code => $_getIZ(0);
  @$pb.TagNumber(1)
  set code($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get subCode => $_getIZ(1);
  @$pb.TagNumber(2)
  set subCode($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSubCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearSubCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get errorMessage => $_getSZ(2);
  @$pb.TagNumber(3)
  set errorMessage($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasErrorMessage() => $_has(2);
  @$pb.TagNumber(3)
  void clearErrorMessage() => clearField(3);

  @$pb.TagNumber(4)
  $0.StaffV1 get staff => $_getN(3);
  @$pb.TagNumber(4)
  set staff($0.StaffV1 v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasStaff() => $_has(3);
  @$pb.TagNumber(4)
  void clearStaff() => clearField(4);
  @$pb.TagNumber(4)
  $0.StaffV1 ensureStaff() => $_ensure(3);
}

class GetStaffV1Request extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetStaffV1Request', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  GetStaffV1Request._() : super();
  factory GetStaffV1Request() => create();
  factory GetStaffV1Request.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetStaffV1Request.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetStaffV1Request clone() => GetStaffV1Request()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetStaffV1Request copyWith(void Function(GetStaffV1Request) updates) => super.copyWith((message) => updates(message as GetStaffV1Request)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetStaffV1Request create() => GetStaffV1Request._();
  GetStaffV1Request createEmptyInstance() => create();
  static $pb.PbList<GetStaffV1Request> createRepeated() => $pb.PbList<GetStaffV1Request>();
  @$core.pragma('dart2js:noInline')
  static GetStaffV1Request getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetStaffV1Request>(create);
  static GetStaffV1Request _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class GetStaffV1Response extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GetStaffV1Response', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'subCode', $pb.PbFieldType.O3)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errorMessage')
    ..aOM<$0.StaffV1>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'staff', subBuilder: $0.StaffV1.create)
    ..hasRequiredFields = false
  ;

  GetStaffV1Response._() : super();
  factory GetStaffV1Response() => create();
  factory GetStaffV1Response.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetStaffV1Response.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GetStaffV1Response clone() => GetStaffV1Response()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GetStaffV1Response copyWith(void Function(GetStaffV1Response) updates) => super.copyWith((message) => updates(message as GetStaffV1Response)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetStaffV1Response create() => GetStaffV1Response._();
  GetStaffV1Response createEmptyInstance() => create();
  static $pb.PbList<GetStaffV1Response> createRepeated() => $pb.PbList<GetStaffV1Response>();
  @$core.pragma('dart2js:noInline')
  static GetStaffV1Response getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetStaffV1Response>(create);
  static GetStaffV1Response _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get code => $_getIZ(0);
  @$pb.TagNumber(1)
  set code($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get subCode => $_getIZ(1);
  @$pb.TagNumber(2)
  set subCode($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSubCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearSubCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get errorMessage => $_getSZ(2);
  @$pb.TagNumber(3)
  set errorMessage($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasErrorMessage() => $_has(2);
  @$pb.TagNumber(3)
  void clearErrorMessage() => clearField(3);

  @$pb.TagNumber(4)
  $0.StaffV1 get staff => $_getN(3);
  @$pb.TagNumber(4)
  set staff($0.StaffV1 v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasStaff() => $_has(3);
  @$pb.TagNumber(4)
  void clearStaff() => clearField(4);
  @$pb.TagNumber(4)
  $0.StaffV1 ensureStaff() => $_ensure(3);
}

class PatchStaffV1Request extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PatchStaffV1Request', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'email')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'avatarUrl')
    ..hasRequiredFields = false
  ;

  PatchStaffV1Request._() : super();
  factory PatchStaffV1Request() => create();
  factory PatchStaffV1Request.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PatchStaffV1Request.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PatchStaffV1Request clone() => PatchStaffV1Request()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PatchStaffV1Request copyWith(void Function(PatchStaffV1Request) updates) => super.copyWith((message) => updates(message as PatchStaffV1Request)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PatchStaffV1Request create() => PatchStaffV1Request._();
  PatchStaffV1Request createEmptyInstance() => create();
  static $pb.PbList<PatchStaffV1Request> createRepeated() => $pb.PbList<PatchStaffV1Request>();
  @$core.pragma('dart2js:noInline')
  static PatchStaffV1Request getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PatchStaffV1Request>(create);
  static PatchStaffV1Request _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get email => $_getSZ(2);
  @$pb.TagNumber(3)
  set email($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasEmail() => $_has(2);
  @$pb.TagNumber(3)
  void clearEmail() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get avatarUrl => $_getSZ(3);
  @$pb.TagNumber(4)
  set avatarUrl($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAvatarUrl() => $_has(3);
  @$pb.TagNumber(4)
  void clearAvatarUrl() => clearField(4);
}

class PatchStaffV1Response extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'PatchStaffV1Response', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'subCode', $pb.PbFieldType.O3)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errorMessage')
    ..hasRequiredFields = false
  ;

  PatchStaffV1Response._() : super();
  factory PatchStaffV1Response() => create();
  factory PatchStaffV1Response.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PatchStaffV1Response.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  PatchStaffV1Response clone() => PatchStaffV1Response()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  PatchStaffV1Response copyWith(void Function(PatchStaffV1Response) updates) => super.copyWith((message) => updates(message as PatchStaffV1Response)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PatchStaffV1Response create() => PatchStaffV1Response._();
  PatchStaffV1Response createEmptyInstance() => create();
  static $pb.PbList<PatchStaffV1Response> createRepeated() => $pb.PbList<PatchStaffV1Response>();
  @$core.pragma('dart2js:noInline')
  static PatchStaffV1Response getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PatchStaffV1Response>(create);
  static PatchStaffV1Response _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get code => $_getIZ(0);
  @$pb.TagNumber(1)
  set code($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get subCode => $_getIZ(1);
  @$pb.TagNumber(2)
  set subCode($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSubCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearSubCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get errorMessage => $_getSZ(2);
  @$pb.TagNumber(3)
  set errorMessage($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasErrorMessage() => $_has(2);
  @$pb.TagNumber(3)
  void clearErrorMessage() => clearField(3);
}

class DeleteStaffV1Request extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DeleteStaffV1Request', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id')
    ..hasRequiredFields = false
  ;

  DeleteStaffV1Request._() : super();
  factory DeleteStaffV1Request() => create();
  factory DeleteStaffV1Request.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteStaffV1Request.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteStaffV1Request clone() => DeleteStaffV1Request()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteStaffV1Request copyWith(void Function(DeleteStaffV1Request) updates) => super.copyWith((message) => updates(message as DeleteStaffV1Request)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteStaffV1Request create() => DeleteStaffV1Request._();
  DeleteStaffV1Request createEmptyInstance() => create();
  static $pb.PbList<DeleteStaffV1Request> createRepeated() => $pb.PbList<DeleteStaffV1Request>();
  @$core.pragma('dart2js:noInline')
  static DeleteStaffV1Request getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteStaffV1Request>(create);
  static DeleteStaffV1Request _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class DeleteStaffV1Response extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'DeleteStaffV1Response', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'proto'), createEmptyInstance: create)
    ..a<$core.int>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'code', $pb.PbFieldType.O3)
    ..a<$core.int>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'subCode', $pb.PbFieldType.O3)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'errorMessage')
    ..hasRequiredFields = false
  ;

  DeleteStaffV1Response._() : super();
  factory DeleteStaffV1Response() => create();
  factory DeleteStaffV1Response.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DeleteStaffV1Response.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DeleteStaffV1Response clone() => DeleteStaffV1Response()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DeleteStaffV1Response copyWith(void Function(DeleteStaffV1Response) updates) => super.copyWith((message) => updates(message as DeleteStaffV1Response)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static DeleteStaffV1Response create() => DeleteStaffV1Response._();
  DeleteStaffV1Response createEmptyInstance() => create();
  static $pb.PbList<DeleteStaffV1Response> createRepeated() => $pb.PbList<DeleteStaffV1Response>();
  @$core.pragma('dart2js:noInline')
  static DeleteStaffV1Response getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DeleteStaffV1Response>(create);
  static DeleteStaffV1Response _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get code => $_getIZ(0);
  @$pb.TagNumber(1)
  set code($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get subCode => $_getIZ(1);
  @$pb.TagNumber(2)
  set subCode($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasSubCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearSubCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get errorMessage => $_getSZ(2);
  @$pb.TagNumber(3)
  set errorMessage($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasErrorMessage() => $_has(2);
  @$pb.TagNumber(3)
  void clearErrorMessage() => clearField(3);
}

class StaffServiceV1Api {
  $pb.RpcClient _client;
  StaffServiceV1Api(this._client);

  $async.Future<ListStaffV1Response> listStaffV1($pb.ClientContext ctx, ListStaffV1Request request) {
    var emptyResponse = ListStaffV1Response();
    return _client.invoke<ListStaffV1Response>(ctx, 'StaffServiceV1', 'ListStaffV1', request, emptyResponse);
  }
  $async.Future<CreateStaffV1Response> createStaffV1($pb.ClientContext ctx, CreateStaffV1Request request) {
    var emptyResponse = CreateStaffV1Response();
    return _client.invoke<CreateStaffV1Response>(ctx, 'StaffServiceV1', 'CreateStaffV1', request, emptyResponse);
  }
  $async.Future<GetStaffV1Response> getStaffV1($pb.ClientContext ctx, GetStaffV1Request request) {
    var emptyResponse = GetStaffV1Response();
    return _client.invoke<GetStaffV1Response>(ctx, 'StaffServiceV1', 'GetStaffV1', request, emptyResponse);
  }
  $async.Future<PatchStaffV1Response> patchStaffV1($pb.ClientContext ctx, PatchStaffV1Request request) {
    var emptyResponse = PatchStaffV1Response();
    return _client.invoke<PatchStaffV1Response>(ctx, 'StaffServiceV1', 'PatchStaffV1', request, emptyResponse);
  }
  $async.Future<DeleteStaffV1Response> deleteStaffV1($pb.ClientContext ctx, DeleteStaffV1Request request) {
    var emptyResponse = DeleteStaffV1Response();
    return _client.invoke<DeleteStaffV1Response>(ctx, 'StaffServiceV1', 'DeleteStaffV1', request, emptyResponse);
  }
}

