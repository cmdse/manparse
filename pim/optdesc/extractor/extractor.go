package extractor

import (
	"github.com/cmdse/core/schema"
)

type optionSynopsis struct {
	expressions []string
	description string
}

// 1 split synopsises
// 2 parse synopsises to option descriptions

type Extractor struct {
	ParseReporter
	extracts         DryOptExtracts
	synopses         []*optionSynopsis
	descriptionModel schema.OptDescriptionModel
}

func NewExtractor(extracts RawOptExtracts) *Extractor {
	return &Extractor{
		ParseReporter:    ParseReporter{},
		extracts:         dryUpOptExtract(extracts),
		descriptionModel: make(schema.OptDescriptionModel, 0, 10),
	}
}

func (extractor *Extractor) handleSynopsisToSplit() {
	var newSynopses []*optionSynopsis
	for _, synopsis := range extractor.synopses {
		if len(synopsis.expressions) == 1 {
			splitted, ok := splitOptExpression(synopsis.expressions[0])
			if ok {
				newSynopses = append(newSynopses,
					&optionSynopsis{
						[]string{splitted[0]},
						synopsis.description,
					},
					&optionSynopsis{
						[]string{splitted[1]},
						synopsis.description,
					})
				continue
			}
		}
		newSynopses = append(newSynopses, synopsis)
	}
	extractor.synopses = newSynopses
}

func (extractor *Extractor) convertExtractsToOptDescription() {
	for _, synopsis := range extractor.synopses {
		optDescription := extractor.extractToOptDescription(synopsis)
		if optDescription != nil {
			extractor.descriptionModel = append(extractor.descriptionModel, optDescription)
		}
	}
}

func (extractor *Extractor) ParseExtracts() schema.OptDescriptionModel {
	extractor.makeOptionSynopsises()
	extractor.handleSynopsisToSplit()
	extractor.convertExtractsToOptDescription()
	return extractor.descriptionModel
}

func (extractor *Extractor) extractToOptDescription(synopsis *optionSynopsis) *schema.OptDescription {
	matchModels := extractor.matchModelsFromSynopsis(synopsis)
	if len(matchModels) > 0 {
		return &schema.OptDescription{
			Description: synopsis.description,
			MatchModels: matchModels,
		}
	} else {
		return nil
	}

}
func (extractor *Extractor) makeOptionSynopsises() {
	optionSynopsises := make([]*optionSynopsis, len(extractor.extracts)+5)
	for i, dry := range extractor.extracts {
		optionSynopsises[i] = &optionSynopsis{
			expressions: splitSynopsis(dry.optSynopsis),
			description: dry.optDescription,
		}
	}
	extractor.synopses = optionSynopsises
}
