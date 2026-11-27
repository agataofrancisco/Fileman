/// Entidade imutável: ficheiro indexado pelo scanner.
enum MediaKind { image, video, audio, document, other }

enum AccessStatus { accessible, partial, inaccessible, protected }

class IndexedFile {
  final String id;
  final String sourceId;
  final String contentUri;
  final String? displayName;
  final String? relativePath;
  final String? mime;
  final int? sizeBytes;
  final DateTime? modifiedAt;
  final DateTime? capturedAt;
  final DateTime? addedAt;
  final int? durationMs;
  final int? width;
  final int? height;
  final MediaKind mediaKind;
  final String? fingerprint;
  final bool protected;
  final AccessStatus accessStatus;
  final int scanGeneration;
  final String? contentSignature;

  const IndexedFile({
    required this.id,
    required this.sourceId,
    required this.contentUri,
    this.displayName,
    this.relativePath,
    this.mime,
    this.sizeBytes,
    this.modifiedAt,
    this.capturedAt,
    this.addedAt,
    this.durationMs,
    this.width,
    this.height,
    this.mediaKind = MediaKind.other,
    this.fingerprint,
    this.protected = false,
    this.accessStatus = AccessStatus.accessible,
    required this.scanGeneration,
    this.contentSignature,
  });
}
