# Navegação — LIVRE

## Fluxo principal
Onboarding → Permissões → Dashboard → Análise → Resultados → Categoria → Pré-visualização/Seleção → Revisão → Confirmação do sistema → Relatório

## Vias laterais (bottom sheets ou ecrãs dedicados)
- WhatsApp (media acessível)
- Duplicados (grupos exatos/semelhantes)
- Aplicações e cache (estatísticas + intents do SO)
- Classificação inteligente (Jev beta, opt-in)
- Definições/Privacidade (idioma, tema, exclusões, consentimento)

## Rotas go_router (sugestão)
- `/` — Dashboard
- `/onboarding` — Boas-vindas
- `/permissions` — Centro de permissões
- `/scan` — Análise (progresso por fase)
- `/results` — Resultados (categorias)
- `/results/:category` — Lista de ficheiros por categoria
- `/whatsapp` — Media WhatsApp
- `/duplicates` — Grupos de duplicados
- `/apps` — Aplicações e cache
- `/classification` — Classificação Jev
- `/review` — Revisão da seleção
- `/cleanup-report` — Relatório pós-operação
- `/settings` — Definições

## Estados por ecrã
Todos os ecrãs de dados devem renderizar: `initial`, `loading`, `partial`, `success`, `empty`, `permission_denied`, `permission_limited`, `offline`, `error`, `retrying`, `cancelled`, `unsupported`.
