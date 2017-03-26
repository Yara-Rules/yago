package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func printVersion() string {
	return fmt.Sprintf("%s %s\nBuild date: %s\nBuild ID: %s\n", Name, Version, BuildDate, BuildID)
}

func printError(msg error) {
	err := make(map[string]string)
	err["error"] = msg.Error()
	j, _ := json.Marshal(err)
	os.Stderr.Write(j)
	os.Exit(1)
}

func errAndExit(msg string) {
	os.Stderr.WriteString(msg + "\n")
	os.Exit(1)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
