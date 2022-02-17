package bgr555

import "image/color"

func ToRGBA(c uint16) color.RGBA {
	return color.RGBA{
		uint8((c & 0b11111) * 0xFF / 0b11111),
		uint8(((c >> 5) & 0b11111) * 0xFF / 0b11111),
		uint8(((c >> 10) & 0b11111) * 0xFF / 0b11111),
		0xFF,
	}
}
