package main

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test_TaskRequestPayload(t *testing.T) {
	t.Run("Should parse bytes as a big.Int", func(t *testing.T) {
		b := parseBigIntToHex(new(big.Int).SetUint64(4))

		i, err := parseHexBytesToBigInt(b)

		assert.Nil(t, err)
		assert.Equal(t, new(big.Int).SetInt64(4), i)
	})

}
