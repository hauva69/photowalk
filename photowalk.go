package main

import (
	"log"
	"image"
	"image/jpeg"
	"io/ioutil"
	"photograph"
	"os"
	"fmt"
	"sort"
	"strings"
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

	photographs := make([]photograph.Photograph, 0)

	for _, fd := range files {
		fmt.Println("hello")
		fn := dn + "/" + fd.Name()

		if strings.Contains(fn, "thumb") {
			continue
		}

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
			photo := photograph.Photograph{fn, size.X, size.Y}
			photographs = append(photographs, photo)
		}
	}

	sort.Reverse(photograph.ByMaximumDimension(photographs))

	for _, p := range photographs {
		fmt.Println("bar")
		fmt.Printf("%v\n", p)
	}
}
