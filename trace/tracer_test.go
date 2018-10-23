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
		tracer.Error("Hello")
		tracer.Warn("Hello")
		tracer.Info("Hello")
		tracer.Trace("Hello")
		if buf.String() != "|erro| Hello\n" {
			t.Errorf("Tracer should not write '%s'", buf.String())
		}
	}
}

func TestNewErr(t *testing.T) {
	var buf bytes.Buffer
	tracer := NewErr(&buf)

	if tracer == nil {
		t.Error("New returned nil")
	} else {
		tracer.Error("Hello")
		tracer.Trace("Hello")
		if buf.String() != "|erro| Hello\n" {
			t.Errorf("Tracer should not write '%s'", buf.String())
		}
	}
}

func TestNewWarn(t *testing.T) {
	var buf bytes.Buffer
	tracer := NewWarn(&buf)

	if tracer == nil {
		t.Error("New returned nil")
	} else {
		tracer.Warn("Hello")
		tracer.Info("Hello")
		tracer.Trace("Hello")
		if buf.String() != "|warn| Hello\n" {
			t.Errorf("Tracer should not write '%s'", buf.String())
		}
	}
}

func TestNewInfo(t *testing.T) {
	var buf bytes.Buffer
	tracer := NewInfo(&buf)

	if tracer == nil {
		t.Error("New returned nil")
	} else {
		tracer.Info("Hello")
		tracer.Trace("Hello")
		if buf.String() != "|info| Hello\n" {
			t.Errorf("Tracer should not write '%s'", buf.String())
		}
	}
}

func TestNewTrace(t *testing.T) {
	var buf bytes.Buffer
	tracer := NewTrace(&buf)

	if tracer == nil {
		t.Error("New returned nil")
	} else {
		tracer.Trace("Hello")
		if buf.String() != "|trac| Hello\n" {
			t.Errorf("Tracer should not write '%s'", buf.String())
		}
	}
}
func TestOff(t *testing.T) {
	silentTracer := Off()
	silentTracer.Error("Hello")
	silentTracer.Warn("Hello")
	silentTracer.Info("Hello")
	silentTracer.Trace("Hello")
}
