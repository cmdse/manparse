package extractor

import (
	"github.com/cmdse/core/schema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Extractor struct", func() {
	Describe("matchModelsFromSynopsisString method", func() {
		DescribeTable("expected outputted MatchModels",
			func(synopsis string, expectedMatchModels ...*schema.MatchModel) {
				extractor := NewExtractor(nil)
				matchModels := extractor.matchModelsFromSynopsisString(synopsis)
				Expect(extractor.Reports).To(HaveLen(0))
				Expect(matchModels).To(HaveLen(len(expectedMatchModels)))
				for i, model := range matchModels {
					expected := expectedMatchModels[i]
					Expect(model.Variant().Name()).To(Equal(expected.Variant().Name()), "variant name")
					Expect(model.Variant()).To(Equal(expected.Variant()), "variant")
					Expect(model.FlagName()).To(Equal(expected.FlagName()), "flag name")
					Expect(model.ParamName()).To(Equal(expected.ParamName()), "param name")
				}
			},
			Entry("should handle a classic POSIX + GNU flag",
				"-L, --dereference",
				schema.NewStandaloneMatchModel(schema.VariantPOSIXShortSwitch, "L"),
				schema.NewStandaloneMatchModel(schema.VariantGNUSwitch, "dereference"),
			),
			Entry("should handle a classic POSIX switch",
				"-p",
				schema.NewStandaloneMatchModel(schema.VariantPOSIXShortSwitch, "p"),
			),
			Entry("should handle a classic X-Toolkit switch",
				"-foo",
				schema.NewStandaloneMatchModel(schema.VariantX2lktSwitch, "foo"),
			),
			Entry("should handle a quoted expression for implicit option assignment",
				"-ldflags 'flag list'",
				schema.NewAssignmentMatchModel(schema.VariantX2lktImplicitAssignment, "ldflags", "flag-list"),
			),
			Entry("should handle a classic implicit option assignment",
				"-interaction mode",
				schema.NewAssignmentMatchModel(schema.VariantX2lktImplicitAssignment, "interaction", "mode"),
			),
			Entry("should handle a classic GNU + POSIX option assignment",
				"-C file, --config-file=file",
				schema.NewAssignmentMatchModel(schema.VariantPOSIXShortAssignment, "C", "file"),
				schema.NewAssignmentMatchModel(schema.VariantGNUExplicitAssignment, "config-file", "file"),
			),
			// SPECIAL CASES
			Entry("should handle an optional option value assignment to the combination of flag + assignment",
				"--context[=CTX]",
				schema.NewStandaloneMatchModel(schema.VariantGNUSwitch, "context"),
				schema.NewAssignmentMatchModel(schema.VariantGNUExplicitAssignment, "context", "CTX"),
			),
			PEntry("should handle a GNU silent assignment + POSIX explicit option assignment such as in 'mv' man page",
				"-t, --target-directory=DIRECTORY",
				schema.NewAssignmentMatchModel(schema.VariantPOSIXShortAssignment, "t", "DIRECTORY"),
				schema.NewAssignmentMatchModel(schema.VariantGNUExplicitAssignment, "target-directory", "DIRECTORY"),
			),
			PEntry("should handle comma-separated list as option assignment values",
				"-m system[,...], --systems=system[,...]",
				schema.NewAssignmentMatchModel(schema.VariantPOSIXShortAssignment, "C", "file"),
				schema.NewAssignmentMatchModel(schema.VariantGNUExplicitAssignment, "config-file", "file"),
			),
		)
	})

})
