# go-zimrss
RSS feed generator for zim notebooks

## Features

It generates a complete RSS feed file, based on the input notebook path and parameters given to the program:

```xml
<?xml version="1.0" encoding="UTF-8"?><rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/">
  <channel>
    <title>Luis GG - Tales and Notes</title>
    <link>https://luisgg.me</link>
    <description>Personal notes and posts on software development, technology and miscellaneous</description>
    <managingEditor>example@gmail.com (Luis Gabriel Gomez)</managingEditor>
    <pubDate>Wed, 23 Dec 2020 21:19:19 -0300</pubDate>
    <item>
      <title>About Me</title>
      <link>https://luisgg.me/About_Me.html</link>
      <description></description>
      <author>Luis Gabriel Gomez</author>
      <pubDate>Wed, 23 Dec 2020 22:19:19 -0300</pubDate>
    </item>
    <item>
      <title>Memory: A Forgotten Topic</title>
      <link>https://luisgg.me/Software_Development/3_-_Articles/Memory,_a_Forgotten_Topic.html</link>
      <description></description>
      <author>Luis Gabriel Gomez</author>
      <pubDate>Wed, 23 Dec 2020 23:09:11 -0300</pubDate>
    </item>
    <item>
      <title>Advanced Generics in .NET</title>
      <link>https://luisgg.me/Software_Development/2_-_langs/1_-_dotNET/Advanced_Generics_in_.NET.html</link>
      <description></description>
      <author>Luis Gabriel Gomez</author>
      <pubDate>Sat, 26 Dec 2020 18:36:45 -0300</pubDate>
    </item>
  </channel>
</rss>
```

It currently lacks descriptions and content (as zim does not handle these as metadata) but once the feed is generated, it can be edited to include these fields if needed

## Usage

Dependencies: go 1.15 or later

Be done in 3 steps

#### 1. Clone this project
```shell
git clone git@github.com:lggomez/go-zimrss.git && cd go-zimrss
```

#### 2. Build
```shell
cd main && go build
```

#### 3. Execute
Example based on my personal site:

```shell
./main -n /home/razael/DEV/lggomez.wiki.zim/Site -t "Luis GG - Tales and Notes" -l "https://luisgg.me/#index" -d "Personal notes and posts on software development, technology and miscellaneous" -a "Luis Gabriel Gomez" -e "example@gmail.com"
```

This will generate the desired `rss.xml` file on the current directory
