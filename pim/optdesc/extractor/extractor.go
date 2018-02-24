package extractor

import (
	"strings"

	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/reporter"
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

func (extractor *Extractor) handleSynopsesToSplit() {
	var newSynopses []*optionSynopsis
	for _, synopsis := range extractor.synopses {
		newSynopses = extractor.splitSynopsisIfOptionalAssignment(synopsis)
	}
	extractor.synopses = newSynopses
}

func formatVariantNames(variants []*schema.OptExpressionVariant) string {
	names := make([]string, len(variants))
	for i, variant := range variants {
		names[i] = variant.Name()
	}
	return strings.Join(names, ", ")
}

func (extractor *Extractor) convertExtractsToOptDescription() {
	for _, synopsis := range extractor.synopses {
		extractor.SetContextf("[extract] Synopsis '%v'", synopsis.raw)
		optDescription := extractor.extractToOptDescription(synopsis)
		if optDescription != nil {
			extractor.descriptionModel = append(extractor.descriptionModel, optDescription)
			variantNames := formatVariantNames(extractor.descriptionModel.Variants())
			extractor.ReportSuccessf("extracted with option expression variant(s) : '%v'", variantNames)
		}
		extractor.RedeemContext()
	}
}
