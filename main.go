package main

import docopt "github.com/docopt/docopt-go"

var (
	// Build variables (set at buildtime by the compiler)
	Name      = "YaGo"
	Version   = "v0.0.0"
	BuildID   = ""
	BuildDate = ""
)

func main() {

	usage := `YaGo - Parsing Yara rules like a Gopher.

Usage:
  yago fileName <fileName> [ --validJSON ]
  yago dirName <dirName> [ --validJSON ]
  yago indexFile <indexFile> [ cwd <path> ] [ --validJSON ]
  yago inputFile <inputFile> outputDir <outputDir> [ --overwrite ] [ --validJSON ]
  yago inputFile <inputFile> outputFile <outputFile> [ --overwrite ] [ --validJSON ]
  yago -h | --help
  yago --version

Options:
  -h --help             Show this screen.
  --overwrite           Overwrites existing files [dafault: false].
  --validJSON           Print rules using a valid JSON format [dafault: false].
  --version             Show version.
`
	version := printVersion()
	arguments, _ := docopt.Parse(usage, nil, true, version, false)

	if arguments["fileName"].(bool) {
		if arguments["<fileName>"].(string) == "" {
			errAndExit("ERROR: You must provide a file.")
		}

		validJSON := arguments["--validJSON"].(bool)
		fileName := arguments["<fileName>"].(string)

		res := processFile(fileName)
		generateOutputFromYara(res, validJSON)

	} else if arguments["dirName"].(bool) {
		if arguments["<dirName>"].(string) == "" {
			errAndExit("ERROR: You must provide a directory.")
		}

		validJSON := arguments["--validJSON"].(bool)
		dirName := arguments["<dirName>"].(string)

		res := processDir(dirName)
		generateOutputFromYara(res, validJSON)

	} else if arguments["indexFile"].(bool) {
		if arguments["<indexFile>"].(string) == "" {
			errAndExit("ERROR: You must provide a index file.")
		}
		cwd := ""
		if arguments["cwd"].(bool) {
			cwd = arguments["<path>"].(string)
		}

		validJSON := arguments["--validJSON"].(bool)
		indexFile := arguments["<indexFile>"].(string)

		res := processIndex(indexFile, cwd)
		generateOutputFromYara(res, validJSON)

	} else if arguments["inputFile"].(bool) {
		if arguments["<inputFile>"].(string) == "" {
			errAndExit("ERROR: You must provide a input file.")
		}

		inputFile := arguments["<inputFile>"].(string)
		validJSON := arguments["--validJSON"].(bool)
		overwrite := arguments["--overwrite"].(bool)

		if arguments["outputDir"].(bool) {
			if arguments["<outputDir>"].(string) == "" {
				errAndExit("ERROR: You must provide a output directory.")
			}

			outputDir := arguments["<outputDir>"].(string)

			res := processInputFile(inputFile, validJSON)
			generateOutputToYaraDir(res, outputDir, overwrite)

		} else if arguments["outputFile"].(bool) {
			if arguments["<outputFile>"].(string) == "" {
				errAndExit("ERROR: You must provide a output file.")
			}

			outputFile := arguments["<outputFile>"].(string)

			res := processInputFile(inputFile, validJSON)
			uniq := unifyRules(res)
			generateOutputToYaraFile(uniq, outputFile, overwrite)
		}

	} else {
		errAndExit("Unexpected argument")
	}
}
