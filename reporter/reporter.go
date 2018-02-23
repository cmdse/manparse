package reporter

import (
	"bytes"
	"fmt"
	"io"

	"github.com/cmdse/manparse/reporter/guesses"
	"github.com/emirpasic/gods/stacks/arraystack"
)

type ParseContext string

type ParseReporter struct {
	contexts  *arraystack.Stack
	Failures  Reports
	Guesses   Reports
	Successes Reports
	writer    io.Writer
}

func NewParseReporter() *ParseReporter {
	return &ParseReporter{
		arraystack.New(),
		make(Reports, 0, 10),
		make(Reports, 0, 10),
		make(Reports, 0, 50),
		nil,
	}
}

func (reporter *ParseReporter) SetWriter(writer io.Writer) {
	reporter.writer = writer
}

func (reporter *ParseReporter) SetContext(context ParseContext) {
	reporter.contexts.Push(context)
}

func (reporter *ParseReporter) SetContextf(context string, args ...interface{}) {
	reporter.contexts.Push(ParseContext(fmt.Sprintf(context, args...)))
}

func (reporter *ParseReporter) RedeemContext() {
	reporter.contexts.Pop()
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

func (reporter *ParseReporter) context(kind string) ParseContext {
	values := reporter.contexts.Values()
	var contextBuffer bytes.Buffer
	contextBuffer.Grow(len(values) * 18)
	contextBuffer.WriteString(kind)
	contextBuffer.WriteString(" ")
	for i, val := range values {
		if context, ok := val.(ParseContext); ok {
			contextBuffer.WriteString(string(context))
			if i < len(values)+1 {
				contextBuffer.WriteString(" → ")
			}
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
