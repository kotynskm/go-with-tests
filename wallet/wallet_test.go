package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	// pull out repeated work and put in helper test
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin){
		// mark as helper
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s expected %s", got, want)
		}
	
	}

	assertError := func(t testing.TB, err error, want error){
		t.Helper()
		// if err == nil {
		// 	t.Error("should have received an error")
		// }

		// fatal will stop the test if it is called
		if err == nil {
			t.Fatal("didn't get error but should have")
		}
		if err != want {
			t.Errorf("got %q wanted %q", err, want)
		}
	}

	assertNoError := func(t testing.TB, err error){
		t.Helper()

		if err != nil {
			t.Fatal("got an error but did not want one")
		}
	}
	t.Run("deposit amount", func(t *testing.T) {
		wallet := Wallet{}
		// then we use our custom type here
		wallet.Deposit(Bitcoin(10))
	
		// got := wallet.Balance()
		fmt.Printf("address of total balance is %v", &wallet.total)
		// then we use our custom type here
		expected := Bitcoin(10)
	
		// if got != expected {
		// 	t.Errorf("got %s expected %s", got, expected)
		// }

		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw amount", func(t *testing.T) {
		wallet := Wallet{total: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(10))

		// got := wallet.Balance()
		expected := Bitcoin(0)

		// if got != expected {
		// 	t.Errorf("got %s expected %s", got, expected)
		// }

		// we want to check also that we did not get an error, which is what we are expecting
		assertNoError(t, err)
		assertBalance(t, wallet, expected)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{startingBalance}

		// we want to return an error if more is taken out than balance allowed
		// now our method must return an error
		err := wallet.Withdraw(Bitcoin(100))

		// assertError(t, err, "cannot withdraw, insufficient funds")
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

		// check if we did not receive an error, if no error then we want to trigger a failure because we should not be able to withdraw more than balance amount
		// if err == nil {
		// 	t.Error("wanted an error but did not get one")
		// }


	})
	
}

