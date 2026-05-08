import 'package:flutter/material.dart';
import 'package:mdm/configs/theme.dart';
import 'package:mdm/providers/task_provider.dart';
import 'package:mdm/widgets/overview_panel.dart';
import 'package:mdm/widgets/task_list_panel.dart';
import 'package:provider/provider.dart';

class OverallPage extends StatefulWidget {
  const OverallPage({super.key});

  @override
  State<OverallPage> createState() => _OverallPageState();
}

class _OverallPageState extends State<OverallPage> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      context.read<TaskProvider>().initialize();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: kLightBackground,
      body: Row(
        crossAxisAlignment: CrossAxisAlignment.stretch,
        children: [
          const Expanded(flex: 1, child: OverviewPanel()),
          Container(width: 1, color: kLightDivider),
          const Expanded(flex: 5, child: TaskListPanel()),
        ],
      ),
    );
  }
}
