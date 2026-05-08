import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';

final Dio client = () {
  final client = Dio();

  // 日志打印
  if (!kReleaseMode) {
    client.interceptors.add(
      LogInterceptor(requestBody: true, responseBody: true, error: true),
    );
  }

  return client;
}();

Future<T> request<T>(
  String method,
  String url, {
  Map<String, dynamic>? queries,
  (bool, dynamic) Function(Response<dynamic>)? respHandler,
  String? contentType,
  Map<String, dynamic>? headers,
  Object? data,
}) async {
  final response = await client.request(
    url,
    options: Options(
      method: method.toUpperCase(),
      contentType: contentType,
      headers: headers,
    ),
    queryParameters: queries,
    data: data,
  );
  if (response.statusCode != 200) {
    throw Exception(
      'http error, code=${response.statusCode}, msg=${response.data}',
    );
  }
  if (respHandler != null) {
    final (ok, respData) = respHandler(response);
    if (ok) return respData;
  }
  return response.data;
}
