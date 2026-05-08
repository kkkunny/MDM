import 'package:flutter/material.dart';
import 'package:mdm/configs/theme.dart';

class AppButton extends StatelessWidget {
  final String text;
  final VoidCallback? onPressed;
  final Color? backgroundColor;
  final Color? textColor;
  final IconData? icon;
  final bool isLoading;
  final bool isOutlined;
  final double? width;
  final double? height;

  const AppButton({
    super.key,
    required this.text,
    this.onPressed,
    this.backgroundColor,
    this.textColor,
    this.icon,
    this.isLoading = false,
    this.isOutlined = false,
    this.width,
    this.height,
  });

  factory AppButton.primary({
    required String text,
    VoidCallback? onPressed,
    IconData? icon,
    bool isLoading = false,
    double? width,
  }) => AppButton(
    text: text,
    onPressed: onPressed,
    backgroundColor: kPrimary,
    textColor: Colors.white,
    icon: icon,
    isLoading: isLoading,
    width: width,
  );

  factory AppButton.danger({
    required String text,
    VoidCallback? onPressed,
    IconData? icon,
    bool isLoading = false,
    double? width,
  }) => AppButton(
    text: text,
    onPressed: onPressed,
    backgroundColor: kError,
    textColor: Colors.white,
    icon: icon,
    isLoading: isLoading,
    width: width,
  );

  factory AppButton.outlined({
    required String text,
    VoidCallback? onPressed,
    Color? color,
    IconData? icon,
    double? width,
  }) => AppButton(
    text: text,
    onPressed: onPressed,
    backgroundColor: Colors.transparent,
    textColor: color ?? Colors.white70,
    icon: icon,
    isOutlined: true,
    width: width,
  );

  @override
  Widget build(BuildContext context) {
    final bgColor = backgroundColor ?? kPrimary;
    final txtColor = textColor ?? Colors.white;

    return Material(
      color: Colors.transparent,
      child: InkWell(
        onTap: isLoading ? null : onPressed,
        borderRadius: BorderRadius.circular(12),
        child: Container(
          width: width,
          height: height ?? 48,
          padding: const EdgeInsets.symmetric(horizontal: 16),
          decoration: BoxDecoration(
            gradient: isOutlined
                ? null
                : LinearGradient(colors: [bgColor, bgColor.withValues(alpha: 0.8)]),
            color: isOutlined ? Colors.transparent : null,
            borderRadius: BorderRadius.circular(12),
            border: isOutlined ? Border.all(color: txtColor.withValues(alpha: 0.3)) : null,
            boxShadow: isOutlined
                ? null
                : [BoxShadow(
                    color: bgColor.withValues(alpha: 0.4),
                    blurRadius: 12,
                    offset: const Offset(0, 4),
                  )],
          ),
          child: Center(
            child: isLoading
                ? SizedBox(
                    width: 20,
                    height: 20,
                    child: CircularProgressIndicator(
                      strokeWidth: 2,
                      valueColor: AlwaysStoppedAnimation(txtColor),
                    ),
                  )
                : Row(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      if (icon != null) ...[
                        Icon(icon, color: txtColor, size: 20),
                        const SizedBox(width: 8),
                      ],
                      Text(text, style: TextStyle(color: txtColor, fontWeight: FontWeight.w600)),
                    ],
                  ),
          ),
        ),
      ),
    );
  }
}
