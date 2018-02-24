package extractor

import "github.com/cmdse/manparse/reporter"

type optionSynopses []*optionSynopsis

func newOptionSynopsesFromDryExtracts(extracts DryOptExtracts, reporter *reporter.ParseReporter) optionSynopses {
	length := len(extracts)
	optionSynopses := make(optionSynopses, length)
	for i, dry := range extracts {
		optionSynopses[i] = &optionSynopsis{
			reporter,
			dry.optSynopsis,
			splitSynopsis(dry.optSynopsis),
			dry.optDescription,
		}
	}
	return optionSynopses

}

func (synopses *optionSynopses) handleSynopsesToSplit() {
	length := len(*synopses)
	var newSynopses = make(optionSynopses, 0, length+5)
	for _, synopsis := range *synopses {
		splitSynpses := synopsis.splitSynopsisIfOptionalAssignment()
		newSynopses = append(newSynopses, splitSynpses...)
	}
	*synopses = newSynopses
}
