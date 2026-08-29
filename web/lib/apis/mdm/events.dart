import 'dart:async';
import 'dart:convert';

import 'package:http/http.dart' as http;

import 'package:mdm/configs/api.dart';
import 'package:mdm/models/vo/task.pb.dart';

/// 订阅服务端 SSE 任务推送，流结束或出错时由调用方负责重连
Stream<TaskEvent> subscribeTaskEvents() async* {
  final client = http.Client();
  try {
    final request = http.Request(
      'GET',
      Uri.parse('${getApiBaseUrl()}/api/task/events'),
    )..headers['Accept'] = 'text/event-stream';

    final response = await client.send(request);
    if (response.statusCode != 200) {
      throw Exception('sse error, code=${response.statusCode}');
    }

    String buffer = '';
    await for (final chunk in response.stream
        .transform(const Utf8Decoder())
        .transform(const LineSplitter())) {
      if (chunk.isEmpty) {
        // 空行表示事件结束
        if (buffer.isEmpty) continue;
        final payload = buffer.startsWith('data: ')
            ? buffer.substring(6)
            : buffer.replaceFirst('data:', '');
        buffer = '';
        if (payload.isEmpty) continue;
        yield TaskEvent.create()..mergeFromProto3Json(jsonDecode(payload));
      } else if (chunk.startsWith(':')) {
        // 心跳注释，忽略
        continue;
      } else {
        buffer = buffer.isEmpty ? chunk : '$buffer\n$chunk';
      }
    }
  } finally {
    client.close();
  }
}