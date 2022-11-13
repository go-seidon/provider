package typeconv_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTypeconv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Typeconv Package")
}
