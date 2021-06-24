# RingBuffer

benchmark:

```bash
$ go test -v -bench=. -benchmem
=== RUN   TestRingBuffer
--- PASS: TestRingBuffer (0.00s)
=== RUN   TestRingBuffer_One
--- PASS: TestRingBuffer_One (0.00s)
goos: windows
goarch: amd64
pkg: github.com/maolonglong/ringbuffer
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz
BenchmarkRingBuffer
BenchmarkRingBuffer-4            4637096               287.8 ns/op             0 B/op          0 allocs/op
BenchmarkSliceQueue
BenchmarkSliceQueue-4            1965145               621.2 ns/op           240 B/op          4 allocs/op
PASS
ok      github.com/maolonglong/ringbuffer       3.936s
```
