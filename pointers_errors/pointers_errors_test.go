package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	result := wallet.Balance()
	expectedResult := Bitcoin(9)
	if expectedResult != result {
		t.Errorf("Expected %s but got %s", expectedResult, result)
	}
}