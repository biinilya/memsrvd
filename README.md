# memsrvd

## https://github.com/bsm/redeo ##

High-performance framework for building redis-protocol compatible TCP servers/services

```
go test github.com/bsm/redeo -bench=. -benchmem
Running Suite: redeo
====================
Random Seed: 1470153349
Will run 34 of 34 specs

••••••••••••••••••••••••••••••••••
Ran 34 of 34 Specs in 0.013 seconds
SUCCESS! -- 34 Passed | 0 Failed | 0 Pending | 0 Skipped PASS
BenchmarkParseRequest_Inline-4        	 1000000	      1741 ns/op	    4304 B/op	       6 allocs/op
BenchmarkParseRequest_Bulk-4          	 1000000	      2249 ns/op	    4357 B/op	      15 allocs/op
BenchmarkResponder_WriteOK-4          	50000000	      24.2 ns/op	      11 B/op	       0 allocs/op
BenchmarkResponder_WriteNil-4         	50000000	      20.7 ns/op	      11 B/op	       0 allocs/op
BenchmarkResponder_WriteInlineString-4	20000000	       335 ns/op	     214 B/op	       0 allocs/op
BenchmarkResponder_WriteString-4      	10000000	       177 ns/op	     216 B/op	       1 allocs/op
BenchmarkResponder_WriteBytes-4       	10000000	       165 ns/op	     216 B/op	       1 allocs/op
BenchmarkResponder_WriteStringBulks-4 	 2000000	      1002 ns/op	    1085 B/op	       6 allocs/op
BenchmarkResponder_WriteBulk-4        	 2000000	       817 ns/op	    1018 B/op	       5 allocs/op
BenchmarkResponder_WriteInt-4         	20000000	       103 ns/op	      33 B/op	       1 allocs/op
ok  	github.com/bsm/redeo	24.978s
```
