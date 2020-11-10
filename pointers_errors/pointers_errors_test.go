package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with sufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(5))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(50))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrorInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, expectedResult Bitcoin) {
	t.Helper()
	result := wallet.Balance()
	if expectedResult != result {
		t.Errorf("Expected %s but got %s", expectedResult, result)
	}
}

func assertError(t *testing.T, got error, wanted error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but didn't get one.")
	}
	if got != wanted {
		t.Errorf("Expected %q but got %q", wanted, got)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but didn't expected one")
	}
}