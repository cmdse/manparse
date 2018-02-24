package pim

import (
	"github.com/cmdse/core/schema"
	"github.com/cmdse/manparse/docbook"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func getMandocbook() *docbook.ManDocBookXml {
	mandoc, _ := docbook.Unmarshal("./test/pdftex.1.xml")
	return mandoc
}

var _ = Describe("ExtractPIMFromDocBook function", func() {
	mandoc := getMandocbook()
	index := 0
	pim := ExtractPIMFromDocBook(mandoc, nil)
	DescribeTable("loading pdftex docbook",
		func(synopsis string, expectedVariants ...*schema.OptExpressionVariant) {
			descriptions := pim.DescriptionModel()
			It("should have the expected number of option definitions", func() {
				Expect(index).To(BeNumerically("<", len(descriptions)))
			})
			descr := descriptions[index]
			index++
			foundVariants := descr.Variants()
			Expect(foundVariants).To(Equal(expectedVariants))
		},
		Entry("with synopsis -draftmode", schema.VariantX2lktSwitch),
		Entry("with synopsis -enc", schema.VariantX2lktSwitch),
		Entry("with synopsis -etex", schema.VariantX2lktSwitch),
		Entry("with synopsis -file-line-error", schema.VariantX2lktSwitch),
		Entry("with synopsis -no-file-line-error", schema.VariantX2lktSwitch),
		Entry("with synopsis -file-line-error-style", schema.VariantX2lktSwitch),
		Entry("with synopsis -fmt format", schema.VariantX2lktImplicitAssignment),
		Entry("with synopsis -halt-on-error", schema.VariantX2lktSwitch),
		Entry("with synopsis -help", schema.VariantX2lktSwitch),
		Entry("with synopsis -interaction mode", schema.VariantX2lktImplicitAssignment),
		Entry("with synopsis -ipc", schema.VariantX2lktSwitch),
		Entry("with synopsis -ipc-start", schema.VariantX2lktSwitch),
		Entry("with synopsis -jobname name", schema.VariantX2lktImplicitAssignment),
		Entry("with synopsis -kpathsea-debug bitmask", schema.VariantX2lktImplicitAssignment),
		Entry("with synopsis -mltex", schema.VariantX2lktSwitch),
		Entry("with synopsis -no-mktex fmt", schema.VariantX2lktSwitch),
		Entry("with synopsis -output-comment string", schema.VariantX2lktImplicitAssignment),
		Entry("with synopsis -output-directory directory", schema.VariantX2lktImplicitAssignment),
		Entry("with synopsis -output-format format", schema.VariantX2lktImplicitAssignment),
	)
})
