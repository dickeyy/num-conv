package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <number>")
		fmt.Println("Example: go run main.go 0xAFB")
		fmt.Println("Example: go run main.go 0b1010")
		fmt.Println("Example: go run main.go 123.45")
		fmt.Println("Example: go run main.go -12.34")
		os.Exit(1)
	}

	input := os.Args[1]
	var decimal float64
	var err error

	// Determine the input format and convert to decimal
	switch {
	case strings.HasPrefix(input, "0x"):
		// Handle hex input
		hexValue, err := strconv.ParseInt(strings.TrimPrefix(input, "0x"), 16, 64)
		if err != nil {
			fmt.Printf("Error: Invalid hex format - %v\n", err)
			os.Exit(1)
		}
		decimal = float64(hexValue)
	case strings.HasPrefix(input, "0b"):
		// Handle binary input
		binaryValue, err := strconv.ParseInt(strings.TrimPrefix(input, "0b"), 2, 64)
		if err != nil {
			fmt.Printf("Error: Invalid binary format - %v\n", err)
			os.Exit(1)
		}
		decimal = float64(binaryValue)
	default:
		// Handle decimal input (including fractional numbers)
		decimal, err = strconv.ParseFloat(input, 64)
	}

	if err != nil {
		fmt.Printf("Error: Invalid input format - %v\n", err)
		os.Exit(1)
	}

	// Convert decimal to other formats
	var hex, binary string
	var isTwosComplement bool
	var isFractional bool = decimal != float64(int64(decimal))

	if decimal < 0 {
		// For negative numbers, we need to calculate the two's complement
		// First convert to positive, then flip bits and add 1
		absValue := uint64(-decimal)
		// Find the minimum number of bits needed
		bits := 0
		temp := absValue
		for temp > 0 {
			bits++
			temp >>= 1
		}
		// Add one bit for the sign
		bits++
		// Calculate two's complement
		mask := (uint64(1) << bits) - 1
		twosComplement := (^absValue & mask) + 1

		hex = fmt.Sprintf("0x%X", twosComplement)
		binary = fmt.Sprintf("0b%b", twosComplement)
		isTwosComplement = true
	} else {
		hex = fmt.Sprintf("0x%X", uint64(decimal))
		binary = fmt.Sprintf("0b%b", uint64(decimal))
		isTwosComplement = false
	}
	decimalStr := fmt.Sprintf("%g", decimal)

	// ANSI escape codes for colors
	green := "\033[32m"
	gray := "\033[90m"
	reset := "\033[0m"

	// Print all formats in green with gray annotations where needed
	fmt.Printf("Decimal: %s%s%s\n", green, decimalStr, reset)

	// Only show annotation for hex if it's two's complement or fractional
	if isTwosComplement || isFractional {
		fmt.Printf("Hex: %s%s%s %s(%s)%s\n",
			green, hex, reset,
			gray, map[bool]string{true: "two's complement", false: "fractional decimal"}[isTwosComplement], reset)
	} else {
		fmt.Printf("Hex: %s%s%s\n", green, hex, reset)
	}

	// Only show annotation for binary if it's two's complement or fractional
	if isTwosComplement || isFractional {
		fmt.Printf("Binary: %s%s%s %s(%s)%s\n",
			green, binary, reset,
			gray, map[bool]string{true: "two's complement", false: "fractional decimal"}[isTwosComplement], reset)
	} else {
		fmt.Printf("Binary: %s%s%s\n", green, binary, reset)
	}
}
