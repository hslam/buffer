# buffer
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/buffer)](https://pkg.go.dev/github.com/hslam/buffer)
[![Build Status](https://github.com/hslam/buffer/workflows/build/badge.svg)](https://github.com/hslam/buffer/actions)
[![codecov](https://codecov.io/gh/hslam/buffer/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/buffer)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/buffer)](https://goreportcard.com/report/github.com/hslam/buffer)
[![LICENSE](https://img.shields.io/github/license/hslam/buffer.svg?style=flat-square)](https://github.com/hslam/buffer/blob/master/LICENSE)

Package buffer implements a variable-sized bytes pool.

## Get started

### Install
```
go get github.com/hslam/buffer
```
### Import
```
import "github.com/hslam/buffer"
```
### Usage
#### Simple Example
```go
package main

import (
	"github.com/hslam/buffer"
)

func main() {
	buf := buffer.GetBuffer(1024)
	buffer.PutBuffer(buf)
}
```

#### Example
```go
package main

import (
	"github.com/hslam/buffer"
)

func main() {
	buffers := buffer.NewBuffers(1024)
	size := 65536

	buf := buffers.GetBuffer(size)
	buffers.PutBuffer(buf)

	p := buffers.AssignPool(size)
	buf = p.GetBuffer(size)
	p.PutBuffer(buf)
}
```

### Benchmark
go test -bench=. -benchmem -benchtime=10s -timeout 30m
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/buffer
BenchmarkAssignPool-8        	1000000000	         4.507 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignSizedPool-8   	699822374	        16.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuffers-8           	254406498	        46.63 ns/op	      24 B/op	       1 allocs/op
BenchmarkSizedBuffer-8       	302837005	        39.32 ns/op	      24 B/op	       1 allocs/op
PASS
ok  	github.com/hslam/buffer	51.466s
```

## License
This package is licensed under a MIT license (Copyright (c) 2021 Meng Huang)

## Author
buffer was written by Meng Huang.


