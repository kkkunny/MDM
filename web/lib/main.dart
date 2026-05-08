import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:provider/provider.dart';
import 'package:mdm/providers/task_provider.dart';
import 'package:mdm/pages/overall_page.dart';
import 'package:mdm/configs/theme.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => TaskProvider(),
      child: GetMaterialApp(
        title: 'Download Manager',
        debugShowCheckedModeBanner: false,
        theme: AppTheme.lightTheme,
        home: const OverallPage(),
      ),
    );
  }
}
