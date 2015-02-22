package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/hauva69/photowalk/logging"
	"github.com/hauva69/photowalk/photograph"
	"io/ioutil"
	"os"
)

func handleFile(sourceDirectory string, file os.FileInfo) {
	fn := fmt.Sprintf("%s/%s", sourceDirectory, file.Name())
	logging.Log.Debug("filename=%s", fn)
	photo := photograph.New()
	err := photo.Load(fn)
	if err != nil {
		logging.Log.Error("%s", err)
	}
	logging.Log.Debug("%d EXIF tags", len(photo.ExifMap))
	for tag, value := range photo.ExifMap {
		fmt.Printf("%s\t%s\n", tag, value)
	}
	logging.Log.Info("photo=%v", photo)
}

func main() {
	usage := `Photowalk.

Usage:
  photowalk list <sourceDir>
  photowalk -h | --help
  photowalk --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, err := docopt.Parse(usage, nil, true, "Photowalk 0.01",
		false)
	if err != nil {
		logging.Log.Fatal(err)
		os.Exit(3)
	}

	fmt.Println(arguments)
	os.Exit(42)
	sourceDirectory := os.Args[1]
	logging.Log.Debug("sourceDirectory=%v", sourceDirectory)
	files, err := ioutil.ReadDir(sourceDirectory)
	if err != nil {
		logging.Log.Fatal("%v", err)
		os.Exit(4)
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
