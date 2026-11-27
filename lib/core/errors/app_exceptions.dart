/// Exceções de domínio da aplicação LIVRE.
class AppException implements Exception {
  final String code;
  final String message;
  const AppException(this.code, this.message);
  @override
  String toString() => 'AppException($code: $message)';
}

class PermissionException extends AppException {
  const PermissionException(String message) : super('PERMISSION_DENIED', message);
}

class StorageUnsupportedException extends AppException {
  const StorageUnsupportedException(String message) : super('UNSUPPORTED', message);
}

class CleanupException extends AppException {
  const CleanupException(String message) : super('CLEANUP_FAILED', message);
}

class AiException extends AppException {
  const AiException(String message) : super('AI_FAILED', message);
}
