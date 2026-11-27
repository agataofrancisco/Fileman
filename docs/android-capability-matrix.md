# Matriz de Capacidades Android

Por API level e classe de permissão. Fontes: documentação oficial Android (developer.android.com) e política Google Play. Não inventar acesso a diretórios privados.

Colunas: Função | Android min | Android target | Permissão | API | Confirmação do sistema? | Dados acessíveis | Dados inacessíveis | Fallback | Publicável na Play? | Evidência/URL | Resultado em device físico

| Função | min | target | Permissão | API | Confirmação sistema? | Acessível | Inacessível | Fallback | Play? | Evidência | Device físico |
|--------|-----|--------|-----------|-----|----------------------|-----------|-------------|----------|-------|-----------|---------------|
| Listar imagens/vídeos/áudio partilhados | 24 | 36 | READ_MEDIA_IMAGES/VIDEO/AUDIO (13+) | MediaStore / ContentResolver | Não (leitura) | Media indexado concedido | `/Android/data` de outras apps, bases de dados privadas | SAF / permissões parciais | Sim | developer.android.com/training/data-storage/shared/media | PENDING |
| Criar lixo (trash) de media | 30 | 36 | Nenhuma (própria URI) | MediaStore.createTrashRequest | Sim (dialog nativo) | URI MediaStore válida | Ficheiros fora do MediaStore | createDeleteRequest (permanente) | Sim | developer.android.com/reference/android/provider/MediaStore | PENDING |
| Eliminar permanente media | 30 | 36 | Nenhuma (própria URI) | MediaStore.createDeleteRequest | Sim (dialog nativo) | URI MediaStore válida | Fora do MediaStore | Informar irreversível | Sim | developer.android.com | PENDING |
| SAF abrir árvore de documentos | 24 | 36 | Nenhuma | ACTION_OPEN_DOCUMENT_TREE | Sim (picker) | Raiz selecionada pelo utilizador | `/Android/data`, `/Android/obb` (não presumir) | Seleção manual de ficheiros | Sim | developer.android.com/training/data-storage/shared/documents-files | PENDING |
| Estatísticas de armazenamento por app | 28 | 36 | PACKAGE_USAGE_STATS (especial) | StorageStatsManager | Não | Estimativas por categoria (quando API permite) | Tamanho exato de ficheiros privados | Abrir definições de armazenamento | Sim (justificar) | developer.android.com/reference/android/app/usage/StorageStatsManager | PENDING |
| Limpar cache externa (app alvo) | 28 | 36 | Nenhuma | StorageManager.ACTION_CLEAR_APP_CACHE | Sim | Cache externa (plataformas suportadas) | Cache interna Chrome, dados privados | ACTION_APPLICATION_DETAILS_SETTINGS | Sim | developer.android.com/reference/android/os/storage/StorageManager | PENDING |
| Abrir definições de armazenamento da app | 24 | 36 | Nenhuma | ACTION_APPLICATION_DETAILS_SETTINGS | Não | Página oficial do SO | Nenhum dado direto | Mensagem explicativa | Sim | developer.android.com | PENDING |
| MANAGE_EXTERNAL_STORAGE | 30 | 36 | Permissão especial | Settings.ACTION_MANAGE_ALL_FILES_ACCESS_PERMISSION | Sim (utilizador concede) | Amplo (sujeito a restrições fabricantes) | Diretórios privados de outras apps | Variante sem a permissão | Só se justificada à Play | developer.android.com/training/data-storage/manage-all-files | PENDING |
| WhatsApp media | 24 | 36 | READ_MEDIA_* | MediaStore / caminho candidato | Não | Media indexado acessível | Bases de dados privadas, conversas cifradas | SAF / seletor do utilizador | Sim | (candidato de descoberta, não garantido) | PENDING |

Versões cobertas na matriz: Android 11 (API 30), 13 (API 33), 14 (API 34), 15 (API 35), 16 (API 36).
Resultado em device físico: **PENDING** — nenhum dispositivo físico conectado neste ambiente.
