package photograph

type Photograph struct {
	OriginalFileName	string
	Width	int
	Height	int
}

func New() *Photograph {
	return new(Photograph)
}
