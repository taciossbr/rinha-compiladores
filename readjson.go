package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type File struct {
	Name       string         `json:"name"`
	Expression map[string]any `json:"expression"`
	Location   Location       `json:"location"`
}

func readFile(filename string) File {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(
		jsonFile)

	var codeFile File
	err = json.Unmarshal(byteValue, &codeFile)

	return codeFile
}
