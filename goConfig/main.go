package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// var data = `
// a: Easy!
// b:
//   c: 2
//   d.test: ["d1", "d2"]
// `

var data, _ = os.ReadFile("conf.yaml")

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int      `yaml:"c"`
		DTest    []string `yaml:"d.test,flow"`
	}
}

func main() {
	// if we use struct containing yaml encoding for yaml formated string
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t after unmarshal:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t after marshal:\n%s\n\n", string(d))
}
