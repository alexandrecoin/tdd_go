package main

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, expectedResult Bitcoin) {
		t.Helper()
		result := wallet.Balance()
		if expectedResult != result {
			t.Errorf("Expected %s but got %s", expectedResult, result)
		}
	}

	assertError := func(t *testing.T, got error, wanted string) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted an error but didn't get one.")
		}
		if got.Error() != wanted {
			t.Errorf("Expected %q but got %q", wanted, got)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(50))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "Withdrawal is too important for current balance")
	})
}