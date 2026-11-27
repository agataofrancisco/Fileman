# PROMPT MESTRE — LIVRE
## Especificação e implementação orientada por OpenSpec para um gestor inteligente de armazenamento Android

> **Como utilizar:** entrega TODO este documento ao teu agente de programação, no diretório raiz do repositório. Este texto é a fonte inicial de requisitos do projeto. Executa-o primeiro em modo **SPEC-FIRST**. Não saltes a especificação para escrever a aplicação diretamente.
>
> **Estado:** requisitos de produto aprovados; detalhes técnicos devem ser verificados e especificados. **Nome LIVRE é provisório.** Idioma da aplicação e da documentação: português de Angola (pt-AO). Identificadores técnicos e nomes de código podem ser em inglês.

---

# 0. PAPEL, MISSÃO E REGRAS DE EXECUÇÃO DO AGENTE

És o **Tech Lead, Product Engineer, Android Storage Specialist, Flutter Architect, Security Engineer, QA Lead e OpenSpec Maintainer** responsável por estruturar e implementar o projeto LIVRE. Não és apenas um gerador de código: deves tomar decisões técnicas fundamentadas, documentar compromissos, reconhecer restrições reais do Android, criar especificações testáveis e executar em mudanças pequenas e verificáveis.

**Missão:** construir um aplicativo Android em Flutter com integrações Kotlin que ajude o utilizador a descobrir onde o armazenamento é consumido, encontrar ficheiros grandes/antigos/duplicados, analisar multimédia acessível do WhatsApp, identificar cache e opções legítimas de manutenção de aplicações, classificar ficheiros com heurísticas locais e o Jev online (opcional), e **apresentar informação e recomendações para que o próprio utilizador escolha o que eliminar**.

## Princípios inegociáveis

1. **O utilizador decide o destino dos seus ficheiros pessoais.** Jev classifica, nunca executa eliminação. Não implementar `autoDelete` para fotografias, vídeos, documentos, conversas, comprovativos ou outros dados pessoais.
2. Só propor «limpeza rápida» para operações claramente conhecidas e autorizadas: cache da própria aplicação, temporários que o nosso app criou, ou mecanismo oficial do sistema com a confirmação apropriada. Mesmo quando o conteúdo é reconstruível, a operação pode causar consumo de dados, bateria e perda temporária de desempenho.
3. **Sem root, sem ADB oculto, sem exploits, sem APIs privadas e sem abusar de AccessibilityService.** Não prometer acesso indiscriminado à cache privada do Chrome nem à base de dados privada do WhatsApp.
4. Trabalhar com **capacidades reais**, dependentes de versão do Android, fabricante, permissões e URI/provider. Não afirmar que uma pasta está acessível sem teste.
5. **Offline-first, privacy-first, least privilege.** Varredura e organização locais por defeito. Conteúdo pessoal não sai do dispositivo sem uma decisão específica do utilizador.
6. Classificar com **Jev como único modelo comercial online inicial**. OCR local e algoritmos determinísticos não são «outros LLMs». Não integrar MiniMax, OpenAI, Gemini, Claude ou outros fornecedores de LLM no MVP. Preparar uma interface interna para eventual substituição sem alterar domínio.
7. Visão do Jev é **experimental**, deve ficar em beta controlada por feature flag; falha de IA não pode bloquear o scanner ou a limpeza manual.
8. A IA e os metadados não são prova de inutilidade. Criar estado `unknown`, tratamento de ambiguidade e regras de proteção.
9. **Nada de métricas ou valores fictícios apresentados como reais.** Distinguir `bytes encontrados`, `bytes selecionados`, `bytes enviados para reciclagem`, `bytes realmente libertados` e `bytes inacessíveis/não medidos`.
10. Não fazer alterações destrutivas ao repositório existente. Identificar primeiro tecnologia, padrões e configurações existentes; propor mudanças compatíveis.
11. Evitar plugins Flutter que afirmem capacidades Android que não possuem. Verificar documentação oficial, licença, manutenção, compatibilidade com compileSdk/targetSdk e necessidade de código nativo.
12. Nada de UI bonita com botões decorativos: cada ação tem estado e fluxo reais, pelo menos mockado e testável antes da integração.
13. Não guardar tokens/chaves Jev no APK; credenciais sensíveis apenas num gateway seguro. Não registar nomes/caminhos pessoais em telemetria ou erros.
14. Separar factos confirmados, hipóteses e itens a validar em dispositivos físicos. Escolher a solução de menor privilégio que satisfaça a função.
15. A documentação deve permanecer sincronizada com o código e com os testes. Nunca marcar uma tarefa como concluída sem evidência concreta.

## Regime de execução e paragem

- **Primeira interação / primeira passagem:** inspecionar repositório, recolher fontes, inicializar/atualizar OpenSpec, criar `proposal`, `specs`, `design`, `tasks`, ADRs e plano de provas técnicas, validar a consistência e **parar antes de aplicar modificações substanciais ao app**. É autorizada apenas a preparação de estrutura documental e scaffolding estritamente necessário, se apropriado.
- **Após aprovação explícita para implementar:** executar os changes de forma incremental, começando pela PoC Android em dispositivo real, e depois aplicar cada change conforme dependências e critérios de aceitação. Dar relatórios objetivos de progresso.
- Se o utilizador disser «implementa tudo» ou «avança com o código», não ficar preso a perguntas não essenciais; avançar por fases, sem saltar testes e sem inventar validações de hardware/loja.
- Se não existir dispositivo real ligado, produzir teste instrumentado e instruções de validação, sinalizando a validação física como **pendente**, não como aprovada.
- Se aparecer um bloqueio crítico, documentar o bloqueio e prosseguir com funcionalidades independentes sem o esconder.

---

# 1. CONTEXTO DO PRODUTO E PROBLEMA A RESOLVER

## 1.1 Problema

Telefones Android de utilização prolongada acumulam dezenas de milhares de imagens, vídeos, áudios, documentos, downloads e outros conteúdos. Aplicações de mensagens, redes sociais e navegadores podem ocupar espaço com media, cache e dados internos. A maioria dos utilizadores não consegue perceber facilmente de onde vem o consumo de armazenamento nem decidir, em lote e com contexto, quais os ficheiros que pretende conservar. Os cleaners tradicionais organizam por extensão, tamanho ou idade, mas não compreendem os ficheiros nem as prioridades pessoais do utilizador.

## 1.2 Proposta de valor

**«Descobre o que está a ocupar espaço, entende o que tens e decide rapidamente o que manter ou remover.»**

A aplicação permite:
- Ver um panorama fiel do armazenamento acessível e das limitações da análise.
- Localizar grandes consumidores de espaço: ficheiros grandes, vídeos, duplicados, downloads, documentos e multimédia do WhatsApp.
- Ordenar/filtrar por categoria, origem, dimensão, datas conhecidas, duplicação, tipo e classificação semântica.
- Classificar fotografias de produtos, comprovativos e outros documentos usando processamento local + Jev opcional.
- Selecionar e eliminar em lote com pré-visualização, avisos, confirmação e operações oficiais do Android.
- Descobrir aplicações que ocupam espaço e abrir as interfaces corretas do sistema para gerir armazenamento e cache de aplicações como Chrome e WhatsApp.
- Operar com Internet fraca, aparelhos de entrada e armazenamento quase cheio.

## 1.3 Público prioritário

Utilizadores de Android, inicialmente em Angola, incluindo equipamentos de entrada (4–8 GB RAM ou menos), armazenamento limitado, conectividade cara/intermitente e uso prolongado de WhatsApp. O app deve usar pt-AO e internacionalização preparada para outras línguas. Não assumir dispositivos topo de gama nem Internet permanente.

## 1.4 Objetivos do MVP

M1. Inventariar e categorizar os ficheiros a que existe acesso autorizado.
M2. Mostrar consumo e *cobertura da análise* de forma compreensível.
M3. Encontrar duplicados exatos e candidatos pela idade/tamanho, sem falsos positivos de duplicação exata.
M4. Organizar multimédia acessível do WhatsApp, distinguindo as áreas inacessíveis.
M5. Apresentar gestão assistida de cache/aplicações sem prometer limpeza privada de terceiros.
M6. Disponibilizar pré-visualização, seleção, revisão e eliminação autorizada.
M7. Disponibilizar classificação Jev opcional como beta com limites, consentimento, cache e tratamento de falhas.
M8. Documentar capacidades Android, medir desempenho e preparar publicação conforme política do Google Play.

## 1.5 Fora do MVP

- Ler conversas ou bases de dados privadas do WhatsApp.
- Recuperar/eliminar dados privados de outras aplicações à força.
- Apagar cookies/sessões/passwords do Chrome diretamente.
- Limpeza universal e automática de «lixo» de terceiros.
- Gestor de permissões por AccessibilityService.
- Root-only functionality.
- Antivírus, aceleração fictícia de RAM, promessa de «aumentar FPS».
- iOS, Windows e macOS nesta primeira fase.
- Chatbot genérico ou assistente conversacional obrigatório.
- Classificação de vídeos completos e áudio integral com LLM.
- Cloud backup, sincronização multidispositivo e publicidade intrusiva.
- Subscrição, pagamentos ou login obrigatório no MVP; preparar arquitetura para monetização futura sem a implementar prematuramente.

---

# 2. DESIGN EXISTENTE — REFERÊNCIA OBRIGATÓRIA

O projeto tem dez ecrãs conceptuais criados anteriormente. **Se estiverem presentes no repositório, inspecioná-los primeiro; não os ignorar nem copiá-los de forma acrítica.** Os ficheiros deverão ficar em `design/references/`:

- `01_boas_vindas.png`
- `02_dashboard.png`
- `03_analise.png`
- `04_resultados.png`
- `05_whatsapp.png`
- `06_duplicados.png`
- `07_classificacao_jev.png`
- `08_aplicacoes_cache.png`
- `09_revisao.png`
- `10_relatorio.png`
- `livre_prototipo.html` (protótipo visual de referência, não implementação final).

Criar um *UI inventory* com componentes visuais de cada ecrã, hierarquia, estados e navegação. A identidade provisória é **LIVRE**, visual dark-first, fundo quase preto/navy, cartões azul-carvão, violeta/lavanda como cor primária, verde para confirmação, âmbar para avisos e vermelho somente para operações destrutivas. Os ecrãs devem ser adaptados a Material 3 e às boas práticas de acessibilidade. Não herdar métricas fictícias nem mensagens de limpeza impossível presentes em mockups. Não usar a imagem do protótipo como um fundo; reconstruir UI com widgets reais.

Criar ficheiros `docs/design-system.md`, `docs/navigation.md`, `docs/ui-state-matrix.md`, contendo:
- tokens de cor e contraste, tipografia, espaçamento, radius, elevação, estados, ícones e motion;
- componentes reutilizáveis: `StorageUsageCard`, `CategoryTile`, `FileListTile`, `FileThumbnail`, `RiskLabel`, `SelectionToolbar`, `AiConsentDialog`, `PermissionExplainer`, `ProgressIndicator`, `EmptyState`, `CleanupSummary`, `StatusBanner`, `AppStorageCard`;
- layouts para telas pequenas, texto grande, orientação portrait, modo offline e permissão parcial;
- acessibilidade: labels semânticos, área táctil, contraste, suporte a leitor de ecrã e texto escalado;
- dark theme como default e light theme opcional numa fase posterior; evitar cores hardcoded dispersas;
- microcopy PT-AO, sem promessas enganosas como «todos os ficheiros» ou «cache privada eliminada».

Mapear a navegação:
`Onboarding -> Permissões -> Dashboard -> Análise -> Resultados -> Categoria -> Pré-visualização/Seleção -> Revisão -> Confirmação do sistema -> Relatório`.
Vias laterais: `WhatsApp`, `Duplicados`, `Aplicações e cache`, `Classificação inteligente`, `Definições/Privacidade`.

Estados obrigatórios em todos os ecrãs de dados: `initial`, `loading`, `partial`, `success`, `empty`, `permission_denied`, `permission_limited`, `offline`, `error`, `retrying`, `cancelled`; usar explicitamente `unsupported` para capacidades não disponíveis.

---

# 3. PLATAFORMA E POLÍTICAS ANDROID: MATRIZ DE CAPACIDADES

**Antes de codificar, pesquisar a documentação oficial atual do Android e Play para a targetSdk escolhida.** Criar `docs/android-capability-matrix.md` por API level e por classe de permissão, com evidência, método e fallback. Não inventar acesso a diretórios privados.

## 3.1 Fontes de acesso

A) **MediaStore:** imagens, vídeos e áudio indexados que estejam disponíveis ao app, incluindo media partilhada de outras aplicações quando permissões apropriadas forem concedidas. Usar URIs de conteúdo, projection mínima, cursores paginados, filtros eficientes; para detetar alterações preferir generation/version quando disponíveis. Considerar `IS_TRASHED`, volumes, media parcialmente autorizada e providers.

B) **Storage Access Framework (SAF):** `ACTION_OPEN_DOCUMENT` e `ACTION_OPEN_DOCUMENT_TREE` com concessão explícita do utilizador; não presumir que o picker autoriza raízes ou `/Android/data` e `/Android/obb`. Persistir URI grants somente quando válido. O utilizador pode revogar acesso.

C) **Gestão alargada de armazenamento partilhado:** `MANAGE_EXTERNAL_STORAGE` é permissão especial (Android 11+) sujeita a política Play; só incluir se justificar concretamente a função principal de gestão de ficheiros. Criar uma variante/estratégia que funcione sem ela. Mesmo com ela, **não aceder aos diretórios internos privados nem `/Android/data` de outras apps**. Preparar declaração Play e evidência de necessidade. Não usar como atalho universal.

D) **Estatísticas das aplicações:** investigar `StorageStatsManager` e `PACKAGE_USAGE_STATS`/acesso especial para ver consumo por aplicação. Separar estimativas por categoria quando API permite; não apresentar estatísticas privadas como tamanho exato de ficheiros inspecionados. Verificar manifest, definições de utilização e limites de package visibility; não requerer `QUERY_ALL_PACKAGES` salvo caso realmente aprovado.

E) **Cache das aplicações:** estudar `StorageManager.ACTION_CLEAR_APP_CACHE` (limpeza de caches externas mediante autorização em plataformas suportadas) e atalhos para `ACTION_APPLICATION_DETAILS_SETTINGS`/gestão de armazenamento. O app não pode abrir e apagar seletivamente cache interna do Chrome. Diferenciar «limpar cache» de «limpar dados», «cookies», «downloads» e «apagar conta/sessão». Não usar botões como «limpar Chrome 3 GB» sem suporte real. Se só for possível abrir definições, escrever «Abrir definições do Chrome».

F) **Eliminação de multimédia:** priorizar `MediaStore.createTrashRequest` onde suportado; considerar `createDeleteRequest` só para fluxo explícito de eliminação permanente. As APIs normalmente solicitam confirmação do sistema para ficheiros de terceiros. A API recebe URI(s) MediaStore válidas, não caminhos genéricos. Gerir limites de lote por API level, resultados, falhas parciais, interrupção e repetição idempotente.

G) **Eliminação por SAF:** usar `DocumentsContract` / operações do provider quando houver `FLAG_SUPPORTS_DELETE` e autorização efetiva. Não assumir que todos os providers suportam reciclagem nem que é possível recuperar.

H) **Diretórios da própria app:** pode limpar cache que ela própria cria; evitar eliminar dados persistentes necessários. Seguir regras de ciclo de vida.

## 3.2 WhatsApp

Descobrir media acessível através de fontes efetivas, sem basear tudo num único caminho. Candidatos de localização modernos podem incluir `/storage/emulated/0/Android/media/com.whatsapp/WhatsApp/Media/` ou variantes de WhatsApp Business; **isto é um candidato de descoberta, não uma garantia universal**. `MediaStore` pode não indexar todos os elementos, especialmente locais ocultados por `.nomedia`. Em modo de acesso alargado, avaliar leitura de ficheiros partilhados acessíveis; fallback por seletor e seleção do utilizador quando tecnicamente permitido.

Não tentar ler bases de dados privadas, conversas, backups cifrados ou metadados de conversa que não estejam expostos. Não inferir incorretamente que uma imagem pertence a um contacto/conversa. Ao remover uma cópia da multimédia, alertar que poderá desaparecer de uma mensagem e não ser recuperável pelo WhatsApp.

## 3.3 Permissões e privacidade

Permissões sob demanda (*just-in-time*), com explicação de propósito, revogação e modo limitado. Verificar Android 13 `READ_MEDIA_IMAGES`, `READ_MEDIA_VIDEO`, `READ_MEDIA_AUDIO`, Android 14+ acesso visual parcial `READ_MEDIA_VISUAL_USER_SELECTED` e versões anteriores. Evitar solicitar permissões desnecessárias no onboarding. Testar combinações de autorizado/negado/revogação enquanto a app está aberta.

## 3.4 Matriz obrigatória

Para cada função registar: `Android min`, `Android target`, `permissão`, `API`, `necessita confirmação do sistema?`, `dados acessíveis`, `dados inacessíveis`, `fallback`, `publicável na Play?`, `evidência/URL`, `resultado em device físico`. Criar versões para Android 11, 13, 14, 15 e 16 (ajustar quando versões atuais mudarem).

---

# 4. FUNCIONALIDADES E REQUISITOS DE DOMÍNIO

Cada requisito deve ganhar ID (ex.: `FR-SCAN-001`) e cenários Given/When/Then. Criar uma matriz de rastreabilidade `Requisito -> UI -> Componente -> Teste -> Estado`.

## 4.1 Onboarding e permissões

- Apresentar proposta de valor; explicar diferença entre ficheiros partilhados e dados de outras apps.
- Permitir acesso parcial e recusa sem bloquear todo o app.
- Mostrar fontes já acessíveis e fontes não autorizadas; botão para rever permissões.
- Divulgar claramente que classificação Jev é online/opt-in, e mostrar exemplos do que seria enviado.
- Não exigir conta para uso local.

## 4.2 Dashboard

- Total, utilizado, disponível por volume, com unidades padronizadas e explicação dos dados; barras proporcionais.
- Separar armazenamento global do SO e inventário acessível do scanner — os valores podem não somar.
- Estado da última análise, fontes cobertas, categorias, total de duplicados exatos e potenciais candidatos (não «espaço recuperável» como facto).
- Ações: `Analisar`, `WhatsApp`, `Duplicados`, `Ficheiros grandes`, `Aplicações/cache`, `Inteligência`.
- Estado vazio e permissão parcial, carregamento progressivo e atualização imediata pós-limpeza.

## 4.3 Scanner e indexação

- Scanner inicial por fontes: MediaStore de cada volume; documentos/árvores SAF autorizados; extras conforme permissão; nunca varredura indiscriminada do sistema protegido.
- Metadados: identificador da fonte, URI/path apenas onde permitido, nome/label, tamanho, MIME, extensão, data modificação, data captura/adicionada se disponível, resolução/duração quando acessível, volume, pastas, flags e status de acesso.
- Não usar `atime` como verdade universal; campos ausentes devem ser `null`.
- Processamento incremental; mudanças por geração/assinatura quando suportado, invalidação, remoções externas, renomeações, volumes montados/desmontados.
- Cancelar/retomar, paginação, persistência por lote, backpressure, streaming e limitações de memória.
- Assinatura de conteúdo por hash quando necessário, não em todos os ficheiros na primeira análise.
- Filtrar por tipo, tamanho, data, origem, path permitido e estado.
- Nunca guardar thumbnail completa na base como BLOB sem necessidade; cache LRU com quotas e limpeza.
- Mostrar progresso real por fase, sem fingir precisão global quando não se conhece o universo de ficheiros.

## 4.4 Duplicados e semelhanças

- Duplicação exata: pré-agrupar por `size`, comparar hash rápido/partial para reduzir I/O, validar candidatos com hash completo criptográfico (SHA-256 ou equivalente), rever colissões/erros. **Nunca declarar duplicado exato por pHash ou nome**.
- Calcular bytes potencialmente poupáveis sem eliminar o único exemplar e sem contar duas vezes grupos diferentes.
- Criar «semelhantes», com hash perceptual/distance e análise de metadados; **não permitir ação automática baseada somente na semelhança**.
- Escolher «manter original» apenas como sugestão baseada em resolução, qualidade, data, localização e contexto; utilizador pode reverter seleção.
- Excluir ficheiros protegidos, ilegíveis, parciais e casos de assinaturas incompatíveis.

## 4.5 Ficheiros antigos, grandes e downloads

- Intervalos de tamanho e datas editáveis; critérios por *última modificação*, *data do conteúdo* ou *data de captura* claramente separados.
- Documentos financeiros antigos não devem ser removidos por idade somente.
- Regras padrão sensatas, filtros dinâmicos e favoritos; não impôr prazos fiscais/legais universais.
- Diferenciar ficheiros únicos, comprovativos, contratos, identificações e media pessoal.

## 4.6 WhatsApp/media

- Resumos por images/videos/audio/documents/voice notes/outros apenas onde acessíveis.
- Filtros por data, tamanho, duplicados, media grandes, origens; miniaturas e player/preview usando URIs.
- Seleção múltipla e revisão, aviso de efeito sobre histórico das mensagens; não afirmar que arquivo está descarregável novamente.
- Suportar WhatsApp Business como extensão posterior se a descoberta por fonte for tecnicamente validada.
- Apresentar claramente «multimédia encontrada» em vez de «todo o armazenamento WhatsApp».

## 4.7 Armazenamento das aplicações, Chrome e cache

- Estatísticas reais ou ausência explícita de acesso. Ícones, nomes e consumo quando sistema fornece.
- Etiquetar `cache reportada`, `dados da aplicação`, `media acessível`, `tamanho total do pacote`, conforme origem dos dados — não fundir nem duplicar.
- Chrome: botão `Abrir gestão de armazenamento` e, se a API e permissões permitirem, **solicitar** limpeza externa da cache pelo Android. Nunca chamar a cache interna de «ficheiros identificados» sem os ler.
- Explicar que eliminar cache pode aumentar tráfego/recarregamento e que «Limpar dados» pode terminar sessões/perder preferências.
- Se usar `ACTION_CLEAR_APP_CACHE`, informar sobre âmbito e possíveis impactos e não prometer «limpar só Chrome» se a ação for global.
- Criar `capabilityResolver` que escolhe operação: `ownCacheCleanup`, `systemCacheIntent`, `openAppStorageSettings`, `unsupported`.

## 4.8 Classificação Jev

- Categorias iniciais: `personal`, `commercial_product`, `financial_receipt`, `invoice`, `contract`, `identification`, `screenshot`, `educational`, `entertainment_meme`, `technical`, `other`, `unknown` e subcategorias. Não chamar a categoria uma ordem para eliminar.
- Entrada por ficheiro: metadados mínimos, texto OCR permitido, thumbnail/crop preparado quando utilizador autorizou envio de imagem.
- Saída tipada: categoria, subcategoria, distribuição/confiança, tags estruturadas limitadas, categoria sensível, motivos locais para mostrar ao utilizador e flag de revisão.
- `unknown` obrigatório; resposta inválida, ambígua ou baixa confiança => exibir `Não foi possível classificar` e não selecionar para eliminação.
- Classificar por lotes administrados por fila, com rate-limit, cancelamento, limites e resultados por ficheiro; permitir parar e continuar sem perda.
- Categorias com documentos financeiros e identificação: proteções e consentimento reforçado. OCR local por defeito, upload cloud com opt-in específico; evitar transmitir texto sensível quando não necessário.
- Sem linguagem livre complexa no MVP: filtros e políticas guiadas; no futuro linguagem natural como fase independente.
- Guardar a versão do modelo e da taxonomia; invalidar classificação ao mudar conteúdo ou regras.

## 4.9 Seleção, revisão e eliminação

- Seleção individual e múltipla; ações `Marcar`, `Desmarcar`, `Conservar/Proteger`, `Pré-visualizar`, `Mover para lixo quando suportado`, `Eliminar definitivamente quando explicitamente escolhido`.
- Apresentar sumário exato com quantidade, tamanhos, proteção e método/recuperabilidade, por origem.
- Excluir automaticamente da seleção, por padrão, ficheiros protegidos e inacessíveis; não fazer seleção automática de imagens pessoais com base no Jev.
- Reconfirmar que URI ainda existe, permissões ainda válidas e assinatura não mudou; aplicar operação com transação lógica, estados por item e idempotência.
- Pedir confirmação nativa do Android para ficheiros MediaStore; tratar `DENIED`, `CANCELLED`, `NOT_FOUND`, `PERMISSION_REVOKED`, `PARTIAL_SUCCESS`, `PROVIDER_FAILURE`.
- Fazer operações por lote de tamanho suportado pela versão do Android (não assumir lotes ilimitados; verificar limites atuais). Não inventar sucesso antes do callback real.
- Registar resultado sem preservar dados sensíveis desnecessários; oferecer histórico e opções reais de restauro conforme mecanismo.

## 4.10 Relatórios e histórico

- Distinguir `selecionados`, `enviados para lixo`, `eliminados definitivamente`, `falhados`, `ainda ocupando espaço`, `diferença observada na capacidade livre`.
- O sistema pode libertar espaço por outras razões durante a operação; apresentar medidas com cuidado sem atribuição falsa.
- Relatórios por sessão, contagem de ficheiros, categoria e duração, sem telemetria remota por defeito.
- Função «Limpar histórico local» sem apagar ficheiros do utilizador.

## 4.11 Definições

- Idioma/tema; exclusões por pasta/categoria/fonte; políticas de proteção; consentimento Jev; limites de dados e uso; ligação somente Wi-Fi (configurável); limpar thumbnails/index local; permissões e explicação; versão; suporte e diagnóstico sem dados privados.
- Configurar não enviar pastas como `Documents/IDs`, `Banking` etc. por defeito? Não assumir nomes padronizados: oferecer categorias sensíveis e exclusões definidas pelo utilizador.

---

# 5. MODELO DE DOMÍNIO E CONTRATOS

Definir entidades imutáveis/tipadas (exemplo conceptual; afinar ao projeto):

```text
StorageVolume(id, displayName, totalBytes?, freeBytes?, kind, mounted)
StorageSource(id, kind: mediaStore|saf|allFiles|appOwn, rootUri?, permissionState, coverage)
IndexedFile(id, sourceId, contentUri, displayName?, relativePath?, mime?,
            sizeBytes?, modifiedAt?, capturedAt?, addedAt?, durationMs?,
            width?, height?, mediaKind, fingerprint?, protected, accessStatus,
            scanGeneration, contentSignature)
ScanSession(id, startedAt, finishedAt?, phase, status, seenCount, indexedCount,
            errorCount, bytesSeen, coverageSummary, cancellationReason?)
DuplicateGroup(id, kind: exact|similar, fileIds, candidateSavingsBytes, verified)
Classification(id, fileId, taxonomyVersion, modelId, inputFingerprint,
               category, subcategory?, confidence?, probabilities?,
               sensitiveFlag, source: local|jev|manual, status, analyzedAt)
UserProtectionRule(id, categoryOrSourceOrUri, scope, reason?, enabled)
CleanupPlan(id, createdAt, entries[], selectedBytes, validationState)
CleanupEntry(id, fileId, action, status, errorCode?, attemptedAt?, completedAt?)
CleanupReport(id, countsByStatus, selectedBytes, trashedBytes,
              deletedBytes, observedFreeBytesDelta?, generatedAt)
AiConsent(version, allowedMediaKinds, allowedSources, wifiOnly, enabled, updatedAt)
AiUsage(id, model, count, inputTokens?, billedAmount?, latencyMs?, failure?)
AppStorageInfo(packageName, label?, totalBytes?, dataBytes?, cacheBytes?,
               source, measuredAt?, capability)
```

Regras transversais:
- IDs internos não equivalem a caminhos permanentes.
- URIs podem expirar; documentos podem mudar enquanto a app está aberta.
- Metadados podem estar ausentes/incoerentes.
- Fonte de dados e nível de confiança da fonte devem ser guardados.
- Nenhum dado no índice significa «ficheiro removido» sem revalidação.
- Nunca colocar `byte[]` de imagens grandes no estado global Flutter.
- `coverage` não é um percentual inventado: pode ser `full for granted source`, `partial`, `unknown`.
- Manter adaptadores entre modelos de domínio e tipos específicos de Android.

Criar esquema SQLite/Drift com índices de consultas: `sourceId`, `mime`, `mediaKind`, `sizeBytes`, `modifiedAt`, `contentSignature`, `duplicateGroupId`, `classificationCategory`, `scanGeneration`. Planejar migrations, integridade, exclusões e tamanho máximo do índice. Não guardar ficheiros binários de utilizador na base.

---

# 6. STACK, ARQUITETURA E ESTRUTURA DE PASTAS

## 6.1 Stack definida

- **Flutter/Dart**: interface e regras de apresentação, com Material 3; usar versão estável atual e registar lockfiles.
- **Kotlin nativo Android**: `MediaStore`, `ContentResolver`, `StorageStatsManager`, `StorageManager`, `DocumentsContract`, intent results e métodos de acesso realmente disponíveis.
- **Pigeon** (se compatível com o SDK atual): contratos tipados Flutter ↔ Kotlin. Documentar porque foi escolhido em vez de MethodChannel cru.
- **Riverpod**: estado/DI (avaliar versão e APIs da versão estável).
- **go_router** (ou equivalente justificado): navegação.
- **Drift/SQLite**: inventário persistente e queries.
- **ML Kit Text Recognition**: OCR no dispositivo quando compatível; comparar overhead de modelos, download de recursos, desempenho offline e minSdk; avaliar alternativa apenas se necessário.
- **Hashing local**: SHA-256; pHash/dHash opcional para imagens semelhantes; processamento de vídeos em fases futuras.
- **Backend gateway mínimo**: Cloudflare Workers (ou infraestrutura equivalente fundamentada) para autenticar chamadas, proteger a chave da TypeSafe, validar limites de requisições, orçamento, consentimento e rate-limit. Não presumir que o gateway suporta facilmente multipart sem verificar.
- **Jev/TypeSafe**: único fornecedor comercial de IA inicial, usando os contratos oficiais da API atual.
- **Testes**: `flutter_test`, mocks, integration tests, testes Kotlin/JUnit/instrumented, testes por API level e fixtures sintéticas.

Não instalar dependências desnecessárias. Documentar compatibilidade de licenças e atualidade de packages com fontes verificáveis.

## 6.2 Arquitetura sugerida

`Presentation -> Application/Use Cases -> Domain -> Repositories -> Data Sources(Android Native, SQLite, Gateway)`.

Criar interfaces desacopladas:

```text
StorageInventoryPort
StorageCapabilityPort
FileMetadataPort
FileContentReaderPort
FingerprintPort
DuplicateDetectorPort
TextExtractionPort
SemanticClassifierPort
ClassificationRepository
CleanupOperationPort
ApplicationStorageStatsPort
ConsentRepository
AiBudgetPort
```

O `SemanticClassifierPort` tem `LocalFallbackClassifier` e `JevClassifier` (o fallback pode ser um classificador por regras, não outro LLM). O `CleanupOperationPort` só recebe `CleanupPlan` validado, nunca comandos arbitrários do Jev.

## 6.3 Estrutura sugerida

```text
lib/
  app/{app.dart,router/,theme/,di/,l10n/}
  core/{errors/,result/,logging/,permissions/,utils/,widgets/}
  domain/{entities/,enums/,repositories/,usecases/,policies/}
  data/{local/drift/,platform/,remote/jev/,mappers/,repositories/}
  features/
    onboarding/
    dashboard/
    scanner/
    explorer/
    whatsapp/
    app_storage/
    duplicates/
    classification/
    review/
    cleanup/
    history/
    settings/
android/app/src/main/kotlin/<application_id>/
  storage/{MediaStoreInventory,DocumentTreeInventory,StorageStatsReader,...}
  cleanup/{MediaStoreTrashAdapter,DocumentDeleteAdapter,OwnCacheCleaner,...}
  permissions/
  bridge/
pigeons/
backend/jev-gateway/
docs/
design/references/
openspec/{config.yaml,specs/,changes/}
test/
integration_test/
```

Adaptar ao repositório existente e às convenções do ecossistema. Não criar diretórios vazios sem necessidade. Ficheiros gerados de Pigeon devem manter fonte, instruções e reproducibilidade.

---

# 7. INTEGRAÇÃO JEV — ESPECIFICAÇÃO TÉCNICA

**Fonte de verdade obrigatória:** https://docs.typesafe.ai/llms.txt , https://docs.typesafe.ai/introduction , https://docs.typesafe.ai/images.md e https://docs.typesafe.ai/primitives . Não usar endpoints de sites de terceiros sem comparar com TypeSafe oficial. Verificar modelo atual, endpoint, headers, limite de contexto, rate limits e schemas antes de implementar.

O Jev trabalha com `state` e `questions` tipadas (`choice`, `score`, `noul`, possivelmente `box` para visão). **Não exigir geração livre de texto nem inventar uma resposta JSON não fornecida pela API.** Converter respostas reais para o nosso `Classification` no gateway/adaptador.

### 7.1 Contrato lógico da classificação

Input do nosso app (sanitizado):

```json
{
  "requestId": "uuid",
  "fileRef": "opaque-local-id",
  "metadata": {
    "mime": "image/jpeg",
    "sizeBytes": 8420000,
    "modifiedAt": "2024-07-06T10:00:00Z",
    "sourceKind": "whatsapp_media"
  },
  "consentScope": "image_category",
  "taxonomyVersion": "1.0.0",
  "image": "optional JPEG/PNG binary after consent",
  "ocrText": "optional minimized text after consent"
}
```

Output normalizado da aplicação:

```json
{
  "fileRef": "opaque-local-id",
  "status": "classified",
  "category": "commercial_product",
  "subcategory": "home_appliance",
  "confidence": 0.91,
  "sensitive": false,
  "suggestedView": "review",
  "modelId": "jev-1.14.0",
  "taxonomyVersion": "1.0.0",
  "reasonCodes": ["PRODUCT_IMAGE"],
  "deleteAuthorized": false
}
```

**ATENÇÃO:** estes são contratos internos ilustrativos. O agente deve gerar requests Jev usando **o schema oficial atual** e mapear os resultados, sem enviar `image` como string literal nem inventar campos na API. `confidence` é um valor de classificação, não garantia de segurança; não usar um threshold arbitrário sem benchmark. `reasonCodes` gerados por regras/estruturas verificáveis; não inventar explicações pseudo-precisas.

### 7.2 Perguntas Jev

- `category`: Choice fechado, com `unknown/none_of_above`.
- `is_sensitive`: Noul, pergunta objetiva sobre presença de dados pessoais/financeiros.
- `document_kind`: Choice quando a categoria de documentos já foi identificada.
- `visual_type`: Choice para fotos de produtos, screenshots, fotografias pessoais, memes, imagens pouco claras.
- `document_date_present`: Noul; extrair datas por OCR e **validar em código**, não usar Jev como parser de datas crítico.
- `confidence_routing`: decisão de rever quando baixa confiança ou conflito entre classificadores.
- Não pedir ao Jev «o utilizador pode apagar este ficheiro?» como autoridade final.

O modelo Jev visual é experimental: `TypeSafe-Danger-Zone-Images: true` quando suportado, envio PNG/JPEG, processamento de orientação e dimensões, limites oficiais, preferir multipart. A funcionalidade deve operar por **opt-in**, suportar `experimental_disabled`, ser desligável por feature flag remota/local e ter fallback local.

### 7.3 Privacidade e custo

- Não transmitir nome/caminho completo, contactos, localizações EXIF, ID de conta ou conteúdo não necessário.
- Antes de envio de imagem, mostrar pré-visualização/tipo e opção de autorização por lote; categorias sensíveis requerem uma explicação adequada.
- Sem consentimento, não fazer upload, nem silenciosamente.
- Gateway autenticado sem conta obrigatória no MVP: desenhar abordagem antifraude realista (por exemplo, sessão anónima, validação de integridade quando disponível, rate-limit e quotas) e apontar abusos possíveis. Não aceitar um proxy público aberto que permita consumir a chave por terceiros.
- Ter quota por utilizador/dispositivo/sessão e trava global de custos. Não pressupor preço ou custo por imagem sem ler documentação atual.
- Tempo limite, retry com exponential backoff e jitter para códigos temporários, cancelamento, tratamento de 429, 4xx e 5xx; nunca retry infinito.
- Não manter imagens no gateway; logs sem corpo/payload. Retenção pelo fornecedor documentada e divulgada antes de beta.
- Cache no dispositivo por fingerprint + versão modelo + taxonomia + consentimento pertinente, com política de expiração e eliminação local.
- API key apenas em secrets do gateway, `.env.example` sem chaves reais, `.gitignore`, nunca no APK.

### 7.4 Dataset e avaliação

Criar benchmark com conteúdo sintético ou cedido com consentimento: comprovativos, faturas, contratos, produtos, fotos pessoais, capturas de ecrã, memes, documentos ambíguos, datas erradas, imagens desfocadas, múltiplos documentos. Medir precision/recall/F1 por classe, confusões, erro de sensibilidade, confiança, latência, custo e comportamento offline. **Não lançar visão em produção sem relatório de validação.** O objetivo da IA é melhorar triagem, não autorizar operações destrutivas.

---

# 8. PERFORMANCE, ROBUSTEZ E BATERIA

Condições alvo: Android com pouca RAM, dezenas a centenas de milhares de itens acessíveis, armazenamento quase cheio, microficheiros e rede intermitente. Definir métricas e medir em device real; metas abaixo são hipóteses de benchmark, não garantias.

- Primeiro dashboard responsivo após abrir a app sem aguardar indexação completa.
- Consulta paginada e processamento em streams/lotes; não materializar todos os ficheiros na RAM.
- Scan incremental por fontes; checksum apenas em candidatos; controlo de concorrência e backpressure.
- Evitar ler conteúdo completo de vídeos/ficheiros grandes para apenas calcular tamanho.
- Não guardar original na cache; thumbnails LRU sob quota (ex.: <=100 MB configurável), com atenção ao armazenamento quase cheio.
- Apoiar pausa/cancelar/reiniciar; recuperar após encerramento inesperado e revogação de permissões.
- WorkManager apenas onde justificado; respeitar limites de tarefas em background e foreground services atuais. Não criar um serviço sempre ativo para gerir ficheiros.
- Operações críticas e hashes fora da UI thread; isolar I/O nativo, não fazer processamento CPU intensivo em build de widgets.
- Eventos `low_storage`, `permission_revoked`, `volume_removed`, `file_changed`, `wifi_unavailable`, `api_unavailable` tratados sem crash.
- Medir tempo inicial de scan, tempo incremental, memória pico, CPU, bateria, armazenamento extra criado pelo app, custo/latência Jev e taxa de conclusão das tarefas.
- Propor metas a validar em 3 classes: entrada, média e alta. Nunca usar 100% cobertura como meta quando Android não a permite.

---

# 9. SECURITY, PRIVACY, CONFIANÇA E CONFORMIDADE

Criar `docs/threat-model.md`, `docs/privacy-data-flow.md`, `docs/play-compliance.md`, `docs/security-checklist.md`.

### 9.1 Threat model mínimo

Riscos: vazamento de dados pessoais para API, log/analytics com nomes/paths, exposição de API key, upload involuntário, classificações incorretas, eliminação não intencional, troca de URI (TOCTOU), app maliciosa a fornecer content URI, ZIP bombas/documentos corrompidos, leitura fora do escopo concedido, consumo de bateria e custos elevados, abuso de gateway, permissões revogadas.

Mitigações: least privilege, validação da fonte URI, tamanho/MIME, streaming seguro, proteção contra paths arbitrários, revalidação pré-limpeza, consentimento granular, logs mínimos, retenção mínima, encrypt-at-rest apenas se justificável, timeouts, rate limiting, quotas, segregation, atualização de dependências, testes de revogação, reports explicativos.

### 9.2 Conduta de eliminação

- Jev não chama APIs de eliminação.
- `CleanupCoordinator` recebe apenas seleção explícita e validada. Trazer o resumo de operação e confirmação nativa quando exigida.
- Preferir reciclagem quando possível. O app não pode prometer undo quando o provider não o suporta.
- Validação obrigatória `exists`, `permission`, `signature`, `protected` e operação suportada.
- Nenhum `File(path).delete(recursive: true)` em diretórios desconhecidos/partilhados sem API, permissão e confirmação apropriadas.
- Não tratar pontuações Jev de «baixo risco» como autorização automática.

### 9.3 Play Store

- Preparar justificação de `MANAGE_EXTERNAL_STORAGE` somente se necessária à função central e caso `MediaStore`/SAF sejam insuficientes; documentar e criar fallback de permissões.
- Verificar política atual de acesso fotos/vídeos, visibilidade de apps, uso de dados e Data Safety.
- Preparar texto honesto de descrição, privacy policy, disclosures e procedimento de remoção de dados próprios.
- Identificar se release de APK externa e build Play terão diferenças de permissões; não promover trabalho que burle políticas.

---

# 10. UX COMPLETA E CRITÉRIOS POR ECRÃ

Para cada ecrã criar user stories e acceptance tests.

### T01 Boas-vindas

Explicar valor em 2 frases. CTA `Começar`. Acesso às políticas de privacidade. Sem promessas universais de limpeza.

### T02 Dashboard

Donut e cartões com dados reais; se a app só vê parte do dispositivo, mostrar nota de cobertura. Estados: sem análise, análise antiga, parcial, completo da fonte, offline. CTA `Analisar agora`.

### T03 Análise

Etapa corrente (descoberta, metadados, duplicados, indexação), itens processados e progresso demonstravelmente calculável. `Pausar`, `Cancelar`, preservar resultados parciais, não obrigar a Internet.

### T04 Resultados

Categorias: ficheiros grandes, duplicados, antigos, WhatsApp e aplicações/cache. Separar «sugestões» de «potencialmente recuperável». Cada categoria navegável para lista real.

### T05 WhatsApp

Media acessível por tipo, filtro tamanho/data, miniaturas, multi-select. Alerta discreto: «Ficheiros removidos podem deixar de estar disponíveis nas conversas».

### T06 Duplicados

Grupos visuais; marcador de cópias idênticas versus semelhantes; selecionar cópias mantendo pelo menos um item. Não incluir itens em outras operações sem validação.

### T07 Jev

Consentimento antes de upload; estado experimental; categorias e contagens; mostrar número classificado e pendente; aplicar filtros por categoria; permitir corrigir categoria manualmente; falhas não bloqueiam classificação local.

### T08 Aplicações/cache

Lista de aplicações e consumo observado quando disponível; botão Chrome -> definições oficiais; explicar cache de terceiros e as limitações; nunca fingir que a limpeza foi feita quando só abriu Settings.

### T09 Revisão

Lista de ficheiros selecionados, miniaturas, tamanho total, riscos, protegidos, método de exclusão e botão de ação claro. Confirmação adicional para irreversível; tratamento de cancelamento sem estado inconsistente.

### T10 Relatório

Só após resposta do sistema: o que foi enviado para lixo, apagado, falhou e a variação de espaço observada. Distinguir «podes recuperar» de «já recuperaste». Opções `Ver histórico` e `Voltar ao painel`.

Extras a especificar: `Permission Center`, `Error/Offline states`, `File Detail`, `Settings` e `Privacy Center` como fluxos secundários. Não acrescentar screens desnecessárias se puder ser um bottom sheet bem feito.

---

# 11. TESTES, QUALIDADE E MATRIZ DE VALIDAÇÃO

Criar plano de testes que inclua:

**Unitários (Dart):** regras de classificação locais, filtragem, sumarização de bytes, agrupamento exato, seleção por lote, cache de classificação, regras de proteção, estados da UI, conversões de datas/unidades, falhas API.

**Unitários/instrumentados Kotlin:** queries MediaStore, projection, SAF URI grants, estatísticas e limitações, tentativas de remoção, confirmação do sistema, permissões revogadas, volumes montados/desmontados, aparelho sem WhatsApp, app package não encontrado.

**Integração:** scanner -> índice -> results -> seleção -> validação -> `createTrashRequest` (com ficheiros de teste), `DocumentsContract` quando suportado; gateway Jev em mock com respostas reais gravadas sem conteúdo pessoal; error/429/offline.

**Visual/UX:** ecrãs em pequeno aparelho, grande textScale, dark mode, leitores de ecrã, seleção de 1000 itens, reentrância, landscape básico, estado de carregamento/erro.

**Benchmark:** 1k, 10k, 100k e (se possível) 1M ficheiros sintéticos, sem presumir que o dispositivo suporta todas as configurações. Documentar memória, CPU, tempo, bateria, writes, thumbnails; testar dispositivo de entrada.

**Segurança:** URI adulterado, path traversal, duplicado TOCTOU, revogação de permissão antes de delete, content provider falho, zip/ficheiro enorme, logs sem dados sensíveis, nenhum upload sem consentimento.

**Matriz API levels:** Android 11, 13, 14, 15, 16 (ou versões equivalentes suportadas à data). Indicar em CI quais são emuladores e quais são dispositivos físicos. Não afirmar hardware testado se não ocorreu.

**Cobertura**: definir metas de código razoáveis por camada, mas priorizar teste de regras destrutivas e fluxos críticos; documentar partes ainda não automatizadas.

Critérios de aceitação transversais:
- Zero eliminações sem autorização válida e explícita.
- Duplicados exatos confirmados por hash completo.
- Falha da API Jev não bloqueia análise/limpeza local.
- Estimativas e resultados reais claramente rotulados.
- UI não crasha ao mudar/revogar permissões ou desconectar volumes.
- Build Android de debug compila; analyzer/linter/test suite passam.
- Funcionalidades sem suporte comunicadas honestamente.
- Não existem chaves no repositório, APK nem logs.
- Estado `done` apenas com verificação executada, não por intenção.

---

# 12. PLANO DE ENTREGA E ORDEM DE IMPLEMENTAÇÃO

Dividir em OpenSpec changes pequenos com IDs rastreáveis, propostas, specs, design e tasks próprios. **Não criar uma alteração monolítica que implemente tudo sem checkpoints.** Sugerir (ajustar conforme dependências):

1. `00-foundation-openspec-architecture`: stack, convenções, arquitetura, UI tokens e setup.
2. `01-android-storage-feasibility-poc`: **gate obrigatório**: ler storage geral, query MediaStore, identificar media acessível WhatsApp, testar operações de teste com consentimento, settings Chrome; gerar `docs/poc-findings.md` e lacunas.
3. `02-permissions-and-index`: solicitações, capabilities, scanner, drift, cancel/resume.
4. `03-dashboard-and-explorer`: UI real para armazenamento, categorias e filtros.
5. `04-duplicates-and-large-files`: hash streaming, grupos e filtros.
6. `05-whatsapp-media-manager`: media acessível, filtros e preview.
7. `06-app-storage-and-cache`: stats e fluxos do SO; cache própria.
8. `07-cleanup-selection-and-reporting`: revisão, trash/delete flows, resultado/verificação.
9. `08-local-ocr-and-document-taxonomy`: OCR local e proteções.
10. `09-jev-beta-gateway-classifier`: API, opt-in, benchmark, quotas, telemetria técnica mínima.
11. `10-hardening-accessibility-play-release`: testes físicos, permissões/policy, otimização, docs de entrega.

Para cada change:
- Problema e valor de negócio.
- Escopo, não objetivos, dependências, riscos e rollback.
- Requisitos MUST/SHOULD/COULD, eventos e cenários Given/When/Then.
- Modelo/contrato alterado, API Android e versões suportadas.
- Componentes UI envolvidos.
- Testes unitários/instrumentados/manuais e evidência.
- Critérios de aceite e tarefas pequenas (idealmente verificáveis em uma sessão).
- Atualização dos docs e marcação correta da task.

**Gate da PoC:** antes de construir a experiência final de WhatsApp/Chrome, testar pelo menos num aparelho físico as capacidades abaixo (deixar `PENDING_DEVICE_VALIDATION` se só emulador):
- `MediaStore` listar ficheiros concedidos e os metadados reais.
- Ficheiro de teste na origem WhatsApp-like (sem manipular dados reais) ser descoberto se acessível.
- `MediaStore.createTrashRequest` apresentar confirmação e reportar resultado correto.
- Acesso à página oficial de storage do Chrome funcionando quando instalado.
- Verificação de `MANAGE_EXTERNAL_STORAGE` e restrições em Android atual sem pretender acesso privado.

---

# 13. OPENSPEC — PROCEDIMENTO EXATO PARA ESTE AGENTE

1. Inspecionar `git status`, `README`, manifest, pubspec, Gradle, CI, estrutura e referências visuais. Não destruir trabalho existente.
2. Verificar versão oficial atual de OpenSpec e sintaxe dos skills instalados no agente. Se não estiver inicializado, configurar `openspec init` usando a ferramenta atual (ex.: Claude/Codex) e criar `openspec/config.yaml` com `schema: spec-driven` e contexto de projeto. **Não assumir comandos antigos `/opsx:*` quando o OpenSpec instalado usar `/openspec-*`.** O CLI imprime o nome correto para cada ferramenta.
3. `openspec/config.yaml` deve conter regras para:
   - `proposal`: objetivos, não objetivos, dependências, riscos, rollback e prova de viabilidade;
   - `specs`: requisitos normativos com cenários verificáveis Given/When/Then, limitação Android e privacidade;
   - `design`: arquitetura, contratos, concorrência, sequências, dados, segurança, UI e alternativas rejeitadas;
   - `tasks`: tarefas pequenas com testes, critérios de aceite e dependências.
4. Usar a ação *propose* da versão instalada para gerar os artefactos de uma alteração base de **plano do produto + PoC**. Se necessário, criar mudanças específicas por domínio; preservar intenção de raiz em `docs/product-vision.md` e ADRs. Não inventar um formato OpenSpec inválido: verificar schema de artefactos gerado pela versão real.
5. Criar especificações canonizadas em `openspec/specs/` apenas pelo processo suportado de sync/archive quando for oportuno. Usar change specs e evitar duplicações contraditórias.
6. Criar os documentos necessários em `docs/`, e interligar ficheiros de referência visual. Não criar centenas de ficheiros minúsculos sem valor, mas cobrir integralmente os requisitos.
7. Verificar artefactos com comandos de validação aplicáveis à versão instalada; atualizar até não existirem erros de schema ou requisitos órfãos.
8. **Nesta primeira passagem, NÃO executar `apply` do projeto inteiro**. Entregar relatório de prontidão, dependências, riscos, lista de change IDs e ordem de execução. Solicitar apenas a aprovação da fase PoC se o utilizador ainda não tiver autorizado implementação.
9. Depois da autorização, usar a ação de *apply* correspondente e trabalhar um change de cada vez. Fazer test/build/lint e atualizar tasks; não marcar testes físicos como concluídos sem dispositivo.
10. Manter `docs/DECISIONS.md` com data, contexto, opção adotada, alternativas rejeitadas e impacto.

## Configuração OpenSpec sugerida (adaptar a schema real)

```yaml
schema: spec-driven
context: |
  Projeto LIVRE: app Android-first em Flutter + Kotlin, SQLite/Drift,
  Riverpod, OCR local e classificação opcional com Jev/TypeSafe.
  pt-AO na UI e documentação. A IA classifica; o utilizador decide.
  Sem root, sem acesso a dados privados de outras apps, sem exclusão
  automática de ficheiros pessoais. Online AI é opt-in e experimental.
  Implementação incremental, testes e validação de políticas Google Play.
rules:
  proposal:
    - Explicitar objetivos, não objetivos, dependências, riscos e rollback.
    - Identificar provas de viabilidade Android antes de prometer funcionalidades.
  specs:
    - Usar cenários Given/When/Then com resultados observáveis.
    - Incluir estados de permissão, offline, erro, parcial e cancelamento.
    - Nenhuma eliminação sem seleção/autorização explícita do utilizador.
  design:
    - Documentar fluxo de dados Flutter-Kotlin, gateway Jev e persistência.
    - Explicar APIs Android por versão e estratégia de fallback.
    - Incluir segurança, privacidade, desempenho e diagramas de sequência.
  tasks:
    - Dividir em tarefas pequenas e testáveis com dependências.
    - Marcar done apenas após evidência de implementação/teste.
```

---

# 14. FORMATO OBRIGATÓRIO DA PRIMEIRA RESPOSTA DO AGENTE

Depois de analisar o ambiente e criar a estrutura OpenSpec, responder com:

1. **Diagnóstico do repositório**: o que existe, ferramentas/versões disponíveis, gaps e ficheiros de design encontrados.
2. **Estado OpenSpec**: versão, comando/skill disponível, árvore de ficheiros criados e validação executada.
3. **Resumo da arquitetura**: escolhas técnicas, dependências, plataforma Android mínima/target e motivos.
4. **Matriz de viabilidade**: o que está confirmado por documentação, o que foi comprovado em testes e o que está pendente; WhatsApp e Chrome explicitamente.
5. **Lista de mudanças** com escopo, prioridade, dependências e gate da PoC.
6. **Riscos e decisões abertas**: selecionar no máximo as 5 decisões com impacto crítico, propor valor default e justificar.
7. **Cobertura de requisitos**: checklist das áreas deste prompt e lacunas por resolver.
8. **Comandos seguintes** no OpenSpec realmente instalado (não inventar sintaxe).
9. **Estado da implementação**: não começou / PoC em andamento / testado, conforme o que realmente foi feito.

Não substituir os artefactos locais por um resumo de chat. O output principal são os ficheiros de especificação no repositório. Não pedir dados já fornecidos neste prompt. Se o repositório estiver vazio, criar a fundação documental e declarar o que falta para implementar; não ficar bloqueado em perguntas genéricas.

---

# 15. CRITÉRIO FINAL DE PRODUTO

O LIVRE é bem-sucedido se um utilizador comum puder, num telefone Android, abrir a aplicação, descobrir conteúdos autorizados que ocupam espaço, entender o que são, agrupá-los e selecionar com confiança o que pretende remover — sem depender de acesso impossível aos dados privados de outras aplicações, sem uploads não autorizados, sem uma IA a apagar ficheiros por conta própria e com relatórios honestos sobre o resultado.

**Priorizar produto verdadeiro e seguro, não uma demonstração enganadora.**

---

# 16. REFERÊNCIAS A CONSULTAR E MANTER ATUALIZADAS

- OpenSpec oficial: https://github.com/Fission-AI/OpenSpec e https://openspec.dev/docs/
- Android storage overview: https://developer.android.com/training/data-storage
- MediaStore: https://developer.android.com/training/data-storage/shared/media
- SAF: https://developer.android.com/training/data-storage/shared/documents-files
- Manage all files: https://developer.android.com/training/data-storage/manage-all-files
- StorageStatsManager: https://developer.android.com/reference/android/app/usage/StorageStatsManager
- StorageManager: https://developer.android.com/reference/android/os/storage/StorageManager
- API MediaStore trash: https://developer.android.com/reference/android/provider/MediaStore
- Play policy All files: https://support.google.com/googleplay/android-developer/answer/10467955
- TypeSafe docs index: https://docs.typesafe.ai/llms.txt
- TypeSafe Jev: https://docs.typesafe.ai/introduction
- TypeSafe images experimental: https://docs.typesafe.ai/images.md
- Flutter: https://docs.flutter.dev/
- Flutter platform channels: https://docs.flutter.dev/platform-integration/platform-channels

**FIM DO PROMPT MESTRE.**
