package log

import (
	"context"

	"github.com/go-kit/log"
	"github.com/pkg/errors"
)

const (
	MsgKey  = "msg"
	ErrKey  = "error"
	NameKey = "name"
)

type logger struct {
	logs      []log.Logger
	prefix    []interface{}
	hasValuer bool
	ctx       context.Context
}

func (c *logger) Log(keyvals ...interface{}) error {
	kvs := make([]interface{}, 0, len(c.prefix)+len(keyvals))
	kvs = append(kvs, c.prefix...)
	if c.hasValuer {
		bindValues(c.ctx, kvs)
	}
	kvs = append(kvs, keyvals...)
	for _, l := range c.logs {
		if err := l.Log(kvs...); err != nil {
			return err
		}
	}
	return nil
}

// ErrMissingValue is appended to keyvals slices with odd length to substitute
// the missing value.
var ErrMissingValue = errors.New("(MISSING)")

// With with logger fields.
func With(l log.Logger, kv ...interface{}) log.Logger {
	if c, ok := l.(*logger); ok {
		kvs := make([]interface{}, 0, len(c.prefix)+len(kv))
		kvs = append(kvs, kv...)
		kvs = append(kvs, c.prefix...)
		if len(kvs)%2 != 0 {
			kvs = append(kvs, ErrMissingValue)
		}
		return &logger{
			logs:      c.logs,
			prefix:    kvs,
			hasValuer: containsValuer(kvs),
			ctx:       c.ctx,
		}
	}

	return &logger{logs: []log.Logger{l}, prefix: kv, hasValuer: containsValuer(kv)}
}

// WithContext returns a shallow copy of l with its context changed
// to ctx. The provided ctx must be non-nil.
func WithContext(ctx context.Context, l log.Logger) log.Logger {
	if c, ok := l.(*logger); ok {
		return &logger{
			logs:      c.logs,
			prefix:    c.prefix,
			hasValuer: c.hasValuer,
			ctx:       ctx,
		}
	}
	return &logger{logs: []log.Logger{l}, ctx: ctx}
}

// MultiLogger wraps multi logger.
func MultiLogger(logs ...log.Logger) log.Logger {
	return &logger{logs: logs}
}
