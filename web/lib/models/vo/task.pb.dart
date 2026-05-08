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

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'task.pbenum.dart';

export 'package:protobuf/protobuf.dart' show GeneratedMessageGenericExtensions;

export 'task.pbenum.dart';

/// 任务信息
class Task extends $pb.GeneratedMessage {
  factory Task({
    $core.String? id,
    $core.String? name,
    TaskPhase? phase,
    $fixnum.Int64? size,
    Category? category,
    $fixnum.Int64? createdAt,
    DownloadStats? downloadStats,
    UploadStats? uploadStats,
  }) {
    final result = create();
    if (id != null) result.id = id;
    if (name != null) result.name = name;
    if (phase != null) result.phase = phase;
    if (size != null) result.size = size;
    if (category != null) result.category = category;
    if (createdAt != null) result.createdAt = createdAt;
    if (downloadStats != null) result.downloadStats = downloadStats;
    if (uploadStats != null) result.uploadStats = uploadStats;
    return result;
  }

  Task._();

  factory Task.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory Task.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Task',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'task'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'id')
    ..aOS(2, _omitFieldNames ? '' : 'name')
    ..e<TaskPhase>(3, _omitFieldNames ? '' : 'phase', $pb.PbFieldType.OE,
        defaultOrMaker: TaskPhase.TpUnknown,
        valueOf: TaskPhase.valueOf,
        enumValues: TaskPhase.values)
    ..a<$fixnum.Int64>(4, _omitFieldNames ? '' : 'size', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOM<Category>(5, _omitFieldNames ? '' : 'category',
        subBuilder: Category.create)
    ..a<$fixnum.Int64>(
        6, _omitFieldNames ? '' : 'createdAt', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOM<DownloadStats>(50, _omitFieldNames ? '' : 'downloadStats',
        subBuilder: DownloadStats.create)
    ..aOM<UploadStats>(51, _omitFieldNames ? '' : 'uploadStats',
        subBuilder: UploadStats.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Task clone() => Task()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Task copyWith(void Function(Task) updates) =>
      super.copyWith((message) => updates(message as Task)) as Task;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Task create() => Task._();
  @$core.override
  Task createEmptyInstance() => create();
  static $pb.PbList<Task> createRepeated() => [] as $pb.PbList<Task>;
  @$core.pragma('dart2js:noInline')
  static Task getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Task>(create);
  static Task? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => $_clearField(2);

  @$pb.TagNumber(3)
  TaskPhase get phase => $_getN(2);
  @$pb.TagNumber(3)
  set phase(TaskPhase value) => $_setField(3, value);
  @$pb.TagNumber(3)
  $core.bool hasPhase() => $_has(2);
  @$pb.TagNumber(3)
  void clearPhase() => $_clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get size => $_getI64(3);
  @$pb.TagNumber(4)
  set size($fixnum.Int64 value) => $_setInt64(3, value);
  @$pb.TagNumber(4)
  $core.bool hasSize() => $_has(3);
  @$pb.TagNumber(4)
  void clearSize() => $_clearField(4);

  @$pb.TagNumber(5)
  Category get category => $_getN(4);
  @$pb.TagNumber(5)
  set category(Category value) => $_setField(5, value);
  @$pb.TagNumber(5)
  $core.bool hasCategory() => $_has(4);
  @$pb.TagNumber(5)
  void clearCategory() => $_clearField(5);
  @$pb.TagNumber(5)
  Category ensureCategory() => $_ensure(4);

  @$pb.TagNumber(6)
  $fixnum.Int64 get createdAt => $_getI64(5);
  @$pb.TagNumber(6)
  set createdAt($fixnum.Int64 value) => $_setInt64(5, value);
  @$pb.TagNumber(6)
  $core.bool hasCreatedAt() => $_has(5);
  @$pb.TagNumber(6)
  void clearCreatedAt() => $_clearField(6);

  @$pb.TagNumber(50)
  DownloadStats get downloadStats => $_getN(6);
  @$pb.TagNumber(50)
  set downloadStats(DownloadStats value) => $_setField(50, value);
  @$pb.TagNumber(50)
  $core.bool hasDownloadStats() => $_has(6);
  @$pb.TagNumber(50)
  void clearDownloadStats() => $_clearField(50);
  @$pb.TagNumber(50)
  DownloadStats ensureDownloadStats() => $_ensure(6);

  @$pb.TagNumber(51)
  UploadStats get uploadStats => $_getN(7);
  @$pb.TagNumber(51)
  set uploadStats(UploadStats value) => $_setField(51, value);
  @$pb.TagNumber(51)
  $core.bool hasUploadStats() => $_has(7);
  @$pb.TagNumber(51)
  void clearUploadStats() => $_clearField(51);
  @$pb.TagNumber(51)
  UploadStats ensureUploadStats() => $_ensure(7);
}

/// 类别
class Category extends $pb.GeneratedMessage {
  factory Category({
    $core.String? name,
    $core.String? icon,
  }) {
    final result = create();
    if (name != null) result.name = name;
    if (icon != null) result.icon = icon;
    return result;
  }

  Category._();

  factory Category.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory Category.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'Category',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'task'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'name')
    ..aOS(2, _omitFieldNames ? '' : 'icon')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Category clone() => Category()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  Category copyWith(void Function(Category) updates) =>
      super.copyWith((message) => updates(message as Category)) as Category;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Category create() => Category._();
  @$core.override
  Category createEmptyInstance() => create();
  static $pb.PbList<Category> createRepeated() => [] as $pb.PbList<Category>;
  @$core.pragma('dart2js:noInline')
  static Category getDefault() =>
      _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Category>(create);
  static Category? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => $_clearField(1);

  @$pb.TagNumber(2)
  $core.String get icon => $_getSZ(1);
  @$pb.TagNumber(2)
  set icon($core.String value) => $_setString(1, value);
  @$pb.TagNumber(2)
  $core.bool hasIcon() => $_has(1);
  @$pb.TagNumber(2)
  void clearIcon() => $_clearField(2);
}

/// 下载状态
class DownloadStats extends $pb.GeneratedMessage {
  factory DownloadStats({
    $fixnum.Int64? speed,
    $fixnum.Int64? size,
  }) {
    final result = create();
    if (speed != null) result.speed = speed;
    if (size != null) result.size = size;
    return result;
  }

  DownloadStats._();

  factory DownloadStats.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory DownloadStats.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'DownloadStats',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'task'),
      createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, _omitFieldNames ? '' : 'speed', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(2, _omitFieldNames ? '' : 'size', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DownloadStats clone() => DownloadStats()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  DownloadStats copyWith(void Function(DownloadStats) updates) =>
      super.copyWith((message) => updates(message as DownloadStats))
          as DownloadStats;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DownloadStats create() => DownloadStats._();
  @$core.override
  DownloadStats createEmptyInstance() => create();
  static $pb.PbList<DownloadStats> createRepeated() =>
      [] as $pb.PbList<DownloadStats>;
  @$core.pragma('dart2js:noInline')
  static DownloadStats getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<DownloadStats>(create);
  static DownloadStats? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get speed => $_getI64(0);
  @$pb.TagNumber(1)
  set speed($fixnum.Int64 value) => $_setInt64(0, value);
  @$pb.TagNumber(1)
  $core.bool hasSpeed() => $_has(0);
  @$pb.TagNumber(1)
  void clearSpeed() => $_clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get size => $_getI64(1);
  @$pb.TagNumber(2)
  set size($fixnum.Int64 value) => $_setInt64(1, value);
  @$pb.TagNumber(2)
  $core.bool hasSize() => $_has(1);
  @$pb.TagNumber(2)
  void clearSize() => $_clearField(2);
}

/// 上传状态
class UploadStats extends $pb.GeneratedMessage {
  factory UploadStats({
    $fixnum.Int64? speed,
    $fixnum.Int64? size,
  }) {
    final result = create();
    if (speed != null) result.speed = speed;
    if (size != null) result.size = size;
    return result;
  }

  UploadStats._();

  factory UploadStats.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory UploadStats.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'UploadStats',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'task'),
      createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, _omitFieldNames ? '' : 'speed', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(2, _omitFieldNames ? '' : 'size', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UploadStats clone() => UploadStats()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  UploadStats copyWith(void Function(UploadStats) updates) =>
      super.copyWith((message) => updates(message as UploadStats))
          as UploadStats;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static UploadStats create() => UploadStats._();
  @$core.override
  UploadStats createEmptyInstance() => create();
  static $pb.PbList<UploadStats> createRepeated() => [] as $pb.PbList<UploadStats>;
  @$core.pragma('dart2js:noInline')
  static UploadStats getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<UploadStats>(create);
  static UploadStats? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get speed => $_getI64(0);
  @$pb.TagNumber(1)
  set speed($fixnum.Int64 value) => $_setInt64(0, value);
  @$pb.TagNumber(1)
  $core.bool hasSpeed() => $_has(0);
  @$pb.TagNumber(1)
  void clearSpeed() => $_clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get size => $_getI64(1);
  @$pb.TagNumber(2)
  set size($fixnum.Int64 value) => $_setInt64(1, value);
  @$pb.TagNumber(2)
  $core.bool hasSize() => $_has(1);
  @$pb.TagNumber(2)
  void clearSize() => $_clearField(2);
}

class ListTasksResponse extends $pb.GeneratedMessage {
  factory ListTasksResponse({
    $core.Iterable<Task>? tasks,
  }) {
    final result = create();
    if (tasks != null) result.tasks.addAll(tasks);
    return result;
  }

  ListTasksResponse._();

  factory ListTasksResponse.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory ListTasksResponse.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'ListTasksResponse',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'task'),
      createEmptyInstance: create)
    ..pc<Task>(1, _omitFieldNames ? '' : 'tasks', $pb.PbFieldType.PM,
        subBuilder: Task.create)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListTasksResponse clone() => ListTasksResponse()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  ListTasksResponse copyWith(void Function(ListTasksResponse) updates) =>
      super.copyWith((message) => updates(message as ListTasksResponse))
          as ListTasksResponse;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static ListTasksResponse create() => ListTasksResponse._();
  @$core.override
  ListTasksResponse createEmptyInstance() => create();
  static $pb.PbList<ListTasksResponse> createRepeated() =>
      [] as $pb.PbList<ListTasksResponse>;
  @$core.pragma('dart2js:noInline')
  static ListTasksResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<ListTasksResponse>(create);
  static ListTasksResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<Task> get tasks => $_getList(0);
}

class CreateTaskRequest extends $pb.GeneratedMessage {
  factory CreateTaskRequest({
    $core.String? link,
    $core.String? name,
    $core.String? category,
  }) {
    final result = create();
    if (link != null) result.link = link;
    if (name != null) result.name = name;
    if (category != null) result.category = category;
    return result;
  }

  CreateTaskRequest._();

  factory CreateTaskRequest.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory CreateTaskRequest.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'CreateTaskRequest',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'task'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'link')
    ..aOS(10, _omitFieldNames ? '' : 'name')
    ..aOS(11, _omitFieldNames ? '' : 'category')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskRequest clone() => CreateTaskRequest()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskRequest copyWith(void Function(CreateTaskRequest) updates) =>
      super.copyWith((message) => updates(message as CreateTaskRequest))
          as CreateTaskRequest;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTaskRequest create() => CreateTaskRequest._();
  @$core.override
  CreateTaskRequest createEmptyInstance() => create();
  static $pb.PbList<CreateTaskRequest> createRepeated() =>
      [] as $pb.PbList<CreateTaskRequest>;
  @$core.pragma('dart2js:noInline')
  static CreateTaskRequest getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<CreateTaskRequest>(create);
  static CreateTaskRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get link => $_getSZ(0);
  @$pb.TagNumber(1)
  set link($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasLink() => $_has(0);
  @$pb.TagNumber(1)
  void clearLink() => $_clearField(1);

  @$pb.TagNumber(10)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(10)
  set name($core.String value) => $_setString(1, value);
  @$pb.TagNumber(10)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(10)
  void clearName() => $_clearField(10);

  @$pb.TagNumber(11)
  $core.String get category => $_getSZ(2);
  @$pb.TagNumber(11)
  set category($core.String value) => $_setString(2, value);
  @$pb.TagNumber(11)
  $core.bool hasCategory() => $_has(2);
  @$pb.TagNumber(11)
  void clearCategory() => $_clearField(11);
}

class CreateTaskResponse extends $pb.GeneratedMessage {
  factory CreateTaskResponse({
    $core.String? id,
  }) {
    final result = create();
    if (id != null) result.id = id;
    return result;
  }

  CreateTaskResponse._();

  factory CreateTaskResponse.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory CreateTaskResponse.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'CreateTaskResponse',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'task'),
      createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'id')
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskResponse clone() => CreateTaskResponse()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  CreateTaskResponse copyWith(void Function(CreateTaskResponse) updates) =>
      super.copyWith((message) => updates(message as CreateTaskResponse))
          as CreateTaskResponse;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static CreateTaskResponse create() => CreateTaskResponse._();
  @$core.override
  CreateTaskResponse createEmptyInstance() => create();
  static $pb.PbList<CreateTaskResponse> createRepeated() =>
      [] as $pb.PbList<CreateTaskResponse>;
  @$core.pragma('dart2js:noInline')
  static CreateTaskResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<CreateTaskResponse>(create);
  static CreateTaskResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String value) => $_setString(0, value);
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => $_clearField(1);
}

class OperateTasksRequest extends $pb.GeneratedMessage {
  factory OperateTasksRequest({
    $core.Iterable<$core.String>? ids,
    Operate? operate,
  }) {
    final result = create();
    if (ids != null) result.ids.addAll(ids);
    if (operate != null) result.operate = operate;
    return result;
  }

  OperateTasksRequest._();

  factory OperateTasksRequest.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory OperateTasksRequest.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'OperateTasksRequest',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'task'),
      createEmptyInstance: create)
    ..pPS(1, _omitFieldNames ? '' : 'ids')
    ..e<Operate>(2, _omitFieldNames ? '' : 'operate', $pb.PbFieldType.OE,
        defaultOrMaker: Operate.OpUnknown,
        valueOf: Operate.valueOf,
        enumValues: Operate.values)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  OperateTasksRequest clone() => OperateTasksRequest()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  OperateTasksRequest copyWith(void Function(OperateTasksRequest) updates) =>
      super.copyWith((message) => updates(message as OperateTasksRequest))
          as OperateTasksRequest;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static OperateTasksRequest create() => OperateTasksRequest._();
  @$core.override
  OperateTasksRequest createEmptyInstance() => create();
  static $pb.PbList<OperateTasksRequest> createRepeated() =>
      [] as $pb.PbList<OperateTasksRequest>;
  @$core.pragma('dart2js:noInline')
  static OperateTasksRequest getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<OperateTasksRequest>(create);
  static OperateTasksRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $pb.PbList<$core.String> get ids => $_getList(0);

  @$pb.TagNumber(2)
  Operate get operate => $_getN(1);
  @$pb.TagNumber(2)
  set operate(Operate value) => $_setField(2, value);
  @$pb.TagNumber(2)
  $core.bool hasOperate() => $_has(1);
  @$pb.TagNumber(2)
  void clearOperate() => $_clearField(2);
}

class StatInfoResponse extends $pb.GeneratedMessage {
  factory StatInfoResponse({
    $fixnum.Int64? taskCount,
    $fixnum.Int64? dlCount,
    $fixnum.Int64? dlSpeed,
    $fixnum.Int64? ulCount,
    $fixnum.Int64? ulSpeed,
  }) {
    final result = create();
    if (taskCount != null) result.taskCount = taskCount;
    if (dlCount != null) result.dlCount = dlCount;
    if (dlSpeed != null) result.dlSpeed = dlSpeed;
    if (ulCount != null) result.ulCount = ulCount;
    if (ulSpeed != null) result.ulSpeed = ulSpeed;
    return result;
  }

  StatInfoResponse._();

  factory StatInfoResponse.fromBuffer($core.List<$core.int> data,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromBuffer(data, registry);
  factory StatInfoResponse.fromJson($core.String json,
          [$pb.ExtensionRegistry registry = $pb.ExtensionRegistry.EMPTY]) =>
      create()..mergeFromJson(json, registry);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(
      _omitMessageNames ? '' : 'StatInfoResponse',
      package: const $pb.PackageName(_omitMessageNames ? '' : 'task'),
      createEmptyInstance: create)
    ..a<$fixnum.Int64>(
        1, _omitFieldNames ? '' : 'taskCount', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(2, _omitFieldNames ? '' : 'dlCount', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(3, _omitFieldNames ? '' : 'dlSpeed', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(4, _omitFieldNames ? '' : 'ulCount', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..a<$fixnum.Int64>(5, _omitFieldNames ? '' : 'ulSpeed', $pb.PbFieldType.OU6,
        defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false;

  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  StatInfoResponse clone() => StatInfoResponse()..mergeFromMessage(this);
  @$core.Deprecated('See https://github.com/google/protobuf.dart/issues/998.')
  StatInfoResponse copyWith(void Function(StatInfoResponse) updates) =>
      super.copyWith((message) => updates(message as StatInfoResponse))
          as StatInfoResponse;

  @$core.override
  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static StatInfoResponse create() => StatInfoResponse._();
  @$core.override
  StatInfoResponse createEmptyInstance() => create();
  static $pb.PbList<StatInfoResponse> createRepeated() =>
      [] as $pb.PbList<StatInfoResponse>;
  @$core.pragma('dart2js:noInline')
  static StatInfoResponse getDefault() => _defaultInstance ??=
      $pb.GeneratedMessage.$_defaultFor<StatInfoResponse>(create);
  static StatInfoResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get taskCount => $_getI64(0);
  @$pb.TagNumber(1)
  set taskCount($fixnum.Int64 value) => $_setInt64(0, value);
  @$pb.TagNumber(1)
  $core.bool hasTaskCount() => $_has(0);
  @$pb.TagNumber(1)
  void clearTaskCount() => $_clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get dlCount => $_getI64(1);
  @$pb.TagNumber(2)
  set dlCount($fixnum.Int64 value) => $_setInt64(1, value);
  @$pb.TagNumber(2)
  $core.bool hasDlCount() => $_has(1);
  @$pb.TagNumber(2)
  void clearDlCount() => $_clearField(2);

  @$pb.TagNumber(3)
  $fixnum.Int64 get dlSpeed => $_getI64(2);
  @$pb.TagNumber(3)
  set dlSpeed($fixnum.Int64 value) => $_setInt64(2, value);
  @$pb.TagNumber(3)
  $core.bool hasDlSpeed() => $_has(2);
  @$pb.TagNumber(3)
  void clearDlSpeed() => $_clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get ulCount => $_getI64(3);
  @$pb.TagNumber(4)
  set ulCount($fixnum.Int64 value) => $_setInt64(3, value);
  @$pb.TagNumber(4)
  $core.bool hasUlCount() => $_has(3);
  @$pb.TagNumber(4)
  void clearUlCount() => $_clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get ulSpeed => $_getI64(4);
  @$pb.TagNumber(5)
  set ulSpeed($fixnum.Int64 value) => $_setInt64(4, value);
  @$pb.TagNumber(5)
  $core.bool hasUlSpeed() => $_has(4);
  @$pb.TagNumber(5)
  void clearUlSpeed() => $_clearField(5);
}

const $core.bool _omitFieldNames =
    $core.bool.fromEnvironment('protobuf.omit_field_names');
const $core.bool _omitMessageNames =
    $core.bool.fromEnvironment('protobuf.omit_message_names');
