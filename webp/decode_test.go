package webp

import (
	"github.com/stretchr/testify/assert"
	"image"
	"image/color"
	"os"
	"testing"
)

func TestDecodeConfig(t *testing.T) {
	type testEntry struct {
		name          string
		file          string
		failed        bool
		width, height int
		model         color.Model
	}

	tests := []testEntry{
		{"valid file", "../test/data/webp-logo-lossy.webp", false, 3000, 2000, color.NRGBAModel},
		{"invalid file ", "../test/data/invalid.webp", true, 0, 0, color.NRGBAModel},
		{"empty file", "../test/data/invalid_empty.webp", true, 0, 0, color.NRGBAModel},
		{"animated file", "../test/data/animated.webp", false, 400, 400, color.NRGBAModel},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Errorf("unable to open file: %v", err)
			}
			//goland:noinspection GoUnhandledErrorResult
			defer f.Close()

			if config, err := DecodeConfig(f); test.failed {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.width, config.Width)
				assert.Equal(t, test.height, config.Height)
				assert.Equal(t, test.model, config.ColorModel)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type testEntry struct {
		name          string
		file          string
		failed        bool
		width, height int
	}

	tests := []testEntry{
		{"valid file", "../test/data/webp-logo-lossy.webp", false, 3000, 2000},
		{"invalid file ", "../test/data/invalid.webp", true, 0, 0},
		{"empty file", "../test/data/invalid_empty.webp", true, 0, 0},
		{"animated file", "../test/data/animated.webp", false, 400, 400},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Errorf("unable to open file: %v", err)
			}
			//goland:noinspection GoUnhandledErrorResult
			defer f.Close()

			img, err := Decode(f)
			if test.failed && err != nil {
				assert.Error(t, err)
			} else if err != nil {
				t.Fatalf("expected error, but got %v", err)
			} else {
				assert.NoError(t, err)
				assert.IsType(t, &image.NRGBA{}, img)
				assert.Equal(t, test.width, img.Bounds().Dx())
				assert.Equal(t, test.height, img.Bounds().Dy())
			}
		})
	}
}
