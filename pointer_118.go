// Copyright © 2022 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

package pointer

// Of is a helper routine that allocates a new any value
// to store v and returns a pointer to it.
func Of[Value any](v Value) *Value {
	return &v
}
