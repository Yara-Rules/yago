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


YaGo is a translation tool which converts Yara rules in JSON format so they could be handled easyly with a NoSQL database, for example.

The way that YaGo works it's really easy, you can just call it by giving a Yara rule as an argument or you can import the modules on your project and use it on your way.

YaGo is written in [Golang](https://golang.org/) so it can run on lots of platforms (see [GOOS and GOARCH](https://golang.org/doc/install/source#environment)). We have provided a Makefile which builds YaGo for the most common platforms, in addition we also provide those binaries as releases [here](https://github.com/Yara-Rules/yago/releases) at Github.com.

# Running YaGo
As it was said before YaGo is a command line tool, but you can use it on your projects by using the modules.

## Command line
When you run YaGo without any argument it returns a help message like this one:

```
YaGo - Parsing Yara rules like a Gopher.

Usage:
  yago fileName <fileName> [ --validJSON ]
  yago dirName <dirName> [ --validJSON ]
  yago indexFile <indexFile> [ cwd <path> ] [ --validJSON ]
  yago inputFile <inputFile> outputDir <outputDir> [ --overwrite ] [ --validJSON ]
  yago inputFile <inputFile> outputFile <outputFile> [ --overwrite ] [ --validJSON ]
  yago -h | --help
  yago --version
```

At the moment YaGo is supporting four input modes. It can parse a single Yara rule file, a directory of Yara rule files, index Yara file, and JSON file.

You can call YaGo by using the next arguments:
* `fileName` which points a Yara rule.
* `dirName` which points a directory with Yara rule files.
* `indexFile` which points a Yara index file.
* `inputFile` which points a JSON file with Yara rules previously parsed.

```
./build/yago fileName ./test/EK_Fragus.yar
```

When parsing an index file with YaGo you can provide the Current Working Directory (CWD) path and that path will be attached at the beginning of each rule path when reading the file.


```
$ cat index.yar
include "rules/EK_Angler.yar"
include "rules/EK_Blackhole.yar"
```

```
./build/yago indexFile index.yar cwd path/with/rules
```

YaGo will look for rules at `path/with/rules/rules/....yar`.

The last argument is `inputFile` that converts rules in JSON format that were previously translated back in Yara rules. This arguments accept two extra arguments which indicate the output is either a directory or file, in case of a file YaGO will merge all rules taking care of import and rule name collitions.

In addition the `inputFile` argument has an `--overwrite` option that overwrite exisitng files on the output directory or file.

Finally, all arguments have a `--validJSON` option. That option tells YaGo to either print out each rule in one line or print out the whole rule set in a file that meets JSON format.

---

A json output looks like

```
{
  "file_name": "rule.yar",
  "imports": null,
  "rules": [
    {
      "name": "Mal_PotPlayer_DLL",
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
      "condition": "uint16 ( 0 ) = = 0x5a4d and filesize < 200KB and $x1 or all of ( $s* )"
    }
  ]
}
```

## Module import
On the other hand, if you would like to use YaGo on your own project, it is as easy as adding the following line in the import section.

```
import (
    "github.com/Yara-Rules/yago/yago"
)
```

You can process the rules using the either parser API or the YaGo's one.

Using parser API

```
package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"

  "github.com/Yara-Rules/yago/yago"
)

func main() {
  file, _ := ioutil.ReadFile("test.yar")

  p := yago.NewParser("InformationName")
  p.SetLogLevel("INFO") // Optionl. Accepts: INFO, WARN, DEBUG
  p.Parse(string(file))

  j, err := json.Marshal(p)
  if err == nil {
    os.Stdout.Write(j)
    os.Stdout.WriteString("\n") // Useful when importing the result in other tools
  } else {
    fmt.Println(err)
  }
}

```

On the other hand, you can use the YaGo API.

```
package main

import "github.com/Yara-Rules/yago/yago"

func main() {
  r := yago.ProcessFile("test.yar")
  yago.GenerateOutputFromYara(r, false)

}

```

# Contribute
If you would like to be part of the Yara comunity or Yara-Rules project you are free to contribute with us in any way. You can send issues or pull requests, by sharing Yara rules, etc.

