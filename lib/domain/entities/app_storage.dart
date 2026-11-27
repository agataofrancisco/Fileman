/// Entidade imutável: estatísticas de armazenamento de uma aplicação.
enum AppStorageSource { system, estimated, unknown }

class AppStorageInfo {
  final String packageName;
  final String? label;
  final int? totalBytes;
  final int? dataBytes;
  final int? cacheBytes;
    final AppStorageSource source;
  final DateTime? measuredAt;
  final String? capability; // ownCacheCleanup | systemCacheIntent | openAppStorageSettings | unsupported

  const AppStorageInfo({
    required this.packageName,
    this.label,
    this.totalBytes,
    this.dataBytes,
    this.cacheBytes,
    required this.source,
    this.measuredAt,
    this.capability,
  });
}
