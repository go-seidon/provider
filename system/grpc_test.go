package system_test

import (
	"fmt"

	"github.com/go-seidon/provider/system"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("Grpc Package", func() {

	Context("FromGrpc function", Label("unit"), func() {
		When("error is not grpc error", func() {
			It("should return default error", func() {
				err := fmt.Errorf("normal error")

				res := system.FromGrpc(err)

				Expect(res).To(Equal(&system.Error{
					Code:    1001,
					Message: "normal error",
				}))
			})
		})

		When("status is internal", func() {
			It("should return error", func() {
				err := status.Error(codes.Internal, "db error")

				res := system.FromGrpc(err)

				Expect(res).To(Equal(&system.Error{
					Code:    1001,
					Message: "failed processing request",
				}))
			})
		})

		When("status is invalid argument", func() {
			It("should return error", func() {
				err := status.Error(codes.InvalidArgument, "name is required field")

				res := system.FromGrpc(err)

				Expect(res).To(Equal(&system.Error{
					Code:    1002,
					Message: "name is required field",
				}))
			})
		})

		When("status is permission denied", func() {
			It("should return error", func() {
				err := status.Error(codes.PermissionDenied, "not allowed to access resource")

				res := system.FromGrpc(err)

				Expect(res).To(Equal(&system.Error{
					Code:    1003,
					Message: "not allowed to access resource",
				}))
			})
		})

		When("status is unauthenticated", func() {
			It("should return error", func() {
				err := status.Error(codes.Unauthenticated, "not authenticated")

				res := system.FromGrpc(err)

				Expect(res).To(Equal(&system.Error{
					Code:    1003,
					Message: "not authenticated",
				}))
			})
		})

		When("status is not found", func() {
			It("should return error", func() {
				err := status.Error(codes.NotFound, "resource not found")

				res := system.FromGrpc(err)

				Expect(res).To(Equal(&system.Error{
					Code:    1004,
					Message: "resource not found",
				}))
			})
		})

		When("status is already exists", func() {
			It("should return error", func() {
				err := status.Error(codes.AlreadyExists, "resource already exists")

				res := system.FromGrpc(err)

				Expect(res).To(Equal(&system.Error{
					Code:    1007,
					Message: "resource already exists",
				}))
			})
		})

		When("status is unavailable", func() {
			It("should return error", func() {
				err := status.Error(codes.Unavailable, "host error")

				res := system.FromGrpc(err)

				Expect(res).To(Equal(&system.Error{
					Code:    1006,
					Message: "failed communicating with the host party",
				}))
			})
		})
	})
})
