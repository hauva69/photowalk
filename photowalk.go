package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/hauva69/photowalk/logging"
	"github.com/hauva69/photowalk/photograph"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getTargetDirectoryNameWithDate(date time.Time,
	targetDirerctory string) string {
	layout := "2006/01/02"
	logging.Log.Debug("layout=%v", layout)
	s := fmt.Sprintf("%s/%s", targetDirerctory, date.Format(layout))
	logging.Log.Debug("targetDirectory=%v", s)

	return s
}

func handleFile(sourceDirectory string, targetDirectory string,
	file os.FileInfo) {
	fn := filepath.Join(sourceDirectory, file.Name())
	logging.Log.Debug("filename=%s", fn)
	photo := photograph.New()
	err := photo.Load(fn)
	if err != nil {
		logging.Log.Error(err.Error())
	}
	logging.Log.Debug("%d EXIF tags", len(photo.ExifMap))
	for tag, value := range photo.ExifMap {
		fmt.Printf("%s\t%s\n", tag, value)
	}
	logging.Log.Debug("photo=%v", photo.OriginalFileName)
	logging.Log.Debug("targetDir=%v", targetDirectory)
	targetDirectory = getTargetDirectoryNameWithDate(photo.Time,
		targetDirectory)
	err = os.MkdirAll(targetDirectory, 0700)
	if err != nil {
		logging.Log.Error(err.Error())
	}
	s := strings.Replace(file.Name(), "_DSC", "", 1)
	s = strings.Replace(s, "DSC_", "", 1)
	ext := filepath.Ext(s)
	s = strings.Replace(s, ext, "", 1)
	ext = strings.ToLower(ext)
	logging.Log.Debug("ext=%s", ext)
	targetFile := fmt.Sprintf("%s_%s_%s%s", photo.Iso8601(), s,
		"makela_ari", ext)
	targetFile = filepath.Join(targetDirectory, targetFile)
	logging.Log.Debug("targetFile=%s", targetFile)
	ioutil.WriteFile(targetFile, photo.Data, 0600)
	targetMd5File := fmt.Sprintf("%s%s", targetFile, ".md5")
	logging.Log.Debug("targetMd5File=%s", targetMd5File)
	ioutil.WriteFile(targetMd5File, []byte(photo.Md5), 0600)
}

func handleDirectoryTree(sourceDirectory string) {
	err := filepath.Walk(sourceDirectory, walkFunc)
	if err != nil {
		logging.Log.Fatal(err)
		os.Exit(5)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		logging.Log.Fatal(err)
		os.Exit(2)
	}
	fmt.Println(path)

	return nil
}

func handleDirectory(sourceDirectory string, targetDirerctory string,
	listOnly bool) {
	files, err := ioutil.ReadDir(sourceDirectory)
	if err != nil {
		logging.Log.Fatal(err.Error())
		os.Exit(4)
	}

	for i := range files {
		f := files[i]
		if f.IsDir() {
			logging.Log.Warning("%s is a directory", f.Name())
		} else {
			photograph.IsPhotographyFile(f.Name())
			if listOnly {
				fmt.Println(filepath.Join(sourceDirectory,
					f.Name()))
			} else {
				handleFile(sourceDirectory,
					targetDirerctory, f)
			}
		}
	}
}

func main() {
	usage := `Photowalk.

Usage:
  photowalk [-r] list <sourceDir>
  photowalk import <sourceDir> <targetDir>
  photowalk -h | --help
  photowalk --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, err := docopt.Parse(usage, nil, true, "Photowalk 0.2",
		false)
	if err != nil {
		logging.Log.Fatal(err)
		os.Exit(3)
	}
	logging.Log.Debug("%v", arguments)

	listOnly := false

	if arguments["list"].(bool) {
		listOnly = true
	}

	targetDirectory := "baz"
	targetDir := arguments["<targetDir>"]
	if targetDir != nil {
		targetDirectory = arguments["<targetDir>"].(string)
	}

	sourceDirectory := arguments["<sourceDir>"].(string)
	logging.Log.Debug("sourceDirectory=%v", sourceDirectory)
	logging.Log.Debug("targetDirectory=%v", targetDirectory)
	if arguments["-r"].(bool) {
		handleDirectoryTree(sourceDirectory)
	} else {
		handleDirectory(sourceDirectory, targetDirectory, listOnly)
	}
}
