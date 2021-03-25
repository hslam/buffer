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
BenchmarkAssignPool-8                          	272420576	        44.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize2-8                 	364565925	        32.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize4-8                 	476849100	        25.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize8-8                 	547407158	        21.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize16-8                	565802832	        21.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize32-8                	692300772	        17.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize64-8                	692946682	        17.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize128-8               	697343150	        17.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize256-8               	707399065	        16.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize512-8               	706726941	        16.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize1024-8              	703273092	        17.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize2048-8              	706539492	        16.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize4096-8              	685306758	        17.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize8192-8              	699159578	        17.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolPageSize16384-8             	710108792	        16.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignSizedPool-8                     	708374830	        16.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuffers-8                             	63472573	       186 ns/op	     187 B/op	       1 allocs/op
BenchmarkBuffersPageSize2-8                    	74870978	       142 ns/op	     136 B/op	       1 allocs/op
BenchmarkBuffersPageSize4-8                    	79931690	       130 ns/op	     146 B/op	       1 allocs/op
BenchmarkBuffersPageSize8-8                    	116281663	       105 ns/op	      91 B/op	       1 allocs/op
BenchmarkBuffersPageSize16-8                   	106896433	        96.2 ns/op	      57 B/op	       1 allocs/op
BenchmarkBuffersPageSize32-8                   	55510353	       206 ns/op	    1167 B/op	       1 allocs/op
BenchmarkBuffersPageSize64-8                   	75555468	       151 ns/op	     600 B/op	       1 allocs/op
BenchmarkBuffersPageSize128-8                  	83812864	       139 ns/op	     317 B/op	       1 allocs/op
BenchmarkBuffersPageSize256-8                  	100000000	       118 ns/op	     175 B/op	       1 allocs/op
BenchmarkBuffersPageSize512-8                  	124124292	        97.0 ns/op	     104 B/op	       1 allocs/op
BenchmarkBuffersPageSize1024-8                 	144947096	        83.3 ns/op	      53 B/op	       1 allocs/op
BenchmarkBuffersPageSize2048-8                 	154371242	        77.6 ns/op	      37 B/op	       1 allocs/op
BenchmarkBuffersPageSize4096-8                 	157944656	        76.1 ns/op	      34 B/op	       1 allocs/op
BenchmarkBuffersPageSize8192-8                 	157254658	        76.3 ns/op	      33 B/op	       1 allocs/op
BenchmarkBuffersPageSize16384-8                	158528670	        75.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffers-8                	85014500	       136 ns/op	     129 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize2-8       	89767378	       128 ns/op	     223 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize4-8       	97410897	       110 ns/op	     166 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize8-8       	138046677	        85.7 ns/op	      76 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize16-8      	138408412	        79.3 ns/op	      69 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize32-8      	173643642	        70.5 ns/op	      74 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize64-8      	84282126	       136 ns/op	     600 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize128-8     	99031294	       123 ns/op	     317 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize256-8     	100000000	       101 ns/op	     175 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize512-8     	143647999	        83.0 ns/op	     104 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize1024-8    	173135672	        69.5 ns/op	      53 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize2048-8    	188456967	        63.8 ns/op	      37 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize4096-8    	193952508	        61.8 ns/op	      34 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize8192-8    	195847075	        61.2 ns/op	      33 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersPageSize16384-8   	198282196	        60.7 ns/op	      32 B/op	       1 allocs/op
BenchmarkSizedBuffer-8                         	270007083	        44.4 ns/op	      32 B/op	       1 allocs/op
PASS
ok  	github.com/hslam/buffer	763.181s
```

## License
This package is licensed under a MIT license (Copyright (c) 2021 Meng Huang)

## Author
buffer was written by Meng Huang.


