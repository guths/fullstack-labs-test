package challenge1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Challenge 1", func() {

	It("Test Case 1", func() {
		values := numbersFractionCalculator([]int{-4, 3, -9, 0, 4, 1})

		Expect(values).Should(HaveKeyWithValue("positives", 0.500000))
		Expect(values).Should(HaveKeyWithValue("negative", 0.333333))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.166667))
	})

	It("Test Case 2", func() {
		values := numbersFractionCalculator([]int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1})

		Expect(values).Should(HaveKeyWithValue("positives", 0.000000))
		Expect(values).Should(HaveKeyWithValue("negative", 1.000000))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.000000))
	})

	It("Test Case 3", func() {
		values := numbersFractionCalculator([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})

		Expect(values).Should(HaveKeyWithValue("positives", 1.000000))
		Expect(values).Should(HaveKeyWithValue("negative", 0.000000))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.000000))
	})

	It("Test Case 4", func() {
		values := numbersFractionCalculator([]int{4, -1, 0})

		Expect(values).Should(HaveKeyWithValue("positives", 0.333333))
		Expect(values).Should(HaveKeyWithValue("negative", 0.333333))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.333333))
	})

	It("Test Case 5", func() {
		values := numbersFractionCalculator([]int{-100, 100, 0, 0, 0, -100, 100, 0,
			-100, 100, 100, 0, 0, 0, 0, -100, -100, -100, 0, -100, 0, 100, 100, -100,
			-100, 100, 100, 100, 100, -100, -100, -100, -100, 100, 0, 0, 100, 0, 0,
			-100, -100, -100, -100, -100, -100, 100, 100, 0, 100, 100, -100, -100,
			-100, 0, 100, -100, 0, 100, 100, -100, 100, -100, 0, -100, -100, 100, 0,
			0, -100, 0, -100, -100, 100, -100, 100, 0, 100, -100, -100, -100, 100, 100,
			100, 100, 0, -100, 0, 100, 100, 100, 0, -100, -100, 0, 0, 100, 0, -100, 100, 100})

		Expect(values).Should(HaveKeyWithValue("positives", 0.340000))
		Expect(values).Should(HaveKeyWithValue("negative", 0.380000))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.280000))
	})

	It("Test Case 6", func() {
		values := numbersFractionCalculator([]int{-6, 1, 79, 17, 64, 94, 87, -77,
			0, -26, 2, -94, 87, -81, -73, -28, 43, 0, -35, 39, -37, -49, -29, 93, 64,
			54, 0, -73, -58, -100, 33, -75, -47, 35, -7, 0, 52, -37, -99, 58, -23, 70,
			-52, 18, 0, -79, -38, 45, -61, 45, 51, -48, 32, 0, -44, -56, 29, -74, -1,
			92, -93, 23, 0, 55, -31, 75, -43, 20, 19, 58, -4, 0, 59, -80, 18, -74, 81,
			63, 62, -92, 0, -23, 7, -91, 22, -1, 38, -73, 79, 0, -2, 90, 80, 74, -74,
			-85, -48, 31, 0, -80})

		Expect(values).Should(HaveKeyWithValue("positives", 0.440000))
		Expect(values).Should(HaveKeyWithValue("negative", 0.450000))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.110000))
	})

	It("Test Case 7", func() {
		values := numbersFractionCalculator([]int{0, -67, -74, -38, -72, -53, 0, -13,
			-95, -91, -100, -59, 0, -10, -68, -71, -62, -21, 0, -42, -57, -16, -66, -23,
			0, -80, -63, -68, -65, -71, 0, -71, -15, -32, -26, -8, 0, -6, -51, -87, -19,
			-71, 0, -93, -26, -35, -56, -89, 0, -21, -74, -39, -57, -8, 0, -69, -29, -24,
			-99, -53, 0, -65, -42, -72, -18, -4, 0, -73, -46, -63, -78, -87})

		Expect(values).Should(HaveKeyWithValue("positives", 0.000000))
		Expect(values).Should(HaveKeyWithValue("negative", 0.833333))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.166667))
	})

	It("Test Case 8", func() {
		values := numbersFractionCalculator([]int{-76, -33, 0, -37, 39, 0, 33, -10, 0, 57,
			59, 0, -14, -57, 0, -96, 45, 0, -60, 82, 0, -94, -31, 0, -68, -53, 0, -57, 100,
			0, 35, 81, 0, -4, 76, 0, 15, 60, 0, 2, -85, 0, 16, -70, 0, 62, -25, 0, 4, -89,
			0, -79, -80, 0, -23, 97, 0, 69, -68, 0, -83, -57, 0, 74, 1, 0, -66, 23, 0, 68,
			-80, 0, 100, 10, 0, -73, -54, 0, 0, 4, 0, -8, -86, 0, 58, -85, 0, -100, -100, 0,
			43, 61, 0, 61, -86, 0, 24, -95, 0, -54})

		Expect(values).Should(HaveKeyWithValue("positives", 0.300000))
		Expect(values).Should(HaveKeyWithValue("negative", 0.360000))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.340000))
	})

	It("Test Case 9", func() {
		values := numbersFractionCalculator([]int{-92, -21, -93, -27, -45, -63, 53, 45, 0, 6,
			-67, 84, 96, 86, 18, 36, 53, 0, 47, 88, 91, -59, 65, 62, 3, 13, 0, -49, -47, 94, -63,
			65, -23, 48, -5, 0, -10, 67, -23, 19, -11, 46, 80, -83, 0, -40, 74, -63, -20, -72, 98,
			-72, 66, 0, -58, -1, 67, -22, 8, -45, 32, -65, 0, -10, -65, -81, -36, -55, -99, -18, -82})

		Expect(values).Should(HaveKeyWithValue("positives", 0.408451))
		Expect(values).Should(HaveKeyWithValue("negative", 0.492958))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.098592))
	})

	It("Test Case 10", func() {
		values := numbersFractionCalculator([]int{0, 100, 35, 0, 94, 40, 42, 87, 59, 0})

		Expect(values).Should(HaveKeyWithValue("positives", 0.700000))
		Expect(values).Should(HaveKeyWithValue("negative", 0.000000))
		Expect(values).Should(HaveKeyWithValue("zeros", 0.300000))
	})

})
