# Conformidade Google Play — LIVRE

## MANAGE_EXTERNAL_STORAGE
- Permissão especial (Android 11+), sujeita a política Play.
- Incluir apenas se justificar concretamente a função principal de gestão de ficheiros e MediaStore/SAF forem insuficientes.
- Criar variante/estratégia que funcione sem ela.
- Mesmo com a permissão, **não aceder a diretórios internos privados nem `/Android/data` de outras apps**.
- Preparar declaração Play e evidência de necessidade. Não usar como atalho universal.

## Políticas a verificar
- Acesso a fotos/vídeos (foto e vídeo do utilizador).
- Visibilidade de aplicações (package visibility; não requerer `QUERY_ALL_PACKAGES` salvo caso aprovado).
- Uso de dados e secção Data Safety do console Play.
- Descrição honesta, privacy policy, disclosures e procedimento de remoção de dados próprios.

## APK externo vs Play
- Identificar diferenças de permissões entre release APK externa e build Play.
- Não promover trabalho que burle políticas do Google Play.

## Declarações necessárias
- Justificativa de `MANAGE_EXTERNAL_STORAGE` (se usada).
- Textos de privacidade em pt-AO.
- Consentimento Jev opt-in divulgado.
- Não prometer limpeza impossível (ex.: cache privada do Chrome, base de dados do WhatsApp).
