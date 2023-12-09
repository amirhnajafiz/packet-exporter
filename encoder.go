package main

import (
	"encoding/json"
)

const unknownLevel = "Unknown"

type encodeLog struct {
	Level string `json:"L"`
}

func EncodeLog(log string) string {
	tmp := new(encodeLog)

	if err := json.Unmarshal([]byte(log), tmp); err != nil {
		return unknownLevel
	}

	return tmp.Level
}
