import 'package:flutter/material.dart';

const kPrimary = Color(0xFF667EEA);
const kPrimaryDark = Color(0xFF764BA2);

const kSuccess = Color(0xFF00FF88);
const kWarning = Color(0xFFFFBE0B);
const kError = Color(0xFFFF6B6B);
const kInfo = Color(0xFF4ECDC4);
const kNeutral = Color(0xFF9CA3AF);

const kLightBackground = Color(0xFFF5F7FA);
const kLightSurface = Color(0xFFFFFFFF);
const kLightSurfaceLight = Color(0xFFF0F2F5);
const kLightText = Color(0xFF1A1A2E);
const kLightTextSecondary = Color(0xFF6B7280);
const kLightDivider = Color(0xFFE5E7EB);

class AppTheme {
  static ThemeData get lightTheme => ThemeData(
    useMaterial3: true,
    brightness: Brightness.light,
    fontFamily: 'Inter',
    scaffoldBackgroundColor: kLightBackground,
    colorScheme: const ColorScheme.light(
      primary: kPrimary,
      secondary: kPrimaryDark,
      surface: kLightSurface,
      error: kError,
    ),
    cardColor: kLightSurface,
    dividerColor: kLightDivider,
    textTheme: _lightTextTheme,
    iconTheme: const IconThemeData(color: kLightText),
    appBarTheme: const AppBarTheme(
      backgroundColor: kLightSurface,
      foregroundColor: kLightText,
      elevation: 0,
      centerTitle: false,
    ),
    elevatedButtonTheme: ElevatedButtonThemeData(
      style: ElevatedButton.styleFrom(
        backgroundColor: kPrimary,
        foregroundColor: Colors.white,
        padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
      ),
    ),
    textButtonTheme: TextButtonThemeData(
      style: TextButton.styleFrom(
        foregroundColor: kLightTextSecondary,
        padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
      ),
    ),
    inputDecorationTheme: InputDecorationTheme(
      filled: true,
      fillColor: kLightSurfaceLight,
      border: OutlineInputBorder(
        borderRadius: BorderRadius.circular(12),
        borderSide: BorderSide.none,
      ),
      enabledBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(12),
        borderSide: BorderSide.none,
      ),
      focusedBorder: OutlineInputBorder(
        borderRadius: BorderRadius.circular(12),
        borderSide: const BorderSide(color: kPrimary),
      ),
      labelStyle: TextStyle(color: kLightTextSecondary),
      hintStyle: TextStyle(color: kLightTextSecondary.withValues(alpha: 0.6)),
    ),
    dialogTheme: DialogThemeData(
      backgroundColor: kLightSurface,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(20)),
      titleTextStyle: const TextStyle(
        color: kLightText,
        fontSize: 16,
        fontWeight: FontWeight.w600,
      ),
      contentTextStyle: TextStyle(
        color: kLightTextSecondary,
        fontSize: 14,
      ),
    ),
    progressIndicatorTheme: const ProgressIndicatorThemeData(color: kPrimary),
  );

  static const TextTheme _lightTextTheme = TextTheme(
    displayLarge: TextStyle(fontSize: 57, fontWeight: FontWeight.bold, color: kLightText),
    displayMedium: TextStyle(fontSize: 45, fontWeight: FontWeight.bold, color: kLightText),
    displaySmall: TextStyle(fontSize: 36, fontWeight: FontWeight.bold, color: kLightText),
    headlineLarge: TextStyle(fontSize: 32, fontWeight: FontWeight.bold, color: kLightText),
    headlineMedium: TextStyle(fontSize: 28, fontWeight: FontWeight.bold, color: kLightText),
    headlineSmall: TextStyle(fontSize: 24, fontWeight: FontWeight.bold, color: kLightText),
    titleLarge: TextStyle(fontSize: 22, fontWeight: FontWeight.w600, color: kLightText),
    titleMedium: TextStyle(fontSize: 16, fontWeight: FontWeight.w600, color: kLightText),
    titleSmall: TextStyle(fontSize: 14, fontWeight: FontWeight.w600, color: kLightText),
    bodyLarge: TextStyle(fontSize: 16, fontWeight: FontWeight.normal, color: kLightText),
    bodyMedium: TextStyle(fontSize: 14, fontWeight: FontWeight.normal, color: kLightText),
    bodySmall: TextStyle(fontSize: 12, fontWeight: FontWeight.normal, color: kLightTextSecondary),
    labelLarge: TextStyle(fontSize: 14, fontWeight: FontWeight.w500, color: kLightText),
    labelMedium: TextStyle(fontSize: 12, fontWeight: FontWeight.w500, color: kLightText),
    labelSmall: TextStyle(fontSize: 11, fontWeight: FontWeight.w500, color: kLightTextSecondary),
  );
}
