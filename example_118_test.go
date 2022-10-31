// Copyright © 2022 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

package pointer_test

import (
	"github.com/xorcare/pointer"
)

func ExampleOf() {
	// Examples for values that have default type casting.
	var _ *int = pointer.Of(0)
	var _ *bool = pointer.Of(true)
	var _ *string = pointer.Of("example")

	// Example of usage with non default type casting.
	var _ *int8 = pointer.Of(int8(0))
	var _ *int8 = pointer.Of[int8](0)
}
