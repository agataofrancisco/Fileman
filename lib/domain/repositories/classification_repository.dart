/// Interface do repositório de classificação.
abstract class ClassificationRepository {
  Future<Classification> classifyFile(String fileId, {bool forceJeV = false});
  Future<Classification?> getClassification(String fileId);
  Future<List<Classification>> listClassifications({int? limit, int? offset});
  Future<void> setManualClassification(String fileId, String category);
  Future<void> clearCache();
}
