import 'package:flutter/material.dart';
import 'package:flutter/foundation.dart' show debugDefaultTargetPlatformOverride;
import 'app/random.dart';

// mobile, web version
// void main() => runApp(new RandomWordsApp());

// desktop version
void main() {
  // https://github.com/flutter/flutter/wiki/Desktop-shells#target-platform-override
  // debugDefaultTargetPlatformOverride = TargetPlatform.fuchsia;

  runApp(new RandomWordsApp());
}
