import 'package:mdm/models/vo/task.pb.dart';

extension TaskPhaseExtension on TaskPhase {
  String get label => switch (this) {
    TaskPhase.TpDownQueued => '等待下载',
    TaskPhase.TpDownRunning => '下载中',
    TaskPhase.TpDownPaused => '暂停下载',
    TaskPhase.TpDownFailed => '下载失败',
    TaskPhase.TpDownCompleted => '下载完成',
    TaskPhase.TpUpQueued => '等待上传',
    TaskPhase.TpUpRunning => '上传中',
    TaskPhase.TpUpPaused => '暂停上传',
    TaskPhase.TpUpFailed => '上传失败',
    TaskPhase.TpUpCompleted => '上传完成',
    _ => 'Unknown',
  };

  bool get isDownload => name.startsWith('TpDown');
}

class DownloadStats {
  final int totalTasks;
  final int downloading;
  final int completed;
  final int paused;
  final int failed;
  final int totalSpeed;
  final int totalDownloaded;
  final int totalSize;

  DownloadStats({
    required this.totalTasks,
    required this.downloading,
    required this.completed,
    required this.paused,
    required this.failed,
    required this.totalSpeed,
    required this.totalDownloaded,
    required this.totalSize,
  });

  factory DownloadStats.fromTasks(List<Task> tasks) => DownloadStats(
    totalTasks: tasks.length,
    downloading: tasks.where((t) => t.phase == TaskPhase.TpDownRunning).length,
    completed: tasks.where((t) => t.phase == TaskPhase.TpDownCompleted).length,
    paused: tasks.where((t) => t.phase == TaskPhase.TpDownQueued).length,
    failed: tasks.where((t) => t.phase == TaskPhase.TpDownFailed).length,
    totalSpeed: tasks
        .where((t) => t.phase == TaskPhase.TpDownRunning)
        .fold(0, (sum, t) => sum + t.downloadStats.speed.toInt()),
    totalDownloaded: tasks.fold(0, (sum, t) => sum + t.downloadStats.size.toInt()),
    totalSize: tasks.fold(0, (sum, t) => sum + t.size.toInt()),
  );
}
