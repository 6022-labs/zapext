package zapext

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestTruncatedEndArray(t *testing.T) {
	var buf bytes.Buffer
	encCfg := zapcore.EncoderConfig{
		MessageKey: "msg",
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encCfg), zapcore.AddSync(&buf), zapcore.InfoLevel)
	logger := zap.New(core)

	logger.Info("test", ArrayTruncateEnd("vals", []int{1, 2, 3, 4}, 2))

	var out map[string]any
	if err := json.Unmarshal(buf.Bytes(), &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	v, ok := out["vals"].([]any)
	assert.True(t, ok)
	assert.Len(t, v, 3)
	assert.Equal(t, 1.0, v[0])
	assert.Equal(t, 2.0, v[1])

	omitted, ok := v[2].(map[string]any)
	assert.True(t, ok)
	assert.Equal(t, float64(2), omitted["_omitted"])
}
