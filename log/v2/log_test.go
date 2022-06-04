package log

import (
	"bytes"
	"testing"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/require"
)

func TestWith(t *testing.T) {
	buf := &bytes.Buffer{}
	logger := log.NewLogfmtLogger(buf)

	kvs := []interface{}{"a", 123}
	lc := log.With(logger, kvs...)
	kvs[1] = 0 // With should copy its key values

	lc = log.With(lc, "b", "c") // With should stack
	err := lc.Log("msg", "message")
	require.NoError(t, err)

	require.Equal(t, "a=123 b=c msg=message\n", buf.String())

	buf.Reset()
	lc = log.WithPrefix(lc, "p", "first")
	err = lc.Log("msg", "message")
	require.NoError(t, err)
	require.Equal(t, "p=first a=123 b=c msg=message\n", buf.String())
}
