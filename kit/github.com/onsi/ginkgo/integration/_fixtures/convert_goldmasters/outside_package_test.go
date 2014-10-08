package tmp_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("something important", func() {

		whatever := &UselessStruct{}
		if whatever.ImportantField != "SECRET_PASSWORD" {
			GinkgoT().Fail()
		}
	})
})

type UselessStruct struct {
	ImportantField string
}
