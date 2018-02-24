package optdesc

import (
	"github.com/cmdse/manparse/docbook"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("extraction functions", func() {
	Describe("findBestMatch function", func() {
		mandoc, _ := docbook.Unmarshal("./test/test-best-match.xml")
		It("should find OPTIONS best match when available", func() {
			match := findBestMatch(mandoc)
			Expect(match).ToNot(BeNil())
			Expect(match.Title).To(Equal("OPTIONS"))
		})
	})
	Describe("ExtractOptDescription function", func() {
		It("should extract an Option description model from the OPTIONS section", func() {
			mandoc, _ := docbook.Unmarshal("./test/test-extract-options.xml")
			model := ExtractOptDescription(mandoc, GinkgoWriter)
			Expect(model).To(HaveLen(11))
		})
		It("should handle optional explicit assignment expressions by splitting synopsis in two", func() {
			mandoc, _ := docbook.Unmarshal("./test/test-optional-explicit-assignment.xml")
			model := ExtractOptDescription(mandoc, GinkgoWriter)
			Expect(model).To(HaveLen(3))
		})
		It("should handle optional implicit assignment expressions by splitting synopsis in two", func() {
			mandoc, _ := docbook.Unmarshal("./test/test-optional-implicit-assignment.xml")
			model := ExtractOptDescription(mandoc, GinkgoWriter)
			Expect(model).To(HaveLen(3))
		})
	})
})
