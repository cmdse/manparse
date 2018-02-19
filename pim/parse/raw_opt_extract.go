package parse

import "github.com/cmdse/manparse/docbook/section"

type rawOptExtract struct {
	// something like
	// -a, --all
	optSynopsis    *section.Node
	optDescription *section.Node
}

type rawOptExtracts []*rawOptExtract
