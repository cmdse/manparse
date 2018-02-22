package extractor

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("splitSynopsis", func() {
	DescribeTable("expected output",
		func(synopsis string, expectedGroups ...string) {
			groups := splitSynopsis(synopsis)
			Expect(groups).To(Equal(expectedGroups))
		},
		Entry("should split a classic POSIX + GNU flag",
			"-L, --dereference",
			"-L",
			"--dereference",
		),
		Entry("should split comma-separated list as option assignment values",
			"-m system[,...], --systems=system[,...]",
			"-m system[,...]",
			"--systems=system[,...]",
		),
		Entry("should split comma-separated list as option assignment values",
			"-C file, --config-file=file",
			"-C file",
			"--config-file=file",
		),
		Entry("should split a GNU silent assignment + POSIX explicit option assignment",
			"-t, --target-directory=DIRECTORY",
			"-t",
			"--target-directory=DIRECTORY",
		),
		Entry("should not split a single X-Toolkit assignment",
			"-interaction mode",
			"-interaction mode",
		),
		Entry("should not split a single quoted expression for implicit option assignment",
			"-ldflags 'flag list'",
			"-ldflags 'flag list'",
		),
		Entry("should not split a single explicit assignment with square brackets",
			"--context[=CTX]",
			"--context[=CTX]",
		),
	)
})
