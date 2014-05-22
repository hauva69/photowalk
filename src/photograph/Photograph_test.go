package photograph

import (
	"testing"
)

func TestNew (t *testing.T) {
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

