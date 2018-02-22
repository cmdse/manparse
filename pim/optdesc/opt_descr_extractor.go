package optdesc

import (
	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/docbook"
	"github.com/cmdse/manparse/docbook/section"
)

const optionSectionName = "OPTIONS"
const descriptionSectionName = "DESCRIPTION"

func secTitleIsOptionsCandidate(sec section.Section) bool {
	return sec.Title == optionSectionName || sec.Title == descriptionSectionName
}

func makeCandidates(docbook *docbook.ManDocBookXml) []*section.Section {
	candidates := make([]*section.Section, 0, 2)
	for _, sec := range docbook.Sections {
		if secTitleIsOptionsCandidate(sec) {
			candidates = append(candidates, &sec)
		}
	}
	return candidates
}

func extractOptDescriptionFromSection(sec *section.Section) (model schema.OptDescriptionModel) {
	var parser SectionParser
	switch sec.Title {
	case optionSectionName:
		parser = OptionSectionParser
	case descriptionSectionName:
		parser = DescriptionSectionParser
	}
	model = parser.ExtractModel(sec)
	return model
}

// This function does its best to extract an option description model from a docbook structure
// Returns nil if cannot find any opt description
func ExtractOptDescription(docbook *docbook.ManDocBookXml) (model schema.OptDescriptionModel) {
	candidates := makeCandidates(docbook)
	var bestMatch *section.Section
	for _, can := range candidates {
		//model = ExtractModel(can)
		bestMatch = can
		if can.Title == optionSectionName {
			// OPTIONS is the best possible match, so just
			// stop parsing other sections
			break
		}
	}
	return extractOptDescriptionFromSection(bestMatch)
}
