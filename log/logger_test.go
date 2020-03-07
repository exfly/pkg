package log

import (
	"bytes"
	"testing"
)

func TestLogger_SetOutput(t *testing.T) {
	l := NewLogger("test")
	output := bytes.NewBuffer(nil)

	l.SetOutput("", output)
	l.Error("test")
	if output.Len() == 0 {
		t.Fatal("should log")
	}
	output.Reset()
}
