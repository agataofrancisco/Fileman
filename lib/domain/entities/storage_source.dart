/// Entidade imutável: fonte de indexação.
enum StorageSourceKind { mediaStore, saf, appOwn, allFiles }

enum PermissionState { granted, denied, limited, unknown, revoked }

class StorageSource {
  final String id;
  final StorageSourceKind kind;
  final String? rootUri;
  final PermissionState permissionState;
  final String coverage; // full | partial | unknown

  const StorageSource({
    required this.id,
    required this.kind,
    this.rootUri,
    required this.permissionState,
    this.coverage = 'unknown',
  });
}
