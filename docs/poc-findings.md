# Prova de Viabilidade (PoC) — Achados

**Estado: PENDING.** Nenhum teste executado ainda.

## Objetivo da PoC (change 01)
Validar capacidades reais de armazenamento Android antes de construir a experiência final:
1. `MediaStore` listar ficheiros concedidos e metadados reais.
2. Ficheiro de teste numa origem tipo WhatsApp (sem manipular dados reais) descoberto se acessível.
3. `MediaStore.createTrashRequest` apresentar confirmação do sistema e reportar resultado correto.
4. Acesso à página oficial de storage do Chrome (quando instalado).
5. Verificação de `MANAGE_EXTERNAL_STORAGE` e restrições em Android atual, sem pretender acesso privado.

## Lacunas a validar em device físico
- [ ] Dispositivo físico conectado (nenhum neste ambiente; `adb devices` vazio).
- [ ] API levels: 30 (11), 33 (13), 34 (14), 35 (15), 36 (16).
- [ ] MediaStore queries e projection.
- [ ] SAF URI grants e `FLAG_SUPPORTS_DELETE`.
- [ ] `createTrashRequest` / `createDeleteRequest` resultados e falhas parciais.
- [ ] `StorageStatsManager` e `PACKAGE_USAGE_STATS`.
- [ ] `ACTION_CLEAR_APP_CACHE` e `ACTION_APPLICATION_DETAILS_SETTINGS` (Chrome).
- [ ] Caminho candidato WhatsApp media (`/storage/emulated/0/Android/media/com.whatsapp/WhatsApp/Media/` ou variantes).

Marca: `PENDING_DEVICE_VALIDATION` até haver dispositivo físico.
