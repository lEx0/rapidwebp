package webp

import (
	"github.com/lEx0/fastwebp/webp/decoder"
	"image"
	"image/color"
	"io"
)

func Decode(r io.Reader) (image.Image, error) {
	d, err := decoder.NewDecoder(decoder.Options{
		ByPassFiltering:        false,
		NoFancyUpsampling:      false,
		Crop:                   nil,
		Scale:                  nil,
		UseThreads:             false,
		DitheringStrength:      0,
		AlphaDitheringStrength: 0,
		Flip:                   false,
	})

	if err != nil {
		return nil, err
	}

	return d.Decode(r)
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	d, err := decoder.NewDecoder(decoder.Options{
		ByPassFiltering:        false,
		NoFancyUpsampling:      false,
		Crop:                   nil,
		Scale:                  nil,
		UseThreads:             false,
		DitheringStrength:      0,
		AlphaDitheringStrength: 0,
		Flip:                   false,
	})

	if err != nil {
		return image.Config{}, err
	}

	features, err := d.GetFeatures(r)

	if err != nil {
		return image.Config{}, err
	}

	return image.Config{
		Width:      features.Width,
		Height:     features.Height,
		ColorModel: color.NRGBAModel,
	}, nil
}
