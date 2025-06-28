package main

import "fmt"

type BankAccount struct {
    balance float64
}

func (b *BankAccount) Deposit(amount float64) {
    b.balance += amount
}

func (b *BankAccount) Withdraw(amount float64) bool {
    if amount > b.balance {
        return false
    }
    b.balance -= amount
    return true
}

func (b BankAccount) Balance() float64 {
    return b.balance
}

func main() {
    account := BankAccount{}
    
    account.Deposit(1000)
    fmt.Println("Баланс после пополнения:", account.Balance())
    
    success := account.Withdraw(500)
    fmt.Println("Снятие 500:", success, "Баланс:", account.Balance())
    
    success = account.Withdraw(600)
    fmt.Println("Снятие 600:", success, "Баланс:", account.Balance())
}