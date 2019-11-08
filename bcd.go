package bcd

import (
	"github.com/albenik/bcd"
)

func BCDAdd(input1 uint8, input2 uint8) (uint8, bool, bool) {
	carry := false
	zero := false

	input1_32 := bcd.ToUint32([]byte{input1})
	input2_32 := bcd.ToUint32([]byte{input2})

	result_32 := input1_32 + input2_32

	if result_32 > 99 {
		carry = true
	}

	result := uint8(bcd.FromUint32(result_32)[3])

	if result == 0 {
		zero = true
	}

	return result, carry, zero
}
