package libwebp

import "testing"

func TestVp8Status_Error(t *testing.T) {
	tests := []struct {
		name string
		v    Vp8Status
		want string
	}{
		{
			name: "out of memory",
			v:    ErrVP8OutOfMemory,
			want: "VP8: out of memory",
		},
		{
			name: "invalid params",
			v:    ErrVP8StatusInvalidParam,
			want: "VP8: invalid parameter",
		},
		{
			name: "bitstream error",
			v:    ErrVP8StatusBitstreamError,
			want: "VP8: bitstream error",
		},
		{
			name: "unsupported feature",
			v:    ErrVP8StatusUnsupportedFeature,
			want: "VP8: unsupported feature",
		},
		{
			name: "suspended",
			v:    ErrVP8StatusSuspended,
			want: "VP8: suspended",
		},
		{
			name: "aborted",
			v:    ErrVP8StatusUserAbort,
			want: "VP8: user abort",
		},
		{
			name: "not enough data",
			v:    ErrVP8StatusNotEnoughData,
			want: "VP8: not enough data",
		},
		{
			name: "unexpeced",
			v:    Vp8Status(100),
			want: "VP8: unknown error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Error(); got != tt.want {
				t.Errorf("Vp8Status.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
