package yago

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Yara-Rules/yago/grammar"
)

const (
	MULTILINE = `\s*/\*([^*]|\*+[^*/])*\*+/\s*`
	INLINE    = `(?m)\s*//.*[\n\r][\n\r]?`
	BLANKS    = `(?m)\s+$`
	QUOTES    = `"`
	MAXBUFF   = 1024 * 1024 // If needed Go will take it form RAM.
)

func NewParser(name string) *grammar.Parser {
	return grammar.New(name)
}

func ProcessFile(fileName string) []*grammar.Parser {
	file, err := ioutil.ReadFile(fileName)
	checkErr(err)

	p := NewParser(path.Base(fileName))
	p.SetLogLevel("INFO")
	p.Parse(string(file))

	var res []*grammar.Parser
	res = append(res, p)
	return res
}

func ProcessDir(dirName string) []*grammar.Parser {
	var res []*grammar.Parser
	fileList := []string{}
	filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})

	for _, filePath := range fileList {
		file, err := ioutil.ReadFile(filePath)
		checkErr(err)

		p := NewParser(path.Base(filePath))
		p.SetLogLevel("INFO")
		p.Parse(string(file))
		res = append(res, p)
	}
	return res
}

func ProcessIndex(indexFile, cwd string) []*grammar.Parser {
	var res []*grammar.Parser
	file, err := ioutil.ReadFile(indexFile)
	checkErr(err)

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
	for _, line := range lines {
		ruleFile := strings.Split(line, " ")
		if len(ruleFile) == 2 {
			rulePath := path.Join(cwd, ruleFile[1])
			if _, err := os.Stat(rulePath); err == nil {

				file, err := ioutil.ReadFile(rulePath)
				checkErr(err)

				p := NewParser(path.Base(rulePath))
				p.SetLogLevel("INFO")
				p.Parse(string(file))
				res = append(res, p)
			} else {
				os.Stdout.WriteString(fmt.Sprintf("WARNING: Rule file %s does not exist. Check the cwd argument.\n", rulePath))
			}
		}
	}
	return res
}

func ProcessInputFile(inputFile string, validJSON bool) []*grammar.Parser {
	var res []*grammar.Parser
	if validJSON {
		file, err := ioutil.ReadFile(inputFile)
		checkErr(err)

		jc := &jsonCloak{}
		err = json.Unmarshal(file, jc)
		if err != nil {
			printError(err)
		}
		for _, r := range jc.Ruleset {
			res = append(res, r)
		}
	} else {
		file, err := os.Open(inputFile)
		checkErr(err)
		defer file.Close()

		var buff []byte

		scanner := bufio.NewScanner(file)
		scanner.Buffer(buff, MAXBUFF)

		var rules *grammar.Parser
		for scanner.Scan() {
			rules = &grammar.Parser{}
			err = json.Unmarshal(scanner.Bytes(), rules)
			if err != nil {
				printError(err)
			}
			res = append(res, rules)
		}
	}
	return res
}

func UnifyRules(rules []*grammar.Parser) unify {
	ruleSet := unify{}

	for _, rule := range rules {
		for _, imp := range rule.Imports {
			ruleSet.addImport(imp)
		}
		for _, r := range rule.Rules {
			ruleSet.addRule(r)
		}
	}
	return ruleSet
}

func GenerateOutputFromYara(res []*grammar.Parser, validJSON bool) {
	if validJSON == true {
		ruleset := map[string][]*grammar.Parser{"ruleset": res}
		j, err := json.Marshal(ruleset)
		if err == nil {
			os.Stdout.Write(j)
		} else {
			printError(err)
		}
	} else {
		for _, r := range res {
			j, err := json.Marshal(r)
			if err == nil {
				os.Stdout.Write(j)
				os.Stdout.WriteString("\n")
			} else {
				printError(err)
			}
		}
	}
}

func GenerateOutputToYaraDir(rules []*grammar.Parser, outputDir string, overwrite bool) {
	for _, rule := range rules {
		savePath := path.Join(outputDir, rule.Name)
		ruleStr := fmt.Sprintf("%s", rule.String())
		if overwrite {
			err := ioutil.WriteFile(savePath, []byte(ruleStr), 0644)
			checkErr(err)
		} else if _, err := os.Stat(savePath); os.IsNotExist(err) {
			err := ioutil.WriteFile(savePath, []byte(ruleStr), 0644)
			checkErr(err)
		}
	}
}

func GenerateOutputToYaraFile(rule unify, outputFile string, overwrite bool) {
	ruleStr := fmt.Sprintf("%s", rule.String())
	if overwrite {
		err := ioutil.WriteFile(outputFile, []byte(ruleStr), 0644)
		checkErr(err)
	} else if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		err := ioutil.WriteFile(outputFile, []byte(ruleStr), 0644)
		checkErr(err)
	}
}
