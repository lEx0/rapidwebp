package libwebp

/*
#cgo CFLAGS: -I../source
#include <stdlib.h>
#include <encode.h>
#include <decode.h>
*/
import "C"
import (
	"errors"
	"image"
	"unsafe"
)

//goland:noinspection GoUnusedConst
const (
	BitstreamFormatUndefined WebPBitstreamFormat = iota
	BitstreamFormatLossy
	BitstreamFormatLossless
)

type (
	// WebPDecoderConfig wrapper for C.WebPDecoderConfig
	WebPDecoderConfig struct {
		internal *C.WebPDecoderConfig

		BypassFiltering   bool // if true, skip the in-loop filtering
		NoFancyUpsampling bool // if true, use faster pointwise upsampler

		UseCropping           bool // if true, cropping is applied first
		CropLeft, CropTop     int  // top-left position for cropping.
		CropWidth, CropHeight int  // Will be snapped to even values, dimension of the cropping area

		UseScaling                bool // if true, scaling is applied afterward
		ScaledWidth, ScaledHeight int  // final resolution

		UseThreads             bool // if true, use multi-threaded decoding
		DitheringStrength      int  // dithering strength (0=Off, 100=full)
		Flip                   bool // if true, flip output vertically
		AlphaDitheringStrength int  // alpha dithering strength in [0..100]
	}
	// WebPBitstreamFeatures wrapper for C.WebPBitstreamFeatures
	WebPBitstreamFeatures struct {
		internal C.WebPBitstreamFeatures

		Width, Height          int
		HasAlpha, HasAnimation bool
		Format                 WebPBitstreamFormat
	}
	WebPBitstreamFormat int
)

// WebPInitDecodeConfig initializes a WebPDecoderConfig with default values.
func WebPInitDecodeConfig(config *WebPDecoderConfig) error {
	config.internal = &C.WebPDecoderConfig{}

	if C.WebPInitDecoderConfig(config.internal) == 0 {
		return errors.New("cannot init decoder config")
	}

	return nil
}

// WebPGetFeatures returns information about the specified WebP image.
func WebPGetFeatures(data []byte) (*WebPBitstreamFeatures, error) {
	if len(data) == 0 {
		return nil, ErrDataIsEmpty
	}

	features := WebPBitstreamFeatures{}

	if status := Vp8Status(
		C.WebPGetFeatures(
			(*C.uint8_t)(&data[0]),
			C.size_t(len(data)),
			&features.internal,
		),
	); status != VP8OK {
		return nil, status
	}

	features.Width = int(features.internal.width)
	features.Height = int(features.internal.height)
	features.HasAlpha = features.internal.has_alpha == 1
	features.HasAnimation = features.internal.has_animation == 1
	features.Format = WebPBitstreamFormat(features.internal.format)

	return &features, nil
}

// WebPDecode decodes the specified WebP image.
func WebPDecode(data []byte, config *WebPDecoderConfig, features *WebPBitstreamFeatures) (image.Image, error) {
	if len(data) == 0 {
		return nil, ErrDataIsEmpty
	} else if config == nil {
		return nil, ErrInvalidDecoderConfig
	}

	config.internal.options.bypass_filtering = bool2CInt(config.BypassFiltering)
	config.internal.options.no_fancy_upsampling = bool2CInt(config.NoFancyUpsampling)
	config.internal.options.use_cropping = bool2CInt(config.UseCropping)
	config.internal.options.crop_left = C.int(config.CropLeft)
	config.internal.options.crop_top = C.int(config.CropTop)
	config.internal.options.crop_width = C.int(config.CropWidth)
	config.internal.options.crop_height = C.int(config.CropHeight)
	config.internal.options.use_scaling = bool2CInt(config.UseScaling)
	config.internal.options.scaled_width = C.int(config.ScaledWidth)
	config.internal.options.scaled_height = C.int(config.ScaledHeight)
	config.internal.options.use_threads = bool2CInt(config.UseThreads)
	config.internal.options.dithering_strength = C.int(config.DitheringStrength)
	config.internal.options.flip = bool2CInt(config.Flip)
	config.internal.options.alpha_dithering_strength = C.int(config.AlphaDitheringStrength)

	config.internal.input = features.internal

	if config.UseScaling {
		config.internal.output.width = config.internal.options.scaled_width
		config.internal.output.height = config.internal.options.scaled_height
	} else if config.UseCropping {
		config.internal.output.width = config.internal.options.crop_width
		config.internal.output.height = config.internal.options.crop_height
	} else {
		config.internal.output.width = config.internal.input.width
		config.internal.output.height = config.internal.input.height
	}

	config.internal.output.colorspace = C.MODE_RGBA
	config.internal.output.is_external_memory = 1

	img := image.NewNRGBA(image.Rectangle{Max: image.Point{
		X: int(config.internal.output.width),
		Y: int(config.internal.output.height),
	}})

	buff := (*C.WebPRGBABuffer)(unsafe.Pointer(&config.internal.output.u[0]))
	buff.stride = C.int(img.Stride)
	buff.rgba = (*C.uint8_t)(&img.Pix[0])
	buff.size = (C.size_t)(len(img.Pix))

	if status := Vp8Status(
		C.WebPDecode(
			(*C.uint8_t)(&data[0]),
			C.size_t(len(data)),
			config.internal,
		),
	); status != VP8OK {
		return nil, status
	}

	return img, nil
}
