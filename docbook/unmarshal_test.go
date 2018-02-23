package docbook

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unmarshal function", func() {
	When("run on options", func() {
		mandoc, err := Unmarshal("./test/doclifter-options.1.xml")
		It("should not throw", func() {
			Expect(err).NotTo(HaveOccurred())
		})
		It("should return a non nil struct pointer", func() {
			Expect(mandoc).NotTo(BeNil())
		})
		It("should handle sections properly", func() {
			Expect(len(mandoc.Sections)).To(Equal(1), "sould be of length 1")
			Expect(mandoc.Sections).To(HaveLen(1))
			Expect(mandoc.Sections[0].Title).To(Equal("OPTIONS"))
		})
	})
	When("run on description", func() {
		mandoc, err := Unmarshal("./test/doclifter-description.1.xml")
		It("should not throw", func() {
			Expect(err).NotTo(HaveOccurred())
		})
		It("should return a non nil struct pointer", func() {
			Expect(mandoc).NotTo(BeNil())
		})
		It("should handle sections properly", func() {
			Expect(len(mandoc.Sections)).To(Equal(1), "sould be of length 1")
			Expect(mandoc.Sections).To(HaveLen(1))
			Expect(mandoc.Sections[0].Title).To(Equal("DESCRIPTION"))
		})
	})
	PWhen("run on full file", func() {
		mandoc, err := Unmarshal("./test/doclifter.1.xml")
		It("should not throw", func() {
			Expect(err).NotTo(HaveOccurred())
		})
		It("should return a non nil struct pointer", func() {
			Expect(mandoc).NotTo(BeNil())
		})
		It("should handle sections properly", func() {
			Expect(len(mandoc.Sections)).To(Equal(2), "sould be of length 2")
			Expect(mandoc.Sections).To(HaveLen(2))
			Expect(mandoc.Sections[0].Title).To(Equal("DESCRIPTION"))
			Expect(mandoc.Sections[1].Title).To(Equal("OPTIONS"))
		})
	})
})
