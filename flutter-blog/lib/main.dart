import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

void main() => runApp(HZBlogApp());

class HZBlogApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'firecat.zhou',
      theme: ThemeData(
        primarySwatch: Colors.teal,
      ),
      home: Scaffold(
        body: Center(
          child: Column(
            mainAxisSize: MainAxisSize.max,
            mainAxisAlignment: MainAxisAlignment.center,
            children: createRowList(),
          ),
        ),
      ),
    );
  }
}

List<Widget> createRowList() {
  return [
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Flexible(
          child: Container(
            padding: EdgeInsets.all(5),
            child: CircleAvatar(
              backgroundImage: AssetImage(
                "assets/images/avatar.jpg",
              ),
              radius: 65.0,
            ),
          ),
        ),
      ],
    ),
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Flexible(
          child: Container(
            padding: EdgeInsets.all(10),
            child: Text(
              "Howard Zhou",
              style: TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 30,
              ),
            ),
          ),
        ),
      ],
    ),
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Flexible(
          child: Container(
            padding: EdgeInsets.all(10),
            child: Text(
              "Software Engineer",
              style: TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 20,
              ),
            ),
          ),
        ),
      ],
    ),
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Container(
          padding: EdgeInsets.all(5),
          child: new InkWell(
              child: new Icon(FontAwesomeIcons.file),
              onTap: () => launch(
                  'https://docs.google.com/document/d/1-s6AhGv09r_OF26UL5pKU5mvQngVoDux0O8Ty-t2b_Y/edit?usp=sharing')),
        ),
        Container(
          padding: EdgeInsets.all(5),
          child: new InkWell(
              child: new Icon(FontAwesomeIcons.linkedin),
              onTap: () => launch('https://www.linkedin.com/in/firecatzhou/')),
        ),
        Container(
          padding: EdgeInsets.all(5),
          child: new InkWell(
              child: new Icon(FontAwesomeIcons.github),
              onTap: () => launch('https://github.com/lovemew67')),
        ),
        Container(
          padding: EdgeInsets.all(5),
          child: new InkWell(
              child: new Icon(FontAwesomeIcons.envelope),
              onTap: () => launch('mailto:firecat.zhou@gmail.com')),
        ),
      ],
    ),
  ];
}
