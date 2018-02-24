package reporter

import (
	"bytes"
	"fmt"
	"io"

	"github.com/cmdse/manparse/reporter/guesses"
)

type ParseContext string

type ParseReporter struct {
	contexts    *ContextQueue
	Failures    Reports
	Guesses     Reports
	Successes   Reports
	Skips       Reports
	writer      io.Writer
	lastContext ParseContext
}

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
	reporter.SetContext(ParseContext(rootContext))
	return reporter
}

func (reporter *ParseReporter) SetWriter(writer io.Writer) {
	reporter.writer = writer
}

func (reporter *ParseReporter) SetContext(context ParseContext) {
	reporter.contexts.Push(context)
	reporter.lastContext = context
}

func (reporter *ParseReporter) SetContextf(context string, args ...interface{}) {
	reporter.SetContext(ParseContext(fmt.Sprintf(context, args...)))
}

func (reporter *ParseReporter) RedeemContext() {
	context, ok := reporter.contexts.Pop()
	if !ok {
		panic("RedeemContext failed because no call to SetContext or SetContextf has preceded\n")
	}
	if reporter.lastContext != context {
		panic(fmt.Sprintf("RedeemContext call mismatched the last set context with SetContext or SetContextf\n\tfound: %v\n\texpected: %v", context, reporter.lastContext))
	} else {
		reporter.lastContext, _ = reporter.contexts.Peek()
	}
}

func (reporter *ParseReporter) ReportGuessf(guess *guesses.Guess, template string, args ...interface{}) {
	reporter.ReportGuess(guess, fmt.Sprintf(template, args...))
}

func (reporter *ParseReporter) ReportGuess(guess *guesses.Guess, message string) {
	reporter.addGuess(fmt.Sprintf("Guess found '%v': %v", guess.Name, message))
}

func (reporter *ParseReporter) ReportSuccessf(template string, args ...interface{}) {
	reporter.addSuccess(fmt.Sprintf(template, args...))
}

func (reporter *ParseReporter) ReportSuccess(message string) {
	reporter.addSuccess(message)
}

func (reporter *ParseReporter) ReportFailuref(template string, args ...interface{}) {
	reporter.addFailure(fmt.Sprintf(template, args...))
}

func (reporter *ParseReporter) ReportFailure(message string) {
	reporter.addFailure(message)
}

func (reporter *ParseReporter) ReportSkipf(template string, args ...interface{}) {
	reporter.addSkip(fmt.Sprintf(template, args...))
}

func (reporter *ParseReporter) ReportSkip(message string) {
	reporter.addSkip(message)
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
