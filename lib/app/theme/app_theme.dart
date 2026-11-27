import 'package:flutter/material.dart';

/// Tokens de cor do Design System LIVRE (dark-first, pt-AO).
abstract class AppColors {
  static const Color background = Color(0xFF0B0F19);
  static const Color surface = Color(0xFF151B2A);
  static const Color primary = Color(0xFF8B5CF6);
  static const Color primaryContainer = Color(0xFF3B0764);
  static const Color success = Color(0xFF22C55E);
  static const Color warning = Color(0xFFF59E0B);
  static const Color error = Color(0xFFEF4444);
  static const Color textPrimary = Color(0xFFFFFFFF);
  static const Color textSecondary = Color(0xFF94A3B8);
  static const Color textMuted = Color(0xFF64748B);
  static const Color border = Color(0xFF1E293B);
}

final ThemeData livreDarkTheme = ThemeData.dark().copyWith(
  scaffoldBackgroundColor: AppColors.background,
  cardColor: AppColors.surface,
  colorScheme: const ColorScheme.dark(
    primary: AppColors.primary,
    surface: AppColors.surface,
    background: AppColors.background,
    error: AppColors.error,
    onPrimary: AppColors.textPrimary,
    onSurface: AppColors.textPrimary,
    onBackground: AppColors.textPrimary,
    onError: AppColors.textPrimary,
  ),
  dividerColor: AppColors.border,
);

final ThemeData livreLightTheme = ThemeData.light();

class AppTheme {
  static ThemeData dark() => livreDarkTheme;
  static ThemeData light() => livreLightTheme;
}
