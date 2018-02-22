package normalize

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func copyStringToInterface(strs []string) []interface{} {
	interfaces := make([]interface{}, len(strs))
	for i := range strs {
		interfaces[i] = strs[i]
	}
	return interfaces
}

var _ = FDescribeTable("normalizeFlag function", func(synopsis string, expectedMatched bool, expectedOutput []string) {
	expression, concreteMatched := normalizeFlag(synopsis)
	Expect(concreteMatched).To(Equal(expectedMatched), "didn't match expected bool return")
	Expect(expression).To(ConsistOf(copyStringToInterface(expectedOutput)...), "didn't match expected expression array")
},
	Entry("when given a synopsis with spaces, it should not succeed",
		"-L, --dereference",
		false,
		[]string{"-L, --dereference"},
	),
	// TODO implement more tests to avoid false positive
	Entry("when given a synopsis matching optional assignment value, it should succeed",
		"--context[=CTX]",
		true,
		[]string{"--context=CTX", "--context"},
	),
)
