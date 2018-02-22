package optdesc

import (
	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/docbook/section"
	"github.com/cmdse/manparse/pim/optdesc/extractor"
)

type SectionParser struct {
	TargetSection     string
	aggregateExtracts func(*SectionParser, *section.Section) extractor.RawOptExtracts
}

func (parser *SectionParser) ExtractModel(sec *section.Section) schema.OptDescriptionModel {
	if parser == nil {
		return nil
	}
	rawExtracts := parser.aggregateExtracts(parser, sec)
	extractor := extractor.NewExtractor(rawExtracts)
	return extractor.ParseExtracts()
}
