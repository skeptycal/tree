package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	c, err := GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.Copyright())
}
