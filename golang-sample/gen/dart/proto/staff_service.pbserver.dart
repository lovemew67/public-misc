///
//  Generated code. Do not modify.
//  source: staff_service.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'staff_service.pb.dart' as $1;
import 'staff_service.pbjson.dart';

export 'staff_service.pb.dart';

abstract class StaffServiceV1ServiceBase extends $pb.GeneratedService {
  $async.Future<$1.ListStaffV1Response> listStaffV1($pb.ServerContext ctx, $1.ListStaffV1Request request);
  $async.Future<$1.CreateStaffV1Response> createStaffV1($pb.ServerContext ctx, $1.CreateStaffV1Request request);
  $async.Future<$1.GetStaffV1Response> getStaffV1($pb.ServerContext ctx, $1.GetStaffV1Request request);
  $async.Future<$1.PatchStaffV1Response> patchStaffV1($pb.ServerContext ctx, $1.PatchStaffV1Request request);
  $async.Future<$1.DeleteStaffV1Response> deleteStaffV1($pb.ServerContext ctx, $1.DeleteStaffV1Request request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'ListStaffV1': return $1.ListStaffV1Request();
      case 'CreateStaffV1': return $1.CreateStaffV1Request();
      case 'GetStaffV1': return $1.GetStaffV1Request();
      case 'PatchStaffV1': return $1.PatchStaffV1Request();
      case 'DeleteStaffV1': return $1.DeleteStaffV1Request();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'ListStaffV1': return this.listStaffV1(ctx, request);
      case 'CreateStaffV1': return this.createStaffV1(ctx, request);
      case 'GetStaffV1': return this.getStaffV1(ctx, request);
      case 'PatchStaffV1': return this.patchStaffV1(ctx, request);
      case 'DeleteStaffV1': return this.deleteStaffV1(ctx, request);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => StaffServiceV1ServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => StaffServiceV1ServiceBase$messageJson;
}

