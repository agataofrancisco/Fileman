/// Entidade imutável: volume de armazenamento.
class StorageVolume {
  final String id;
  final String displayName;
  final int? totalBytes;
  final int? freeBytes;
  final String kind; // internal, primary, sd_card, usb, etc.
  final bool mounted;

  const StorageVolume({
    required this.id,
    required this.displayName,
    this.totalBytes,
    this.freeBytes,
    required this.kind,
    required this.mounted,
  });
}
