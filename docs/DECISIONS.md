# Registro de Decisões (ADRs leves)

Formato: Data | Contexto | Opção adotada | Alternativas rejeitadas | Impacto

## 2026-11-27 — Stack de backend
- Contexto: O prompt mestre sugeria Cloudflare Workers para o gateway Jev. O utilizador solicitou backend em Go.
- Opção adotada: Backend gateway em Go (Go stdlib HTTP + JSON), implantável em qualquer infra (VPS, Cloud Run, etc.). Substitui Cloudflare Workers.
- Alternativas rejeitadas: Cloudflare Workers (limite de multipart a verificar, lock-in de fornecedor).
- Impacto: `backend/jev-gateway/` em Go; contrato gateway ↔ Flutter mantém-se JSON; Chave Jev apenas no backend Go, nunca no APK.

## 2026-11-27 — application_id e pacote Kotlin
- Contexto: Identificador da aplicação Android.
- Opção adotada: `ao.libre.fileman` (application_id) e pacote Kotlin `ao.libre.fileman`.
- Alternativas rejeitadas: `com.fileman.libre`, `ao.guitchi.libre`.
- Impacto: Manifesto, Gradle, bridge Kotlin e pasta nativa usam este pacote.

## 2026-11-27 — minSdk / targetSdk / compileSdk
- Contexto: Versões alvo do Android.
- Opção adotada: minSdk 24, targetSdk 36, compileSdk 36.
- Alternativas rejeitadas: minSdk inferior a 24 (não suporta scoped storage bem).
- Impacto: Cobertura Android 7.0+; matriz de capacidades para API 30 (11), 33 (13), 34 (14), 35 (15), 36 (16).
