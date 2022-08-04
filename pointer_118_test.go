// Copyright © 2022 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

package pointer

import (
	"os"
	"testing"
	"time"
)

func TestOf(t *testing.T) {
	equal(t, true, *Of(true))
	equal(t, false, *Of(false))
	equal(t, byte(1), *Of(byte(1)))
	equal(t, []byte{1}, *Of([]byte{1}))
	equal(t, complex64(1), *Of(complex64(1)))
	equal(t, complex128(1), *Of(complex128(1)))
	equal(t, float32(1.1), *Of(float32(1.1)))
	equal(t, float64(1.1), *Of(float64(1.1)))
	equal(t, int(1), *Of(int(1)))
	equal(t, int8(1), *Of(int8(1)))
	equal(t, int16(1), *Of(int16(1)))
	equal(t, int32(1), *Of(int32(1)))
	equal(t, int64(1), *Of(int64(1)))
	equal(t, rune(1), *Of(rune(1)))
	equal(t, string("pointer"), *Of(string("pointer")))
	equal(t, uint(1), *Of(uint(1)))
	equal(t, uint8(1), *Of(uint8(1)))
	equal(t, uint16(1), *Of(uint16(1)))
	equal(t, uint32(1), *Of(uint32(1)))
	equal(t, uint64(1), *Of(uint64(1)))
	equal(t, uintptr(1), *Of(uintptr(1)))
	equal(t, time.Time{}.Add(time.Hour), *Of(time.Time{}.Add(time.Hour)))
	equal(t, 127*time.Hour, *Of(127 * time.Hour))
	equal(t, os.Interrupt, *Of(os.Interrupt))
}
