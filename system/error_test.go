package system_test

import (
	"github.com/go-seidon/provider/system"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error Package", func() {

	Context("Error function", Label("unit"), func() {
		When("success create error message", func() {
			It("should return result", func() {
				err := system.NewError(1001, "network error")

				Expect(err.Error()).To(Equal("network error"))
			})
		})
	})

})
