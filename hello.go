package main

import "fmt"

type User struct {
	Name    string
	Balance float64
}

type FoodOrder struct {
	User     User
	Order    []string
	Subtotal float64
}

var menu = map[string]float64{
	"Caviar":  600,
	"Boba":    2.49,
	"Burrito": 1.99,
	"Pasta":   2.99,
}

func calculateSubtotal(order []string) (float64, error) {
	var subtotal float64
	for _, item := range order {
		if price, ok := menu[item]; ok {
			subtotal += price
		} else {
			return 0, fmt.Errorf("invalid item: %s", item)
		}
	}
	return subtotal, nil
}

func main() {
	var orders []FoodOrder

	// Prompt the user for their name
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	// Prompt the user for their balance
	var balance float64
	fmt.Print("Enter your balance: ")
	fmt.Scanln(&balance)

	// Create a user object from the struct
	user := User{Name: name, Balance: balance}

	// Prompt the user for their food order
	fmt.Println("Enter your food order, or 'done' to finish:")
	var orderItem string
	var order []string
	for {
		fmt.Scanln(&orderItem)
		if orderItem == "done" {
			break
		}
		order = append(order, orderItem)
	}

	// Calculate the subtotal of the food order
	subtotal, err := calculateSubtotal(order)
	if err != nil {
		fmt.Println(err)
		return
	}

	if user.Balance < subtotal {
		fmt.Println("Insufficient funds")
		return
	}

	// Create a food order object from the struct
	foodOrder := FoodOrder{User: user, Order: order, Subtotal: subtotal}

	// Add the food order to the pending orders
	orders = append(orders, foodOrder)

	// Display a message confirming the order
	fmt.Println(orders)
}
