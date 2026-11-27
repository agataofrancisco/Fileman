# Fluxo de Dados e Privacidade — LIVRE

## Princípios
- Offline-first, privacy-first, least privilege.
- Conteúdo pessoal não sai do dispositivo sem decisão específica do utilizador.
- Não guardar tokens/chaves Jev no APK.
- Não registar nomes/caminhos pessoais em telemetria ou erros.

## Fluxo local (sem rede)
Scanner → Indexação SQLite/Drift (metadados apenas) → UI → Seleção do utilizador → Operação oficial do Android (trash/delete).

## Fluxo Jev (opt-in, experimental)
1. Utilizador abre Classificação Jev e lê o `AiConsentDialog` (o que será enviado, que é online, opt-in).
2. Consentimento gravado no `AiConsent` (versão, tipos de mídia, fontes, Wi-Fi only, habilitado).
3. App prepara input sanitizado: `requestId`, `fileRef` (ID opaco local), metadados mínimos, OCR local opcional, imagem cortada apenas se utilizador autorizou envio de imagem.
4. Nunca enviar: nome/caminho completo, contactos, localizações EXIF, ID de conta, conteúdo não necessário.
5. Requisição enviada ao gateway Go (autenticado). Gateway valida consentimento, rate-limit, quota e encaminha para Jev/TypeSafe.
6. Resposta mapeada para `Classification` local (categoria, subcategoria, confiança, sensível, motivo).
7. Resultado guardado em cache local por fingerprint + versão modelo + taxonomia.
8. Classificação nunca autoriza eliminação; `deleteAuthorized` é sempre `false`.

## Armazenamento local
- Base de dados Drift/SQLite no diretório da aplicação (não partilhado).
- Nenhum ficheiro binário de utilizador guardado na base.
- Thumbnails em cache LRU com quota (ex.: ≤100 MB configurável).
- Histórico de limpeza local, sem telemetria remota por defeito.
- Função «Limpar histórico local» não apaga ficheiros do utilizador.

## Retenção
- Imagens não retidas no gateway; logs sem corpo/payload.
- Retenção do fornecedor divulgada antes do beta Jev.
