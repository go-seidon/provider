package datetime_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDatetime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Datetime Package")
}
