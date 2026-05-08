import 'package:flutter/material.dart';
import 'package:mdm/configs/theme.dart';
import 'package:mdm/models/vo/task.pb.dart';

Color getStatusColor(TaskPhase phase) {
  final suffix = phase.name.replaceFirst('TpDown', '').replaceFirst('TpUp', '');
  return switch (suffix) {
    'Queued' => kNeutral,
    'Running' => kInfo,
    'Paused' => kWarning,
    'Failed' => kError,
    'Completed' => kSuccess,
    _ => kNeutral,
  };
}

Color getStatusBackgroundColor(TaskPhase phase) => getStatusColor(phase).withValues(alpha: 0.1);

Color getStatusBorderColor(TaskPhase phase) => getStatusColor(phase).withValues(alpha: 0.3);

LinearGradient getPrimaryGradient() => const LinearGradient(colors: [kPrimary, kPrimaryDark]);
