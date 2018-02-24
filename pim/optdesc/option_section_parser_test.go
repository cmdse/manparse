package optdesc

import (
	"github.com/cmdse/manparse/docbook"
	"github.com/cmdse/manparse/docbook/section"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func getOptionsSection(path string, finder func(*docbook.ManDocBookXml) *section.Section) *section.Section {
	mandoc, _ := docbook.Unmarshal(path)
	section := finder(mandoc)
	return section
}

var _ = Describe("OptionSectionParser", func() {
	Describe("aggregateExtracts method", func() {
		// TODO add reporter to log when no pattern is met
		It("should handle sibling pattern", func() {
			section := getOptionsSection("./test/test-extract-options.xml", findBestMatch)
			rawExtracts := OptionSectionParser.AggregateExtracts(section)
			Expect(rawExtracts).To(HaveLen(11))
		})
		It("should handle varlistentry pattern", func() {
			section := getOptionsSection("./test/test-varlist-pattern.xml", findBestMatch)
			rawExtracts := OptionSectionParser.AggregateExtracts(section)
			Expect(rawExtracts).To(HaveLen(11))
		})
	})
	Describe("bubbleNodes method", func() {
		It("should work when no subsection is met", func() {
			section := getOptionsSection("./test/test-extract-options.xml", findBestMatch)
			flattenChildren := bubbleNodes(section.Children, isSection)
			Expect(flattenChildren).To(HaveLen(23))
		})
		FIt("should work with nested subsections", func() {
			findCommands := func(docbook *docbook.ManDocBookXml) *section.Section {
				return findSection(docbook, "COMMANDS")
			}
			section := getOptionsSection("./test/test-nested-subsections.xml", findCommands)
			flattenChildren := bubbleNodes(section.Children, isSection)
			Expect(flattenChildren).To(HaveLen(11))
		})
	})
})
