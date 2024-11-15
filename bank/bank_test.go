package bank_test

import (
	"exercise_testing/bank"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestData struct {
	Spec   string // deskripsi dari test case
	Input  float64
	Expect float64
}

var _ = Describe("Bank", func() {
	var account *bank.Account

	BeforeEach(func() {
		account = bank.NewAccount("123456789", "Salwa Az-Zahra", 50.00, 2024)
	})

	// USING TESTDATA
	Describe("when depositing an amount", func() {
		var testsSuccess = []TestData{
			{"must add balance with amount 175.00", 175.00, 225.00},
			{"must add balance with amount 25.00", 25.00, 75.00},
			{"must add balance with amount 62.25", 62.25, 112.25},
		}

		for _, test := range testsSuccess {
			Context(test.Spec, func() {
				It("should increase the balance by the given amount", func() {
					result := account.Deposit(test.Input)

					Expect(result).To(BeNil())
					Expect(account.GetBalance()).To(Equal(test.Expect))
				})
			})
		}

		var testsFail = []TestData{
			{"should return an error if the amount is zero", 0, 50.00},
			{"should return an error if the amount is negative", -100.00, 50.00},
		}

		for _, test := range testsFail {
			It(test.Spec, func() {
				result := account.Deposit(test.Input)

				Expect(result).To(HaveOccurred())
				Expect(account.GetBalance()).To(Equal(test.Expect))
			})
		}
	})

	Describe("when withdrawing an amount", func() {
		var testsSuccess = []TestData{
			{"must add balance with amount 25.00", 25.00, 25.00},
			{"must add balance with amount 17.50", 17.50, 32.50},
			{"must add balance with amount 10.58", 10.58, 39.42},
		}

		for _, test := range testsSuccess {
			Context(test.Spec, func() {
				It("should decrease the balance by the given amount", func() {
					result := account.Withdraw(test.Input)

					Expect(result).To(BeNil())
					Expect(account.GetBalance()).To(Equal(test.Expect))
				})
			})
		}

		var tests = []TestData{
			{"should return an error if the amount is zero", 0, 50.00},
			{"should return an error if the amount is negative", -25.00, 50.00},
		}

		for _, test := range tests {
			It(test.Spec, func() {
				result := account.Withdraw(test.Input)

				Expect(result).To(HaveOccurred())
				Expect(account.GetBalance()).To(Equal(test.Expect))
			})
		}
	})

	Describe("when getting the balance", func() {
		It("should return the current balance", func() {
			result := account.GetBalance()

			Expect(result).To(Equal(50.00))
		})
	})

	// Describe("when depositing an amount", func() {
	// 	It("should increase the balance by the given amount", func() {
	// 		result := account.Deposit(100.50)

	// 		Expect(result).To(BeNil())
	// 		Expect(account.GetBalance()).To(Equal(150.50))
	// 	})

	// 	It("should return an error if the amount is zero", func() {
	// 		result := account.Deposit(0)

	// 		Expect(result).To(HaveOccurred())
	// 		Expect(account.GetBalance()).To(Equal(50.00))
	// 	})

	// 	It("should return an error if the amount is negative", func() {
	// 		result := account.Deposit(-100.00)

	// 		Expect(result).To(HaveOccurred())
	// 		Expect(account.GetBalance()).To(Equal(50.00))
	// 	})
	// })

	// Describe("when withdrawing an amount", func() {
	// 	It("should decrease the balance by the given amount", func() {
	// 		result := account.Withdraw(15.00)

	// 		Expect(result).To(BeNil())
	// 		Expect(account.GetBalance()).To(Equal(35.00))
	// 	})

	// 	It("should return an error if the amount is zero", func() {
	// 		result := account.Withdraw(0)

	// 		Expect(result).To(HaveOccurred())
	// 		Expect(account.GetBalance()).To(Equal(50.00))
	// 	})

	// 	It("should return an error if the amount is negative", func() {
	// 		result := account.Withdraw(-25.00)

	// 		Expect(result).To(HaveOccurred())
	// 		Expect(account.GetBalance()).To(Equal(50.00))
	// 	})
	// })

	// Describe("when getting the balance", func() {
	// 	It("should return the current balance", func() {
	// 		result := account.GetBalance()

	// 		Expect(result).To(Equal(50.00))
	// 	})
	// })

	AfterEach(func() {
		account.Reset()
	})
})
