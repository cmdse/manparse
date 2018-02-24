package optdesc

import (
	"github.com/cmdse/manparse/docbook"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("extraction functions", func() {
	Describe("findBestMatch function", func() {
		mandoc, _ := docbook.Unmarshal("./test/doclifter.1.xml")
		It("should find OPTIONS best match when available", func() {
			match := findBestMatch(mandoc)
			Expect(match).ToNot(BeNil())
			Expect(match.Title).To(Equal("OPTIONS"))
		})
	})
	Describe("ExtractOptDescription function", func() {
		mandoc, _ := docbook.Unmarshal("./test/doclifter-options.1.xml")
		It("should extract an Option description model from the OPTIONS section", func() {
			model := ExtractOptDescription(mandoc, GinkgoWriter)
			Expect(model).To(HaveLen(11))
		})
	})
})
