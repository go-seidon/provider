package typeconv_test

import (
	"github.com/go-seidon/provider/typeconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Integer Package", func() {

	Context("Int function", Label("unit"), func() {
		When("input is positive", func() {
			It("should return positive", func() {
				res := typeconv.Int(2)

				expectRes := 2
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				res := typeconv.Int(-2)

				expectRes := -2
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				res := typeconv.Int(0)

				expectRes := 0
				Expect(res).To(Equal(&expectRes))
			})
		})
	})

	Context("IntVal function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return zero", func() {
				res := typeconv.IntVal(nil)

				Expect(res).To(Equal(0))
			})
		})

		When("input is positive", func() {
			It("should return positive", func() {
				input := 2
				res := typeconv.IntVal(&input)

				Expect(res).To(Equal(2))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				input := -2
				res := typeconv.IntVal(&input)

				Expect(res).To(Equal(-2))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				input := 0
				res := typeconv.IntVal(&input)

				Expect(res).To(Equal(0))
			})
		})
	})

	Context("Int32 function", Label("unit"), func() {
		When("input is positive", func() {
			It("should return positive", func() {
				res := typeconv.Int32(2)

				expectRes := int32(2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				res := typeconv.Int32(-2)

				expectRes := int32(-2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				res := typeconv.Int32(0)

				expectRes := int32(0)
				Expect(res).To(Equal(&expectRes))
			})
		})
	})

	Context("Int32Val function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return zero", func() {
				res := typeconv.Int32Val(nil)

				Expect(res).To(Equal(int32(0)))
			})
		})

		When("input is positive", func() {
			It("should return positive", func() {
				input := int32(2)
				res := typeconv.Int32Val(&input)

				Expect(res).To(Equal(int32(2)))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				input := int32(-2)
				res := typeconv.Int32Val(&input)

				Expect(res).To(Equal(int32(-2)))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				input := int32(0)
				res := typeconv.Int32Val(&input)

				Expect(res).To(Equal(int32(0)))
			})
		})
	})

	Context("Int64 function", Label("unit"), func() {
		When("input is positive", func() {
			It("should return positive", func() {
				res := typeconv.Int64(2)

				expectRes := int64(2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				res := typeconv.Int64(-2)

				expectRes := int64(-2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				res := typeconv.Int64(0)

				expectRes := int64(0)
				Expect(res).To(Equal(&expectRes))
			})
		})
	})

	Context("Int64Val function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return zero", func() {
				res := typeconv.Int64Val(nil)

				Expect(res).To(Equal(int64(0)))
			})
		})

		When("input is positive", func() {
			It("should return positive", func() {
				input := int64(2)
				res := typeconv.Int64Val(&input)

				Expect(res).To(Equal(int64(2)))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				input := int64(-2)
				res := typeconv.Int64Val(&input)

				Expect(res).To(Equal(int64(-2)))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				input := int64(0)
				res := typeconv.Int64Val(&input)

				Expect(res).To(Equal(int64(0)))
			})
		})
	})

})
