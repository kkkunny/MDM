import 'package:flutter/material.dart';
import 'package:mdm/apis/mdm/task.dart';
import 'package:mdm/configs/theme.dart';
import 'package:mdm/models/task.dart';
import 'package:mdm/models/vo/task.pb.dart';
import 'package:mdm/utils/color_helper.dart';
import 'package:mdm/utils/format_helper.dart';
import 'package:mdm/utils/icon_helper.dart';

class TaskCard extends StatefulWidget {
  final Task task;
  final bool isSelected;
  final VoidCallback onTap;

  const TaskCard({super.key, required this.task, required this.isSelected, required this.onTap});

  @override
  State<TaskCard> createState() => _TaskCardState();
}

class _TaskCardState extends State<TaskCard> with SingleTickerProviderStateMixin {
  bool _isHovered = false;
  late final AnimationController _controller = AnimationController(
    duration: const Duration(milliseconds: 150),
    vsync: this,
  );
  late final Animation<double> _scaleAnimation = Tween<double>(begin: 1.0, end: 1.02).animate(
    CurvedAnimation(parent: _controller, curve: Curves.easeOut),
  );

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  void _operate(Operate op) {
    operateTasks(OperateTasksRequest(ids: [widget.task.id], operate: op));
  }

  @override
  Widget build(BuildContext context) {
    final statusColor = getStatusColor(widget.task.phase);

    return MouseRegion(
      onEnter: (_) {
        setState(() => _isHovered = true);
        _controller.forward();
      },
      onExit: (_) {
        setState(() => _isHovered = false);
        _controller.reverse();
      },
      child: AnimatedBuilder(
        animation: _scaleAnimation,
        builder: (context, child) {
          return Transform.scale(
            scale: _scaleAnimation.value,
            child: GestureDetector(
              onTap: widget.onTap,
              child: AnimatedContainer(
                duration: const Duration(milliseconds: 200),
                padding: const EdgeInsets.all(20),
                decoration: BoxDecoration(
                  gradient: widget.isSelected
                      ? LinearGradient(colors: [
                          kPrimary.withValues(alpha: 0.2),
                          kPrimaryDark.withValues(alpha: 0.1),
                        ])
                      : null,
                  color: widget.isSelected ? null : kLightSurface,
                  borderRadius: BorderRadius.circular(16),
                  border: Border.all(
                    color: widget.isSelected
                        ? kPrimary.withValues(alpha: 0.5)
                        : _isHovered
                            ? kLightDivider.withValues(alpha: 0.8)
                            : kLightDivider,
                    width: widget.isSelected ? 2 : 1,
                  ),
                  boxShadow: _isHovered
                      ? [BoxShadow(
                          color: kPrimary.withValues(alpha: 0.1),
                          blurRadius: 20,
                          offset: const Offset(0, 8),
                        )]
                      : null,
                ),
                child: Column(
                  children: [
                    _buildHeader(statusColor),
                    const SizedBox(height: 16),
                    _buildProgressSection(statusColor),
                    const SizedBox(height: 16),
                    _buildFooter(statusColor),
                  ],
                ),
              ),
            ),
          );
        },
      ),
    );
  }

  Widget _buildHeader(Color statusColor) {
    return Row(
      children: [
        _buildFileIcon(statusColor),
        const SizedBox(width: 16),
        Expanded(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                widget.task.name,
                style: const TextStyle(color: kLightText, fontSize: 16, fontWeight: FontWeight.w600),
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
              ),
              const SizedBox(height: 4),
              Row(
                children: [
                  _buildStatusBadge(statusColor),
                  const SizedBox(width: 8),
                  if (widget.task.category.name.isNotEmpty) _buildCategoryBadge(),
                ],
              ),
            ],
          ),
        ),
        _buildActionButtons(statusColor),
      ],
    );
  }

  Widget _buildFileIcon(Color statusColor) {
    return Container(
      width: 48,
      height: 48,
      decoration: BoxDecoration(
        color: statusColor.withValues(alpha: 0.1),
        borderRadius: BorderRadius.circular(12),
      ),
      child: Icon(IconHelper.getFileIcon(widget.task.name), color: statusColor, size: 24),
    );
  }

  Widget _buildStatusBadge(Color statusColor) {
    final isRunning = widget.task.phase == TaskPhase.TpDownRunning;
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 4),
      decoration: BoxDecoration(
        color: getStatusBackgroundColor(widget.task.phase),
        borderRadius: BorderRadius.circular(20),
        border: Border.all(color: getStatusBorderColor(widget.task.phase)),
      ),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          if (isRunning)
            SizedBox(
              width: 12,
              height: 12,
              child: CircularProgressIndicator(
                strokeWidth: 2,
                valueColor: AlwaysStoppedAnimation(statusColor),
              ),
            )
          else
            Container(
              width: 8,
              height: 8,
              decoration: BoxDecoration(color: statusColor, shape: BoxShape.circle),
            ),
          const SizedBox(width: 6),
          Text(
            widget.task.phase.label,
            style: TextStyle(color: statusColor, fontSize: 12, fontWeight: FontWeight.w500),
          ),
        ],
      ),
    );
  }

  Widget _buildCategoryBadge() {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 4),
      decoration: BoxDecoration(
        color: kLightSurfaceLight,
        borderRadius: BorderRadius.circular(20),
      ),
      child: Text(
        widget.task.category.name,
        style: const TextStyle(color: kLightTextSecondary, fontSize: 12),
      ),
    );
  }

  Widget _buildActionButtons(Color statusColor) {
    final phase = widget.task.phase;
    return Row(
      children: [
        if (phase == TaskPhase.TpDownQueued || phase == TaskPhase.TpDownRunning)
          _buildIconButton(icon: Icons.pause_rounded, color: kWarning, onPressed: () => _operate(Operate.OpPause)),
        if (phase == TaskPhase.TpDownPaused)
          _buildIconButton(icon: Icons.play_arrow_rounded, color: kInfo, onPressed: () => _operate(Operate.OpResume)),
        if (phase == TaskPhase.TpDownFailed)
          _buildIconButton(icon: Icons.refresh_rounded, color: kPrimary, onPressed: () => _operate(Operate.OpRetry)),
        const SizedBox(width: 8),
        _buildIconButton(icon: Icons.delete_outline_rounded, color: kError, onPressed: () => _operate(Operate.OpDelete)),
      ],
    );
  }

  Widget _buildIconButton({
    required IconData icon,
    required Color color,
    required VoidCallback onPressed,
  }) {
    return Tooltip(
      message: icon == Icons.pause_rounded ? '暂停'
          : icon == Icons.play_arrow_rounded ? '继续'
          : icon == Icons.refresh_rounded ? '重试'
          : '删除',
      child: Material(
        color: Colors.transparent,
        child: InkWell(
          onTap: onPressed,
          borderRadius: BorderRadius.circular(10),
          child: Container(
            width: 40,
            height: 40,
            decoration: BoxDecoration(
              color: color.withValues(alpha: 0.1),
              borderRadius: BorderRadius.circular(10),
            ),
            child: Icon(icon, color: color, size: 20),
          ),
        ),
      ),
    );
  }

  Widget _buildProgressSection(Color statusColor) {
    final downloaded = widget.task.downloadStats.size;
    final total = widget.task.size;
    final progressValue = total > 0 ? downloaded.toDouble() / total.toDouble() : 0.0;

    return Column(
      children: [
        LinearProgressIndicator(
          minHeight: 8,
          borderRadius: BorderRadius.circular(6),
          value: progressValue,
          backgroundColor: kLightDivider,
          valueColor: AlwaysStoppedAnimation(
            widget.task.phase == TaskPhase.TpDownRunning ? kInfo : statusColor,
          ),
        ),
        const SizedBox(height: 12),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text(
              FormatHelper.formatProgress(downloaded.toDouble(), total.toDouble()),
              style: TextStyle(color: statusColor, fontSize: 14, fontWeight: FontWeight.w600),
            ),
            Text(
              '${FormatHelper.formatSize(downloaded.toInt())} / ${FormatHelper.formatSize(total.toInt())}',
              style: const TextStyle(color: kLightTextSecondary, fontSize: 13),
            ),
          ],
        ),
      ],
    );
  }

  Widget _buildFooter(Color statusColor) {
    return Row(
      children: [
        if (widget.task.phase == TaskPhase.TpDownRunning) ...[
          Icon(Icons.speed_rounded, color: kLightTextSecondary, size: 16),
          const SizedBox(width: 6),
          Text(
            '${FormatHelper.formatSpeed(widget.task.downloadStats.speed.toDouble())}/s',
            style: const TextStyle(color: kLightTextSecondary, fontSize: 13),
          ),
          const SizedBox(width: 16),
          Icon(Icons.timer_outlined, color: kLightTextSecondary, size: 16),
          const SizedBox(width: 6),
          Text(
            FormatHelper.formatDuration(
              widget.task.downloadStats.speed == 0
                  ? Duration.zero
                  : Duration(seconds: ((widget.task.size.toDouble() - widget.task.downloadStats.size.toDouble()) / widget.task.downloadStats.speed.toDouble()).toInt()),
            ),
            style: const TextStyle(color: kLightTextSecondary, fontSize: 13),
          ),
        ],
        const Spacer(),
        Icon(Icons.access_time_rounded, color: kLightTextSecondary, size: 14),
        const SizedBox(width: 4),
        Text(
          FormatHelper.formatDateTime(
            DateTime.fromMillisecondsSinceEpoch(widget.task.createdAt.toInt() * 1000),
          ),
          style: const TextStyle(color: kLightTextSecondary, fontSize: 12),
        ),
      ],
    );
  }
}
