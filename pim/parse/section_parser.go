package parse

import (
	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/docbook/section"
)

type SectionParser struct {
	TargetSection     string
	aggregateExtracts func(*SectionParser, *section.Section) rawOptExtracts
}

func (parser *SectionParser) ExtractModel(sec *section.Section) schema.OptDescriptionModel {
	if parser == nil {
		return nil
	}
	rawExtracts := parser.aggregateExtracts(parser, sec)
	return parser.parseExtracts(rawExtracts)
}

func (parser *SectionParser) parseExtracts(extracts rawOptExtracts) schema.OptDescriptionModel {
	var model = make(schema.OptDescriptionModel, 0, 10)
	for _, extract := range extracts {
		model = append(model, extractToOptDescription(extract))
	}
	return model
}
