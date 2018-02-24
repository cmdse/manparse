package optdesc

var DescriptionSectionParser = SectionParser{
	TargetSection:     descriptionSectionName,
	aggregateExtracts: OptionSectionParser.aggregateExtracts,
}
