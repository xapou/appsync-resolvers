package resolvers

import (
	"encoding/json"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Payload", func() {
	Context("Invalid JSON", func() {
		message := payload{json.RawMessage(`{""name": "example"}`)}
		example := resolver{func(args struct {
			Name string `json:"name"`
		}) error {
			return nil
		}}

		args, err := message.parse(reflect.TypeOf(example.function).In(0))

		It("should error", func() {
			Expect(err).To(HaveOccurred())
		})

		It("should return nil", func() {
			Expect(args).To(BeZero())
		})
	})

	Context("Valid JSON and resolver with 1 parameter", func() {
		argumentsMessage := payload{json.RawMessage(`{"name": "example"}`)}
		example := resolver{func(args struct {
			Name string `json:"name"`
		}) error {
			return nil
		}}

		args, err := argumentsMessage.parse(reflect.TypeOf(example.function).In(0))

		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return struct", func() {
			Expect(args).NotTo(BeNil())
		})

		It("should parse data", func() {
			Expect(args.FieldByName("Name").String()).To(Equal("example"))
		})
	})

	Context("Valid JSON and resolver with 2 parameters", func() {
		argumentsMessage := payload{json.RawMessage(`{"name": "example"}`)}
		identityMessage := payload{json.RawMessage(`{"username": "me@example.com"}`)}
		example := resolver{func(args struct {
			Name string `json:"name"`
		}, identity struct {
			Username string `json:"username"`
		}) error {
			return nil
		}}

		arguments, err := argumentsMessage.parse(reflect.TypeOf(example.function).In(0))
		identity, err := identityMessage.parse(reflect.TypeOf(example.function).In(1))

		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return struct", func() {
			Expect(arguments).NotTo(BeNil())
			Expect(identity).NotTo(BeNil())
		})

		It("should parse data", func() {
			Expect(arguments.FieldByName("Name").String()).To(Equal("example"))
			Expect(identity.FieldByName("Username").String()).To(Equal("me@example.com"))
		})
	})

	Context("Valid JSON and resolver with 3 parameters", func() {
		argumentsMessage := payload{json.RawMessage(`{"name": "example"}`)}
		sourceMessage := payload{json.RawMessage(`{"parentId": 1234}`)}
		identityMessage := payload{json.RawMessage(`{"username": "me@example.com"}`)}
		example := resolver{func(args struct {
			Name string `json:"name"`
		}, src struct {
			ParentID int `json:"parentId"`
		}, identity struct {
			Username string `json:"username"`
		}) error {
			return nil
		}}

		arguments, err := argumentsMessage.parse(reflect.TypeOf(example.function).In(0))
		source, err := sourceMessage.parse(reflect.TypeOf(example.function).In(1))
		identity, err := identityMessage.parse(reflect.TypeOf(example.function).In(2))

		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return struct", func() {
			Expect(arguments).NotTo(BeNil())
			Expect(source).NotTo(BeNil())
			Expect(identity).NotTo(BeNil())
		})

		It("should parse data", func() {
			Expect(arguments.FieldByName("Name").String()).To(Equal("example"))
			Expect(source.FieldByName("ParentID").Int()).To(Equal(int64(1234)))
			Expect(identity.FieldByName("Username").String()).To(Equal("me@example.com"))
		})
	})
})
