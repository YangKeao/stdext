# stdext

`stdext` is a library for golang that provides commonly requested functionality that the author expects to exist in the standard library but currently does not. The library is still in development, and tweaking with the golang internal is always dangerous. Use at your own risk.

## goroutine

This library provides a way to get runtime information or tweaking the go runtime. Here are the introduction to several important/funny functions.

### `goroutine.Id()`

The ability to get routine id is frequently requested by golang developers. It has been implemented several times, but only the slow way (decoding the result of `runtime.Stack`) works fine now.

This implementation decodes the ELF information of `/proc/self/exe` to obtain an offset of `runtime.tlsg` relative to the TLS pointer and the offset of the goid field in the `struct G`. It also uses an inlined assembly to load `FSBASE`. This implementation is 600x faster than decoding the result of `runtime.Stack`. Here is the result of benchmark:

```
BenchmarkGetGoroutineID-16          182718586   6.637 ns/op
BenchmarkGetGoroutineIDLegacy-16    275188      3888 ns/op
```

Please note that the current implementation has limitations and only supports linux and amd64.

## Tested Go Version

The library has been tested with the following go version:

- 1.20.1

The author is currently working on implementing automatic tests for the library, but this is proving to be difficult due to the need to use additional tools such as `delve` to verify the correctness of the return value.
