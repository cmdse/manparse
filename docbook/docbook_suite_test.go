package docbook_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDocbook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docbook Suite")
}
