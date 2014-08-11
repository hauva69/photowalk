package main

import (
	"fmt"
	"github.com/hauva69/photowalk/logging"
	"github.com/hauva69/photowalk/photograph"
	"github.com/rwcarlsen/goexif/exif"
	"io/ioutil"
	"os"
)

func handleFile(sourceDirectory string, file os.FileInfo) {
	fn := fmt.Sprintf("%s/%s", sourceDirectory, file.Name())
	logging.Log.Debug("filename=%s", fn)
	fd, err := os.Open(fn)
	if err != nil {
		logging.Log.Error("%v", err)
	}
	defer fd.Close()
	exifData, err := exif.Decode(fd)
	if err != nil {
		logging.Log.Error("%v", err)
	}
	photo := photograph.New()
	logging.Log.Info("%v", photo)
	logging.Log.Info("%v", exifData.Walk(photo))
}

func main() {
	if 2 != len(os.Args) {
		msg := "Source directorys as a command line argument " +
			"required."
		logging.Log.Fatal(msg)
		os.Exit(2)
	}
	sourceDirectory := os.Args[1]
	logging.Log.Debug("sourceDirectory=%v", sourceDirectory)
	files, err := ioutil.ReadDir(sourceDirectory)
	if err != nil {
		logging.Log.Fatal("%v", err)
		os.Exit(1)
	}

	for i := range files {
		f := files[i]
		if f.IsDir() {
			logging.Log.Warning("%s is a directory", f)
		} else {
			handleFile(sourceDirectory, f)
		}
	}
}
