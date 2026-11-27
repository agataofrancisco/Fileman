# Matriz de Estados de UI — LIVRE

Cada ecrã de dados implementa explicitamente estes estados:

| Estado | Significado | Ação do utilizador |
|--------|-------------|--------------------|
| `initial` | Sem dados ainda, primeira abertura | `Analisar agora` / `Conceder permissões` |
| `loading` | Indexação/scanner em curso | `Pausar`, `Cancelar` (resultados parciais preservados) |
| `partial` | Fonte parcialmente autorizada ou análise incompleta | Revisar fontes cobertas, rever permissões |
| `success` | Análise concluída para a(s) fonte(s) | Navegar para resultados |
| `empty` | Nenhum ficheiro encontrado na fonte | Mensagem orientadora, não inventar dados |
| `permission_denied` | Permissão negada | Botão revisar permissões, modo limitado |
| `permission_limited` | Acesso visual parcial / árvore parcial | Mostrar o que está coberto e o que não está |
| `offline` | Sem rede (Jev indisponível) | Operações locais continuam; Jev desativado com aviso |
| `error` | Falha scanner/API/IO | `Tentar novamente`, detalhes sem dados sensíveis |
| `retrying` | Tentativa automática em curso (backoff) | Cancelar opcional |
| `cancelled` | Operação cancelada pelo utilizador | Estado limpo, sem dados inconsistentes |
| `unsupported` | Capacidade não disponível no dispositivo/versão | Mensagem honesta, alternativa se existir |

Microcopy: sem promessas enganosas. «Unsupported» comunica honestamente a limitação.
