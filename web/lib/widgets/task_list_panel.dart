import 'package:flutter/material.dart';
import 'package:mdm/apis/mdm/task.dart';
import 'package:mdm/configs/theme.dart';
import 'package:mdm/models/vo/task.pb.dart';
import 'package:mdm/providers/task_provider.dart';
import 'package:mdm/widgets/add_task_dialog.dart';
import 'package:mdm/widgets/common/app_button.dart';
import 'package:mdm/widgets/common/app_dialog.dart';
import 'package:mdm/widgets/task_card.dart';
import 'package:provider/provider.dart';

class TaskListPanel extends StatelessWidget {
  const TaskListPanel({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: kLightBackground,
      child: Column(
        children: [
          _buildToolbar(context),
          Expanded(child: _buildTaskList(context)),
        ],
      ),
    );
  }

  Widget _buildToolbar(BuildContext context) {
    return Consumer<TaskProvider>(
      builder: (context, provider, _) {
        return Container(
          padding: const EdgeInsets.all(20),
          decoration: const BoxDecoration(
            color: kLightSurface,
            border: Border(bottom: BorderSide(color: kLightDivider)),
          ),
          child: Column(
            children: [
              Row(
                children: [
                  Expanded(child: _buildSearchBar(context, provider)),
                  const SizedBox(width: 16),
                  _buildAddButton(context),
                ],
              ),
              if (provider.hasSelection) ...[
                const SizedBox(height: 16),
                _buildSelectionToolbar(context, provider),
              ],
            ],
          ),
        );
      },
    );
  }

  Widget _buildSearchBar(BuildContext context, TaskProvider provider) {
    return Container(
      height: 48,
      decoration: BoxDecoration(
        color: kLightSurfaceLight,
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: kLightDivider),
      ),
      child: TextField(
        onChanged: provider.setSearchQuery,
        style: const TextStyle(color: kLightText),
        decoration: InputDecoration(
          hintText: '搜索任务...',
          hintStyle: TextStyle(color: kLightTextSecondary),
          prefixIcon: Icon(Icons.search_rounded, color: kLightTextSecondary),
          border: InputBorder.none,
          contentPadding: const EdgeInsets.symmetric(horizontal: 16, vertical: 14),
        ),
      ),
    );
  }

  Widget _buildAddButton(BuildContext context) {
    return AppButton.primary(
      text: '添加任务',
      icon: Icons.add_rounded,
      onPressed: () => showDialog(
        context: context,
        builder: (context) => AddTaskDialog(
          action: (AddTaskData data) async {
            await createTask(CreateTaskRequest(link: data.link, name: data.taskName));
          },
        ),
      ),
    );
  }

  Widget _buildSelectionToolbar(BuildContext context, TaskProvider provider) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
      decoration: BoxDecoration(
        color: kPrimary.withValues(alpha: 0.1),
        borderRadius: BorderRadius.circular(12),
        border: Border.all(color: kPrimary.withValues(alpha: 0.3)),
      ),
      child: Row(
        children: [
          Text(
            '${provider.selectedTaskIds.length} 已选择',
            style: const TextStyle(color: kPrimary, fontWeight: FontWeight.w600),
          ),
          const Spacer(),
          _buildToolbarButton(
            icon: Icons.pause_rounded,
            label: '全部暂停',
            onTap: () {},
          ),
          const SizedBox(width: 12),
          _buildToolbarButton(
            icon: Icons.delete_rounded,
            label: '删除',
            color: kError,
            onTap: () => _showDeleteConfirmDialog(context, provider),
          ),
          const SizedBox(width: 12),
          _buildToolbarButton(
            icon: Icons.close_rounded,
            label: '取消',
            onTap: () => provider.clearSelection(),
          ),
        ],
      ),
    );
  }

  Widget _buildToolbarButton({
    required IconData icon,
    required String label,
    required VoidCallback onTap,
    Color color = Colors.white70,
  }) {
    return Material(
      color: Colors.transparent,
      child: InkWell(
        onTap: onTap,
        borderRadius: BorderRadius.circular(8),
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
          child: Row(
            children: [
              Icon(icon, color: color, size: 18),
              const SizedBox(width: 6),
              Text(label, style: TextStyle(color: color, fontSize: 13)),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildTaskList(BuildContext context) {
    return Consumer<TaskProvider>(
      builder: (context, provider, _) {
        if (provider.isLoading) {
          return const Center(child: CircularProgressIndicator(color: kPrimary));
        }

        if (provider.error != null) {
          return Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Icon(Icons.error_outline_rounded, size: 64, color: Colors.black12),
                const SizedBox(height: 16),
                Text('加载任务失败', style: TextStyle(color: Colors.black38, fontSize: 16)),
                const SizedBox(height: 8),
                TextButton(
                  onPressed: () => provider.fetchTasks(),
                  child: const Text('重试'),
                ),
              ],
            ),
          );
        }

        if (provider.tasks.isEmpty) {
          return Center(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Icon(Icons.download_done_rounded, size: 80, color: Colors.black12),
                const SizedBox(height: 20),
                Text('暂无任务', style: TextStyle(color: Colors.black38, fontSize: 18, fontWeight: FontWeight.w500)),
                const SizedBox(height: 8),
                Text('点击"添加任务"开始下载', style: TextStyle(color: Colors.black26, fontSize: 14)),
              ],
            ),
          );
        }

        return ListView.builder(
          padding: const EdgeInsets.all(20),
          itemCount: provider.tasks.length,
          itemBuilder: (context, index) {
            final task = provider.tasks[index];
            return Padding(
              padding: const EdgeInsets.only(bottom: 12),
              child: TaskCard(
                task: task,
                isSelected: provider.selectedTaskIds.contains(task.id),
                onTap: () => provider.toggleSelection(task.id),
              ),
            );
          },
        );
      },
    );
  }

  void _showDeleteConfirmDialog(BuildContext context, TaskProvider provider) {
    AppDialog.showConfirmDialog(
      context,
      title: '删除下载？',
      content: '确定要删除 ${provider.selectedTaskIds.length} 个选中的下载吗？',
      confirmText: '删除',
      cancelText: '取消',
      confirmColor: kError,
    ).then((confirmed) {
      if (confirmed == true) provider.deleteSelected();
    });
  }
}
