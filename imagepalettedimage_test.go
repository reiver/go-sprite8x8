package sprite8x8_test

import (
	"github.com/reiver/go-sprite8x8"

	"image"

	"testing"
)

func TestImagePalettedImage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var x image.PalettedImage = sprite8x8.Paletted{}

	if nil == x {
		t.Error("This should never happen")
	}
}
