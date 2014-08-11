package photograph

import (
	"fmt"
	"github.com/hauva69/photowalk/logging"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

type ExifMap map[exif.FieldName]*tiff.Tag

type Photograph struct {
	OriginalFileName string
	Width            int
	Height           int
	ExifMap
}

type ByMaximumDimension []Photograph

func New() *Photograph {
	p := new(Photograph)
	return p
}

func (p *Photograph) getMaximumDimension() int {
	if p.Width > p.Height {
		return p.Width
	} else {
		return p.Height
	}
}

// Walk implements exif.Walker interface.
func (p *Photograph) Walk(field exif.FieldName, tag *tiff.Tag) error {
	logging.Log.Info("%v: %v", field, tag)
	// p.ExifMap[field] = tag

	return nil
}

// String returns the fields of the Photograph as a string.
func (p *Photograph) String() string {
	return fmt.Sprintf("%v\t%d\t%d", p.OriginalFileName, p.Width, p.Height)
}

func (a ByMaximumDimension) Len() int {
	return len(a)
}

// http://golang.org/pkg/sort/#Float64Slice.Less
// http://golang.org/pkg/sort/#example__sortKeys

func (a ByMaximumDimension) Less(i, j int) bool {
	iMax := a[i].getMaximumDimension()
	jMax := a[j].getMaximumDimension()

	return iMax < jMax
}

func (a ByMaximumDimension) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
