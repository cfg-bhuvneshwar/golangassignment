package main

import (
	"fmt"
	"os"
)

type Credential struct {
	userId   string
	password string
}

var balance int

func UserInputForBankOperations() {
	fmt.Println("Select the operation you want to do :")
	fmt.Println("d -> Deposit Money")
	fmt.Println("w -> Withdraw Money")
	fmt.Println("b -> Request balance")
	fmt.Println("e -> Exit")

	var operation string
	fmt.Scan(&operation)

	if operation == "d" {
		DepositMoney()
	} else if operation == "w" {
		WithdrawMoney()
	} else if operation == "b" {
		RequestBalance()
	} else if operation == "e" {
		os.Exit(1)
	}
}

func WithdrawMoney() {
	if balance == 0 {
		fmt.Println("Insufficent balance")
		UserInputForBankOperations()
	} else {
		fmt.Print("Enter Amount you want to withdraw : ")
		fmt.Scan(&balance)

		fmt.Println("Amount withdraw successfully. Your current balance is : ", balance)
		UserInputForBankOperations()
	}
}

func DepositMoney() {
	fmt.Print("Enter Amount you want to deposit : ")
	fmt.Scan(&balance)

	fmt.Println("Amount deposit successfully. Your current balance is : ", balance)
	UserInputForBankOperations()
}

func RequestBalance() {
	fmt.Println("Current Balance is : ", balance)
	UserInputForBankOperations()
}

func main() {

	fmt.Println("Hi Welcome!!")
	fmt.Println("Please select an option from the menu below")

	var option string

	fmt.Println("l  -> Login")
	fmt.Println("c -> Create New Account")
	fmt.Println("q -> Quit")
	fmt.Scan(&option)

	var id string
	var password string

	if option == "l" {
		fmt.Print("Please enter userid : ")
		fmt.Scan(&id)

		fmt.Print("Please enter your password : ")
		fmt.Scan(&password)

		s := Credential{userId: id, password: password}

		fmt.Println("Your userId is", s.userId, "and password is", password)

		UserInputForBankOperations()

	} else if option == "c" {
		fmt.Print("Please enter userid : ")
		fmt.Scan(&id)

		fmt.Print("Please enter your password : ")
		fmt.Scan(&password)

		s := Credential{userId: id, password: password}

		fmt.Println("Your Account has been created with userId ", s.userId)
	} else if option == "q" {
		os.Exit(1)
	}
}
