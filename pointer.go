// Copyright © 2019, Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pointer

import "reflect"

// Any is a helper routine that allocates a new interface value
// to store v and returns a pointer to it.
func Any(v interface{}) interface{} {
	r := reflect.New(reflect.TypeOf(v))
	reflect.ValueOf(r.Interface()).Elem().Set(reflect.ValueOf(v))
	return r.Interface()
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Byte is a helper routine that allocates a new byte value
// to store v and returns a pointer to it.
func Byte(v byte) *byte { return &v }

// Complex64 is a helper routine that allocates a new complex64 value
// to store v and returns a pointer to it.
func Complex64(v complex64) *complex64 { return &v }

// Complex128 is a helper routine that allocates a new complex128 value
// to store v and returns a pointer to it.
func Complex128(v complex128) *complex128 { return &v }

// Float32 is a helper routine that allocates a new float32 value
// to store v and returns a pointer to it.
func Float32(v float32) *float32 { return &v }

// Float64 is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int8 is a helper routine that allocates a new int8 value
// to store v and returns a pointer to it.
func Int8(v int8) *int8 { return &v }

// Int16 is a helper routine that allocates a new int16 value
// to store v and returns a pointer to it.
func Int16(v int16) *int16 { return &v }

// Int32 is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it.
func Int32(v int32) *int32 { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Rune is a helper routine that allocates a new rune value
// to store v and returns a pointer to it.
func Rune(v rune) *rune { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Uint is a helper routine that allocates a new uint value
// to store v and returns a pointer to it.
func Uint(v uint) *uint { return &v }

// Uint8 is a helper routine that allocates a new uint8 value
// to store v and returns a pointer to it.
func Uint8(v uint8) *uint8 { return &v }

// Uint16 is a helper routine that allocates a new uint16 value
// to store v and returns a pointer to it.
func Uint16(v uint16) *uint16 { return &v }

// Uint32 is a helper routine that allocates a new uint32 value
// to store v and returns a pointer to it.
func Uint32(v uint32) *uint32 { return &v }

// Uint64 is a helper routine that allocates a new uint64 value
// to store v and returns a pointer to it.
func Uint64(v uint64) *uint64 { return &v }

// Uintptr is a helper routine that allocates a new uintptr value
// to store v and returns a pointer to it.
func Uintptr(v uintptr) *uintptr { return &v }
