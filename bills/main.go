package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Println("create a new bill name:")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("choose option (a - add item , s - save bill , t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("item name:", reader)
		price, _ := getInput("item price:", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added-", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("enter tip amount ($):", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip should be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added -", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the file", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}

}

func main() {
	mybill := createBill()
	promptOptions(mybill)

	fmt.Println(mybill)

	// mybill := newBill("sehajal's Bill")
	// mybill.addItem("soup", 3.50)
	// mybill.addItem("pasta", 6.50)
	// mybill.addItem("pizza", 11.50)
	// mybill.addItem("kitkat", 0.50)
	// mybill.updateTip(10)

	// mybill.format()
	// fmt.Println(mybill.format())
}
