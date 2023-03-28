package challenge2

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Challenge2", func() {

	It("Test Case 1", func() {
		value, err := diceFacesCalculator(6, 6, 6)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).To(Equal(18))
	})

	It("Test Case 2", func() {
		value, err := diceFacesCalculator(5, 5, 5)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).To(Equal(15))
	})

	It("Test Case 3", func() {
		value, err := diceFacesCalculator(1, 2, 3)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).To(Equal(3))
	})

	It("Test Case 4", func() {
		value, err := diceFacesCalculator(1, 3, 1)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).To(Equal(2))
	})

	It("Test Case 5", func() {
		value, err := diceFacesCalculator(3, 5, 3)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).To(Equal(6))
	})

	It("Test Case 6", func() {
		value, err := diceFacesCalculator(6, 5, 4)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).To(Equal(6))
	})

	It("Test Case 7", func() {
		value, err := diceFacesCalculator(4, 5, 6)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).To(Equal(6))
	})

	It("Test Case 8", func() {
		_, err := diceFacesCalculator(7, 6, 5)

		Expect(err).Should(MatchError("Dice out of number range"))
	})

	It("Test Case 9", func() {
		_, err := diceFacesCalculator(0, 6, 5)

		Expect(err).Should(MatchError("Dice out of number range"))
	})

	It("Test Case 10", func() {
		_, err := diceFacesCalculator(-1, 2, 3)

		Expect(err).Should(MatchError("Dice out of number range"))
	})

})
