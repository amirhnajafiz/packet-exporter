package main

import "encoding/json"

const unknownLevel = "Unknown"

type encode struct {
	Level string `json:"L"`
}

func encodeLog(log string) string {
	tmp := new(encode)

	if err := json.Unmarshal([]byte(log), tmp); err != nil {
		return unknownLevel
	}

	return tmp.Level
}
