package crypto_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSHA1(t *testing.T) {
	assert.EqualValues(t, SHA1("sha1 this string"), "cf23df2207d99a74fbe169e3eba035e633b65d94")
}

