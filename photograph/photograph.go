package photograph

import (
	"fmt"
	"github.com/hauva69/photowalk/logging"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/rwcarlsen/goexif/tiff"
	"os"
)

type ExifMap map[exif.FieldName]*tiff.Tag

type Photograph struct {
	OriginalFileName string
	Width            int
	Height           int
	ExifMap
}

// New returns a pointer to a new Photograph.
func New() *Photograph {
	p := new(Photograph)
	p.ExifMap = make(ExifMap)
	return p
}

// Load initializes a Photograph from a file.
func (p *Photograph) Load(fileName string) error {
	p.OriginalFileName = fileName
	fd, err := os.Open(fileName)
	if err != nil {
		logging.Log.Error("%v", err)
		return err
	}
	defer fd.Close()
	exifData, err := exif.Decode(fd)
	if err != nil {
		logging.Log.Error("%v", err)
		return err
	}
	for _, maker := range mknote.All {
		err = maker.Parse(exifData)
		if err != nil {
			logging.Log.Error("v", err)
			return err
		}
	}
	err = exifData.Walk(p)
	if err != nil {
		logging.Log.Error("%v", err)
		return err
	}

	return nil
}

// Walk implements exif.Walker interface and initializes the Photograph
// from the EXIF data.
func (p *Photograph) Walk(field exif.FieldName, tag *tiff.Tag) error {
	logging.Log.Info("%v: %v", field, tag)
	p.ExifMap[field] = tag

	return nil
}

// String returns the fields of the Photograph as a string.
func (p *Photograph) String() string {
	return fmt.Sprintf("%v\t%d\t%d", p.OriginalFileName, p.Width, p.Height)
}
