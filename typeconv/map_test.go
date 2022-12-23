package typeconv_test

import (
	"github.com/go-seidon/provider/typeconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map Package", func() {

	Context("MapVal function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.MapVal(nil)

				Expect(res).To(Equal(map[string]interface{}{}))
			})
		})

		When("input is empty", func() {
			It("should return empty", func() {
				input := map[string]interface{}{}
				res := typeconv.MapVal(input)

				Expect(res).To(Equal(map[string]interface{}{}))
			})
		})

		When("input is not empty", func() {
			It("should return result", func() {
				input := map[string]interface{}{
					"key": "value",
				}
				res := typeconv.MapVal(input)

				Expect(res).To(Equal(map[string]interface{}{
					"key": "value",
				}))
			})
		})

		When("input contain nil value", func() {
			It("should return result", func() {
				input := map[string]interface{}{
					"key": "value",
					"nil": nil,
				}
				res := typeconv.MapVal(input)

				Expect(res).To(Equal(map[string]interface{}{
					"key": "value",
				}))
			})
		})
	})
})
