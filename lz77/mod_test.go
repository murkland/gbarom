package lz77_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/yumland/gbarom/lz77"
)

func TestLZ77Decompress(t *testing.T) {
	out, err := lz77.Decompress(bytes.NewBuffer([]byte{16, 37, 0, 0, 0, 104, 101, 108, 108, 111, 32, 116, 104, 0, 101, 114, 101, 33, 49, 50, 51, 52, 0, 53, 54, 49, 50, 51, 52, 53, 54, 128, 144, 11, 10, 0}))
	if err != nil {
		t.Fatalf("Decompress(): %s (got %v)", err, string(out))
	}
	if expected := []byte("hello there!123456123456123456123456\n"); !reflect.DeepEqual(out, expected) {
		t.Fatalf("expected %v, got %v", expected, out)
	}
}
