package main

import (
	"log"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"fmt"
)

func getImage(fn string) (image.Image, error) {
	r, err := os.Open(fn)

	if err != nil {
		return nil, err
	}

	im, err := jpeg.Decode(r)

	if err != nil {
		return nil, err
	}

	return im, nil
}

func main() {
	dn := "/work/uploads"
	files, err := ioutil.ReadDir(dn)

	if err != nil {
		log.Fatal(err)
	}

	for _, fd := range files {
		fn := dn + "/" + fd.Name()
		fi, err := os.Stat(fn)

		if (err != nil) {
			log.Println(err)
		} else if !fi.Mode().IsRegular() {
			continue
		}

		im, err := getImage(fn)

		if err != nil {
			log.Printf("%s: %q\n", fn, err)
		} else {
			size := im.Bounds().Size()
			fmt.Printf("%s\t%d\t%d\n", fn, size.X, size.Y)
		}
	}
}
