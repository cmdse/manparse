package reporter

type Report struct {
	Kind    *Kind
	Message string
	context *ParseContext
	offset  int
}

type Reports []*Report

func (report *Report) PrettyPrint() string {
	return report.Kind.Sprint(report.Message)
}

func (reports *Reports) addReport(reporter *ParseReporter, message string, kind *Kind) {
	lastContext := reporter.contextTree.lastContext
	report := &Report{
		Kind:    kind,
		Message: message,
		context: lastContext,
		offset:  0,
	}
	lastContext.AddChildReport(report)
	*reports = append(*reports, report)
}
