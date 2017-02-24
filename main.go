package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/Yara-Rules/yago/parser"
)

const (
	MULTILINE = `\s*/\*([^*]|\*+[^*/])*\*+/\s*`
	INLINE    = `(?m)\s*//.*[\n\r][\n\r]?`
	BLANKS    = `(?m)\s+$`
	QUOTES    = `"`
)

func processFile(fileName *string) []*parser.Parser {
	file, err := ioutil.ReadFile(*fileName)
	if err != nil {
		log.Fatal(err)
	}

	p := parser.New(path.Base(*fileName))
	p.SetLogLevel("INFO")
	p.Parse(string(file))

	var res []*parser.Parser
	res = append(res, p)
	return res
}

func processDir(dirName *string) []*parser.Parser {
	var res []*parser.Parser
	fileList := []string{}
	filepath.Walk(*dirName, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})

	for _, filePath := range fileList {
		file, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}
		p := parser.New(path.Base(filePath))
		p.Parse(string(file))
		res = append(res, p)
	}
	return res
}

func processIndex(indexFile, indexCwd *string) []*parser.Parser {
	var res []*parser.Parser
	file, err := ioutil.ReadFile(*indexFile)
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(MULTILINE)
	index := re.ReplaceAllString(string(file), "")

	re = regexp.MustCompile(INLINE)
	index = re.ReplaceAllString(index, "")
	index = re.ReplaceAllString(index, "")

	re = regexp.MustCompile(BLANKS)
	index = re.ReplaceAllString(index, "")

	re = regexp.MustCompile(QUOTES)
	index = re.ReplaceAllString(index, "")

	lines := strings.Split(index, "\n")
	for _, value := range lines {
		ruleFile := strings.Split(value, " ")
		if len(ruleFile) == 2 {
			rulePath := path.Join(*indexCwd, ruleFile[1])
			if _, err := os.Stat(rulePath); err == nil {
				file, err := ioutil.ReadFile(rulePath)
				if err != nil {
					log.Fatal(err)
				}
				p := parser.New(path.Base(rulePath))
				p.Parse(string(file))
				res = append(res, p)
			}
		}
	}

	return res
}

func main() {

	fileName := flag.String("fileName", "", "Yara file you want to parse")
	dirName := flag.String("dirName", "", "Directory with a set of yara rules")
	indexFile := flag.String("indexFile", "", "Yara index file")
	indexCwd := flag.String("cwd", ".", "CWD from the Yara rules will be imported")
	format := flag.String("format", "json", "Format you want the Yara rule to be parsed")
	flag.Parse()

	if len(*fileName) == 0 && len(*dirName) == 0 && len(*indexFile) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if len(*fileName) != 0 && (len(*dirName) != 0 || len(*indexFile) != 0) ||
		len(*dirName) != 0 && (len(*fileName) != 0 || len(*indexFile) != 0) ||
		len(*indexFile) != 0 && (len(*fileName) != 0 || len(*dirName) != 0) {
		os.Stderr.WriteString("ERR: You must provide only one of fileName, dirName or indexFile.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var p []*parser.Parser

	if len(*fileName) != 0 {
		p = processFile(fileName)
	} else if len(*dirName) != 0 {
		p = processDir(dirName)
	} else if len(*indexFile) != 0 {
		p = processIndex(indexFile, indexCwd)
	}

	for _, r := range p {
		if *format == "json" {
			j, err := json.Marshal(r)
			if err == nil {
				os.Stdout.Write(j)
			}
		} else if *format == "bson" {
			j, err := bson.Marshal(r)
			if err == nil {
				os.Stdout.Write(j)
			}
		}
	}
	return
}
