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

func help() {
	fmt.Println("FIXME")
}

func main() {
	if 2 != len(os.Args) {
		msg := "Source directory as a command line argument " +
			"required."
		logging.Log.Fatal(msg)
		os.Exit(2)
	}

	usage := `Photowalk.

Usage:
  photowalk list
  photowalk -h | --help
  photowalk --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.Parse(usage, nil, true, "Photowalk 0.01", false)
	fmt.Println(arguments)

	command := os.Args[1]

	switch {
	default:
		help()
	case command == "list":
		logging.Log.Debug("command=%s", command)
	}
	os.Exit(42)
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
