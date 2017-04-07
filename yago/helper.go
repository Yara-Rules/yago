package yago

import (
	"encoding/json"
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printError(msg error) {
	err := make(map[string]string)
	err["error"] = msg.Error()
	j, _ := json.Marshal(err)
	os.Stderr.Write(j)
	os.Exit(1)
}
