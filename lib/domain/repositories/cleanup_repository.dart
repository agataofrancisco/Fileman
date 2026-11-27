/// Interface do repositório de operações de limpeza.
abstract class CleanupRepository {
  Future<CleanupReport> executePlan(CleanupPlan plan);
  Future<CleanupReport?> getReport(String id);
  Future<List<CleanupReport>> listReports({int? limit, int? offset});
  Future<void> clearLocalHistory();
}
