package resolvers

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Invocation", func() {
	Context("With Arguments", func() {
		data := invocation{
			Resolve: "example.resolver",
			Context: context{
				Arguments: json.RawMessage(`{ "foo": "bar" }`),
			},
		}

		It("should detect data", func() {
			Expect(data.Context.Arguments).To(Equal(json.RawMessage(`{ "foo": "bar" }`)))
		})
	})

	Context("With Source", func() {
		data := invocation{
			Resolve: "example.resolver",
			Context: context{
				Source: json.RawMessage(`{ "bar": "foo" }`),
			},
		}

		It("should detect data", func() {
			Expect(data.Context.Source).To(Equal(json.RawMessage(`{ "bar": "foo" }`)))
		})
	})
})
