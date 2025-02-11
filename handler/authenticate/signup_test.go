package authenticate

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Signup", func() {
	Context("When the password is valid", func() {
		It("should return true", func() {
			Expect(validatePassword("!Ac123456")).ToNot(HaveOccurred())
		})
	})

	Context("When the password is missing lower case character", func() {
		It("should return false", func() {
			Expect(validatePassword("!A123456")).To(HaveOccurred())
		})
	})

	Context("When the password is missing upper case character", func() {
		It("should return false", func() {
			Expect(validatePassword("!c123456")).To(HaveOccurred())
		})
	})

	Context("When the password is missing digit", func() {
		It("should return false", func() {
			Expect(validatePassword("!Abcbcbcbcb")).To(HaveOccurred())
		})
	})

	Context("When the password is missing special character", func() {
		It("should return false", func() {
			Expect(validatePassword("Ac123456")).To(HaveOccurred())
		})
	})

	Context("When the password length is smaller than 8", func() {
		It("should return false", func() {
			Expect(validatePassword("Ac12345")).To(HaveOccurred())
		})
	})

	Context("When the password length is longer than 16", func() {
		It("should return false", func() {
			Expect(validatePassword("Ac123461111111111111111")).To(HaveOccurred())
		})
	})
})
