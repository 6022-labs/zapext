package zapext

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TruncatedEndArray implements zapcore.ArrayMarshaler, allowing slices to be
// truncated to a maximum length when logged. If the slice is longer than Max,
// an object with an "_omitted" field is appended to indicate how many elements
// were excluded.
type TruncatedEndArray[T any] struct {
	Slice []T
	Max   int
}

// MarshalLogArray encodes the slice to the provided ArrayEncoder while
// respecting the configured maximum length.
func (t TruncatedEndArray[T]) MarshalLogArray(ae zapcore.ArrayEncoder) error {
	n := len(t.Slice)
	limit := n
	if t.Max > 0 && n > t.Max {
		limit = t.Max
	}

	for i := 0; i < limit; i++ {
		// AppendReflected handles arbitrary element types.
		ae.AppendReflected(t.Slice[i])
	}

	if limit < n {
		// Indicate how many elements were omitted.
		ae.AppendObject(zapcore.ObjectMarshalerFunc(func(oe zapcore.ObjectEncoder) error {
			oe.AddInt("_omitted", n-limit)
			return nil
		}))
	}

	return nil
}

// ArrayTruncateEnd constructs a zap Field that logs the provided slice, truncating it
// to at most max elements.
func ArrayTruncateEnd[T any](key string, s []T, max int) zapcore.Field {
	return zap.Array(key, TruncatedEndArray[T]{Slice: s, Max: max})
}
