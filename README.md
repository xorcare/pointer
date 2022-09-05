# pointer

[![Go workflow status badge](https://github.com/xorcare/pointer/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/xorcare/pointer/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/xorcare/pointer/badge.svg?branch=main)](https://codecov.io/gh/xorcare/pointer)
[![Go Report Card](https://goreportcard.com/badge/github.com/xorcare/pointer)](https://goreportcard.com/report/github.com/xorcare/pointer)
[![GoDoc](https://godoc.org/github.com/xorcare/pointer?status.svg)](https://pkg.go.dev/github.com/xorcare/pointer)

Package pointer contains helper routines for simplifying the creation
of optional fields of basic type.

## Installation

```bash
go get github.com/xorcare/pointer
```

## Generics

Starting with Go 18+, you can use the `Of` method to get pointers to values of any types.

## Examples

Examples of using the library are presented on [godoc.org][GDE]
and in the [source library code][SCE].

## FAQ

| Question                                                              | Source                                               |
|-----------------------------------------------------------------------|------------------------------------------------------|
| How to set **bool** pointer in a struct literal or variable?          | `var _ *bool = pointer.Bool(true)`                   |
| How to set **byte** pointer in a struct literal or variable?          | `var _ *byte = pointer.Byte(1)`                      |
| How to set **complex64** pointer in a struct literal or variable?     | `var _ *complex64 = pointer.Complex64(1.1)`          |
| How to set **complex128** pointer in a struct literal or variable?    | `var _ *complex128 = pointer.Complex128(1.1)`        |
| How to set **float32** pointer in a struct literal or variable?       | `var _ *float32 = pointer.Float32(1.1)`              |
| How to set **float64** pointer in a struct literal or variable?       | `var _ *float64 = pointer.Float64(1.1)`              |
| How to set **int** pointer in a struct literal or variable?           | `var _ *int = pointer.Int(1)`                        |
| How to set **int8** pointer in a struct literal or variable?          | `var _ *int8 = pointer.Int8(8)`                      |
| How to set **int16** pointer in a struct literal or variable?         | `var _ *int16 = pointer.Int16(16)`                   |
| How to set **int32** pointer in a struct literal or variable?         | `var _ *int32 = pointer.Int32(32)`                   |
| How to set **int64** pointer in a struct literal or variable?         | `var _ *int64 = pointer.Int64(64)`                   |
| How to set **rune** pointer in a struct literal or variable?          | `var _ *rune = pointer.Rune(1)`                      |
| How to set **string** pointer in a struct literal or variable?        | `var _ *string = pointer.String("ptr")`              |
| How to set **uint** pointer in a struct literal or variable?          | `var _ *uint = pointer.Uint(1)`                      |
| How to set **uint8** pointer in a struct literal or variable?         | `var _ *uint8 = pointer.Uint8(8)`                    |
| How to set **uint16** pointer in a struct literal or variable?        | `var _ *uint16 = pointer.Uint16(16)`                 |
| How to set **uint32** pointer in a struct literal or variable?        | `var _ *uint32 = pointer.Uint32(32)`                 |
| How to set **uint64** pointer in a struct literal or variable?        | `var _ *uint64 = pointer.Uint64(64)`                 |
| How to set **time.Time** pointer in a struct literal or variable?     | `var _ *time.Time = pointer.Time(time.Now())`        |
| How to set **time.Duration** pointer in a struct literal or variable? | `var _ *time.Duration = pointer.Duration(time.Hour)` |

## License

© 2019-2020,2022 Vasiliy Vasilyuk <xorcare@gmail.com>

Released under the [BSD 3-Clause License].

[BSD 3-Clause License]: https://github.com/xorcare/pointer/blob/master/LICENSE 'BSD 3-Clause "New" or "Revised" License'
[GDE]: https://godoc.org/github.com/xorcare/pointer#pkg-examples 'Examples of using package pointer'
[SCE]: https://github.com/xorcare/pointer/blob/master/example_test.go 'Source code examples of using package pointer'
