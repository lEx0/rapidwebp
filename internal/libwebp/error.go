package libwebp

import "errors"

var (
	ErrDataIsEmpty          = errors.New("data is empty")
	ErrInvalidDecoderConfig = errors.New("invalid decoder config")
)

const (
	VP8OK Vp8Status = iota
	ErrVP8OutOfMemory
	ErrVP8StatusInvalidParam
	ErrVP8StatusBitstreamError
	ErrVP8StatusUnsupportedFeature
	ErrVP8StatusSuspended
	ErrVP8StatusUserAbort
	ErrVP8StatusNotEnoughData
)

// Vp8Status is the status of the VP8 decoder.
type Vp8Status int

// Error returns the error message of the VP8 decoder.
func (e Vp8Status) Error() string {
	switch e {
	case ErrVP8OutOfMemory:
		return "VP8: out of memory"
	case ErrVP8StatusInvalidParam:
		return "VP8: invalid parameter"
	case ErrVP8StatusBitstreamError:
		return "VP8: bitstream error"
	case ErrVP8StatusUnsupportedFeature:
		return "VP8: unsupported feature"
	case ErrVP8StatusSuspended:
		return "VP8: suspended"
	case ErrVP8StatusUserAbort:
		return "VP8: user abort"
	case ErrVP8StatusNotEnoughData:
		return "VP8: not enough data"
	default:
		return "VP8: unknown error"
	}
}
