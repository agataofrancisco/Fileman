/// Entidades de limpeza (plano, entrada e relatório).
enum CleanupAction { trash, deletePermanent }

enum CleanupStatus { pending, validated, success, failed, cancelled, skipped }

class CleanupEntry {
  final String id;
  final String fileId;
  final CleanupAction action;
  final CleanupStatus status;
  final String? errorCode;
  final DateTime? attemptedAt;
  final DateTime? completedAt;

  const CleanupEntry({
    required this.id,
    required this.fileId,
    required this.action,
    required this.status,
    this.errorCode,
    this.attemptedAt,
    this.completedAt,
  });
}

class CleanupPlan {
  final String id;
  final DateTime createdAt;
  final List<CleanupEntry> entries;
  final int selectedBytes;
  final String validationState; // valid | invalid | needsRevalidation

  const CleanupPlan({
    required this.id,
    required this.createdAt,
    required this.entries,
    required this.selectedBytes,
    this.validationState = 'valid',
  });
}

class CleanupReport {
  final String id;
  final Map<CleanupStatus, int> countsByStatus;
  final int selectedBytes;
  final int trashedBytes;
  final int deletedBytes;
  final int? observedFreeBytesDelta;
  final DateTime generatedAt;

  const CleanupReport({
    required this.id,
    required this.countsByStatus,
    required this.selectedBytes,
    required this.trashedBytes,
    required this.deletedBytes,
    this.observedFreeBytesDelta,
    required this.generatedAt,
  });
}
