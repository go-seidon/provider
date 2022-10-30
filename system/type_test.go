package system_test

import (
	"github.com/go-seidon/provider/system"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type Package", func() {

	Context("String function", Label("unit"), func() {
		When("value is not empty", func() {
			It("should return result", func() {
				res := system.String("value")

				value := "value"
				Expect(res).To(Equal(&value))
			})
		})

		When("value is empty", func() {
			It("should return result", func() {
				res := system.String("")

				empty := ""
				Expect(res).To(Equal(&empty))
			})
		})
	})

	Context("Int function", Label("unit"), func() {
		When("value is not empty", func() {
			It("should return result", func() {
				res := system.Int(10)

				value := 10
				Expect(res).To(Equal(&value))
			})
		})

		When("value is empty", func() {
			It("should return result", func() {
				res := system.Int(0)

				empty := 0
				Expect(res).To(Equal(&empty))
			})
		})
	})

	Context("Int32 function", Label("unit"), func() {
		When("value is not empty", func() {
			It("should return result", func() {
				res := system.Int32(10)

				value := int32(10)
				Expect(res).To(Equal(&value))
			})
		})

		When("value is empty", func() {
			It("should return result", func() {
				res := system.Int32(0)

				empty := int32(0)
				Expect(res).To(Equal(&empty))
			})
		})
	})

	Context("Int64 function", Label("unit"), func() {
		When("value is not empty", func() {
			It("should return result", func() {
				res := system.Int64(10)

				value := int64(10)
				Expect(res).To(Equal(&value))
			})
		})

		When("value is empty", func() {
			It("should return result", func() {
				res := system.Int64(0)

				empty := int64(0)
				Expect(res).To(Equal(&empty))
			})
		})
	})

	Context("Float32 function", Label("unit"), func() {
		When("value is not empty", func() {
			It("should return result", func() {
				res := system.Float32(10)

				value := float32(10)
				Expect(res).To(Equal(&value))
			})
		})

		When("value is empty", func() {
			It("should return result", func() {
				res := system.Float32(0)

				empty := float32(0)
				Expect(res).To(Equal(&empty))
			})
		})
	})

	Context("Float64 function", Label("unit"), func() {
		When("value is not empty", func() {
			It("should return result", func() {
				res := system.Float64(10)

				value := float64(10)
				Expect(res).To(Equal(&value))
			})
		})

		When("value is empty", func() {
			It("should return result", func() {
				res := system.Float64(0)

				empty := float64(0)
				Expect(res).To(Equal(&empty))
			})
		})
	})

	Context("Bool function", Label("unit"), func() {
		When("value is not true", func() {
			It("should return result", func() {
				res := system.Bool(true)

				true := bool(true)
				Expect(res).To(Equal(&true))
			})
		})

		When("value is false", func() {
			It("should return result", func() {
				res := system.Bool(false)

				false := bool(false)
				Expect(res).To(Equal(&false))
			})
		})
	})

})
