# Design System — LIVRE

Identidade provisória: **LIVRE**. Dark-first por defeito. Light theme opcional numa fase posterior.

## Tokens de cor (pt-AO microcopy, sem promessas enganosas)

| Token | Uso | Cor (dark) | Contraste |
|-------|-----|------------|-----------|
| `background` | Fundo ecrãs | Quase preto / navy (`#0B0F19`) | — |
| `surface` | Cartões | Azul-carvão (`#151B2A`) | ≥ 4.5:1 vs texto |
| `primary` | Violeta/lavanda primária | `#8B5CF6` | ≥ 4.5:1 texto branco |
| `primary_container` | Estados selecionados | `#3B0764` | — |
| `success` | Confirmação, recuperável | Verde (`#22C55E`) | — |
| `warning` | Avisos, cache | Âmbar (`#F59E0B`) | — |
| `error` | Apenas operações destrutivas | Vermelho (`#EF4444`) | — |
| `text_primary` | Texto principal | `#FFFFFF` | — |
| `text_secondary` | Texto secundário | `#94A3B8` | ≥ 3:1 |
| `text_muted` | Legendos, cobertura | `#64748B` | ≥ 3:1 |
| `border` | Divisórias | `#1E293B` | — |
| `danger_destructive` | Botão eliminação irreversível | Vermelho apenas no botão de ação | — |

## Tipografia
- Família: Inter (regular 400, bold 700). Texto escalado respeitado.
- Títulos ecrã: `headlineMedium` / `headlineSmall`.
- Labels cartão: `bodyLarge` ou `titleMedium`.
- Microcopy estados: `bodyMedium`.

## Espaçamento, radius, elevação
- Espaçamento base: 8px grid (4, 8, 12, 16, 20, 24, 32).
- Radius cartões: 12px (médio), 16px (ecrã cheio).
- Elevação: 0 (cartões superfície), 1–2 (tooltips/bottom sheets), 4 (dialogs).

## Estados obrigatórios em ecrãs de dados
`initial`, `loading`, `partial`, `success`, `empty`, `permission_denied`, `permission_limited`, `offline`, `error`, `retrying`, `cancelled`, `unsupported` (capacidade não disponível).

## Componentes reutilizáveis
- `StorageUsageCard` — total/utilizado/disponível por volume, barra proporcional.
- `CategoryTile` — categoria, contagem, bytes, ícone, navegação.
- `FileListTile` — miniatura, nome, tamanho, data, fonte, estado de proteção.
- `FileThumbnail` — miniatura via URI, cache LRU, sem BLOB na base.
- `RiskLabel` — etiqueta de sensibilidade/proteção (semântica, leitor de ecrã).
- `SelectionToolbar` — marcar/desmarcar, conservar/proteger, pré-visualizar, ação destrutiva.
- `AiConsentDialog` — consentimento Jev, opt-in, exemplo do que é enviado.
- `PermissionExplainer` — explicação just-in-time, revogação, modo limitado.
- `ProgressIndicator` — progresso real por fase, sem precisão global fingida.
- `EmptyState` — vazio, sem análise, sem dados da fonte.
- `CleanupSummary` — quantidade, tamanhos, proteção, método, recuperabilidade.
- `StatusBanner` — offline, permissão parcial, erro, Jev indisponível.
- `AppStorageCard` — app, ícone, consumo observado, fonte dos dados, botão de ação.

## Acessibilidade
- Labels semânticos em todos os elementos interativos.
- Área táctil mínima 48×48 dp.
- Contraste WCAG AA mínimo.
- Suporte a leitor de ecrã (TalkBack).
- Texto escalado (font scale).
- Modo offline e permissão parcial comunicados de forma clara.

## Microcopy PT-AO
Não usar promessas enganosas como «todos os ficheiros» ou «cache privada eliminada». Usar «ficheiros a que tenho acesso», «cache reportada», «espaço observado».

## Navegação (mapeamento)
`Onboarding -> Permissões -> Dashboard -> Análise -> Resultados -> Categoria -> Pré-visualização/Seleção -> Revisão -> Confirmação do sistema -> Relatório`.
Vias laterais: `WhatsApp`, `Duplicados`, `Aplicações e cache`, `Classificação inteligente`, `Definições/Privacidade`.
