package log_test

import (
	"fmt"
	"testing"
	"time"

	logv2 "github.com/exfly/pkg/log/v2"
	"github.com/go-kit/log"
)

func TestValueBinding(t *testing.T) {
	t.Parallel()
	var output []interface{}

	logger := log.Logger(log.LoggerFunc(func(keyvals ...interface{}) error {
		output = keyvals
		return nil
	}))

	start := time.Date(2015, time.April, 25, 0, 0, 0, 0, time.UTC)
	now := start
	mocktime := func() time.Time {
		now = now.Add(time.Second)
		return now
	}

	lc := logv2.With(logger, "ts", logv2.Timestamp(mocktime), "caller", logv2.DefaultCaller)

	lc.Log("foo", "bar")
	timestamp, ok := output[1].(time.Time)
	if !ok {
		t.Fatalf("want time.Time, have %T", output[1])
	}
	if want, have := start.Add(time.Second), timestamp; want != have {
		t.Errorf("output[1]: want %v, have %v", want, have)
	}
	if want, have := "valuer_test.go:30", fmt.Sprint(output[3]); want != have {
		t.Errorf("output[3]: want %s, have %s", want, have)
	}

	// A second attempt to confirm the bindings are truly dynamic.
	lc.Log("foo", "bar")
	timestamp, ok = output[1].(time.Time)
	if !ok {
		t.Fatalf("want time.Time, have %T", output[1])
	}
	if want, have := start.Add(2*time.Second), timestamp; want != have {
		t.Errorf("output[1]: want %v, have %v", want, have)
	}
	if want, have := "valuer_test.go:43", fmt.Sprint(output[3]); want != have {
		t.Errorf("output[3]: want %s, have %s", want, have)
	}
}
