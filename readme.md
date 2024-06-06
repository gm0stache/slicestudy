# slicestudy

### slice init test
just had to see for myself how stark the difference between different alloc and
indexing options is.
find the relevant file [here](sliceinit_test.go).

### append test
manipulating slices affects the backing array and therefore all other slices
pointing to the same backing array.
find the relevant file [here](append_test.go).

# how to run
### regular unit tests
- `go test ./...`

### benchmarks
- simple: `go test -bench=.`
- multiple executions: `go test -bench=. -count=10`
