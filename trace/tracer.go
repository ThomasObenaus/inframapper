// Package trace is a simple helper for tracing
package trace

import (
	"fmt"
	"io"
)

// LogLevel represents a loglevel
type LogLevel int

const (
	// LLError loglevel error
	LLError LogLevel = iota

	// LLWarn loglevel warn
	LLWarn

	// LLInfo loglevel info
	LLInfo

	// LLTrace print all
	LLTrace
)

// Tracer is an interface for tracing
type Tracer interface {
	Error(...interface{})
	Warn(...interface{})
	Info(...interface{})
	Trace(...interface{})
}

type tracer struct {
	out      io.Writer
	loglevel LogLevel
}

func (t *tracer) Trace(a ...interface{}) {
	if t.loglevel < LLTrace {
		return
	}
	fmt.Fprint(t.out, "|trac| ")
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (t *tracer) Info(a ...interface{}) {
	if t.loglevel < LLInfo {
		return
	}
	fmt.Fprint(t.out, "|info| ")
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (t *tracer) Warn(a ...interface{}) {
	if t.loglevel < LLWarn {
		return
	}
	fmt.Fprint(t.out, "|warn| ")
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

func (t *tracer) Error(a ...interface{}) {
	fmt.Fprint(t.out, "|erro| ")
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New returns a new tracer object on loglevel error
func New(w io.Writer) Tracer {
	return NewCfgTracer(w, LLError)
}

// NewErr returns a new tracer object on loglevel error
func NewErr(w io.Writer) Tracer {
	return NewCfgTracer(w, LLError)
}

// NewWarn returns a new tracer object on loglevel warn
func NewWarn(w io.Writer) Tracer {
	return NewCfgTracer(w, LLWarn)
}

// NewInfo returns a new tracer object on loglevel info
func NewInfo(w io.Writer) Tracer {
	return NewCfgTracer(w, LLInfo)
}

// NewTrace returns a new tracer object on loglevel trace
func NewTrace(w io.Writer) Tracer {
	return NewCfgTracer(w, LLTrace)
}

// NewCfgTracer returns a new tracer object whose loglevel can be configured
func NewCfgTracer(w io.Writer, logLevel LogLevel) Tracer {
	return &tracer{
		out:      w,
		loglevel: logLevel,
	}
}

type nilTracer struct{}

func (t *nilTracer) Error(a ...interface{}) {}
func (t *nilTracer) Warn(a ...interface{})  {}
func (t *nilTracer) Info(a ...interface{})  {}
func (t *nilTracer) Trace(a ...interface{}) {}

// Off returns a silent tracer object
func Off() Tracer {
	return &nilTracer{}
}
