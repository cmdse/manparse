package reporter

import (
	"fmt"
	"io"
	"strings"

	"github.com/cmdse/manparse/reporter/guesses"
)

// ParseReporter is a struct allowing to report parsing events with a dynamic scope for advanced pretty-print.
type ParseReporter struct {
	contextTree *ContextTree
	Failures    Reports
	Guesses     Reports
	Successes   Reports
	Skips       Reports
	writer      io.Writer
}

// NewParseReporter creates an instance of ParseReporter
func NewParseReporter(rootContext string) *ParseReporter {
	reporter := &ParseReporter{
		NewContextTree(rootContext),
		make(Reports, 0, 10),
		make(Reports, 0, 10),
		make(Reports, 0, 50),
		make(Reports, 0, 50),
		nil,
	}
	return reporter
}

// WriteReports pretty-print reports to the configured writer if any.
func (reporter *ParseReporter) WriteReports() {
	if reporter.writer != nil {
		fmt.Fprint(reporter.writer, reporter.contextTree.PrettyPrint())
	}
}

// SetWriter sets the writer to which childReports should be written to.
func (reporter *ParseReporter) SetWriter(writer io.Writer) {
	reporter.writer = writer
}

// SetContextf sets current context which is a dynamic logging scope.
// It can be called like fmt.Sprint, and returns the ParseContext instance
// which should be given to ReleaseContext later.
//
// Typical usage :
//
//  context := SetContextf("new context %v", stringer)
//  defer ReleaseContext(context)
func (reporter *ParseReporter) SetContextf(context string, args ...interface{}) *ParseContext {
	return reporter.contextTree.AddContext(fmt.Sprintf(context, args...))
}

// ReleaseContext release current context.
// It should be passed the ParseContext instance returned by SetContextf
//
// Typical usage :
//
//  context := SetContextf("new context %v", stringer)
//  defer ReleaseContext(context)
func (reporter *ParseReporter) ReleaseContext(contextToRel *ParseContext) {
	context := reporter.contextTree.lastContext
	if context != contextToRel {
		panic(fmt.Sprintf("ReleaseContext call mismatched the last set context with SetContext or SetContextf\n\tfound: %v\n\texpected: %v", contextToRel, context))
	}
	reporter.contextTree.ReleaseContext()
}

// ReportGuessf childReports a Guess
// It can be called like fmt.Sprint
func (reporter *ParseReporter) ReportGuessf(guess *guesses.Guess, template string, args ...interface{}) {
	message := fmt.Sprintf(template, args...)
	reporter.addGuess(fmt.Sprintf("%v\n%v", guess.Name, message))
}

// ReportSuccessf childReports a parsing success
// It can be called like fmt.Sprint
func (reporter *ParseReporter) ReportSuccessf(template string, args ...interface{}) {
	reporter.addSuccess(fmt.Sprintf(template, args...))
}

// ReportSuccessf childReports a parsing failure (unexpected course of events). Use ReportSkipf
// when the error is recoverable.
// It can be called like fmt.Sprint
func (reporter *ParseReporter) ReportFailuref(template string, args ...interface{}) {
	reporter.addFailure(fmt.Sprintf(template, args...))
}

// ReportSkipf childReports a skipped element during parsing.
// It can be called like fmt.Sprint
func (reporter *ParseReporter) ReportSkipf(template string, args ...interface{}) {
	reporter.addSkip(fmt.Sprintf(template, args...))
}

func (reporter *ParseReporter) lastContext() *ParseContext {
	return reporter.contextTree.lastContext
}

func (reporter *ParseReporter) addFailure(message string) {
	reporter.Failures.addReport(reporter, message, KindFailure)
}

func (reporter *ParseReporter) addGuess(message string) {
	reporter.Guesses.addReport(reporter, message, KindGuess)
}

func (reporter *ParseReporter) addSuccess(message string) {
	reporter.Successes.addReport(reporter, message, KindSuccess)
}

func (reporter *ParseReporter) addSkip(message string) {
	reporter.Skips.addReport(reporter, message, KindSkip)
}

func (reporter *ParseReporter) printf(offset int, template string, args ...interface{}) {
	offsetTemplate := fmt.Sprintf("%v%v\n", strings.Repeat("  ", offset), template)
	if reporter.writer != nil {
		fmt.Fprintf(reporter.writer, offsetTemplate, args...)
	}
}
