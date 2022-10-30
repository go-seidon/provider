package system_test

import (
	"github.com/go-seidon/provider/system"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Success Package", func() {

	Context("Constructor function", Label("unit"), func() {
		When("success create object", func() {
			It("should return result", func() {
				res := system.NewSuccess(1000, "success")

				Expect(res).To(Equal(system.SystemSuccess{
					Code:    1000,
					Message: "success",
				}))
			})
		})
	})

})
