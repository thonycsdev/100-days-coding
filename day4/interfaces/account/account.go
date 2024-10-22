package account

import (
	"fmt"
	"github.com/beevik/guid"
	"go-challange/interfaces/logs"
	"time"
)

type Transaction interface {
	make_transaction(credit_account *Account)
	account_deposit(amount float64)
	account_withdraw(amount float64)
	show_balance()
}

type Account struct {
	Id         string    `json:"id"`
	Owner_name string    `json:"owner_name"`
	Balance    float64   `json:"balance"`
	Created_at time.Time `json:"created_at"`
}

func New(owner_name string, balance float64) *Account {
	account := Account{Id: guid.New().StringUpper(), Owner_name: owner_name, Balance: balance, Created_at: time.Now()}
	logs.CreateLog("Account Created", &account)
	return &account
}

func (a *Account) MakeTransaction(credit_account *Account) {

}

func (a *Account) AccountDeposit(amount float64) {
	a.Balance += amount

}

func (a *Account) AccountWithdraw(amount float64) {

}

func (a Account) ShowBalance() {
	fmt.Printf("Account balance: %.2f\n", a.Balance)
}

func (a Account) ShowAccount() {

	fmt.Printf("Account Id: %s\n", a.Id)
	fmt.Printf("Account Owner: %s\n", a.Owner_name)
	fmt.Printf("Account balance: %.2f\n", a.Balance)
	fmt.Printf("Created At: %s\n", a.Created_at)
}
