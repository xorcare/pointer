// Copyright © 2022 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Deprecated: Use built-in `new(expr)` in Go 1.26+.
// Only for Go 1.18–1.25.
//go:build go1.26

package pointer

// Of is a helper routine that allocates a new any value
// to store v and returns a pointer to it.
func Of[Value any](v Value) *Value {
	return new(v)
}
