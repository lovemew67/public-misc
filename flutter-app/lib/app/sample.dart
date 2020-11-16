import 'package:flutter/material.dart';

class SampleApp extends StatelessWidget {
  Widget build(BuildContext context) {
    return new Container(
      child: new Center(
        child: new Text(
          'Hello World',
          textDirection: TextDirection.ltr,
          style: new TextStyle(fontSize: 40.0, color: Colors.black87),
        ),
      ),
    );
  }
}
