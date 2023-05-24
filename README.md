# gonce
Generic `Once.Do` with returning an element and repeat in case of error.
Similar to [sync.Once](https://pkg.go.dev/sync#Once). 

## Usage

```go
o := Once[int64]{}
res1, err := o.Do(func() (result int64, err error) {
    return rand.Int63(), nil
})
if err != nil {
    panic(err)
}
res2, err := o.Do(func() (result int64, err error) {
    return rand.Int63(), nil
})
if err != nil {
    panic(err)
}
fmt.Printf("res1: (%T); res2: (%T); res1 == res2: %v", res1, res2, res1 == res2)
// Output: res1: (int64); res2: (int64); res1 == res2: true
```

## Benchmarks

```shell
% go test -bench=. -benchmem ./...
goos: darwin
goarch: arm64
pkg: github.com/sv-tools/gonce
BenchmarkSyncOnce-8     1000000000               0.1814 ns/op          0 B/op          0 allocs/op
BenchmarkOnce-8         1000000000               0.4282 ns/op          0 B/op          0 allocs/op
```
