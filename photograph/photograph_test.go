package photograph

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	const image = "/path/to/image/photowalk.png"
	const width, height = 13, 42
	im := New()
	im.OriginalFileName = image
	im.Width = width
	im.Height = height
	fmt.Printf("im=%v\n", im)

	if im.OriginalFileName != image {
		t.Errorf("OriginalFileName=%s, want %s",
			im.OriginalFileName, image)
	}

	if im.Width != width {
		t.Errorf("Width=%v, want %d", im.Width, width)
	}

	if im.Height != height {
		t.Errorf("Height=%v, want %d", im.Height, height)
	}
}

func TestString(t *testing.T) {
	photo := New()
	expected := "/images/1.jpg\t1000\t2000"
	got := fmt.Sprintf("%v", photo)
	if got != expected {
		t.Errorf("expected = %v, want %v", expected, got)
	}
}
