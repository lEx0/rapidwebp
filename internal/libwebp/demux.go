package libwebp

import "C"

/*
#cgo CFLAGS: -I../source
#include <stdlib.h>
#include <demux.h>

static WebPDemuxer* newWebPDemuxer(const uint8_t* data, size_t data_size) {
	WebPData webp_data = {};
	WebPDataInit(&webp_data);

	webp_data.bytes = data;
	webp_data.size = data_size;

	return WebPDemux(&webp_data);
}

static uint8_t* getWebPDemuxerFirstFrameBytes(const uint8_t* in, size_t in_size, size_t* out_size) {
	WebPData webp_data = {};
	WebPDataInit(&webp_data);

	webp_data.bytes = in;
	webp_data.size = in_size;

	WebPDemuxer* demuxer = WebPDemux(&webp_data);

	if (demuxer == NULL) {
		return NULL;
	}

	WebPIterator iter = {};
	if (WebPDemuxGetFrame(demuxer, 1, &iter)) {
		*out_size = iter.fragment.size;
		uint8_t* out = malloc(*out_size);
		memcpy(out, iter.fragment.bytes, *out_size);
		WebPDemuxReleaseIterator(&iter);
		WebPDemuxDelete(demuxer);
		return out;
	}

	return NULL;
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

func WebPDemuxerFirstFrameBytes(data []byte) ([]byte, error) {
	outSize := C.size_t(0)
	out := C.getWebPDemuxerFirstFrameBytes(
		(*C.uint8_t)(&data[0]),
		C.size_t(len(data)),
		&outSize,
	)

	if out == nil {
		return nil, errors.New("cannot decode animated WebP file")
	}

	defer C.free(unsafe.Pointer(out))

	return C.GoBytes(unsafe.Pointer(out), C.int(outSize)), nil
}
