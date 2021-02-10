package main

import (
	"fmt"
	"io/ioutil"

	uyaml "github.com/buildkite/yaml"
	"github.com/go-vela/types/yaml"
)

func main() {
	var good yaml.Build
	input, err := ioutil.ReadFile("stages.yml")
	if err != nil {
		panic(err)
	}
	err = uyaml.Unmarshal([]byte(input), &good)
	if err != nil {
		panic(err)
	}

	for _, stage := range good.Stages {
		fmt.Println(stage.Name)
	}

	out, err := uyaml.Marshal(good)
	if err != nil {
		panic(err)
	}

	var bad yaml.Build
	err = uyaml.Unmarshal([]byte(out), &bad)
	if err != nil {
		panic(err)
	}

	for _, stage := range bad.Stages {
		fmt.Println(stage.Name)
	}
}
