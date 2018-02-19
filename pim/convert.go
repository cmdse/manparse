package pim

import (
	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/docbook"
)

// Extract a PIM from a docbook formatted man-page
func ExtractPIMFromDocBook(docbook *docbook.ManDocBookXml) *schema.ProgramInterfaceModel {
	optionDescrModel := ExtractOptDescription(docbook)
	return schema.NewProgramInterfaceModel(optionDescrModel.Variants(), optionDescrModel)
}
