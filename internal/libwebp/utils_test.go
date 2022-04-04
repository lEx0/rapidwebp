package libwebp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool2int(t *testing.T) {
	assert.Equal(t, toCInt(1), bool2CInt(true))
	assert.Equal(t, toCInt(0), bool2CInt(false))
}
