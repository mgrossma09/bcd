package bcd

import (
	"testing"

	"gotest.tools/assert"
)

func TestBCDAdd(t *testing.T) {
	test_cases := []struct {
		input1    uint8
		input2    uint8
		expResult uint8
		expCarry  bool
		expZero   bool
	}{
		{0x1, 0x1, 0x2, false, false},
		{0x2, 0x2, 0x4, false, false},
		{0x15, 0x20, 0x35, false, false},
		{0x15, 0x27, 0x42, false, false},
		{0x22, 0x77, 0x99, false, false},
		{0x23, 0x77, 0x00, true, true},
		{0x24, 0x77, 0x01, true, false},
		{0x0, 0x0, 0x0, false, true},

		// leave out illegal BCD inputs for now
		// {0x15, 0x2f, 0x4a, false, false},
		// {0x15, 0x2e, 0x49, false, false},
		// {0x1f, 0x2f, 0x44, false, false},
		// {0x1e, 0xf5, 0x79, true, false},
	}

	for _, tc := range test_cases {
		result, carry, zero := BCDAdd(tc.input1, tc.input2)
		assert.Equal(t, tc.expResult, result)
		assert.Equal(t, tc.expCarry, carry)
		assert.Equal(t, tc.expZero, zero)
	}
}
