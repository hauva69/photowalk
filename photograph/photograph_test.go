package photograph

import (
	"fmt"
	"sort"
	"testing"
)

func TestNew(t *testing.T) {
	const image = "/path/to/image/photowalk.png"
	const width, height = 13, 42
	im := New()

	if im.OriginalFileName != "" {
		t.Errorf("OriginalFileName = %v, want an empty string", im.OriginalFileName)
	}

	if im.Width != 0 {
		t.Errorf("Width = %v, want 0", im.Width)
	}

	if im.Height != 0 {
		t.Errorf("Height = %v, want 0", im.Height)
	}

	im = &Photograph{image, width, height}

	if im.OriginalFileName != image {
		t.Errorf("OriginalFileName = %v, want %v", im.OriginalFileName, image)
	}

	if im.Width != width {
		t.Errorf("Width = %v, want %v", im.Width, width, height)
	}

	if im.Height != height {
		t.Errorf("Height = %v, want %v", im.Height, width, height)
	}
}

func TestSortByMaximumDimension(t *testing.T) {
	images := []Photograph{
		{"/images/3.nef", 3000, 356},
		{"/images/1.jpg", 123, 1000},
		{"/images/2.png", 223, 2000},
	}

	sortedImages := []Photograph{
		{"/images/1.jpg", 123, 1000},
		{"/images/2.png", 223, 2000},
		{"/images/3.nef", 3000, 356},
	}

	sort.Sort(ByMaximumDimension(images))

	for i := 0; i < len(images); i++ {
		if images[i] != sortedImages[i] {
			t.Errorf("images[%d] = %v, want %v", i, images[i], sortedImages[i])
		}
	}
}

func TestString(t *testing.T) {
	photo := Photograph{"/images/1.jpg", 1000, 2000}
	expected := "/images/1.jpg\t1000\t2000"
	got := fmt.Sprintf("%v", photo)
	if got != expected {
		t.Errorf("expected = %v, want %v", expected, got)
	}
}
