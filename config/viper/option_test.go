package viper_test

import (
	"github.com/go-seidon/provider/config/viper"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logging Package", func() {

	Context("WithFileName function", Label("unit"), func() {
		When("parameter is specified", func() {
			It("should return result", func() {
				opt := viper.WithFileName(".env")
				var res viper.ConfigOption
				opt(&res)

				Expect(res.FileName).To(Equal(".env"))
			})
		})
	})

})
