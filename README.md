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
#### Example
```go
package main

import (
	"github.com/hslam/buffer"
)

func main() {
	buffers := buffer.NewBuffers(64)
	size := 1024
	p := buffers.AssignPool(size)
	buf := p.GetBuffer(size)
	p.PutBuffer(buf)
}
```

### Benchmark
go test -bench=. -benchmem -benchtime=10s -timeout 30m
```
goos: darwin
goarch: amd64
pkg: github.com/hslam/buffer
BenchmarkAssignPool-8                      	309195586	        37.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit2-8                 	426532843	        28.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit4-8                 	549478910	        21.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit8-8                 	611556474	        19.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit16-8                	630507688	        19.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit32-8                	652233843	        18.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit64-8                	634555623	        18.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit128-8               	630217266	        17.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit256-8               	600531289	        19.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit512-8               	618427552	        19.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit1024-8              	669723151	        18.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit2048-8              	666245268	        17.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit4096-8              	600138476	        19.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit8192-8              	606141609	        19.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignPoolUnit16384-8             	732857472	        16.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAssignSizedPool-8                 	764906876	        15.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuffers-8                         	83001972	       200 ns/op	     374 B/op	       1 allocs/op
BenchmarkBuffersUnit2-8                    	89647047	       137 ns/op	     212 B/op	       1 allocs/op
BenchmarkBuffersUnit4-8                    	98385249	       123 ns/op	     140 B/op	       1 allocs/op
BenchmarkBuffersUnit8-8                    	116590702	        98.7 ns/op	      78 B/op	       1 allocs/op
BenchmarkBuffersUnit16-8                   	112042692	        90.8 ns/op	      54 B/op	       1 allocs/op
BenchmarkBuffersUnit32-8                   	55108789	       209 ns/op	    1167 B/op	       1 allocs/op
BenchmarkBuffersUnit64-8                   	73719296	       157 ns/op	     600 B/op	       1 allocs/op
BenchmarkBuffersUnit128-8                  	82647715	       143 ns/op	     318 B/op	       1 allocs/op
BenchmarkBuffersUnit256-8                  	100000000	       120 ns/op	     175 B/op	       1 allocs/op
BenchmarkBuffersUnit512-8                  	100000000	       101 ns/op	     104 B/op	       1 allocs/op
BenchmarkBuffersUnit1024-8                 	139521472	        87.0 ns/op	      52 B/op	       1 allocs/op
BenchmarkBuffersUnit2048-8                 	149084665	        81.2 ns/op	      37 B/op	       1 allocs/op
BenchmarkBuffersUnit4096-8                 	150741712	        79.6 ns/op	      34 B/op	       1 allocs/op
BenchmarkBuffersUnit8192-8                 	152904382	        78.2 ns/op	      33 B/op	       1 allocs/op
BenchmarkBuffersUnit16384-8                	164429362	        72.5 ns/op	      32 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffers-8            	80713873	       151 ns/op	     177 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit2-8       	98389815	       121 ns/op	     222 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit4-8       	113768036	       103 ns/op	     150 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit8-8       	136869103	        83.2 ns/op	      93 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit16-8      	146680798	        77.4 ns/op	      88 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit32-8      	59754711	       196 ns/op	    1167 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit64-8      	82703866	       139 ns/op	     600 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit128-8     	94845117	       126 ns/op	     318 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit256-8     	100000000	       103 ns/op	     176 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit512-8     	142224250	        84.4 ns/op	     104 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit1024-8    	168944704	        70.7 ns/op	      53 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit2048-8    	184570704	        65.4 ns/op	      38 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit4096-8    	189369930	        64.3 ns/op	      34 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit8192-8    	190938460	        63.9 ns/op	      33 B/op	       1 allocs/op
BenchmarkAssignPoolAndBuffersUnit16384-8   	203900736	        58.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkSizedBuffer-8                     	269845520	        44.4 ns/op	      32 B/op	       1 allocs/op
PASS
ok  	github.com/hslam/buffer	739.532s
```

## License
This package is licensed under a MIT license (Copyright (c) 2021 Meng Huang)

## Author
buffer was written by Meng Huang.


