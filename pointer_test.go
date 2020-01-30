// Copyright © 2019-2020 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pointer

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func equal(t testing.TB, expected, actual interface{}) {
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}

	t.Logf("expected: %q, actual: %q", fmt.Sprint(expected), fmt.Sprint(actual))

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("values obtained are not equal, expected: %v, actual: %v", expected, actual)
	}
}

func TestAny(t *testing.T) {
	var _ *bool = Any(true).(*bool)
	var _ *byte = Any(byte(1)).(*byte)
	var _ *complex64 = Any(complex64(1.1)).(*complex64)
	var _ *complex128 = Any(complex128(1.1)).(*complex128)
	var _ *float32 = Any(float32(1.1)).(*float32)
	var _ *float64 = Any(float64(1.1)).(*float64)
	var _ *int = Any(int(1)).(*int)
	var _ *int8 = Any(int8(8)).(*int8)
	var _ *int16 = Any(int16(16)).(*int16)
	var _ *int32 = Any(int32(32)).(*int32)
	var _ *int64 = Any(int64(64)).(*int64)
	var _ *rune = Any(rune(1)).(*rune)
	var _ *string = Any("ptr").(*string)
	var _ *uint = Any(uint(1)).(*uint)
	var _ *uint8 = Any(uint8(8)).(*uint8)
	var _ *uint16 = Any(uint16(16)).(*uint16)
	var _ *uint32 = Any(uint32(32)).(*uint32)
	var _ *uint64 = Any(uint64(64)).(*uint64)
	var _ *uintptr = Any(uintptr(64)).(*uintptr)

	equal(t, true, *Any(true).(*bool))
	equal(t, false, *Any(false).(*bool))
	equal(t, true, *(*Any(Bool(true)).(**bool)))
	equal(t, int64(1), *Any(int64(1)).(*int64))
	equal(t, uint64(1), *Any(uint64(1)).(*uint64))
	equal(t, struct{}{}, *Any(struct{}{}).(*struct{}))
	equal(t, string("pointer"), *Any(string("pointer")).(*string))
	equal(t, reflect.Method{}, *Any(reflect.Method{}).(*reflect.Method))
	equal(t, reflect.Method{Name: "name"}, *Any(reflect.Method{Name: "name"}).(*reflect.Method))
}

func TestBool(t *testing.T) {
	equal(t, true, *Bool(true))
	equal(t, false, *Bool(false))
}

func TestByte(t *testing.T) {
	equal(t, byte(0), *Byte(byte(0)))
	equal(t, byte(1), *Byte(byte(1)))
}

func TestComplex64(t *testing.T) {
	equal(t, complex64(0), *Complex64(complex64(0)))
	equal(t, complex64(1), *Complex64(complex64(1)))
}

func TestComplex128(t *testing.T) {
	equal(t, complex128(0), *Complex128(complex128(0)))
	equal(t, complex128(1), *Complex128(complex128(1)))
}

func TestFloat32(t *testing.T) {
	equal(t, float32(-1.1), *Float32(float32(-1.1)))
	equal(t, float32(-1), *Float32(float32(-1)))
	equal(t, float32(0.0), *Float32(float32(0.0)))
	equal(t, float32(1), *Float32(float32(1)))
	equal(t, float32(1.1), *Float32(float32(1.1)))
}

func TestFloat64(t *testing.T) {
	equal(t, float64(-1.1), *Float64(float64(-1.1)))
	equal(t, float64(-1), *Float64(float64(-1)))
	equal(t, float64(0.0), *Float64(float64(0.0)))
	equal(t, float64(1), *Float64(float64(1)))
	equal(t, float64(1.1), *Float64(float64(1.1)))
}

func TestInt(t *testing.T) {
	equal(t, int(-1), *Int(int(-1)))
	equal(t, int(0), *Int(int(0)))
	equal(t, int(1), *Int(int(1)))
}

func TestInt8(t *testing.T) {
	equal(t, int8(-1), *Int8(int8(-1)))
	equal(t, int8(0), *Int8(int8(0)))
	equal(t, int8(1), *Int8(int8(1)))
}

func TestInt16(t *testing.T) {
	equal(t, int16(-1), *Int16(int16(-1)))
	equal(t, int16(0), *Int16(int16(0)))
	equal(t, int16(1), *Int16(int16(1)))
}

func TestInt32(t *testing.T) {
	equal(t, int32(-1), *Int32(int32(-1)))
	equal(t, int32(0), *Int32(int32(0)))
	equal(t, int32(1), *Int32(int32(1)))
}

func TestInt64(t *testing.T) {
	equal(t, int64(-1), *Int64(int64(-1)))
	equal(t, int64(0), *Int64(int64(0)))
	equal(t, int64(1), *Int64(int64(1)))
}

func TestRune(t *testing.T) {
	equal(t, rune(0), *Rune(rune(0)))
	equal(t, rune(1), *Rune(rune(1)))
}

func TestString(t *testing.T) {
	equal(t, string(""), *String(string("")))
	equal(t, string("pointer"), *String(string("pointer")))
}

func TestUint(t *testing.T) {
	equal(t, uint(0), *Uint(uint(0)))
	equal(t, uint(1), *Uint(uint(1)))
}

func TestUint8(t *testing.T) {
	equal(t, uint8(0), *Uint8(uint8(0)))
	equal(t, uint8(1), *Uint8(uint8(1)))
}

func TestUint16(t *testing.T) {
	equal(t, uint16(0), *Uint16(uint16(0)))
	equal(t, uint16(1), *Uint16(uint16(1)))
}

func TestUint32(t *testing.T) {
	equal(t, uint32(0), *Uint32(uint32(0)))
	equal(t, uint32(1), *Uint32(uint32(1)))
}

func TestUint64(t *testing.T) {
	equal(t, uint64(0), *Uint64(uint64(0)))
	equal(t, uint64(1), *Uint64(uint64(1)))
}

func TestUintptr(t *testing.T) {
	equal(t, uintptr(0), *Uintptr(uintptr(0)))
	equal(t, uintptr(1), *Uintptr(uintptr(1)))
}

func TestTime(t *testing.T) {
	equal(t, time.Time{}, *Time(time.Time{}))
	equal(t, time.Time{}.Add(time.Hour), *Time(time.Time{}.Add(time.Hour)))
}

func TestDuration(t *testing.T) {
	equal(t, -1*time.Hour, *Duration(-1 * time.Hour))
	equal(t, time.Duration(0), *Duration(time.Duration(0)))
	equal(t, 127*time.Hour, *Duration(127 * time.Hour))
}
