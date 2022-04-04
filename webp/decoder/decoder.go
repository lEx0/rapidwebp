package decoder

import (
	"github.com/lEx0/rapidwebp/internal/libwebp"
	"image"
	"io"
	"io/ioutil"
)

type (
	Decoder struct {
		options Options
		config  *libwebp.WebPDecoderConfig
	}
	Options struct {
		ByPassFiltering        bool
		NoFancyUpsampling      bool
		Crop                   *image.Rectangle
		Scale                  *image.Rectangle
		UseThreads             bool
		DitheringStrength      int
		AlphaDitheringStrength int
		Flip                   bool
	}
	Features struct {
		Width  int
		Height int

		HasAlpha     bool
		HasAnimation bool
	}
)

func NewDecoder(options Options) (d *Decoder, err error) {
	d = &Decoder{options: options}
	d.config = &libwebp.WebPDecoderConfig{
		BypassFiltering:        d.options.ByPassFiltering,
		NoFancyUpsampling:      d.options.NoFancyUpsampling,
		UseThreads:             d.options.UseThreads,
		DitheringStrength:      d.options.DitheringStrength,
		AlphaDitheringStrength: d.options.AlphaDitheringStrength,
		Flip:                   d.options.Flip,
	}

	if d.options.Crop != nil {
		d.config.CropTop = d.options.Crop.Min.Y
		d.config.CropLeft = d.options.Crop.Min.X
		d.config.CropWidth = d.options.Crop.Max.X - d.options.Crop.Min.X
		d.config.CropHeight = d.options.Crop.Max.Y - d.options.Crop.Min.Y
	}

	if d.options.Scale != nil {
		d.config.ScaledWidth = d.options.Scale.Max.X
		d.config.ScaledHeight = d.options.Scale.Max.Y
	}

	err = libwebp.WebPInitDecodeConfig(d.config)

	return
}

func (d *Decoder) GetFeatures(r io.Reader) (f *Features, err error) {
	if data, err := ioutil.ReadAll(r); err != nil {
		return nil, err
	} else if features, err := libwebp.WebPGetFeatures(data); err != nil {
		return nil, err
	} else {
		return &Features{
			Width:  features.Width,
			Height: features.Height,

			HasAlpha:     features.HasAlpha,
			HasAnimation: features.HasAnimation,
		}, nil
	}
}

func (d *Decoder) Decode(r io.Reader) (image.Image, error) {
	if data, err := ioutil.ReadAll(r); err != nil {
		return nil, err
	} else if features, err := libwebp.WebPGetFeatures(data); err != nil {
		return nil, err
	} else {
		return libwebp.WebPDecode(data, d.config, features)
	}
}
