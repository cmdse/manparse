package extractor

import "fmt"

type ParseContext string

type Report struct {
	context ParseContext
	Message string
}
type ParseReporter struct {
	context ParseContext
	Reports []*Report
}

func (reporter *ParseReporter) addReport(message string) {
	reporter.Reports = append(reporter.Reports, &Report{
		context: reporter.context,
		Message: message,
	})
}

func (reporter *ParseReporter) SetContext(context ParseContext) {
	reporter.context = context
}

func (reporter *ParseReporter) Report(message string) {
	reporter.addReport(message)
}

func (reporter *ParseReporter) Reportf(template string, args ...interface{}) {
	reporter.addReport(fmt.Sprintf(template, args...))
}

func (reporter *ParseReporter) Errorf(template string, args ...interface{}) {
	reporter.addReport(fmt.Sprintf(template, args...))
}
