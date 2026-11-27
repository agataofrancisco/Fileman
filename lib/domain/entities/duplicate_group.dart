/// Entidade imutável: grupo de duplicados (exato ou semelhante).
enum DuplicateKind { exact, similar }

class DuplicateGroup {
  final String id;
  final DuplicateKind kind;
  final List<String> fileIds;
  final int candidateSavingsBytes;
  final bool verified; // true apenas para exact com hash completo válido

  const DuplicateGroup({
    required this.id,
    required this.kind,
    required this.fileIds,
    required this.candidateSavingsBytes,
    this.verified = false,
  });
}
