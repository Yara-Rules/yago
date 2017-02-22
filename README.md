```
 /{}     /{}         /{}{}{}
|  {}   /{}/        /{}__  {}
 \  {} /{}//{}{}{} | {}  \__/  /{}{}{}
  \  {}{}/|____  {}| {} /{}{} /{}__  {}
   \  {}/  /{}{}{}{| {}|_  {}| {}  \ {}
    | {}  /{}__  {}| {}  \ {}| {}  | {}
    | {} |  {}{}{}}|  {}{}{}/|  {}{}{}/
    |__/  \_______/ \______/  \______/



```


YaGo is a translation tool which converts Yara rules in JSON format so they could be handeled easyly with a NoSQL database, for example.

The way that YaGo works its really easy, you can just call de tool by giving a Yara rule as an argument or you can import the modules on your project and use it on your way.

YaGo is written in [Golang](https://golang.org/) so it can run on lots of platforms (see [GOOS and GOARCH](https://golang.org/doc/install/source#environment)). We have provide a Makefile which build YaGo on the most common platforms, in addition we provide those binaries as releases [here](https://github.com/Yara-Rules/yago/releases) at Github.com.

# Running YaGo
As sed before YaGo is a command line tool, but you can use it on your projects by using the modules.

## Command line
When you run YaGo without any argument it returns a help message like this one:

```
./build/yago
  -cwd string
      CWD from the Yara rules will be imported (default ".")
  -dirName string
      Directory with a set of yara rules
  -fileName string
      Yara file you want to parse
  -format string
      Format you want the Yara rule to be parsed (default "json")
  -indexFile string
      Yara index file
```

At the moment YaGo is supporting three input modes. It can parse a sigle Yara rule file, a directory of Yara rule files, and index Yara file.

You can call YaGo using the next flags:
* `-fileName` which points a Yara rule.
* `-dirName` which points a directory with Yara rule files.
* `-indexFile` which points a Yara index file.

```
./build/yago -fileName test/EK_Fragus.yar
```

When parsing a index file with YaGo you can provide a Current Working Directory (CWD) path and it will be attached at the begining of each rule path when reading the file.


```
$ cat index.yar
include "rules/EK_Angler.yar"
include "rules/EK_Blackhole.yar"
```

```
./build/yago -indexFile index.yar -cwd path/with/rules
```

YaGo will look for rules at `path/with/rules/rules/`.

The last option is the flag `-format`. At the moment YaGo is only supporting two formats: `json`, and `bson`. By default YaGol will use `json`.


A json output looks like

```
{
  "file_name": "rule.yar",
  "imports": null,
  "rules": [
    {
      "name": "Mal\_PotPlayer\_DLL",
      "global": false,
      "private": false,
      "tags": [
        "dll"
      ],
      "meta": {
        "author": "Florian Roth",
        "date": "2016-05-25",
        "description": "Detects a malicious PotPlayer.dll",
        "hash1": "705409bc11fb45fa3c4e2fa9dd35af7d4613e52a713d9c6ea6bc4baff49aa74a",
        "reference": "https://goo.gl/13Wgy1",
        "score": "70"
      },
      "strings": [
        {
          "name": "$x1",
          "value": "C:\\\\Users\\\\john\\\\Desktop\\\\PotPlayer\\\\Release\\\\PotPlayer.pdb",
          "modifers": [
            "fullword",
            "ascii"
          ]
        },
        {
          "name": "$s3",
          "value": "PotPlayer.dll",
          "modifers": [
            "fullword",
            "ascii"
          ]
        },
        {
          "name": "$s4",
          "value": "\\\\update.dat",
          "modifers": [
            "fullword",
            "ascii"
          ]
        }
      ],
      "condition": "uint16 ( 0 ) = = 0x5a4d and filesize < 200KB and $x1 or all of ( $s * )"
    }
  ]
}
```

## Module import
On the other hand, if you would like to use YaGo on your own project it is as easy as adding the following line in the import section.

```
import (
    "github.com/Yara-Rules/yago/parser"
)
```

Onece imported you can parse Yara rules by instantiating a new parser.

```
p := parser.New("FileName")
```

And parse the Yara rule by calling `Parse` method and giving the rules as string.

```
p.Parse(string(file))
```

Puting all together

```
filePath := "path/to/yara/rule"

file, err := ioutil.ReadFile(filePath)
if err != nil {
    log.Fatal(err)
}

p := parser.New(path.Base(filePath))

p.Parse(string(file))
j, err := json.Marshal(p)
if err == nil {
    os.Stdout.Write(j)
}
```

# Contribute
If you would like to be part of the Yara comunity or Yara-Rules project you are free to contribute with us in any maner. Yo can send issues o pull requests, by sharing Yara rules, etc.

# Changelog
Version: **0.1.0**
Initial release.

# License
YaGo, tool to translate Yara rules into JSON format.
Copyright (C) 2017 Jaume Martín <jaumemartin@protonmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
