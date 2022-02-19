package bgr555_test

import (
	"image/color"
	"testing"

	"github.com/yumland/gbarom/bgr555"
)

func TestToRGBA(t *testing.T) {
	o := bgr555.ToRGBA(0x4F0F)
	if expected := (color.RGBA{0x7B, 0xC5, 0x9C, 0xFF}); expected != o {
		t.Errorf("expected %v, got %v", expected, o)
	}
}
