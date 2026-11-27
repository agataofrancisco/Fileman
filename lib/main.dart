import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'app/app.dart';

void main() {
  runApp(const LivreApp());
}

class LivreApp extends StatelessWidget {
  const LivreApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'LIVRE',
      localizationsDelegates: const [
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
      ],
      supportedLocales: const [Locale('pt', 'AO'), Locale('pt', 'PT'), Locale('en', 'US')],
      locale: const Locale('pt', 'AO'),
      themeMode: ThemeMode.dark,
      darkTheme: AppTheme.dark,
      home: const AppShell(),
      debugShowCheckedModeBanner: false,
    );
  }
}

class AppShell extends StatelessWidget {
  const AppShell({super.key});

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Center(child: Text('LIVRE — scaffolding inicial. Em construção.')),
    );
  }
}
