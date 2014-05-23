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

	photographs := make([]photograph.Photograph, 30000, 30000)

	for i := 0; i < len(files); i++ {
		fd := files[i]
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
			photo := photograph.Photograph{fn, size.X, size.Y}
			photographs[i] = photo
				}
	}

	sort.Reverse(photograph.ByMaximumDimension(photographs))

	for _, p := range photographs {
		fmt.Printf("%v\n", p)
	}
}
