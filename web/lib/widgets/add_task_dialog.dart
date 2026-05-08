import 'package:flutter/material.dart';
import 'package:mdm/configs/theme.dart';

class AddTaskData {
  final String link;
  final String taskName;

  AddTaskData(this.link, {this.taskName = ''});
}

class AddTaskDialog extends StatelessWidget {
  final void Function(AddTaskData data)? action;

  const AddTaskDialog({super.key, this.action});

  @override
  Widget build(BuildContext context) {
    final taskNameCtl = TextEditingController();
    final linkCtl = TextEditingController();

    return AlertDialog(
      backgroundColor: kLightSurface,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20)),
      title: const Text('添加新任务', style: TextStyle(color: kLightText)),
      content: SizedBox(
        width: 400,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextField(
              controller: linkCtl,
              style: const TextStyle(color: kLightText),
              decoration: InputDecoration(
                labelText: '下载链接',
                labelStyle: TextStyle(color: kLightTextSecondary),
                hintText: 'https://example.com/file.zip',
                hintStyle: TextStyle(color: kLightTextSecondary.withValues(alpha: 0.6)),
                filled: true,
                fillColor: kLightSurfaceLight,
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(12),
                  borderSide: BorderSide.none,
                ),
                prefixIcon: Icon(Icons.link_rounded, color: kLightTextSecondary),
              ),
            ),
            const SizedBox(height: 16),
            TextField(
              controller: taskNameCtl,
              style: const TextStyle(color: kLightText),
              decoration: InputDecoration(
                labelText: '任务名（可选）',
                labelStyle: TextStyle(color: kLightTextSecondary),
                filled: true,
                fillColor: kLightSurfaceLight,
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(12),
                  borderSide: BorderSide.none,
                ),
                prefixIcon: Icon(Icons.edit_rounded, color: kLightTextSecondary),
              ),
            ),
          ],
        ),
      ),
      actions: [
        TextButton(
          onPressed: () => Navigator.pop(context),
          child: Text('取消', style: TextStyle(color: kLightTextSecondary)),
        ),
        ElevatedButton(
          onPressed: () {
            if (linkCtl.text.isEmpty) return;
            action?.call(AddTaskData(linkCtl.text, taskName: taskNameCtl.text));
            Navigator.pop(context);
          },
          style: ElevatedButton.styleFrom(
            backgroundColor: kPrimary,
            shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
          ),
          child: const Text('下载'),
        ),
      ],
    );
  }
}
