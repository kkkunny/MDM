import 'dart:async';
import 'package:flutter/foundation.dart';
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

const _refreshInterval = Duration(seconds: 5);

class TaskProvider extends ChangeNotifier {
  List<Task> _tasks = [];
  bool _isLoading = false;
  String? _error;
  FilterType _currentFilter = FilterType.dlRunning;
  String _searchQuery = '';
  Set<String> _selectedTaskIds = {};
  StreamSubscription? _pollSubscription;

  List<Task> get tasks => _filteredTasks;
  bool get isLoading => _isLoading;
  String? get error => _error;
  FilterType get currentFilter => _currentFilter;
  String get searchQuery => _searchQuery;
  Set<String> get selectedTaskIds => _selectedTaskIds;
  bool get hasSelection => _selectedTaskIds.isNotEmpty;

  DownloadStats get stats => DownloadStats.fromTasks(_tasks);

  List<Task> get _filteredTasks {
    var filtered = _tasks.where((t) => _currentFilter.phases.contains(t.phase)).toList();
    if (_searchQuery.isNotEmpty) {
      filtered = filtered
          .where((t) => t.name.toLowerCase().contains(_searchQuery.toLowerCase()))
          .toList();
    }
    return filtered;
  }

  Future<void> initialize() async {
    await fetchTasks();
    _startPolling();
  }

  Future<void> fetchTasks() async {
    _isLoading = true;
    _error = null;
    notifyListeners();

    try {
      _tasks = (await listTasks()).tasks;
    } catch (e) {
      _error = e.toString();
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  void _startPolling() {
    _pollSubscription?.cancel();
    _pollSubscription = Stream.periodic(_refreshInterval, (_) => listTasks()).listen(
      (futureResp) async {
        _tasks = (await futureResp).tasks;
        notifyListeners();
      },
      onError: (e) {
        _error = e.toString();
        notifyListeners();
      },
    );
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
    _pollSubscription?.cancel();
    super.dispose();
  }
}
