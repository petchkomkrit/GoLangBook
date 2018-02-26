package main

import (
	"fmt"
)

type VendingMachine struct {
	insertedMoney int
	coins         map[string]int
	items         map[string]int
}

func (m VendingMachine) insertMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) insertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
}

func (m *VendingMachine) selectItem(item string) string {
	change := m.insertedMoney - m.items[item]
	return item + m.change(change)
}

func (m VendingMachine) change(c int) string {
	var str string
	values := [...]int{10, 5, 2, 1}
	coins := [...]string{"T", "F", "TW", "O"}

	for i := 0; c > 0; i++ {
		if c >= values[i] {
			str += ", " + coins[i]
			c -= values[i]
			i--
		}
	}

	return str
}

func (m *VendingMachine) coinReturn() string {
	coins := m.change(m.insertedMoney)
	m.insertedMoney = 0
	return coins[2:len(coins)]
}

func main() {
	var initialCoins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
	var initialItems = map[string]int{"SD": 18, "CC": 12, "DW": 7}
	vm := VendingMachine{coins: initialCoins, items: initialItems}

	fmt.Println("Inserted Money : ", vm.insertMoney())
	vm.insertCoin("T")
	vm.insertCoin("F")
	vm.insertCoin("TW")
	vm.insertCoin("O")
	vm.insertCoin("T")
	vm.insertCoin("T")
	fmt.Println("Inserted Money : ", vm.insertMoney())

	can := vm.selectItem("SD")
	fmt.Println(can)

	/* fmt.Println("Inserted Money : ", vm.insertMoney())
	coin := vm.coinReturn()
	fmt.Println(coin) */
}
