package optdesc

import (
	"io"

	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/docbook/section"
	"github.com/cmdse/manparse/pim/optdesc/extractor"
)

type SectionParser struct {
	TargetSection     string
	aggregateExtracts func(*SectionParser, *section.Section) extractor.RawOptExtracts
}

func (parser *SectionParser) ExtractModel(sec *section.Section, writer io.Writer) schema.OptDescriptionModel {
	if parser == nil {
		return nil
	}
	rawExtracts := parser.aggregateExtracts(parser, sec)
	extractor := extractor.NewExtractor(rawExtracts)
	extractor.SetWriter(writer)
	return extractor.ParseExtracts()
}
