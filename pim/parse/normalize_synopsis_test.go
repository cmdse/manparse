package parse

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("normalization functions", func() {
	Describe("splitSynopsisToParts function", func() {
		DescribeTable("expected output", func(synopsis string, parts ...string) {
			Expect(splitSynopsisToParts(synopsis)).To(Equal(parts))
		},
			Entry("should return one part when synopsis is one-part", "-foo", "-foo"),
			Entry("should return two part when synopsis is two-parts", "-foo bar", "-foo", "bar"),
			Entry("should return two part when synopsis is two-parts and quoted value", "-ldflags 'flag list'", "-ldflags", "'flag list'"),
		)
	})
	Describe("findSeparator function", func() {
		DescribeTable("expected output", func(synopsis string, sep string) {
			Expect(findSeparator(synopsis)).To(Equal(sep))
		},
			Entry("should handle '=' sign", "-foo=bar", "="),
			Entry("should handle one space", "-foo bar", " "),
			Entry("should handle multiple spaces", "-foo  bar", "  "),
		)
	})
	Describe("normalizeOptSynopsis function", func() {
		DescribeTable("expected output", func(synopsis string, expectedNormalized string) {
			Expect(normalizeOptSynopsis(synopsis)).To(Equal(expectedNormalized))
		},
			Entry("should normalize single-quoted expressions", "-ldflags 'flag list'", "-ldflags flag-list"),
			Entry("should normalize double-quoted expressions", "-ldflags \"flag list\"", "-ldflags flag-list"),
			Entry("should not normalize non-quoted expressions", "-m system[,...]", "-m system[,...]"),
		)

	})

})
