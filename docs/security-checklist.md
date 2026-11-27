# Security Checklist — LIVRE

- [ ] Chave Jev apenas no backend Go (secrets), nunca no APK nem repositório.
- [ ] `.env.example` sem chaves reais; `.gitignore` inclui `.env`.
- [ ] Sem nomes/caminhos pessoais em logs ou telemetria.
- [ ] Consentimento Jev gravado antes de qualquer upload; sem upload silencioso.
- [ ] `deleteAuthorized` sempre `false` na classificação Jev.
- [ ] Revalidação pré-limpeza: `exists`, `permission`, `signature`, `protected`.
- [ ] Preferir `createTrashRequest`; `createDeleteRequest` só para eliminação permanente explícita.
- [ ] Nenhum `File(path).delete(recursive: true)` em diretórios desconhecidos/partilhados.
- [ ] Validação de URI, tamanho e MIME antes de leitura (proteção contra path traversal / ZIP bomba).
- [ ] Streaming seguro; limites de tamanho de ficheiro processado.
- [ ] Timeouts e retry com exponential backoff + jitter no gateway (sem retry infinito).
- [ ] Rate limit e quota por utilizador/dispositivo/sessão; trava global de custos.
- [ ] Cache de classificação com expiração e eliminação local.
- [ ] Testes de revogação de permissão e volume removido durante operação.
- [ ] Build de debug compila; analyzer/linter/test suite passam.
- [ ] Funcionalidades sem suporte comunicadas honestamente (`unsupported`).
- [ ] Estado `done` apenas com verificação executada, não por intenção.
