// Package trace is a simple helper for tracing
package trace

import (
	"fmt"
	"io"
)

// Tracer is an interface for tracing
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off returns a silent tracer object
func Off() Tracer {
	return &nilTracer{}
}

// New returns a new tracer object
func New(w io.Writer) Tracer {
	return &tracer{
		out: w,
	}
}
