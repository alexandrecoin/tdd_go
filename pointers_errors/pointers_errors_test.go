package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		result := wallet.Balance()
		expectedResult := Bitcoin(10)
		if expectedResult != result {
			t.Errorf("Expected %s but got %s", expectedResult, result)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))
		result := wallet.Balance()
		expectedResult := Bitcoin(5)
		if expectedResult != result {
			t.Errorf("Expected %s but got %s", expectedResult, result)
		}
	})
}