import 'dart:async';
import 'package:flutter/foundation.dart';
import 'package:mdm/apis/mdm/events.dart';
import 'package:mdm/apis/mdm/task.dart';
import 'package:mdm/models/vo/task.pb.dart' hide DownloadStats;
import 'package:mdm/models/task.dart';

enum FilterType {
  dlRunning,
  dlCompleted,
  dlPaused,
  dlFailed,
  ulRunning,
  ulCompleted,
  ulPaused,
  ulFailed,
}

extension FilterTypeExtension on FilterType {
  List<TaskPhase> get phases => switch (this) {
    FilterType.dlRunning => [TaskPhase.TpDownRunning, TaskPhase.TpDownQueued],
    FilterType.dlCompleted => [TaskPhase.TpDownCompleted],
    FilterType.dlPaused => [TaskPhase.TpDownPaused],
    FilterType.dlFailed => [TaskPhase.TpDownFailed],
    FilterType.ulRunning => [TaskPhase.TpUpRunning, TaskPhase.TpUpQueued],
    FilterType.ulCompleted => [TaskPhase.TpUpCompleted],
    FilterType.ulPaused => [TaskPhase.TpUpPaused],
    FilterType.ulFailed => [TaskPhase.TpUpFailed],
  };
}

const _reconnectDelay = Duration(seconds: 3);

class TaskProvider extends ChangeNotifier {
  List<Task> _tasks = [];
  bool _isLoading = true;
  String? _error;
  FilterType _currentFilter = FilterType.dlRunning;
  String _searchQuery = '';
  Set<String> _selectedTaskIds = {};
  StreamSubscription? _eventSubscription;
  Timer? _reconnectTimer;
  bool _disposed = false;

  List<Task> get tasks => _filteredTasks;
  bool get isLoading => _isLoading;
  String? get error => _error;
  FilterType get currentFilter => _currentFilter;
  String get searchQuery => _searchQuery;
  Set<String> get selectedTaskIds => _selectedTaskIds;
  bool get hasSelection => _selectedTaskIds.isNotEmpty;
  bool get isConnected => !_isLoading && _error == null;

  DownloadStats get stats => DownloadStats.fromTasks(_tasks);

  int getTaskCount(FilterType filter) {
    return _tasks.where((t) => filter.phases.contains(t.phase)).length;
  }

  List<Task> get _filteredTasks {
    var filtered = _tasks.where((t) => _currentFilter.phases.contains(t.phase)).toList();
    if (_searchQuery.isNotEmpty) {
      filtered = filtered
          .where((t) => t.name.toLowerCase().contains(_searchQuery.toLowerCase()))
          .toList();
    }
    return filtered;
  }

  void initialize() {
    _startEvents();
  }

  /// 手动重连（错误重试按钮）
  void retry() {
    if (_disposed) return;
    _reconnectTimer?.cancel();
    _isLoading = true;
    _error = null;
    notifyListeners();
    _startEvents();
  }

  void _startEvents() {
    _reconnectTimer?.cancel();
    _eventSubscription?.cancel();
    _eventSubscription = subscribeTaskEvents().listen(
      (event) {
        _applyEvent(event);
        notifyListeners();
      },
      onError: (e) {
        _error = e.toString();
        _isLoading = false;
        notifyListeners();
        _scheduleReconnect();
      },
      onDone: () {
        _scheduleReconnect();
      },
      cancelOnError: true,
    );
  }

  void _applyEvent(TaskEvent event) {
    switch (event.type) {
      case TaskEventType.TetFull:
        _tasks = event.tasks;
      case TaskEventType.TetUpsert:
        final byId = {for (final t in _tasks) t.id: t};
        for (final t in event.tasks) {
          byId[t.id] = t;
        }
        for (final id in event.removedIds) {
          byId.remove(id);
        }
        _tasks = byId.values.toList();
      case TaskEventType.TetUnknown:
        return;
    }
    _tasks.sort((a, b) => b.createdAt.compareTo(a.createdAt));
    _isLoading = false;
    _error = null;
  }

  void _scheduleReconnect() {
    if (_disposed) return;
    _eventSubscription?.cancel();
    _eventSubscription = null;
    _reconnectTimer?.cancel();
    _reconnectTimer = Timer(_reconnectDelay, () {
      if (_disposed) return;
      _startEvents();
    });
  }

  Future<void> deleteSelected({bool deleteFile = false}) async {
    await operateTasks(OperateTasksRequest(
      ids: _selectedTaskIds.toList(),
      operate: Operate.OpDelete,
    ));
    _tasks.removeWhere((t) => _selectedTaskIds.contains(t.id));
    _selectedTaskIds.clear();
    notifyListeners();
  }

  void setFilter(FilterType filter) {
    _currentFilter = filter;
    notifyListeners();
  }

  void setSearchQuery(String query) {
    _searchQuery = query;
    notifyListeners();
  }

  void toggleSelection(String taskId) {
    if (_selectedTaskIds.contains(taskId)) {
      _selectedTaskIds.remove(taskId);
    } else {
      _selectedTaskIds.add(taskId);
    }
    notifyListeners();
  }

  void toggleSelectAll() {
    if (_selectedTaskIds.length == _filteredTasks.length) {
      _selectedTaskIds.clear();
    } else {
      _selectedTaskIds = _filteredTasks.map((t) => t.id).toSet();
    }
    notifyListeners();
  }

  void clearSelection() {
    _selectedTaskIds.clear();
    notifyListeners();
  }

  @override
  void dispose() {
    _disposed = true;
    _reconnectTimer?.cancel();
    _eventSubscription?.cancel();
    super.dispose();
  }
}
