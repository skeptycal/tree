package main

import (
	"fmt"

	"github.com/skeptycal/tree"

	log "github.com/sirupsen/logrus"
)

func main() {
	c, err := tree.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Copyright())
}
