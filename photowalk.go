package main

import (
	"github.com/hauva69/photowalk/logging"
	"image"
	"image/jpeg"
	"io/ioutil"
	"github.com/hauva69/photowalk/photograph"
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
	dn := "examples"
	files, err := ioutil.ReadDir(dn)

	if err != nil {
		logging.Log.Fatal(err)
	}

	photographs := make([]photograph.Photograph, 0)

	for _, fd := range files {
		fn := dn + "/" + fd.Name()

		if strings.Contains(fn, "thumb") {
			continue
		}

		fi, err := os.Stat(fn)

		if (err != nil) {
			logging.Log.Error("%v", err)
		} else if !fi.Mode().IsRegular() {
			continue
		}

		im, err := getImage(fn)

		if err != nil {
			logging.Log.Error("%s: %q", fn, err)
		} else {
			size := im.Bounds().Size()
			photo := photograph.Photograph{fn, size.X, size.Y}
			photographs = append(photographs, photo)
		}
	}

	sort.Reverse(photograph.ByMaximumDimension(photographs))

	for _, p := range photographs {
		fmt.Printf("%v\n", p)
	}
}
