package main

import (
	"log"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
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
		im, err := getImage(fn)

		if err != nil {
			log.Printf("%s: %q\n", fn, err)
		} else {
			size := im.Bounds().Size()
			log.Printf("%s: width=%d height=%d", fn, size.X, size.Y)
		}
	}
}
