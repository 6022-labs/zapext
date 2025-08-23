package zapext

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TruncatedStartArray implements zapcore.ArrayMarshaler, allowing slices to be
// truncated from the start to a maximum length when logged. If the slice is longer than Max,
// an object with an "_omitted" field is prepended to indicate how many elements were excluded.
type TruncatedStartArray[T any] struct {
	Slice []T
	Max   int
}

// MarshalLogArray encodes the slice to the provided ArrayEncoder while
// respecting the configured maximum length, truncating from the start.
func (t TruncatedStartArray[T]) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	n := len(t.Slice)
	start := 0
	if t.Max > 0 && n > t.Max {
		start = n - t.Max
	}

	for i := start; i < n; i++ {
		ae.AppendReflected(t.Slice[i])
	}

	if start > 0 {
		// Indicate how many elements were omitted at the start.
		ae.AppendObject(zapcore.ObjectMarshalerFunc(func(oe zapcore.ObjectEncoder) error {
			oe.AddInt("_omitted", start)
			return nil
		}))
	}

	return nil
}

// ArrayTruncateStart constructs a zap Field that logs the provided slice, truncating it
// from the start to at most max elements.
func ArrayTruncateStart[T any](key string, s []T, max int) zapcore.Field {
	return zap.Array(key, TruncatedStartArray[T]{Slice: s, Max: max})
}
