package reporter

import (
	"github.com/cmdse/manparse/reporter/guesses"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseReporter", func() {
	Describe("SetWriter", func() {
		reporter := NewParseReporter("foo")
		It("should not panic when passed context does match previous SetContextf returned context", func() {
			setWriter := func() {
				reporter.SetWriter(GinkgoWriter)
			}
			Expect(setWriter).NotTo(Panic())
		})
	})
	Describe("SetContextf method", func() {
		reporter := NewParseReporter("foo")
		reporter.SetWriter(GinkgoWriter)
		context := reporter.SetContextf("this is context")
		It("should add context", func() {
			Expect(reporter.lastContext()).To(Equal(context))
		})
	})
	Describe("ReleaseContext method", func() {
		reporter1 := NewParseReporter("foo")
		reporter1.SetWriter(GinkgoWriter)
		causePanic1 := func() {
			reporter1.ReleaseContext(nil)
		}
		It("should panic when called while context queue is empty", func() {
			Expect(causePanic1).To(Panic())
		})
		reporter2 := NewParseReporter("foo")
		reporter2.SetWriter(GinkgoWriter)
		reporter2.SetContextf("this is context")
		causePanic2 := func() {
			reporter2.SetContextf("Context A")
			contextB := NewParseContext("Context B", nil, 1)
			reporter2.ReleaseContext(contextB)
		}
		It("should panic when passed context does not match previous SetContextf returned context", func() {
			Expect(causePanic2).To(Panic())
		})
		reporter3 := NewParseReporter("foo")
		reporter3.SetWriter(GinkgoWriter)
		reporter3.SetContextf("this is context")
		noPanic := func() {
			contextA := reporter3.SetContextf("Context A")
			reporter3.ReleaseContext(contextA)
		}
		It("should not panic when passed context does match previous SetContextf returned context", func() {
			Expect(noPanic).ToNot(Panic())
		})
	})
	Describe("ReportGuessf method", func() {
		reporter := NewParseReporter("foo")
		reporter.ReportGuessf(guesses.OptionalImplicitAssignment, "hey, guess !!!")
		It("Should increment the number of guesses", func() {
			Expect(reporter.Guesses).To(HaveLen(1))
		})
	})
	Describe("ReportSuccessf method", func() {
		reporter := NewParseReporter("foo")
		reporter.ReportSuccessf("hey, failure !!!")
		It("Should increment the number of successes", func() {
			Expect(reporter.Successes).To(HaveLen(1))
		})
	})
	Describe("ReportFailuref method", func() {
		reporter := NewParseReporter("foo")
		reporter.ReportFailuref("hey, failure !!!")

		It("Should increment the number of failures", func() {
			Expect(reporter.Failures).To(HaveLen(1))
		})
	})
	Describe("ReportSkipf method", func() {
		reporter := NewParseReporter("foo")
		reporter.ReportSkipf("hey, skip !!!")
		It("Should increment the number of skips", func() {
			Expect(reporter.Skips).To(HaveLen(1))
		})
	})
})
