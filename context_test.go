package resolvers

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Invocation", func() {
	Context("With Arguments", func() {
		data := context{
			Info: info{
				FieldName: "example.resolver",
			},
			Arguments: json.RawMessage(`{ "foo": "bar" }`),
		}

		It("should detect data", func() {
			Expect(data.payload()).To(Equal(json.RawMessage(`{ "foo": "bar" }`)))
		})

		It("should detect resolver", func() {
			Expect(data.resolver()).To(Equal(`example.resolver`))
		})
	})

	Context("With Headers", func() {
		data := context{
			Request: request{
				Headers: json.RawMessage(`{ "foo": "bar" }`),
			},
		}

		It("should detect headers", func() {
			Expect(data.headers()).To(Equal(json.RawMessage(`{ "foo": "bar" }`)))
		})
	})

	Context("With Identity", func() {
		identity := &Identity{
			AccountId: "123",
		}
		data := context{
			Identity: identity,
		}

		It("should detect identity", func() {
			Expect(data.identity()).To(Equal(identity))
		})
	})
})
