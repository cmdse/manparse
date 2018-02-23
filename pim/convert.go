package pim

import (
	"io"

	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/docbook"
	"github.com/cmdse/manparse/pim/optdesc"
)

// Extract a PIM from a docbook formatted man-page
func ExtractPIMFromDocBook(docbook *docbook.ManDocBookXml, writer io.Writer) *schema.ProgramInterfaceModel {
	optionDescrModel := optdesc.ExtractOptDescription(docbook, writer)
	return schema.NewProgramInterfaceModel(optionDescrModel.Variants(), optionDescrModel)
}
