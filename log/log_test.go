package log_test

import (
	"bytes"
	"fmt"

	"github.com/exfly/pkg/log"
)

func ExampleLog() {
	l := log.NewLogger("test")
	output := bytes.NewBuffer(nil)

	l.SetOutput("", output)
	l.Error("test")
	output.Reset()
	fmt.Println(string(output.Bytes()))
}
