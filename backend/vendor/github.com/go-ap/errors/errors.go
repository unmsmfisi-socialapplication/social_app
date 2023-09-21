package errors

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Export a number of functions or variables from package errors.
var (
	As     = errors.As
	Is     = errors.Is
	Unwrap = errors.Unwrap
)

// IncludeBacktrace is a static variable that decides if when creating an error we store the backtrace with it.
var IncludeBacktrace = true

// Err is our custom error type that can store backtrace, file and line number
type Err struct {
	m string
	c error
	t stack
}

func (e Err) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		io.WriteString(s, e.m)
		switch {
		case s.Flag('+'):
			if e.c != nil {
				io.WriteString(s, ": ")
				io.WriteString(s, fmt.Sprintf("%+s", e.c))
			}
		}
	case 'v':
		e.Format(s, 's')
		switch {
		case s.Flag('+'):
			if e.t != nil {
				io.WriteString(s, "\n\t")
				e.t.Format(s, 'v')
			}
		}
	}
}

// Error implements the error interface
func (e Err) Error() string {
	if IncludeBacktrace {
		return e.m
	}
	s := strings.Builder{}
	s.WriteString(e.m)
	if ch := errors.Unwrap(e); ch != nil {
		s.WriteString(": ")
		s.WriteString(ch.Error())
	}
	return s.String()
}

// Unwrap implements the errors.Wrapper interface
func (e Err) Unwrap() error {
	return e.c
}

// StackTrace returns the stack trace as returned by the debug.Stack function
func (e Err) StackTrace() StackTrace {
	return e.t.StackTrace()
}

// Annotatef wraps an error with new message
func Annotatef(e error, s string, args ...interface{}) *Err {
	err := wrap(e, s, args...)
	return &err
}

// Newf creaates a new error
func Newf(s string, args ...interface{}) *Err {
	err := wrap(nil, s, args...)
	return &err
}

// Errorf is an alias for Newf
func Errorf(s string, args ...interface{}) error {
	err := wrap(nil, s, args...)
	return &err
}

// As implements support for errors.As
func (e *Err) As(err interface{}) bool {
	switch x := err.(type) {
	case **Err:
		*x = e
	case *Err:
		*x = *e
	default:
		return false
	}
	return true
}

type StackTracer interface {
	StackTrace() StackTrace
}

// ancestorOfCause returns true if the caller looks to be an ancestor of the given stack
// trace. We check this by seeing whether our stack prefix-matches the cause stack, which
// should imply the error was generated directly from our goroutine.
func ancestorOfCause(ourStack stack, causeStack StackTrace) bool {
	// Stack traces are ordered such that the deepest frame is first. We'll want to check
	// for prefix matching in reverse.
	//
	// As an example, imagine we have a prefix-matching stack for ourselves:
	// [
	//   "github.com/go-ap/processing/processing.Validate",
	//   "testing.tRunner",
	//   "runtime.goexit"
	// ]
	//
	// We'll want to compare this against an error cause that will have happened further
	// down the stack. An example stack trace from such an error might be:
	// [
	//   "github.com/go-ap/errors/errors.New",
	//   "testing.tRunner",
	//   "runtime.goexit"
	// ]
	//
	// Their prefix matches, but we'll have to handle the match carefully as we need to match
	// from back to forward.

	// We can't possibly prefix match if our stack is larger than the cause stack.
	if len(ourStack) > len(causeStack) {
		return false
	}

	// We know the sizes are compatible, so compare program counters from back to front.
	for idx := 0; idx < len(ourStack); idx++ {
		if ourStack[len(ourStack)-1] != (uintptr)(causeStack[len(causeStack)-1]) {
			return false
		}
	}

	return true
}

func wrap(e error, s string, args ...interface{}) Err {
	err := Err{
		c: e,
		m: fmt.Sprintf(s, args...),
	}
	if IncludeBacktrace {
		causeStackTracer := new(StackTracer)
		// If our cause has set a stack trace, and that trace is a child of our own function
		// as inferred by prefix matching our current program counter stack, then we only want
		// to decorate the error message rather than add a redundant stack trace.
		stack := callers(2)
		if !(As(e, causeStackTracer) && ancestorOfCause(*stack, (*causeStackTracer).StackTrace())) {
			err.t = *stack
		}
	}
	return err
}
