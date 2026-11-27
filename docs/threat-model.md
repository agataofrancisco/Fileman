# Threat Model — LIVRE

## Riscos identificados
1. Vazamento de dados pessoais para API Jev/gateway.
2. Logs/analytics com nomes/caminhos pessoais.
3. Exposição da chave da API Jev (no APK, repositório ou logs).
4. Upload involuntário de imagens sem consentimento.
5. Classificações incorretas da IA usadas como autorização de eliminação.
6. Eliminação não intencional de ficheiros pessoais.
7. Troca de URI (TOCTOU) entre validação e operação.
8. App maliciosa fornecendo content URI.
9. Documentos corrompidos / ZIP bombas a exaurir memória.
10. Leitura fora do escopo de permissão concedido.
11. Consumo excessivo de bateria e custos elevados de gateway.
12. Abuso do gateway por terceiros (chave comprometida).
13. Permissões revogadas durante operação em curso.

## Mitigações
- Least privilege: apenas MediaStore/SAF/concedido; nunca diretórios privados de outras apps.
- Validação da fonte URI, tamanho e MIME antes de leitura.
- Streaming seguro; limites de tamanho; proteção contra paths arbitrários.
- Revalidação pré-limpeza: `exists`, `permission`, `signature`, `protected`, operação suportada.
- Consentimento granular para envio Jev; categorias sensíveis exigem explicação adequada.
- Logs mínimos: nenhum nome/caminho pessoal em telemetria ou erros.
- Retenção mínima; cache de classificação com expiração e eliminação local.
- Chave Jev apenas em secrets do gateway Go; nunca no APK.
- Timeouts, rate limiting, quotas por utilizador/dispositivo/sessão, trava global de custos.
- Segregação gateway: não expor proxy público aberto.
- Testes de revogação de permissão e volume removido.
- Relatórios explicativos sem atribuição falsa de espaço libertado.

## Conduta de eliminação
- Jev nunca chama APIs de eliminação.
- `CleanupCoordinator` recebe apenas seleção explícita e validada.
- Preferir reciclagem (`createTrashRequest`) quando suportado.
- Sem promessa de undo quando o provider não suporta.
- Nenhum `File(path).delete(recursive: true)` em diretórios desconhecidos/partilhados sem API, permissão e confirmação.
- Pontuação Jev de «baixo risco» não é autorização automática.
