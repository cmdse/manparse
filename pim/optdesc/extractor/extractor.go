package extractor

import (
	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/reporter"
)

type Extractor struct {
	*reporter.ParseReporter
	extracts         DryOptExtracts
	synopses         optionSynopses
	descriptionModel schema.OptDescriptionModel
}

func NewExtractor(extracts RawOptExtracts, rootContext string) *Extractor {
	return &Extractor{
		ParseReporter:    reporter.NewParseReporter(rootContext),
		extracts:         dryUpOptExtract(extracts),
		descriptionModel: make(schema.OptDescriptionModel, 0, 10),
	}
}

func (extractor *Extractor) ParseExtracts() schema.OptDescriptionModel {
	extractor.makeOptionSynopses()
	extractor.handleSynopsesToSplit()
	extractor.convertSynopsesToOptDescription()
	return extractor.descriptionModel
}

func (extractor *Extractor) makeOptionSynopses() {
	extractor.synopses = newOptionSynopsesFromDryExtracts(extractor.extracts, extractor.ParseReporter)
}

func (extractor *Extractor) handleSynopsesToSplit() {
	extractor.synopses.handleSynopsesToSplit()
}

func (extractor *Extractor) convertSynopsesToOptDescription() {
	for _, synopsis := range extractor.synopses {
		optDescription := synopsis.toOptDescription()
		if optDescription != nil {
			extractor.descriptionModel = append(extractor.descriptionModel, optDescription)
		}
	}
}
