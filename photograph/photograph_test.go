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

	if im.OriginalFileName != image {
		t.Errorf("OriginalFileName = %s, want %s",
			im.OriginalFileName, image)
	}

	if im.Width != width {
		t.Errorf("Width = %v, want %d", im.Width, width)
	}

	if im.Height != height {
		t.Errorf("Height = %v, want 0", im.Height, height)
	}

	im = New()

	if im.OriginalFileName != image {
		t.Errorf("OriginalFileName = %v, want %v", im.OriginalFileName, image)
	}

	if im.Width != width {
		t.Errorf("Width = %v, want %v", im.Width, width)
	}

	if im.Height != height {
		t.Errorf("Height = %v, want %v", im.Height, height)
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
