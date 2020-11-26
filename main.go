package main

import (
	"fmt"
	"strconv"
)

func MoneyChangerLogic(price int, pay int) []string {
	returnMoney := pay - price
	allMoneyType := [9]int{1000, 500, 100, 50, 20, 10, 5, 2, 1}
	msgList := []string{}
	msg := ""
	if price < pay {
		fmt.Println("customer will get change", returnMoney, "Baht")
		for i := 0; i < len(allMoneyType); i++ {
			if returnMoney/allMoneyType[i] >= 1 {
				x := "bill"
				if allMoneyType[i] <= 10 {
					x = "coin"
				}
				msg = strconv.Itoa(allMoneyType[i]) + " is " + strconv.Itoa(returnMoney/allMoneyType[i]) + x
				msgList = append(msgList, msg)
				returnMoney %= allMoneyType[i]
			}
		}
		return msgList
	} else if price == pay {
		return msgList
	} else {
		return msgList
	}
}
