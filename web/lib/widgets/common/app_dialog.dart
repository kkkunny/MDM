import 'package:flutter/material.dart';
import 'package:mdm/configs/theme.dart';

class AppDialog extends StatelessWidget {
  final String title;
  final String? content;
  final List<Widget> actions;
  final Widget? customContent;

  const AppDialog({
    super.key,
    required this.title,
    this.content,
    required this.actions,
    this.customContent,
  });

  static Future<bool?> showConfirmDialog(
    BuildContext context, {
    required String title,
    required String content,
    String confirmText = '确认',
    String cancelText = '取消',
    Color confirmColor = kPrimary,
  }) {
    return showDialog<bool>(
      context: context,
      builder: (context) => AppDialog(
        title: title,
        content: content,
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context, false),
            child: Text(cancelText, style: TextStyle(color: kLightTextSecondary)),
          ),
          ElevatedButton(
            onPressed: () => Navigator.pop(context, true),
            style: ElevatedButton.styleFrom(
              backgroundColor: confirmColor,
              shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
            ),
            child: Text(confirmText),
          ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      backgroundColor: kLightSurface,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20)),
      title: Text(title, style: const TextStyle(color: kLightText)),
      content: customContent ??
          (content != null ? Text(content!, style: TextStyle(color: kLightTextSecondary)) : null),
      actions: actions,
    );
  }
}
