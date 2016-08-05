# memsrvd

## Assumptions
* We are to create custom in-memory cache, so:
  1. We have some very special requirements
  2. Performance is important
  3. No general solution is possible
* We are to have telnet-like or HTTP-based protocol, so:
  1. It is not wise to reimplement the wheel, common is better then special
  2. Performance matters, otherwise we don't need a custom solution
* I need to write some code here, it would be a bad test case otherwise
  
## Solutions
* We don't need to actually implement anything without sample use-cases, so just create a mock storage
* Redis protocol is simple, with good performance and telnet-like at the same time. So fits perfectly.
* Golang API client is any redis client
* Make commands compatible with Redis
* If we need persistence, scaling, auth, we should consider of using [ledigo](https://github.com/siddontang/ledisdb) with custom storage plugin
* No assumptions about consistency model, so let the server be eventually-consistent
* No bulk operations requested, so no bulk support, just the native Redis pipeline

## Progress
* redis-cli compat
* redis-benchmark compat (NOTE: only a few commands are implemented)
* [TODO] Server cli options and help 
* [TODO] Tests
* [TODO] Benchmarks

## Commands implemented
* [PING](http://redis.io/commands/ping) message
* [GET](http://redis.io/commands/get) key
* [TODO TTL][SET](http://redis.io/commands/set) key value [EX seconds] [PX milliseconds]
  (NX|XX options are not supported)
* [DEL](http://redis.io/commands/del) key [key ...]
* [TODO TTL][EXPIRE](http://redis.io/commands/expire) key seconds?
* [HGET](http://redis.io/commands/hget) key field
* [HSET](http://redis.io/commands/hset) key field value
* [HDEL](http://redis.io/commands/hdel) key field [field ...]
* [HKEYS](http://redis.io/commands/hkeys) key
* [TODO][LPUSH](http://redis.io/commands/lpush) key value [value ...]
* [TODO][LPOP](http://redis.io/commands/lpop) key
* [TODO][LLEN](http://redis.io/commands/llen) key

Commands are compatible with redis-3.2

## Redis-benchmark

### Pipeline 1
```
memsrvd: redis-benchmark -p 16379 -P 1 -n 100000 -r 123123 set k__rand_int__ test
    ====== set k__rand_int__ test ======
      100000 requests completed in 1.59 seconds
      50 parallel clients
      3 bytes payload
      keep alive: 1
    
    98.87% <= 1 milliseconds
    99.93% <= 2 milliseconds
    99.98% <= 3 milliseconds
    99.99% <= 4 milliseconds
    100.00% <= 5 milliseconds
    100.00% <= 6 milliseconds
    100.00% <= 23 milliseconds
    62932.66 requests per second

redis: redis-benchmark -p 6379 -P 1 -n 100000 -r 123123 set k__rand_int__ test
    ====== set k__rand_int__ test ======
      100000 requests completed in 1.08 seconds
      50 parallel clients
      3 bytes payload
      keep alive: 1
    
    99.97% <= 2 milliseconds
    100.00% <= 2 milliseconds
    92764.38 requests per second
```

### Pipeline 10
```
memsrvd: redis-benchmark -p 16379 -P 10 -n 100000 -r 123123 set k__rand_int__ test
    ====== set k__rand_int__ test ======
     100000 requests completed in 0.66 seconds
     50 parallel clients
     3 bytes payload
     keep alive: 1
    
    1.74% <= 1 milliseconds
    11.15% <= 2 milliseconds
    68.59% <= 3 milliseconds
    90.08% <= 4 milliseconds
    96.31% <= 5 milliseconds
    98.42% <= 6 milliseconds
    98.93% <= 7 milliseconds
    99.32% <= 8 milliseconds
    99.54% <= 9 milliseconds
    99.58% <= 10 milliseconds
    99.62% <= 11 milliseconds
    99.76% <= 12 milliseconds
    99.91% <= 13 milliseconds
    99.94% <= 14 milliseconds
    99.98% <= 15 milliseconds
    100.00% <= 15 milliseconds
    151285.92 requests per second
   
redis: redis-benchmark -p 6379 -P 10 -n 100000 -r 123123 set k__rand_int__ test
    ====== set k__rand_int__ test ======
      100000 requests completed in 0.24 seconds
      50 parallel clients
      3 bytes payload
      keep alive: 1
    
    41.49% <= 1 milliseconds
    99.02% <= 2 milliseconds
    99.70% <= 3 milliseconds
    100.00% <= 3 milliseconds
    416666.69 requests per second
```
## Third-parties

### [Redeo](https://github.com/bsm/redeo)
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

### [Ctrie](github.com/Workiva/go-datastructures)
A concurrent, lock-free hash array mapped trie with efficient non-blocking snapshots. For lookups, Ctries have comparable performance to concurrent skip lists and concurrent hashmaps. One key advantage of Ctries is they are dynamically allocated. Memory consumption is always proportional to the number of keys in the Ctrie, while hashmaps typically have to grow and shrink. Lookups, inserts, and removes are O(logn).

One interesting advantage Ctries have over traditional concurrent data structures is support for lock-free, linearizable, constant-time snapshots. Most concurrent data structures do not support snapshots, instead opting for locks or requiring a quiescent state. This allows Ctries to have O(1) iterator creation and clear operations and O(logn) size retrieval.
```
go test github.com/Workiva/go-datastructures/trie/ctrie -bench=. -benchmem
PASS
BenchmarkInsert-4          	 3000000	       524 ns/op	     192 B/op	       8 allocs/op
BenchmarkLookup-4          	10000000	       191 ns/op	      52 B/op	       2 allocs/op
BenchmarkRemove-4          	10000000	       181 ns/op	      52 B/op	       2 allocs/op
BenchmarkSnapshot-4        	 3000000	       420 ns/op	     176 B/op	       7 allocs/op
BenchmarkReadOnlySnapshot-4	 5000000	       340 ns/op	     136 B/op	       5 allocs/op
ok  	github.com/Workiva/go-datastructures/trie/ctrie	11.213s
```