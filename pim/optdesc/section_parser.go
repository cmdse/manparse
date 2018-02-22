package optdesc

import (
	"fmt"
	"os"

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
	return parser.parseExtracts(rawExtracts)
}

func (parser *SectionParser) parseExtracts(extracts extractor.RawOptExtracts) schema.OptDescriptionModel {
	var model = make(schema.OptDescriptionModel, 0, 10)
	for _, extract := range extracts {
		optDescription, err := extractor.ExtractToOptDescription(extract)
		if err != nil {
			model = append(model, optDescription)
		} else {
			fmt.Fprint(os.Stderr, err)
		}
	}
	return model
}
