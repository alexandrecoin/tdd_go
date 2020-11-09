package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	result := wallet.Balance()
	expectedResult := Bitcoin(10)
	if expectedResult != result {
		t.Errorf("Expected %d but got %d", expectedResult, result)
	}
}