import 'package:flutter/material.dart';
import 'package:mdm/configs/theme.dart';
import 'package:mdm/providers/task_provider.dart';
import 'package:mdm/utils/color_helper.dart';
import 'package:mdm/utils/format_helper.dart';
import 'package:provider/provider.dart';

class OverviewPanel extends StatelessWidget {
  const OverviewPanel({super.key});

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      padding: const EdgeInsets.all(20),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          _buildSpeedCard(context),
          const SizedBox(height: 20),
          _buildFilterSection(context),
        ],
      ),
    );
  }

  Widget _buildSpeedCard(BuildContext context) {
    return Consumer<TaskProvider>(
      builder: (context, provider, _) {
        final stats = provider.stats;
        return Container(
          padding: const EdgeInsets.all(20),
          decoration: BoxDecoration(
            gradient: getPrimaryGradient(),
            borderRadius: BorderRadius.circular(20),
            boxShadow: [
              BoxShadow(
                color: kPrimary.withValues(alpha: 0.3),
                blurRadius: 20,
                offset: const Offset(0, 10),
              ),
            ],
          ),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  const Text('下载速度', style: TextStyle(color: Colors.white70, fontSize: 14)),
                  Container(
                    padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                    decoration: BoxDecoration(
                      color: Colors.white.withValues(alpha: 0.2),
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: Row(
                      children: [
                        Container(
                          width: 8,
                          height: 8,
                          decoration: BoxDecoration(
                            color: stats.downloading > 0 ? kSuccess : Colors.white54,
                            shape: BoxShape.circle,
                          ),
                        ),
                        const SizedBox(width: 6),
                        Text(
                          stats.downloading > 0 ? '已连接' : '已断联',
                          style: const TextStyle(color: Colors.white, fontSize: 12),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 16),
              Row(
                crossAxisAlignment: CrossAxisAlignment.end,
                children: [
                  Text(
                    FormatHelper.formatSpeed(stats.totalSpeed.toDouble()),
                    style: const TextStyle(color: Colors.white, fontSize: 36, fontWeight: FontWeight.bold),
                  ),
                  const Padding(
                    padding: EdgeInsets.only(bottom: 6, left: 4),
                    child: Text('/s', style: TextStyle(color: Colors.white70, fontSize: 16)),
                  ),
                ],
              ),
              const SizedBox(height: 16),
              LinearProgressIndicator(
                minHeight: 6,
                borderRadius: BorderRadius.circular(8),
                value: stats.totalSize > 0 ? stats.totalDownloaded / stats.totalSize : 0.0,
                backgroundColor: Colors.white.withValues(alpha: 0.2),
                valueColor: const AlwaysStoppedAnimation(Colors.white),
              ),
              const SizedBox(height: 6),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Text(FormatHelper.formatSize(stats.totalDownloaded),
                      style: TextStyle(color: Colors.white60, fontSize: 12)),
                  Text(FormatHelper.formatSize(stats.totalSize),
                      style: TextStyle(color: Colors.white60, fontSize: 12)),
                ],
              ),
            ],
          ),
        );
      },
    );
  }

  Widget _buildFilterSection(BuildContext context) {
    return Consumer<TaskProvider>(
      builder: (context, provider, _) {
        return Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            ExpansionTile(
              initiallyExpanded: true,
              leading: const Icon(Icons.download_rounded),
              title: const Text('下载任务'),
              children: FilterType.values.getRange(0, 4)
                  .map((filter) => _buildFilterItem(context, filter, provider))
                  .toList(),
            ),
            const SizedBox(height: 12),
            ExpansionTile(
              initiallyExpanded: true,
              leading: const Icon(Icons.upload_rounded),
              title: const Text('上传任务'),
              children: FilterType.values.skip(4)
                  .map((filter) => _buildFilterItem(context, filter, provider))
                  .toList(),
            ),
          ],
        );
      },
    );
  }

  static const _filterIcons = {
    FilterType.dlRunning: Icons.download_for_offline_rounded,
    FilterType.dlCompleted: Icons.check_circle_rounded,
    FilterType.dlPaused: Icons.pause_circle_rounded,
    FilterType.dlFailed: Icons.pause_circle_rounded,
    FilterType.ulRunning: Icons.upload,
    FilterType.ulCompleted: Icons.check_circle_rounded,
    FilterType.ulPaused: Icons.pause_circle_rounded,
    FilterType.ulFailed: Icons.pause_circle_rounded,
  };

  static const _filterLabels = {
    FilterType.dlRunning: '下载中',
    FilterType.dlCompleted: '已完成',
    FilterType.dlPaused: '暂停中',
    FilterType.dlFailed: '已失败',
    FilterType.ulRunning: '上传中',
    FilterType.ulCompleted: '已完成',
    FilterType.ulPaused: '暂停中',
    FilterType.ulFailed: '已失败',
  };

  static Color _filterColor(FilterType filter) {
    return switch (filter) {
      FilterType.dlFailed || FilterType.ulFailed => kError,
      FilterType.dlPaused || FilterType.ulPaused => kWarning,
      FilterType.dlCompleted || FilterType.ulCompleted => kSuccess,
      _ => kInfo,
    };
  }

  Widget _buildFilterItem(BuildContext context, FilterType filter, TaskProvider provider) {
    final isSelected = provider.currentFilter == filter;
    final color = _filterColor(filter);
    final icon = Icon(_filterIcons[filter], color: color, size: 20);

    return Padding(
      padding: const EdgeInsets.only(bottom: 8),
      child: Material(
        color: Colors.transparent,
        child: InkWell(
          onTap: () => provider.setFilter(filter),
          borderRadius: BorderRadius.circular(12),
          child: AnimatedContainer(
            duration: const Duration(milliseconds: 200),
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
            decoration: BoxDecoration(
              color: isSelected ? color.withValues(alpha: 0.15) : Colors.transparent,
              borderRadius: BorderRadius.circular(12),
              border: Border.all(
                color: isSelected ? color.withValues(alpha: 0.3) : Colors.transparent,
              ),
            ),
            child: Row(
              children: [
                icon,
                const SizedBox(width: 12),
                Text(
                  _filterLabels[filter]!,
                  style: TextStyle(
                    color: isSelected ? color : Colors.black54,
                    fontWeight: isSelected ? FontWeight.w600 : FontWeight.normal,
                  ),
                ),
                const Spacer(),
                if (isSelected) Icon(Icons.chevron_right_rounded, color: color, size: 20),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
