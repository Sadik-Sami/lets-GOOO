package main

import "fmt"

func main() {
	// ======================================
	// 🧠 INTEGER TYPES (SIGNED & UNSIGNED)
	// ======================================

	/*
		Go provides several integer types of different sizes.
		Each type defines how many bits are used to represent the number,
		and whether it can be negative or not.
	*/

	// SIGNED INTEGERS (can store negative and positive)
	var i8 int8 = -128                   // range: -128 to +127
	var i16 int16 = -32768               // range: -32,768 to +32,767
	var i32 int32 = -2147483648          // range: -2,147,483,648 to +2,147,483,647
	var i64 int64 = -9223372036854775808 // range: -(2^63) to (2^63 - 1)

	// UNSIGNED INTEGERS (only positive)
	var u8 uint8 = 255                    // range: 0 to 255
	var u16 uint16 = 65535                // range: 0 to 65,535
	var u32 uint32 = 4294967295           // range: 0 to 4,294,967,295
	var u64 uint64 = 18446744073709551615 // range: 0 to 18,446,744,073,709,551,615

	/*
		The size determines how many bits are reserved for storing the number.
		Each bit doubles the range.
	*/

	// The default `int` and `uint` types depend on architecture:
	// - On 32-bit systems → int = int32, uint = uint32
	// - On 64-bit systems → int = int64, uint = uint64
	var archInt int = -123456
	var archUint uint = 987654

	// ======================================
	// 🧮 FLOATING POINT TYPES
	// ======================================

	/*
		float32 → 4 bytes (single precision)
		float64 → 8 bytes (double precision)
		The larger the size, the more decimal precision.
	*/

	var f32 float32 = 3.1415926535         // ~6–7 digits precision
	var f64 float64 = 3.141592653589793238 // ~15–16 digits precision

	// ======================================
	// 🔢 COMPLEX NUMBERS
	// ======================================

	/*
		Go supports complex numbers out of the box!
		complex64  = real(float32) + imaginary(float32)
		complex128 = real(float64) + imaginary(float64)
	*/

	var c64 complex64 = 2 + 3i
	var c128 complex128 = 2.5 + 4.7i

	// ======================================
	// 🧩 CHARACTER TYPES
	// ======================================

	/*
		byte → alias for uint8 → represents a single ASCII character (8 bits)
		rune → alias for int32 → represents a Unicode code point (32 bits)
	*/

	var ch byte = 'A'  // ASCII character (1 byte)
	var uni rune = '♄' // Unicode character (4 bytes)

	// ======================================
	// ⚙️ BOOLEAN TYPE
	// ======================================

	var flag bool = true // stored as 1 byte internally

	// ======================================
	// 🔤 STRINGS
	// ======================================

	var name string = "Sadik Al Sami"
	/*
		Strings in Go are sequences of bytes (UTF-8 encoded).
		They are immutable — you cannot modify characters directly.
	*/

	// ======================================
	// 🧾 OUTPUT EXAMPLES
	// ======================================

	fmt.Println("Signed Integers:", i8, i16, i32, i64)
	fmt.Println("Unsigned Integers:", u8, u16, u32, u64)
	fmt.Println("Architecture Dependent Int:", archInt, archUint)
	fmt.Println("Floats:", f32, f64)
	fmt.Println("Complex:", c64, c128)
	fmt.Println("Characters:", ch, string(ch), uni, string(uni))
	fmt.Println("Boolean:", flag)
	fmt.Println("String:", name)

	fmt.Println("\n🔍 Using Printf format verbs:")
	fmt.Printf("int8: %d | int16: %d | int32: %d | int64: %d\n", i8, i16, i32, i64)
	fmt.Printf("uint8: %d | uint16: %d | uint32: %d | uint64: %d\n", u8, u16, u32, u64)
	fmt.Printf("float32: %.6f | float64: %.15f\n", f32, f64)
	fmt.Printf("complex64: %v | complex128: %v\n", c64, c128)
	fmt.Printf("byte: %c | rune: %c\n", ch, uni)
	fmt.Printf("bool: %v | string: %s\n", flag, name)
	fmt.Printf("Type of 'uni' = %T | Type of 'ch' = %T | Type of 'flag' = %T\n", uni, ch, flag)

	// ======================================
	// 💾 MEMORY INSIGHT
	// ======================================

	/*
		Each data type’s size directly affects how much memory is reserved
		in RAM for that variable.

		On a 32-bit system:
			- Memory cells are 4 bytes each.
			- So int32, float32, or rune fit into one cell.
			- int64, float64 take two cells.

		On a 64-bit system:
			- Memory cells are 8 bytes each.
			- So int64, float64, rune each fit comfortably in one cell.

		Stack vs Heap:
			- Local, short-lived values (numbers, bools) → stack
			- Strings, slices, maps (dynamic structures) → heap
	*/

	/*
		🧩 Recap:
		--------------------------------------------------------
		Type        | Size (bytes) | Description
		-------------|--------------|----------------------------
		int8         | 1 byte       | signed small integer
		int16        | 2 bytes      | medium integer
		int32 (rune) | 4 bytes      | large integer / Unicode
		int64        | 8 bytes      | very large integer
		uint8 (byte) | 1 byte       | unsigned small integer
		uint16       | 2 bytes      | unsigned medium integer
		uint32       | 4 bytes      | unsigned large integer
		uint64       | 8 bytes      | unsigned very large integer
		float32      | 4 bytes      | single precision float
		float64      | 8 bytes      | double precision float
		complex64    | 8 bytes      | complex with float32 parts
		complex128   | 16 bytes     | complex with float64 parts
		bool         | 1 byte       | true/false
		byte         | 1 byte       | alias for uint8
		rune         | 4 bytes      | alias for int32
		string       | variable     | sequence of bytes (UTF-8)
	*/
}
