// This is a generated file - do not edit.
//
// Generated from task.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use taskPhaseDescriptor instead')
const TaskPhase$json = {
  '1': 'TaskPhase',
  '2': [
    {'1': 'TpUnknown', '2': 0},
    {'1': 'TpDownQueued', '2': 51},
    {'1': 'TpDownRunning', '2': 52},
    {'1': 'TpDownPaused', '2': 53},
    {'1': 'TpDownFailed', '2': 54},
    {'1': 'TpDownCompleted', '2': 55},
    {'1': 'TpUpQueued', '2': 101},
    {'1': 'TpUpRunning', '2': 102},
    {'1': 'TpUpPaused', '2': 103},
    {'1': 'TpUpFailed', '2': 104},
    {'1': 'TpUpCompleted', '2': 105},
  ],
};

/// Descriptor for `TaskPhase`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List taskPhaseDescriptor = $convert.base64Decode(
    'CglUYXNrUGhhc2USDQoJVHBVbmtub3duEAASEAoMVHBEb3duUXVldWVkEDMSEQoNVHBEb3duUn'
    'VubmluZxA0EhAKDFRwRG93blBhdXNlZBA1EhAKDFRwRG93bkZhaWxlZBA2EhMKD1RwRG93bkNv'
    'bXBsZXRlZBA3Eg4KClRwVXBRdWV1ZWQQZRIPCgtUcFVwUnVubmluZxBmEg4KClRwVXBQYXVzZW'
    'QQZxIOCgpUcFVwRmFpbGVkEGgSEQoNVHBVcENvbXBsZXRlZBBp');

@$core.Deprecated('Use operateDescriptor instead')
const Operate$json = {
  '1': 'Operate',
  '2': [
    {'1': 'OpUnknown', '2': 0},
    {'1': 'OpDelete', '2': 1},
    {'1': 'OpResume', '2': 2},
    {'1': 'OpPause', '2': 3},
    {'1': 'OpRetry', '2': 4},
  ],
};

/// Descriptor for `Operate`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List operateDescriptor = $convert.base64Decode(
    'CgdPcGVyYXRlEg0KCU9wVW5rbm93bhAAEgwKCE9wRGVsZXRlEAESDAoIT3BSZXN1bWUQAhILCg'
    'dPcFBhdXNlEAMSCwoHT3BSZXRyeRAE');

@$core.Deprecated('Use taskDescriptor instead')
const Task$json = {
  '1': 'Task',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    {
      '1': 'phase',
      '3': 3,
      '4': 1,
      '5': 14,
      '6': '.task.TaskPhase',
      '10': 'phase'
    },
    {'1': 'size', '3': 4, '4': 1, '5': 4, '10': 'size'},
    {
      '1': 'category',
      '3': 5,
      '4': 1,
      '5': 11,
      '6': '.task.Category',
      '9': 0,
      '10': 'category',
      '17': true
    },
    {'1': 'created_at', '3': 6, '4': 1, '5': 4, '10': 'createdAt'},
    {
      '1': 'download_stats',
      '3': 50,
      '4': 1,
      '5': 11,
      '6': '.task.DownloadStats',
      '9': 1,
      '10': 'downloadStats',
      '17': true
    },
    {
      '1': 'upload_stats',
      '3': 51,
      '4': 1,
      '5': 11,
      '6': '.task.UploadStats',
      '9': 2,
      '10': 'uploadStats',
      '17': true
    },
  ],
  '8': [
    {'1': '_category'},
    {'1': '_download_stats'},
    {'1': '_upload_stats'},
  ],
};

/// Descriptor for `Task`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List taskDescriptor = $convert.base64Decode(
    'CgRUYXNrEg4KAmlkGAEgASgJUgJpZBISCgRuYW1lGAIgASgJUgRuYW1lEiUKBXBoYXNlGAMgAS'
    'gOMg8udGFzay5UYXNrUGhhc2VSBXBoYXNlEhIKBHNpemUYBCABKARSBHNpemUSLwoIY2F0ZWdv'
    'cnkYBSABKAsyDi50YXNrLkNhdGVnb3J5SABSCGNhdGVnb3J5iAEBEh0KCmNyZWF0ZWRfYXQYBi'
    'ABKARSCWNyZWF0ZWRBdBI/Cg5kb3dubG9hZF9zdGF0cxgyIAEoCzITLnRhc2suRG93bmxvYWRT'
    'dGF0c0gBUg1kb3dubG9hZFN0YXRziAEBEjkKDHVwbG9hZF9zdGF0cxgzIAEoCzIRLnRhc2suVX'
    'Bsb2FkU3RhdHNIAlILdXBsb2FkU3RhdHOIAQFCCwoJX2NhdGVnb3J5QhEKD19kb3dubG9hZF9z'
    'dGF0c0IPCg1fdXBsb2FkX3N0YXRz');

@$core.Deprecated('Use categoryDescriptor instead')
const Category$json = {
  '1': 'Category',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    {'1': 'icon', '3': 2, '4': 1, '5': 9, '10': 'icon'},
  ],
};

/// Descriptor for `Category`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List categoryDescriptor = $convert.base64Decode(
    'CghDYXRlZ29yeRISCgRuYW1lGAEgASgJUgRuYW1lEhIKBGljb24YAiABKAlSBGljb24=');

@$core.Deprecated('Use downloadStatsDescriptor instead')
const DownloadStats$json = {
  '1': 'DownloadStats',
  '2': [
    {'1': 'speed', '3': 1, '4': 1, '5': 4, '10': 'speed'},
    {'1': 'size', '3': 2, '4': 1, '5': 4, '10': 'size'},
  ],
};

/// Descriptor for `DownloadStats`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List downloadStatsDescriptor = $convert.base64Decode(
    'Cg1Eb3dubG9hZFN0YXRzEhQKBXNwZWVkGAEgASgEUgVzcGVlZBISCgRzaXplGAIgASgEUgRzaX'
    'pl');

@$core.Deprecated('Use uploadStatsDescriptor instead')
const UploadStats$json = {
  '1': 'UploadStats',
  '2': [
    {'1': 'speed', '3': 1, '4': 1, '5': 4, '10': 'speed'},
    {'1': 'size', '3': 2, '4': 1, '5': 4, '10': 'size'},
  ],
};

/// Descriptor for `UploadStats`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List uploadStatsDescriptor = $convert.base64Decode(
    'CgtVcGxvYWRTdGF0cxIUCgVzcGVlZBgBIAEoBFIFc3BlZWQSEgoEc2l6ZRgCIAEoBFIEc2l6ZQ'
    '==');

@$core.Deprecated('Use listTasksResponseDescriptor instead')
const ListTasksResponse$json = {
  '1': 'ListTasksResponse',
  '2': [
    {'1': 'tasks', '3': 1, '4': 3, '5': 11, '6': '.task.Task', '10': 'tasks'},
  ],
};

/// Descriptor for `ListTasksResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List listTasksResponseDescriptor = $convert.base64Decode(
    'ChFMaXN0VGFza3NSZXNwb25zZRIgCgV0YXNrcxgBIAMoCzIKLnRhc2suVGFza1IFdGFza3M=');

@$core.Deprecated('Use createTaskRequestDescriptor instead')
const CreateTaskRequest$json = {
  '1': 'CreateTaskRequest',
  '2': [
    {'1': 'link', '3': 1, '4': 1, '5': 9, '10': 'link'},
    {'1': 'name', '3': 10, '4': 1, '5': 9, '10': 'name'},
    {
      '1': 'category',
      '3': 11,
      '4': 1,
      '5': 9,
      '9': 0,
      '10': 'category',
      '17': true
    },
  ],
  '8': [
    {'1': '_category'},
  ],
};

/// Descriptor for `CreateTaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTaskRequestDescriptor = $convert.base64Decode(
    'ChFDcmVhdGVUYXNrUmVxdWVzdBISCgRsaW5rGAEgASgJUgRsaW5rEhIKBG5hbWUYCiABKAlSBG'
    '5hbWUSHwoIY2F0ZWdvcnkYCyABKAlIAFIIY2F0ZWdvcnmIAQFCCwoJX2NhdGVnb3J5');

@$core.Deprecated('Use createTaskResponseDescriptor instead')
const CreateTaskResponse$json = {
  '1': 'CreateTaskResponse',
  '2': [
    {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `CreateTaskResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createTaskResponseDescriptor =
    $convert.base64Decode('ChJDcmVhdGVUYXNrUmVzcG9uc2USDgoCaWQYASABKAlSAmlk');

@$core.Deprecated('Use operateTasksRequestDescriptor instead')
const OperateTasksRequest$json = {
  '1': 'OperateTasksRequest',
  '2': [
    {'1': 'ids', '3': 1, '4': 3, '5': 9, '10': 'ids'},
    {
      '1': 'operate',
      '3': 2,
      '4': 1,
      '5': 14,
      '6': '.task.Operate',
      '10': 'operate'
    },
  ],
};

/// Descriptor for `OperateTasksRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List operateTasksRequestDescriptor = $convert.base64Decode(
    'ChNPcGVyYXRlVGFza3NSZXF1ZXN0EhAKA2lkcxgBIAMoCVIDaWRzEicKB29wZXJhdGUYAiABKA'
    '4yDS50YXNrLk9wZXJhdGVSB29wZXJhdGU=');

@$core.Deprecated('Use statInfoResponseDescriptor instead')
const StatInfoResponse$json = {
  '1': 'StatInfoResponse',
  '2': [
    {'1': 'task_count', '3': 1, '4': 1, '5': 4, '10': 'taskCount'},
    {'1': 'dl_count', '3': 2, '4': 1, '5': 4, '10': 'dlCount'},
    {'1': 'dl_speed', '3': 3, '4': 1, '5': 4, '10': 'dlSpeed'},
    {'1': 'ul_count', '3': 4, '4': 1, '5': 4, '10': 'ulCount'},
    {'1': 'ul_speed', '3': 5, '4': 1, '5': 4, '10': 'ulSpeed'},
  ],
};

/// Descriptor for `StatInfoResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List statInfoResponseDescriptor = $convert.base64Decode(
    'ChBTdGF0SW5mb1Jlc3BvbnNlEh0KCnRhc2tfY291bnQYASABKARSCXRhc2tDb3VudBIZCghkbF'
    '9jb3VudBgCIAEoBFIHZGxDb3VudBIZCghkbF9zcGVlZBgDIAEoBFIHZGxTcGVlZBIZCgh1bF9j'
    'b3VudBgEIAEoBFIHdWxDb3VudBIZCgh1bF9zcGVlZBgFIAEoBFIHdWxTcGVlZA==');
