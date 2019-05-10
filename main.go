package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("expand: ")

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(os.ExpandEnv(string(data)))
}
