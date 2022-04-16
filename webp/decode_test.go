package webp

import (
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

			if config, err := DecodeConfig(f); test.failed && err == nil {
				t.Error(err)
			} else if !test.failed && err != nil {
				t.Error(err)
			} else if err != nil {
			} else if test.width != config.Width || test.height != config.Height {
				t.Errorf("expected width: %d, height: %d, got width: %d, height: %d", test.width, test.height, config.Width, config.Height)
			} else if test.model != config.ColorModel {
				t.Errorf("expected model: %v, got model: %v", test.model, config.ColorModel)
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
			if test.failed && err == nil {
				t.Error(err)
			} else if !test.failed && err != nil {
				t.Error(err)
			} else if err != nil {
			} else if test.width != img.Bounds().Dx() || test.height != img.Bounds().Dy() {
				t.Errorf("expected width: %d, height: %d, got width: %d, height: %d", test.width, test.height, img.Bounds().Dx(), img.Bounds().Dy())
			}
		})
	}
}
