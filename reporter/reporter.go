package reporter

import (
	"bytes"
	"fmt"
	"io"

	"github.com/cmdse/manparse/reporter/guesses"
)

// ParseContext is the scope in which reported events occurred.
type ParseContext string

// ParseReporter is a struct allowing to report parsing events with a dynamic scope for advanced pretty-print.
type ParseReporter struct {
	contexts    *ContextQueue
	Failures    Reports
	Guesses     Reports
	Successes   Reports
	Skips       Reports
	writer      io.Writer
	lastContext ParseContext
}

// NewParseReporter creates an instance of ParseReporter
func NewParseReporter(rootContext string) *ParseReporter {
	reporter := &ParseReporter{
		NewContextQueue(),
		make(Reports, 0, 10),
		make(Reports, 0, 10),
		make(Reports, 0, 50),
		make(Reports, 0, 50),
		nil,
		"",
	}
	reporter.SetContextf(rootContext)
	return reporter
}

// SetWriter sets the writer to which reports should be written to.
func (reporter *ParseReporter) SetWriter(writer io.Writer) {
	reporter.writer = writer
}

// SetContextf sets current context which is a dynamic logging scope.
// It can be called like fmt.Sprintf, and returns the ParseContext instance
// which should be given to ReleaseContext later.
//
// Typical usage :
//
//  context := SetContextf("new context %v", stringer)
//  defer ReleaseContext(context)
func (reporter *ParseReporter) SetContextf(context string, args ...interface{}) ParseContext {
	pcontext := ParseContext(fmt.Sprintf(context, args...))
	reporter.contexts.Push(pcontext)
	reporter.lastContext = pcontext
	return pcontext
}

// ReleaseContext release current context.
// It should be passed the ParseContext instance returned by SetContextf
//
// Typical usage :
//
//  context := SetContextf("new context %v", stringer)
//  defer ReleaseContext(context)
func (reporter *ParseReporter) ReleaseContext(contextToRel ParseContext) {
	context, ok := reporter.contexts.Pop()
	if !ok {
		panic("ReleaseContext failed because there were more calls to ReleaseContext than SetContextf \n")
	}
	if reporter.lastContext != contextToRel {
		panic(fmt.Sprintf("ReleaseContext call mismatched the last set context with SetContext or SetContextf\n\tfound: %v\n\texpected: %v", context, reporter.lastContext))
	} else {
		reporter.lastContext, _ = reporter.contexts.Peek()
	}
}

// ReportGuessf reports a Guess
// It can be called like fmt.Sprintf
func (reporter *ParseReporter) ReportGuessf(guess *guesses.Guess, template string, args ...interface{}) {
	message := fmt.Sprintf(template, args...)
	reporter.addGuess(fmt.Sprintf("Guess found '%v': %v", guess.Name, message))
}

// ReportSuccessf reports a parsing success
// It can be called like fmt.Sprintf
func (reporter *ParseReporter) ReportSuccessf(template string, args ...interface{}) {
	reporter.addSuccess(fmt.Sprintf(template, args...))
}

// ReportSuccessf reports a parsing failure (unexpected course of events). Use ReportSkipf
// when the error is recoverable.
// It can be called like fmt.Sprintf
func (reporter *ParseReporter) ReportFailuref(template string, args ...interface{}) {
	reporter.addFailure(fmt.Sprintf(template, args...))
}

// ReportSkipf reports a skipped element during parsing.
// It can be called like fmt.Sprintf
func (reporter *ParseReporter) ReportSkipf(template string, args ...interface{}) {
	reporter.addSkip(fmt.Sprintf(template, args...))
}

func (reporter *ParseReporter) context(kind string) ParseContext {
	values := reporter.contexts.Values()
	var contextBuffer bytes.Buffer
	//contextBuffer.Grow(len(values) * 18)
	contextBuffer.WriteString(kind)
	contextBuffer.WriteString(" ")
	for i, context := range values {
		contextBuffer.WriteString(string(context))
		if i < len(values)+1 {
			contextBuffer.WriteString(" → ")
		}

	}
	return ParseContext(contextBuffer.String())
}

func (reporter *ParseReporter) addFailure(message string) {
	reporter.Failures.addReport(reporter, message, "[FAILURE]")
}

func (reporter *ParseReporter) addGuess(message string) {
	reporter.Guesses.addReport(reporter, message, "[GUESS]")
}

func (reporter *ParseReporter) addSuccess(message string) {
	reporter.Successes.addReport(reporter, message, "[SUCCESS]")
}

func (reporter *ParseReporter) addSkip(message string) {
	reporter.Skips.addReport(reporter, message, "[SKIP]")
}
