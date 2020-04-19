import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

void main() => runApp(HZBlogApp());

class HZBlogApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Howard Zhou - Backend Developer',
      theme: ThemeData(
        primarySwatch: Colors.teal,
      ),
      home: Scaffold(
        body: ListView(
          children: createList(),
        ),
      ),
    );
  }
}

List<Widget> createList() {
  return [
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Image.asset(
          'assets/images/background.jpg',
          height: 500.0
        ),
      ],
    ),
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Container(
          padding: EdgeInsets.all(5),
          child: CircleAvatar(
            backgroundImage: AssetImage(
              "assets/images/avatar.jpg",
            ),
            radius: 65.0,
          ),
        ),
      ],
    ),
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Container(
          padding: EdgeInsets.all(10),
          child: Text(
            "Howard Zhou",
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 30,
            ),
          ),
        )
      ],
    ),
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Container(
          padding: EdgeInsets.all(10),
          child: Text(
            "Backend Engineer",
            style: TextStyle(
              fontWeight: FontWeight.bold,
              fontSize: 20,
            ),
          ),
        )
      ],
    ),
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Container(
          padding: EdgeInsets.all(10),
          child: Text(
            "Graduated from Information Management, National Taiwan University, specialized on Big Data, Cloud Computing, Information and Cloud Security and Virtualization, etc.\n\nPassionate to evolve in the development of Internet of Things and third-party payment in Taiwan.",
            style: TextStyle(
              fontSize: 15,
            ),
          ),
          constraints: BoxConstraints(
            maxWidth: 320.0,
          ),
        )
      ],
    ),
    Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        Container(
          padding: EdgeInsets.all(5),
          child: new InkWell(
              child: new Icon(FontAwesomeIcons.envelope),
              onTap: () => launch('mailto:firecat.zhou@gmail.com')
          ),
        ),
        Container(
          padding: EdgeInsets.all(5),
          child: new InkWell(
              child: new Icon(FontAwesomeIcons.linkedin),
              onTap: () => launch('https://www.linkedin.com/in/firecatzhou/')
          ),
        ),
        Container(
          padding: EdgeInsets.all(5),
          child: new InkWell(
              child: new Icon(FontAwesomeIcons.github),
              onTap: () => launch('https://github.com/lovemew67')
          ),
        ),
      ],
    ),
  ];
}
