package mongo_test

import (
	"fmt"
	"testing"

	"github.com/go-seidon/provider/mongo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMongo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mongo Package")
}

var _ = Describe("Client Package", func() {
	Context("NewClient function", Label("unit"), func() {
		When("db mode is not valid", func() {
			It("should return error", func() {
				res, err := mongo.NewClient(
					mongo.WithConfig(mongo.ClientConfig{
						DbName: "db_name",
						DbMode: "invalid",
					}),
				)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("mode is not supported")))
			})
		})

		When("auth mode is not valid", func() {
			It("should return error", func() {
				res, err := mongo.NewClient(
					mongo.WithConfig(mongo.ClientConfig{
						DbName:   "db_name",
						AuthMode: "invalid",
					}),
					mongo.UsingStandalone("host", 27017),
				)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("auth is not supported")))
			})
		})

		When("success create standalone client", func() {
			It("should return result", func() {
				res, err := mongo.NewClient(
					mongo.WithBasicAuth("user", "pw", "db_name"),
					mongo.WithConfig(mongo.ClientConfig{
						DbName: "db_name",
					}),
					mongo.UsingStandalone("host", 27017),
				)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

		When("success create replica client", func() {
			It("should return result", func() {
				res, err := mongo.NewClient(
					mongo.WithBasicAuth("user", "pw", "db_name"),
					mongo.WithConfig(mongo.ClientConfig{
						DbName: "db_name",
					}),
					mongo.UsingReplication("rs-name", []string{"h1:27030", "h2:27031"}),
				)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
})
