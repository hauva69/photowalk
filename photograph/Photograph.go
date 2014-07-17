package photograph

import "fmt"

type Photograph struct {
	OriginalFileName string
	Width            int
	Height           int
}

type ByMaximumDimension []Photograph

func New() *Photograph {
	return new(Photograph)
}

func (p Photograph) getMaximumDimension() int {
	if p.Width > p.Height {
		return p.Width
	} else {
		return p.Height
	}
}

func (p Photograph) String() string {
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
