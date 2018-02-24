package optdesc

import (
	"github.com/cmdse/manparse/docbook"
	"github.com/cmdse/manparse/docbook/section"
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func getOptionsSection() *section.Section {
	mandoc, _ := docbook.Unmarshal("./test/doclifter-options.1.xml")
	section := findBestMatch(mandoc)
	return section
}

var _ = Describe("OptionSectionParser", func() {
	Describe("aggregateExtracts method", func() {
		section := getOptionsSection()
		// TODO test with entrylist pattern
		// TODO add reporter to log when no pattern is met
		It("should convert to raw when meeting the sibling pattern", func() {
			rawExtracts := OptionSectionParser.AggregateExtracts(section)
			Expect(rawExtracts).To(HaveLen(11))
		})
	})
	Describe("bubbleNodes method", func() {
		section := getOptionsSection()
		// TODO test with nested subsections
		It("should work when no subsection is met", func() {
			flattenChildren := bubbleNodes(section.Children, isSection)
			Expect(flattenChildren).To(HaveLen(23))
		})
	})
})
