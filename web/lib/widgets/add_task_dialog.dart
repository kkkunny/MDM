import 'package:flutter/material.dart';
import 'package:mdm/configs/theme.dart';

class AddTaskData {
  final List<String> links;
  final String taskName;
  final String category;

  AddTaskData(this.links, {this.taskName = '', this.category = ''});
}

class AddTaskDialog extends StatefulWidget {
  final void Function(AddTaskData data)? action;

  const AddTaskDialog({super.key, this.action});

  @override
  State<AddTaskDialog> createState() => _AddTaskDialogState();
}

class _AddTaskDialogState extends State<AddTaskDialog> {
  final _linkCtrls = <TextEditingController>[TextEditingController()];
  final _categoryCtl = TextEditingController();
  final _taskNameCtl = TextEditingController();

  @override
  void dispose() {
    for (final c in _linkCtrls) {
      c.dispose();
    }
    _categoryCtl.dispose();
    _taskNameCtl.dispose();
    super.dispose();
  }

  void _addLinkField() {
    setState(() {
      _linkCtrls.add(TextEditingController());
    });
  }

  void _removeLinkField(int index) {
    setState(() {
      _linkCtrls[index].dispose();
      _linkCtrls.removeAt(index);
    });
  }

  List<String> _getValidLinks() {
    return _linkCtrls.map((c) => c.text.trim()).where((s) => s.isNotEmpty).toList();
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      backgroundColor: kLightSurface,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20)),
      title: const Text('添加新任务', style: TextStyle(color: kLightText)),
      content: SizedBox(
        width: 400,
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            ...List.generate(_linkCtrls.length, (i) {
              final ctl = _linkCtrls[i];
              return Padding(
                padding: EdgeInsets.only(bottom: i < _linkCtrls.length - 1 ? 8 : 0),
                child: Row(
                  children: [
                    Expanded(
                      child: TextField(
                        controller: ctl,
                        style: const TextStyle(color: kLightText),
                        decoration: InputDecoration(
                          labelText: i == 0 ? '下载链接' : '备选链接 ${i + 1}',
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
                    ),
                    if (_linkCtrls.length > 1)
                      SizedBox(
                        width: 32,
                        child: IconButton(
                          icon: Icon(Icons.close_rounded, color: kLightTextSecondary, size: 18),
                          onPressed: () => _removeLinkField(i),
                          padding: EdgeInsets.zero,
                        ),
                      ),
                  ],
                ),
              );
            }),
            const SizedBox(height: 8),
            Align(
              alignment: Alignment.centerLeft,
              child: TextButton.icon(
                onPressed: _addLinkField,
                icon: Icon(Icons.add_rounded, color: kPrimary, size: 18),
                label: Text('添加备选链接', style: TextStyle(color: kPrimary, fontSize: 13)),
                style: TextButton.styleFrom(padding: EdgeInsets.zero),
              ),
            ),
            const SizedBox(height: 16),
            TextField(
              controller: _categoryCtl,
              style: const TextStyle(color: kLightText),
              decoration: InputDecoration(
                labelText: '类别（可选）',
                labelStyle: TextStyle(color: kLightTextSecondary),
                filled: true,
                fillColor: kLightSurfaceLight,
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(12),
                  borderSide: BorderSide.none,
                ),
                prefixIcon: Icon(Icons.category_rounded, color: kLightTextSecondary),
              ),
            ),
            const SizedBox(height: 16),
            TextField(
              controller: _taskNameCtl,
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
            final links = _getValidLinks();
            if (links.isEmpty) return;
            widget.action?.call(AddTaskData(links, taskName: _taskNameCtl.text, category: _categoryCtl.text));
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
