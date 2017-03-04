package args_test

import (
	"fmt"

	. "github.com/dpb587/blob-receipt/cli/args"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeyValue", func() {
	var subject KeyValue

	Describe("UnmarshalFlag", func() {
		BeforeEach(func() {
			subject = KeyValue{}
		})

		It("splits by =", func() {
			err := subject.UnmarshalFlag("key1=value1")

			Expect(err).ToNot(HaveOccurred())
			Expect(subject.Key).To(Equal("key1"))
			Expect(subject.Value).To(Equal("value1"))
		})

		It("errors when = is missing", func() {
			err := subject.UnmarshalFlag("value1")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Expected key=value format: value1"))
		})

		It("keeps subsequent = in the value", func() {
			err := subject.UnmarshalFlag("key1=subkey1=subvalue1")

			Expect(err).ToNot(HaveOccurred())
			Expect(subject.Key).To(Equal("key1"))
			Expect(subject.Value).To(Equal("subkey1=subvalue1"))
		})
	})

	Describe("String", func() {
		It("stringifies", func() {
			subject = KeyValue{
				Key:   "key1",
				Value: "value1",
			}

			Expect(fmt.Sprintf("%s", subject)).To(Equal("key1=value1"))
		})
	})
})
