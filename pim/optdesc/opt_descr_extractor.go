package optdesc

import (
	"io"

	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/docbook"
	"github.com/cmdse/manparse/docbook/section"
)

const optionSectionName = "OPTIONS"
const descriptionSectionName = "DESCRIPTION"

func secTitleIsOptionsCandidate(sec *section.Section) bool {
	return sec.Title == optionSectionName || sec.Title == descriptionSectionName
}

func makeCandidates(docbook *docbook.ManDocBookXml) []*section.Section {
	candidates := make([]*section.Section, 0, 2)
	for _, sec := range docbook.Sections {
		if secTitleIsOptionsCandidate(sec) {
			candidates = append(candidates, sec)
		}
	}
	return candidates
}

func extractOptDescriptionFromSection(sec *section.Section, writer io.Writer) (model schema.OptDescriptionModel) {
	var parser SectionParser
	switch sec.Title {
	case optionSectionName:
		parser = OptionSectionParser
	case descriptionSectionName:
		parser = DescriptionSectionParser
	default:
		panic("unsupported section")
	}
	model = parser.ExtractModel(sec, writer)
	return model
}

func findBestMatch(docbook *docbook.ManDocBookXml) *section.Section {
	candidates := makeCandidates(docbook)
	var bestMatch *section.Section
	for _, can := range candidates {
		bestMatch = can
		if can.Title == optionSectionName {
			// OPTIONS is the best possible match, so just
			// stop parsing other sections
			break
		}
	}
	return bestMatch
}

func findSection(docbook *docbook.ManDocBookXml, sectionName string) *section.Section {
	candidates := docbook.Sections
	for _, can := range candidates {
		if can.Title == sectionName {
			return can
		}
	}
	return nil
}

// This function does its best to extract an option description model from a docbook structure.
// Writer argument can be either nil (no logging) or an io.Writer to write failures and guesses.
// Returns nil if cannot find any opt description
func ExtractOptDescription(docbook *docbook.ManDocBookXml, writer io.Writer) (model schema.OptDescriptionModel) {
	var bestMatch = findBestMatch(docbook)
	if bestMatch != nil {
		return extractOptDescriptionFromSection(bestMatch, writer)
	} else {
		return nil
	}
}
