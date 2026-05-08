import 'package:flutter/material.dart';

class IconHelper {
  static IconData getFileIcon(String fileName) {
    final ext = fileName.split('.').last.toLowerCase();
    return switch (ext) {
      'zip' || 'rar' || '7z' || 'tar' || 'gz' => Icons.folder_zip_rounded,
      'exe' || 'dmg' || 'pkg' || 'msi' => Icons.install_desktop_rounded,
      'iso' => Icons.album_rounded,
      'pdf' => Icons.picture_as_pdf_rounded,
      'mp4' || 'mkv' || 'avi' || 'mov' => Icons.video_file_rounded,
      'mp3' || 'wav' || 'flac' => Icons.audio_file_rounded,
      'jpg' || 'jpeg' || 'png' || 'gif' || 'webp' => Icons.image_rounded,
      _ => Icons.insert_drive_file_rounded,
    };
  }
}
