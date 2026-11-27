/// Result tipado para operações que podem falhar, sem exceções no fluxo feliz.
sealed class Result<T> {
  const Result();
}

final class Success<T> extends Result<T> {
  final T value;
  const Success(this.value);
}

final class Failure<T> extends Result<T> {
  final AppError error;
  const Failure(this.error);
}

class AppError {
  final String code;
  final String message;
  final bool isRetryable;
  final bool isDestructive;
  const AppError({
    required this.code,
    required this.message,
    this.isRetryable = false,
    this.isDestructive = false,
  });
}
