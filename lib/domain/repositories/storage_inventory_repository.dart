/// Interface do repositório de inventário de armazenamento.
abstract class StorageInventoryRepository {
  Future<List<StorageVolume>> listVolumes();
  Future<List<StorageSource>> listSources();
  Future<void> indexSource(StorageSource source);
  Future<List<IndexedFile>> queryFiles({
    String? sourceId,
    String? mime,
    int? minSize,
    int? maxSize,
    DateTime? before,
    DateTime? after,
    int? limit,
    int? offset,
  });
  Future<IndexedFile?> getFile(String id);
}
