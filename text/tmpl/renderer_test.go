package tmpl_test

import (
	. "github.com/bbuck/dragon-mud/text/tmpl"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var templateWithBraces = "{{This}} should have {braces}"

var _ = Describe("Renderer", func() {
	Describe("Render", func() {
		Context("with simple templates", func() {
			var (
				r     Renderer
				data1 = map[string]interface{}{
					"Name": "World",
				}
				data2 = map[string]interface{}{
					"Name": "Mundo",
				}
				result1, result2 string
				err1, err2       error
			)

			Register(testTemplate, "test")
			r, _ = Template("test")

			BeforeEach(func() {
				result1, err1 = r.Render(data1)
				result2, err2 = r.Render(data2)
			})

			It("does not fail", func() {
				Ω(err1).Should(BeNil())
				Ω(err2).Should(BeNil())
			})

			It("renders to a string", func() {
				Ω(result1).Should(Equal("Hello, World!"))
			})

			It("renders with different results", func() {
				Ω(result2).Should(Equal("Hello, Mundo!"))
			})
		})

		Context("with braces that shouldn't be parsed", func() {
			var (
				r      Renderer
				result string
				err    error
				data   = map[string]interface{}{
					"This": "This",
				}
			)

			Register(templateWithBraces, "otherTest")
			r, _ = Template("otherTest")

			BeforeEach(func() {
				result, err = r.Render(data)
			})

			It("doesn't fail", func() {
				Ω(err).Should(BeNil())
			})

			It("renders correctly", func() {
				Ω(result).Should(Equal("This should have {braces}"))
			})
		})

		Context("when given a struct instead of a map", func() {
			var (
				r      Renderer
				result string
				err    error
				data   = struct {
					Name string
				}{
					Name: "Mundo",
				}
			)

			BeforeEach(func() {
				Register(testTemplate, "test")
				r, _ = Template("test")
				result, err = r.Render(data)
			})

			It("doesn't fail to render", func() {
				Ω(err).Should(BeNil())
			})

			It("renders the correct string", func() {
				Ω(result).Should(Equal("Hello, Mundo!"))
			})
		})
	})
})
