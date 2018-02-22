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
