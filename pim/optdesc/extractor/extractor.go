package extractor

import (
	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/reporter"
	"github.com/cmdse/manparse/reporter/guesses"
)

type optionSynopsis struct {
	raw         string
	expressions []string
	description string
}

type Extractor struct {
	*reporter.ParseReporter
	extracts         DryOptExtracts
	synopses         []*optionSynopsis
	descriptionModel schema.OptDescriptionModel
}

func NewExtractor(extracts RawOptExtracts) *Extractor {
	return &Extractor{
		ParseReporter:    reporter.NewParseReporter(),
		extracts:         dryUpOptExtract(extracts),
		descriptionModel: make(schema.OptDescriptionModel, 0, 10),
	}
}

func (extractor *Extractor) ParseExtracts() schema.OptDescriptionModel {
	extractor.makeOptionSynopses()
	extractor.handleSynopsisToSplit()
	extractor.convertExtractsToOptDescription()
	return extractor.descriptionModel
}

func (extractor *Extractor) makeOptionSynopses() {
	optionSynopsises := make([]*optionSynopsis, len(extractor.extracts), len(extractor.extracts)+5)
	for i, dry := range extractor.extracts {
		optionSynopsises[i] = &optionSynopsis{
			raw:         dry.optSynopsis,
			expressions: splitSynopsis(dry.optSynopsis),
			description: dry.optDescription,
		}
	}
	extractor.synopses = optionSynopsises
}

func (extractor *Extractor) handleSynopsisToSplit() {
	var newSynopses []*optionSynopsis
	for _, synopsis := range extractor.synopses {
		extractor.SetContextf("In synopsis %v", synopsis.raw)
		if len(synopsis.expressions) == 1 {
			expr := synopsis.expressions[0]
			split, ok := splitOptExpression(expr)
			if ok {
				extractor.ReportGuessf(
					guesses.OptionalImplicitAssignment,
					"I found an option expression '%v' witch looked like an optional option assignment, so I split it to two synopsis.",
					expr)
				newSynopses = append(newSynopses,
					&optionSynopsis{
						split[0],
						[]string{split[0]},
						synopsis.description,
					},
					&optionSynopsis{
						split[1],
						[]string{split[1]},
						synopsis.description,
					})
				continue
			}
		}
		newSynopses = append(newSynopses, synopsis)
		extractor.RedeemContext()
	}
	extractor.synopses = newSynopses
}

func (extractor *Extractor) convertExtractsToOptDescription() {
	for _, synopsis := range extractor.synopses {
		extractor.SetContextf("Synopsis %v ", synopsis.raw)
		optDescription := extractor.extractToOptDescription(synopsis)
		if optDescription != nil {
			extractor.descriptionModel = append(extractor.descriptionModel, optDescription)
			extractor.ReportSuccessf("successfully extracted with option expressions variants : %v", extractor.descriptionModel.Variants())
		}
		extractor.RedeemContext()
	}
}
