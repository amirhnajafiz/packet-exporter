package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type EncodeLog struct {
	Level string `json:"L"`
}

func main() {
	FILE := flag.String("file", "log.txt", "log file path")

	flag.Parse()

	content, _ := os.ReadFile(*FILE)
	logs := strings.Split(string(content), "\n")

	for _, log := range logs {
		if len(log) < 2 {
			continue
		}

		tmp := new(EncodeLog)

		if err := json.Unmarshal([]byte(log), tmp); err != nil {
			panic(err)
		}

		fmt.Println(tmp.Level)
	}
}
