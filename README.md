# zapext

Extension helpers for [zap](https://github.com/uber-go/zap).

## Installation

To install this package, use Go modules:

```sh
go get github.com/6022-labs/zapext
```

## Truncated arrays

### Truncated end array

The `TruncatedEndArray` type implements `zapcore.ArrayMarshaler` and allows logging
only the first `N` elements of a slice. Use the `ArrayTruncateEnd` helper to construct
a field:

```go
logger.Info("values", ArrayTruncateEnd("values", []int{1,2,3,4}, 2))
// values: [1,2,{"_omitted":2}]
```

### Truncated start array

The `TruncatedStartArray` type implements `zapcore.ArrayMarshaler` and allows logging
only the last `N` elements of a slice. Use the `ArrayTruncateStart` helper to construct
a field:

```go
logger.Info("values", ArrayTruncateStart("values", []int{1,2,3,4}, 2))
// values: [3,4,{"_omitted":2}]
```
