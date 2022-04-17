package main

import (
	"log"
)

func main() {
	r := newRoot()

	err := r.Register()
	if err != nil {
		log.Fatal(err)
	}

	defer r.Release()

	err = r.Resolve()
	if err != nil {
		log.Fatal(err)
	}
}
