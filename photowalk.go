package main

import (
	"log"
	"io/ioutil"
)

func main() {
	dn := "/work/uploads"
	files, err := ioutil.ReadDir(dn)

	if err != nil {
		log.Fatal(err)
	}

	for _, fd := range files {
		log.Println(fd.Name())
	}
}
