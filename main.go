package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"

	"gopkg.in/mgo.v2/bson"

	"github.com/Yara-Rules/yago/parser"
)

func main() {

	fileName := flag.String("fileName", "", "Yara file you want to parse")
	format := flag.String("format", "json", "Format you want the Yara rule to be parsed")
	flag.Parse()

	if len(*fileName) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	p := parser.New(path.Base(*fileName))
	p.SetLogLevel("INFO")

	p.Parse(string(file))
	if *format == "json" {
		j, err := json.Marshal(p)
		if err == nil {
			os.Stdout.Write(j)
			return
		}
	} else if *format == "bson" {
		j, err := bson.Marshal(p)
		if err == nil {
			os.Stdout.Write(j)
			return
		}
	}
}
