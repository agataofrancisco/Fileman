# Visão do Produto — LIVRE

## Problema
Telefones Android de utilização prolongada acumulam dezenas de milhares de imagens, vídeos, áudios, documentos, downloads e outros conteúdos. A maioria dos utilizadores não consegue perceber facilmente de onde vem o consumo de armazenamento nem decidir, em lote e com contexto, quais os ficheiros que pretende conservar.

## Proposta de valor
«Descobre o que está a ocupar espaço, entende o que tens e decide rapidamente o que manter ou remover.»

## Público prioritário
Utilizadores de Android, inicialmente em Angola, incluindo equipamentos de entrada (4–8 GB RAM ou menos), armazenamento limitado, conectividade cara/intermitente e uso prolongado de WhatsApp.

## Princípios inegociáveis
1. O utilizador decide o destino dos seus ficheiros pessoais. Jev classifica, nunca executa eliminação.
2. Só propor «limpeza rápida» para operações claramente conhecidas e autorizadas (cache da própria app, temporários que o app criou, ou mecanismo oficial do sistema com confirmação).
3. Sem root, sem ADB oculto, sem exploits, sem APIs privadas e sem abusar de AccessibilityService.
4. Trabalhar com capacidades reais, dependentes de versão do Android, fabricante, permissões e URI/provider.
5. Offline-first, privacy-first, least privilege.
6. Classificar com Jev como único modelo comercial online inicial. OCR local e algoritmos determinísticos não são «outros LLMs».
7. Visão do Jev é experimental, beta controlada por feature flag; falha de IA não pode bloquear o scanner nem a limpeza manual.
8. A IA e os metadados não são prova de inutilidade. Criar estado `unknown`, tratamento de ambiguidade e regras de proteção.
9. Nada de métricas ou valores fictícios apresentados como reais.
10. Nada de UI bonita com botões decorativos: cada ação tem estado e fluxo reais.

## Objetivos do MVP (M1–M8)
Consultar secção 1.4 do PROMPT_MESTRE_OPENSPEC_LIVRE.md.

## Fora do MVP
Consultar secção 1.5 do PROMPT_MESTRE_OPENSPEC_LIVRE.md.

## Critério final de produto
Um utilizador comum pode, num telefone Android, abrir a aplicação, descobrir conteúdos autorizados que ocupam espaço, entender o que são, agrupá-los e selecionar com confiança o que pretende remover — sem depender de acesso impossível aos dados privados de outras aplicações, sem uploads não autorizados, sem uma IA a apagar ficheiros por conta própria e com relatórios honestos sobre o resultado.
