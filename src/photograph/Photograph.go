package photograph

type Photograph struct {
	OriginalFileName	string
	Width	int
	Height	int
}

type ByMaximumDimension []Photograph

func New() *Photograph {
	return new(Photograph)
}

func (a ByMaximumDimension) Len() int {
	return len(a)
}

// http://golang.org/pkg/sort/#Float64Slice.Less

func (a ByMaximumDimension) Less(i, j, int) bool {

}
