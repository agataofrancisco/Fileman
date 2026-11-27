import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

/// Definição das rotas da aplicação LIVRE.
///
/// Fluxo principal:
/// Onboarding -> Permissões -> Dashboard -> Análise -> Resultados -> Categoria
/// -> Pré-visualização/Seleção -> Revisão -> Confirmação do sistema -> Relatório.
///
/// Vias laterais: WhatsApp, Duplicados, Aplicações e cache,
/// Classificação inteligente, Definições/Privacidade.
final GoRouter livreRouter = GoRouter(
  initialLocation: '/',
  routes: [
    GoRoute(
      path: '/',
      builder: (context, state) => const _PlaceholderScreen('Dashboard'),
    ),
    GoRoute(
      path: '/onboarding',
      builder: (context, state) => const _PlaceholderScreen('Boas-vindas'),
    ),
    GoRoute(
      path: '/permissions',
      builder: (context, state) => const _PlaceholderScreen('Permissões'),
    ),
    GoRoute(
      path: '/scan',
      builder: (context, state) => const _PlaceholderScreen('Análise'),
    ),
    GoRoute(
      path: '/results',
      builder: (context, state) => const _PlaceholderScreen('Resultados'),
    ),
    GoRoute(
      path: '/results/:category',
      builder: (context, state) => const _PlaceholderScreen('Categoria'),
    ),
    GoRoute(
      path: '/whatsapp',
      builder: (context, state) => const _PlaceholderScreen('WhatsApp'),
    ),
    GoRoute(
      path: '/duplicates',
      builder: (context, state) => const _PlaceholderScreen('Duplicados'),
    ),
    GoRoute(
      path: '/apps',
      builder: (context, state) => const _PlaceholderScreen('Aplicações e cache'),
    ),
    GoRoute(
      path: '/classification',
      builder: (context, state) => const _PlaceholderScreen('Classificação'),
    ),
    GoRoute(
      path: '/review',
      builder: (context, state) => const _PlaceholderScreen('Revisão'),
    ),
    GoRoute(
      path: '/cleanup-report',
      builder: (context, state) => const _PlaceholderScreen('Relatório'),
    ),
    GoRoute(
      path: '/settings',
      builder: (context, state) => const _PlaceholderScreen('Definições'),
    ),
  ],
);

class _PlaceholderScreen extends StatelessWidget {
  final String title;
  const _PlaceholderScreen(this.title);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: Text(title)),
      body: Center(child: Text('$title — em implementação.')),
    );
  }
}
