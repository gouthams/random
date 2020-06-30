package main

import (
	"fmt"
	"github.com/gouthams/helperUtils/phoneUtils"
	logger "github.com/sirupsen/logrus"
)

func main() {

	//Default level to INFO
	logger.SetLevel(logger.InfoLevel)

	//Sample phone number check
	phoneNumber := "sds254-7096!-sdsd"
	logger.Printf("Is phone number: %s easy to dial? ->  %t\n", phoneNumber, phoneUtils.IsEasyDial(phoneNumber))

	//Get the phone number from the user
	var inputNumber string
	fmt.Println("Enter the number to check whether it is easy dial number")
	fmt.Scanln(&inputNumber)
	logger.Printf("Is phone number: %s easy to dial? ->  %t\n", inputNumber, phoneUtils.IsEasyDial(inputNumber))

}
