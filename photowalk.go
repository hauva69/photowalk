package main

import (
	"fmt"
	"github.com/hauva69/photowalk/logging"
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
	_, err = exif.Decode(fd)
	if err != nil {
		logging.Log.Error("%v", err)
	}
}

func main() {
	if 2 != len(os.Args) {
		msg := "Source directorys as a command line argument " + 
			"required."
		logging.Log.Fatal(msg)
		os.Exit(2)
	}
	sourceDirectory := os.Args[1]
	logging.Log.Debug("src=%v", sourceDirectory)
	files, err := ioutil.ReadDir(sourceDirectory)
	if err != nil {
		logging.Log.Fatal("%v", err)
		os.Exit(1)
	}

	for i := range files {
		f := files[i]
		logging.Log.Debug("file=%s", f.Name())
		if f.IsDir() {
			logging.Log.Warning("%s is a directory", f)
		} else {
			handleFile(sourceDirectory, f)
		}
	}
}
