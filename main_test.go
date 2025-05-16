package main

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		input            string
		decimal          float64
		hex              string
		binary           string
		isTwosComplement bool
		isFractional     bool
	}{
		{"0", 0, "0x0", "0b0", false, false},
		// decimal positive full number
		{"123", 123, "0x7B", "0b1111011", false, false},
		{"95913", 95913, "0x176A9", "0b10111011010101001", false, false},
		// decimal positive fractional number
		{"123.45", 123.45, "0x7B", "0b1111011", false, true},
		{"59.301", 59.301, "0x3B", "0b111011", false, true},
		// decimal negative full number
		{"-123", -123, "0x85", "0b10000101", true, false},
		{"-95913", -95913, "0x28957", "0b101000100101010111", true, false},
		// binary
		{"0b1010", 10, "0xA", "0b1010", false, false},
		{"0b1111", 15, "0xF", "0b1111", false, false},
		// hex
		{"0xAFB", 2811, "0xAFB", "0b101011111011", false, false},
		{"0xFFFF", 65535, "0xFFFF", "0b1111111111111111", false, false},
	}

	for _, test := range tests {
		decimal, hex, binary, isTwosComplement, isFractional, err := convert(test.input)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if decimal != test.decimal {
			t.Errorf("Decimal: %v, Expected: %v", decimal, test.decimal)
		}
		if hex != test.hex {
			t.Errorf("Hex: %v, Expected: %v", hex, test.hex)
		}
		if binary != test.binary {
			t.Errorf("Binary: %v, Expected: %v", binary, test.binary)
		}
		if isTwosComplement != test.isTwosComplement {
			t.Errorf("IsTwosComplement: %v, Expected: %v", isTwosComplement, test.isTwosComplement)
		}
		if isFractional != test.isFractional {
			t.Errorf("IsFractional: %v, Expected: %v", isFractional, test.isFractional)
		}
	}
}
