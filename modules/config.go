package modules

import (
	"os"
	"fmt"
	"encoding/json"
)

type Configuration struct {
	GithubApiKey	string
}

var (
	Config Configuration = Configuration{}
)

func ConfigLoad(filename string) {
	file, err := os.Open(filename)
	
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	e := decoder.Decode(&Config)
	
	if e != nil {
		fmt.Println("error:", e)
	}
}