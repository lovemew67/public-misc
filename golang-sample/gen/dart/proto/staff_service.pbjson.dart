///
//  Generated code. Do not modify.
//  source: staff_service.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'staff.pbjson.dart' as $0;

const ListStaffV1Request$json = const {
  '1': 'ListStaffV1Request',
  '2': const [
    const {'1': 'offset', '3': 1, '4': 1, '5': 5, '10': 'offset'},
    const {'1': 'limit', '3': 2, '4': 1, '5': 5, '10': 'limit'},
  ],
};

const ListStaffV1Response$json = const {
  '1': 'ListStaffV1Response',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 5, '10': 'code'},
    const {'1': 'sub_code', '3': 2, '4': 1, '5': 5, '10': 'subCode'},
    const {'1': 'error_message', '3': 3, '4': 1, '5': 9, '10': 'errorMessage'},
    const {'1': 'total', '3': 4, '4': 1, '5': 5, '10': 'total'},
    const {'1': 'staff', '3': 5, '4': 3, '5': 11, '6': '.proto.StaffV1', '10': 'staff'},
  ],
};

const CreateStaffV1Request$json = const {
  '1': 'CreateStaffV1Request',
  '2': const [
    const {'1': 'staff', '3': 1, '4': 1, '5': 11, '6': '.proto.StaffV1', '10': 'staff'},
  ],
};

const CreateStaffV1Response$json = const {
  '1': 'CreateStaffV1Response',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 5, '10': 'code'},
    const {'1': 'sub_code', '3': 2, '4': 1, '5': 5, '10': 'subCode'},
    const {'1': 'error_message', '3': 3, '4': 1, '5': 9, '10': 'errorMessage'},
    const {'1': 'staff', '3': 4, '4': 1, '5': 11, '6': '.proto.StaffV1', '10': 'staff'},
  ],
};

const GetStaffV1Request$json = const {
  '1': 'GetStaffV1Request',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

const GetStaffV1Response$json = const {
  '1': 'GetStaffV1Response',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 5, '10': 'code'},
    const {'1': 'sub_code', '3': 2, '4': 1, '5': 5, '10': 'subCode'},
    const {'1': 'error_message', '3': 3, '4': 1, '5': 9, '10': 'errorMessage'},
    const {'1': 'staff', '3': 4, '4': 1, '5': 11, '6': '.proto.StaffV1', '10': 'staff'},
  ],
};

const PatchStaffV1Request$json = const {
  '1': 'PatchStaffV1Request',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'email', '3': 3, '4': 1, '5': 9, '10': 'email'},
    const {'1': 'avatar_url', '3': 4, '4': 1, '5': 9, '10': 'avatarUrl'},
  ],
};

const PatchStaffV1Response$json = const {
  '1': 'PatchStaffV1Response',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 5, '10': 'code'},
    const {'1': 'sub_code', '3': 2, '4': 1, '5': 5, '10': 'subCode'},
    const {'1': 'error_message', '3': 3, '4': 1, '5': 9, '10': 'errorMessage'},
  ],
};

const DeleteStaffV1Request$json = const {
  '1': 'DeleteStaffV1Request',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

const DeleteStaffV1Response$json = const {
  '1': 'DeleteStaffV1Response',
  '2': const [
    const {'1': 'code', '3': 1, '4': 1, '5': 5, '10': 'code'},
    const {'1': 'sub_code', '3': 2, '4': 1, '5': 5, '10': 'subCode'},
    const {'1': 'error_message', '3': 3, '4': 1, '5': 9, '10': 'errorMessage'},
  ],
};

const StaffServiceV1ServiceBase$json = const {
  '1': 'StaffServiceV1',
  '2': const [
    const {'1': 'ListStaffV1', '2': '.proto.ListStaffV1Request', '3': '.proto.ListStaffV1Response', '4': const {}},
    const {'1': 'CreateStaffV1', '2': '.proto.CreateStaffV1Request', '3': '.proto.CreateStaffV1Response', '4': const {}},
    const {'1': 'GetStaffV1', '2': '.proto.GetStaffV1Request', '3': '.proto.GetStaffV1Response', '4': const {}},
    const {'1': 'PatchStaffV1', '2': '.proto.PatchStaffV1Request', '3': '.proto.PatchStaffV1Response', '4': const {}},
    const {'1': 'DeleteStaffV1', '2': '.proto.DeleteStaffV1Request', '3': '.proto.DeleteStaffV1Response', '4': const {}},
  ],
};

const StaffServiceV1ServiceBase$messageJson = const {
  '.proto.ListStaffV1Request': ListStaffV1Request$json,
  '.proto.ListStaffV1Response': ListStaffV1Response$json,
  '.proto.StaffV1': $0.StaffV1$json,
  '.proto.CreateStaffV1Request': CreateStaffV1Request$json,
  '.proto.CreateStaffV1Response': CreateStaffV1Response$json,
  '.proto.GetStaffV1Request': GetStaffV1Request$json,
  '.proto.GetStaffV1Response': GetStaffV1Response$json,
  '.proto.PatchStaffV1Request': PatchStaffV1Request$json,
  '.proto.PatchStaffV1Response': PatchStaffV1Response$json,
  '.proto.DeleteStaffV1Request': DeleteStaffV1Request$json,
  '.proto.DeleteStaffV1Response': DeleteStaffV1Response$json,
};

