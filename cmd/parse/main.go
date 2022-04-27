package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aluttik/go-crossplane"
)

func main() {
	payload, err := crossplane.Parse(os.Args[1], &crossplane.ParseOptions{})
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(b))

	f, err := os.Create("nginx.json")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(string(b))

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
