package mysql_test

import (
	"testing"

	"github.com/go-seidon/provider/mysql"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMySQL(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MySQL Package")
}

var _ = Describe("Client Package", func() {
	Context("NewClient function", Label("unit"), func() {
		When("success create client", func() {
			It("should return result", func() {
				res, err := mysql.NewClient(
					mysql.WithAuth("user", "pw"),
					mysql.WithConfig(mysql.ClientConfig{
						DbName: "db_name",
					}),
					mysql.WithLocation("host", 3306),
				)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

		When("parse time is used", func() {
			It("should return result", func() {
				res, err := mysql.NewClient(
					mysql.WithAuth("user", "pw"),
					mysql.WithConfig(mysql.ClientConfig{
						DbName: "db_name",
					}),
					mysql.WithLocation("host", 3306),
					mysql.ParseTime(),
				)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
