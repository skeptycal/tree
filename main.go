package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/zsh/tree"
)

func main() {
	c, err := tree.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Copyright())
}
