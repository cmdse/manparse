package extractor

import (
	"github.com/cmdse/core/schema"
)

// 1 split synopsises
// 2 parse synopsises to option descriptions

type Extractor struct {
	ParseReporter
	extracts RawOptExtracts
}

func NewExtractor(extracts RawOptExtracts) *Extractor {
	return &Extractor{
		ParseReporter: ParseReporter{},
		extracts:      extracts,
	}
}

func (extractor *Extractor) ParseExtracts() schema.OptDescriptionModel {
	extracts := extractor.extracts
	var model = make(schema.OptDescriptionModel, 0, 10)
	for _, extract := range extracts {
		optDescription := extractor.extractToOptDescription(extract)
		if optDescription != nil {
			model = append(model, optDescription)
		}
	}
	return model
}

func (extractor *Extractor) extractToOptDescription(extract *RawOptExtract) *schema.OptDescription {
	matchModels := extractor.matchModelsFromSynopsisString(extract.optSynopsis.Flatten())
	if len(matchModels) > 0 {
		return &schema.OptDescription{
			Description: extract.optDescription.Flatten(),
			MatchModels: matchModels,
		}
	} else {
		return nil
	}

}
