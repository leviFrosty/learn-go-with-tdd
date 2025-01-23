package pointersErrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		err := wallet.Withdraw(10)
		assetNoError(t, err)
		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(100)
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrorInsufficientFunds)

	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("Got '%s', want '%s'", got, want)
	}
}

func assetNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't receive an error but wanted one")
	}
	if got != want {
		t.Errorf("got error message '%q' want '%q'", got, want)
	}
}
