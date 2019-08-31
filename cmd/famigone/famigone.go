package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/kanataxa/famigone/pkg/bus"
	"github.com/kanataxa/famigone/pkg/cassettie"
)

func main() {
	c, err := cassettie.New(filepath.Join("testdata", "hello_world.nes"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bus.New(c))
}
