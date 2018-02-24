package reporter

import "fmt"

type Report struct {
	Kind    string
	Message string
	context ParseContext
}

type Reports []*Report

func (report *Report) Inline() string {
	return fmt.Sprintf("%v %v", report.context, report.Message)
}

func (reports *Reports) addReport(reporter *ParseReporter, message string, kind string) {
	report := &Report{
		Kind:    kind,
		Message: message,
		context: reporter.context(kind),
	}
	*reports = append(*reports, report)
	if reporter.writer != nil {
		fmt.Fprintf(reporter.writer, "%v\n", report.Inline())
	}
}
