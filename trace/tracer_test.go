package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("New returned nil")
	} else {
		tracer.Trace("Hello")
		if buf.String() != "Hello\n" {
			t.Errorf("Tracer should not write '%s'", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	silentTracer := Off()
	silentTracer.Trace("Hello")
}
