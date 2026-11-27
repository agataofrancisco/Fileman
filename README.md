# LIVRE — Gestor Inteligente de Armazenamento Android

App Android-first em Flutter + Kotlin nativo que ajuda o utilizador a descobrir onde o armazenamento é consumido, encontrar ficheiros grandes/antigos/duplicados, analisar multimédia acessível do WhatsApp e classificar documentos com processamento local + classificação Jev/TypeSafe opcional (beta).

> **O utilizador decide o destino dos seus ficheiros pessoais.** Jev classifica, nunca executa eliminação. Sem `autoDelete` de dados pessoais.

## Idioma
UI e documentação em português de Angola (pt-AO). Identificadores técnicos em inglês.

## Stack
- Flutter/Dart (Material 3, Riverpod, go_router, Drift/SQLite)
- Kotlin nativo Android (MediaStore, SAF, StorageStatsManager, StorageManager, DocumentsContract)
- **Go** backend gateway (Jev/TypeSafe autenticado, rate-limit, quotas, consentimento)
- ML Kit Text Recognition (OCR local)

## Estado
Primeira passagem (SPEC-FIRST): estrutura documental, OpenSpec e scaffolding criados. Implementação ainda não iniciada. Ver `docs/` e `openspec/`.

## Licença
Privado — uso interno.
