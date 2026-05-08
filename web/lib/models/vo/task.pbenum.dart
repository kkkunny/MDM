// This is a generated file - do not edit.
//
// Generated from task.proto.

// @dart = 3.3

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: curly_braces_in_flow_control_structures
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

/// 任务阶段
class TaskPhase extends $pb.ProtobufEnum {
  static const TaskPhase TpUnknown =
      TaskPhase._(0, _omitEnumNames ? '' : 'TpUnknown');
  static const TaskPhase TpDownQueued =
      TaskPhase._(51, _omitEnumNames ? '' : 'TpDownQueued');
  static const TaskPhase TpDownRunning =
      TaskPhase._(52, _omitEnumNames ? '' : 'TpDownRunning');
  static const TaskPhase TpDownPaused =
      TaskPhase._(53, _omitEnumNames ? '' : 'TpDownPaused');
  static const TaskPhase TpDownFailed =
      TaskPhase._(54, _omitEnumNames ? '' : 'TpDownFailed');
  static const TaskPhase TpDownCompleted =
      TaskPhase._(55, _omitEnumNames ? '' : 'TpDownCompleted');
  static const TaskPhase TpUpQueued =
      TaskPhase._(101, _omitEnumNames ? '' : 'TpUpQueued');
  static const TaskPhase TpUpRunning =
      TaskPhase._(102, _omitEnumNames ? '' : 'TpUpRunning');
  static const TaskPhase TpUpPaused =
      TaskPhase._(103, _omitEnumNames ? '' : 'TpUpPaused');
  static const TaskPhase TpUpFailed =
      TaskPhase._(104, _omitEnumNames ? '' : 'TpUpFailed');
  static const TaskPhase TpUpCompleted =
      TaskPhase._(105, _omitEnumNames ? '' : 'TpUpCompleted');

  static const $core.List<TaskPhase> values = <TaskPhase>[
    TpUnknown,
    TpDownQueued,
    TpDownRunning,
    TpDownPaused,
    TpDownFailed,
    TpDownCompleted,
    TpUpQueued,
    TpUpRunning,
    TpUpPaused,
    TpUpFailed,
    TpUpCompleted,
  ];

  static final $core.Map<$core.int, TaskPhase> _byValue =
      $pb.ProtobufEnum.initByValue(values);
  static TaskPhase? valueOf($core.int value) => _byValue[value];

  const TaskPhase._(super.value, super.name);
}

/// 操作行为
class Operate extends $pb.ProtobufEnum {
  static const Operate OpUnknown =
      Operate._(0, _omitEnumNames ? '' : 'OpUnknown');
  static const Operate OpDelete =
      Operate._(1, _omitEnumNames ? '' : 'OpDelete');
  static const Operate OpResume =
      Operate._(2, _omitEnumNames ? '' : 'OpResume');
  static const Operate OpPause = Operate._(3, _omitEnumNames ? '' : 'OpPause');
  static const Operate OpRetry = Operate._(4, _omitEnumNames ? '' : 'OpRetry');

  static const $core.List<Operate> values = <Operate>[
    OpUnknown,
    OpDelete,
    OpResume,
    OpPause,
    OpRetry,
  ];

  static final $core.List<Operate?> _byValue =
      $pb.ProtobufEnum.$_initByValueList(values, 4);
  static Operate? valueOf($core.int value) =>
      value < 0 || value >= _byValue.length ? null : _byValue[value];

  const Operate._(super.value, super.name);
}

const $core.bool _omitEnumNames =
    $core.bool.fromEnvironment('protobuf.omit_enum_names');
