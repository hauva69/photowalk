package photograph

import (
	"fmt"
	"log"
	"testing"
	"time"
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
	const timeFormat = "2006-01-02 15:04:05"
	const timeString = "2006-01-02 15:04:05"
	photo := New()
	photo.OriginalFileName = "/images/1.jpg"
	photo.Width = 1000
	photo.Height = 2000
	tmp, err := time.Parse(timeFormat, timeString)
	if err != nil {
		log.Fatal("Unable to parse date.")
	} else {
		photo.Time = tmp
	}
	expected := "/images/1.jpg\t1000\t2000\t" + timeString +
		" +0000 UTC"
	got := fmt.Sprintf("%v", photo)
	if got != expected {
		t.Errorf("expected = %v, want %v", expected, got)
	}
}
