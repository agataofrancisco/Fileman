/// Entidade imutável: resultado de classificação (local ou Jev).
enum ClassificationSource { local, jev, manual }

enum ClassificationStatus { classified, unknown, failed, pending }

class Classification {
  final String id;
  final String fileId;
  final String taxonomyVersion;
  final String modelId;
  final String inputFingerprint;
  final String category;
  final String? subcategory;
  final double? confidence;
  final bool sensitiveFlag;
  final ClassificationSource source;
  final ClassificationStatus status;
  final DateTime analyzedAt;
  final List<String>? reasonCodes;
  final bool deleteAuthorized; // sempre false por regra de negócio

  const Classification({
    required this.id,
    required this.fileId,
    required this.taxonomyVersion,
    required this.modelId,
    required this.inputFingerprint,
    required this.category,
    this.subcategory,
    this.confidence,
    this.sensitiveFlag = false,
    required this.source,
    required this.status,
    required this.analyzedAt,
    this.reasonCodes,
    this.deleteAuthorized = false,
  });
}
