package photograph

import (
	"fmt"
	"github.com/hauva69/photowalk/logging"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/rwcarlsen/goexif/tiff"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type ExifMap map[exif.FieldName]*tiff.Tag

type Photograph struct {
	OriginalFileName string
	Data             []byte
	Time             time.Time
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
	exif.RegisterParsers(mknote.All...)
	exifData, err := exif.Decode(fd)
	if err != nil {
		logging.Log.Error("%v", err)
		return err
	}
	err = exifData.Walk(p)
	if err != nil {
		logging.Log.Error("%v", err)
		return err
	}
	offset, err := fd.Seek(0, 0)
	if err != nil {
		logging.Log.Error("%v", err)
		return err
	}
	logging.Log.Debug("offset=%v", offset)
	p.Data, err = ioutil.ReadAll(fd)
	if err != nil {
		logging.Log.Error("%v", err)
		return err
	}

	return nil
}

// Return the Time member as a ISO 8601 string
func (p *Photograph) Iso8601() string {
	layout := "2006-01-02"

	return p.Time.Format(layout)
}

// Walk implements exif.Walker interface and initializes the Photograph
// from the EXIF data.
func (p *Photograph) Walk(field exif.FieldName, tag *tiff.Tag) error {
	logging.Log.Info("%v: %v", field, tag)
	p.ExifMap[field] = tag
	if "DateTime" == field {
		const timeFormat = "2006:01:02 15:04:05"
		s, err := tag.StringVal()
		if err != nil {
			return err
		}
		p.Time, err = time.Parse(timeFormat, s)
		if err != nil {
			return err
		}
	}
	return nil
}

// String returns the fields of the Photograph as a string.
func (p *Photograph) String() string {
	return fmt.Sprintf("%s\t%d\t%d\t%v", p.OriginalFileName, p.Width,
		p.Height, p.Time)
}

func IsPhotographyFile(filename string) bool {
	extension := filepath.Ext(filename)
	logging.Log.Debug("extension=%s", extension)
	switch extension {
	case ".nef", ".NEF":
		return true
	}

	return false
}
