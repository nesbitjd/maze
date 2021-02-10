package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("maze.txt")

	if err != nil {
		log.Fatal(err)
	}

	mazemap := compress(string(content))

	fmt.Println(mazemap)

	fmt.Println(len(mazemap))

	fmt.Printf("The third index in mazemap is %v", mazemap[2])
}

func compress(dirtymap string) string {
	tidymap := strings.Replace(dirtymap, "  ", " ", -1)
	tidymap = strings.Replace(tidymap, "--", "-", -1)

	return tidymap
}
