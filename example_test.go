// Copyright © 2019, Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pointer_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xorcare/pointer"
)

func ExampleAny() {
	type example struct {
		Bool      bool
		BoolPtr   *bool
		Uint      uint
		UintPtr   *uint
		String    string
		StringPtr *string
	}
	encoder := json.NewEncoder(os.Stdout)

	fmt.Print("Example use universal function: ")
	_ = encoder.Encode(example{
		Bool:      true,
		BoolPtr:   pointer.Any(true).(*bool),
		Uint:      32,
		UintPtr:   pointer.Any(uint(64)).(*uint),
		String:    "xor",
		StringPtr: pointer.Any("care").(*string),
	})

	fmt.Print("Example use specific functions: ")
	_ = encoder.Encode(example{
		Bool:      true,
		BoolPtr:   pointer.Bool(true),
		Uint:      32,
		UintPtr:   pointer.Uint(64),
		String:    "xor",
		StringPtr: pointer.String("care"),
	})

	fmt.Print("Example of using a non-standard hack: ")
	_ = encoder.Encode(example{
		Bool:      true,
		BoolPtr:   &[]bool{true}[0],
		Uint:      32,
		UintPtr:   &[]uint{64}[0],
		String:    "xor",
		StringPtr: &[]string{"care"}[0],
	})

	// Output:
	// Example use universal function: {"Bool":true,"BoolPtr":true,"Uint":32,"UintPtr":64,"String":"xor","StringPtr":"care"}
	// Example use specific functions: {"Bool":true,"BoolPtr":true,"Uint":32,"UintPtr":64,"String":"xor","StringPtr":"care"}
	// Example of using a non-standard hack: {"Bool":true,"BoolPtr":true,"Uint":32,"UintPtr":64,"String":"xor","StringPtr":"care"}
}

func ExampleBool() {
	s := struct{ Pointer *bool }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Bool(true)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]bool{true}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *bool <nil>
	// The value set by the library: *bool true
	// The value set in elegant way: *bool true
}

func ExampleByte() {
	s := struct{ Pointer *byte }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Byte(127)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]byte{255}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *uint8 <nil>
	// The value set by the library: *uint8 127
	// The value set in elegant way: *uint8 255
}

func ExampleComplex64() {
	s := struct{ Pointer *complex64 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Complex64(complex64(55.751878))
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]complex64{complex64(48.752467)}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *complex64 <nil>
	// The value set by the library: *complex64 (55.751877+0i)
	// The value set in elegant way: *complex64 (48.75247+0i)
}

func ExampleComplex128() {
	s := struct{ Pointer *complex128 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Complex128(complex128(55.753184))
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]complex128{complex128(48.743702)}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *complex128 <nil>
	// The value set by the library: *complex128 (55.753184+0i)
	// The value set in elegant way: *complex128 (48.743702+0i)
}

func ExampleFloat32() {
	s := struct{ Pointer *float32 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Float32(float32(55.753184))
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]float32{float32(48.743702)}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *float32 <nil>
	// The value set by the library: *float32 55.753185
	// The value set in elegant way: *float32 48.743702
}

func ExampleFloat64() {
	s := struct{ Pointer *float64 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Float64(float64(55.753184))
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]float64{float64(48.743702)}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *float64 <nil>
	// The value set by the library: *float64 55.753184
	// The value set in elegant way: *float64 48.743702
}

func ExampleInt() {
	s := struct{ Pointer *int }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Int(1)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]int{2}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *int <nil>
	// The value set by the library: *int 1
	// The value set in elegant way: *int 2
}

func ExampleInt8() {
	s := struct{ Pointer *int8 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Int8(63)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]int8{127}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *int8 <nil>
	// The value set by the library: *int8 63
	// The value set in elegant way: *int8 127
}

func ExampleInt16() {
	s := struct{ Pointer *int16 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Int16(16383)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]int16{32767}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *int16 <nil>
	// The value set by the library: *int16 16383
	// The value set in elegant way: *int16 32767
}

func ExampleInt32() {
	s := struct{ Pointer *int32 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Int32(1073741823)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]int32{2147483647}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *int32 <nil>
	// The value set by the library: *int32 1073741823
	// The value set in elegant way: *int32 2147483647
}

func ExampleInt64() {
	s := struct{ Pointer *int64 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Int64(4611686018427387903)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]int64{9223372036854775807}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *int64 <nil>
	// The value set by the library: *int64 4611686018427387903
	// The value set in elegant way: *int64 9223372036854775807
}

func ExampleRune() {
	s := struct{ Pointer *rune }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Rune(120)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %3[2]d %[2]q", s.Pointer, *s.Pointer))

	s.Pointer = &[]rune{99}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %3[2]d %[2]q", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *int32 <nil>
	// The value set by the library: *int32 120 'x'
	// The value set in elegant way: *int32  99 'c'
}

func ExampleString() {
	s := struct{ Pointer *string }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.String("xor")
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]string{"care"}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *string <nil>
	// The value set by the library: *string xor
	// The value set in elegant way: *string care
}

func ExampleUint() {
	s := struct{ Pointer *uint }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Uint(1)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]uint{2}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *uint <nil>
	// The value set by the library: *uint 1
	// The value set in elegant way: *uint 2
}

func ExampleUint8() {
	s := struct{ Pointer *uint8 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Uint8(63)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]uint8{127}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *uint8 <nil>
	// The value set by the library: *uint8 63
	// The value set in elegant way: *uint8 127
}

func ExampleUint16() {
	s := struct{ Pointer *uint16 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Uint16(16383)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]uint16{32767}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *uint16 <nil>
	// The value set by the library: *uint16 16383
	// The value set in elegant way: *uint16 32767
}

func ExampleUint32() {
	s := struct{ Pointer *uint32 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Uint32(1073741823)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]uint32{2147483647}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *uint32 <nil>
	// The value set by the library: *uint32 1073741823
	// The value set in elegant way: *uint32 2147483647
}

func ExampleUint64() {
	s := struct{ Pointer *uint64 }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Uint64(4611686018427387903)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]uint64{9223372036854775807}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *uint64 <nil>
	// The value set by the library: *uint64 4611686018427387903
	// The value set in elegant way: *uint64 9223372036854775807
}

func ExampleUintptr() {
	s := struct{ Pointer *uintptr }{}
	fmt.Println(fmt.Sprintf("Zero pointer value: %T %v", s.Pointer, s.Pointer))

	s.Pointer = pointer.Uintptr(4611686018427387903)
	fmt.Println(fmt.Sprintf("The value set by the library: %T %v", s.Pointer, *s.Pointer))

	s.Pointer = &[]uintptr{9223372036854775807}[0]
	fmt.Println(fmt.Sprintf("The value set in elegant way: %T %v", s.Pointer, *s.Pointer))

	// Output:
	// Zero pointer value: *uintptr <nil>
	// The value set by the library: *uintptr 4611686018427387903
	// The value set in elegant way: *uintptr 9223372036854775807
}
