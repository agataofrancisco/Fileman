/// Interface do repositório de consentimento (Jev e permissões).
abstract class ConsentRepository {
  Future<bool> isJeVEnabled();
  Future<void> setJeVEnabled(bool enabled);
  Future<bool> hasImageUploadConsent();
  Future<void> grantImageUploadConsent();
  Future<void> revokeImageUploadConsent();
}
