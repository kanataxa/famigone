package main

import (
	"log"
	"path/filepath"

	"github.com/kanataxa/famigone/pkg/runner"
)

func main() {

	runner, err := runner.New(filepath.Join("testdata", "nestest.nes"))
	if err != nil {
		log.Fatal(err)
	}
	runner.Run()
}
