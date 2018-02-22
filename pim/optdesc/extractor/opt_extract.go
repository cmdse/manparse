package extractor

import "github.com/cmdse/manparse/docbook/section"

type RawOptExtract struct {
	// something like
	// -a, --all
	optSynopsis    *section.Node
	optDescription *section.Node
}

func NewRawOptExtract(optSynopsis *section.Node, optDescription *section.Node) *RawOptExtract {
	return &RawOptExtract{
		optSynopsis,
		optDescription,
	}
}

type RawOptExtracts []*RawOptExtract

type DryOptExtract struct {
	optSynopsis    string
	optDescription string
}

type DryOptExtracts []*DryOptExtract

func dryUpOptExtract(rawOptExtracts RawOptExtracts) DryOptExtracts {
	dried := make(DryOptExtracts, len(rawOptExtracts))
	for i, raw := range rawOptExtracts {
		dried[i] = &DryOptExtract{
			optSynopsis:    raw.optSynopsis.Flatten(),
			optDescription: raw.optDescription.Flatten(),
		}
	}
	return dried
}
